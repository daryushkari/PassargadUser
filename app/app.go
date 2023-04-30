package app

import (
	"PassargadUser/config"
	"PassargadUser/delivery/rest"
	"PassargadUser/entities/domain"
	"PassargadUser/pkg/sqlite"
	"PassargadUser/repository"
	"github.com/gin-gonic/gin"
	"log"
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

	r := gin.Default()
	//r.Use(gin.Recovery())
	//r.Use(gin.Logger())
	r.POST("/create", rest.CreateUser)
	r.POST("/login", rest.LoginUser)
	r.Run()

	//lis, err := net.Listen("tcp", cnf.ExternalExpose.GrpcPort)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//s := userDelivery.Server{}
	//
	//grpcServer := grpc.NewServer()
	//
	//userPb.RegisterUserServer(grpcServer, &s)
	//
	//if err = grpcServer.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//}
}
