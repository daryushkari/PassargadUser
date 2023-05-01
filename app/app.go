package app

import (
	userPb "PassargadUser/api/pb/proto"
	"PassargadUser/app/middleware"
	"PassargadUser/config"
	grpcGateway "PassargadUser/delivery/grpc"
	"PassargadUser/delivery/rest"
	"PassargadUser/entities/domain"
	"PassargadUser/pkg/jtrace"
	"PassargadUser/pkg/postgresql"
	"PassargadUser/repository"
	"context"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

const (
	JaegerURL = "http://localhost:14268/api/traces"
)

func InitApp() {
	cfg, err := config.Get("./config.json", config.ProductionEnv)
	if err != nil {
		log.Fatalf("config not available with error: %v", err.Error())
	}

	err = config.SetSecret("secret-config.json")
	if err != nil {
		log.Fatalf("secret config not available with error: %v", err.Error())
	}

	err, sDB := postgresql.Get(cfg)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err.Error())
	}

	err, tp := jtrace.TracerProvider(JaegerURL)
	if err != nil {
		log.Fatal(err)
	}
	otel.SetTracerProvider(tp)
	defer tp.Shutdown(context.Background())

	repository.UsrRepo.InitDB(sDB)

	err = sDB.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("migration failed: %v", err.Error())
	}

	ginLogPath := "gin.log"
	f, err := os.Create(ginLogPath)
	if err != nil {
		log.Fatalf("log file creation failed: %v", err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f)

	go RouteGRPC(cfg)

	r := gin.Default()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())
	r.Use(middleware.JWTVerify())
	AddUserRouter(r)
	err = r.Run()
	if err != nil {
		log.Println("rest server failed to run")
	}
}

func RouteGRPC(cnf *config.EnvConfig) {
	log.Println("starting grpc server at ", cnf.ExternalExpose.GrpcPort)
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
