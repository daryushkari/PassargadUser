package app

import (
	userPb "PassargadUser/api/pb/proto"
	"PassargadUser/app/middleware"
	"PassargadUser/config"
	grpcGateway "PassargadUser/delivery/grpc"
	"PassargadUser/delivery/rest"
	"PassargadUser/entities/domain"
	"PassargadUser/pkg/sqlite"
	"PassargadUser/repository"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net"
)

func InitApp() {
	cfg, err := config.Get("./config.json", config.ProductionEnv)
	if err != nil {
		log.Fatalf("config not available with error: %v", err.Error())
	}
	err, sDB := sqlite.Get(cfg.Database.Name)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err.Error())
	}

	repository.UsrRepo.InitDB(sDB)

	err = sDB.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("migration failed: %v", err.Error())
	}

	go RouteGRPC(cfg)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.JWTVerify())
	AddUserRouter(r)
	r.Run()
}

func RouteGRPC(cnf *config.EnvConfig) {
	lis, err := net.Listen("tcp", cnf.ExternalExpose.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpcGateway.Server{}
	grpcServer := grpc.NewServer()
	userPb.RegisterUserServer(grpcServer, &s)

	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func AddUserRouter(r *gin.Engine) {
	userRouter := r.Group("/user")
	{
		userRouter.POST("/create", rest.CreateUser)
		userRouter.POST("/login", rest.LoginUser)
		userRouter.GET("/info", rest.GetUserInfo)
		userRouter.POST("/update", rest.UpdateUser)
		userRouter.DELETE("/delete", rest.DeleteUser)
	}
}
