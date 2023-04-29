package app

import (
	"PassargadUser/config"
	"PassargadUser/domain"
	"PassargadUser/pkg/sqlite"
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

	err = sDB.AutoMigrate(&domain.User{})
	if err != nil {
		log.Fatalf("migration failed: %v", err.Error())
	}

	r := gin.Default()
	r.GET("/tab", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nine",
		})
	})
	r.Run()
	// listen and serve on 0.0.0.0:8080

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
