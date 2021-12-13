package main

import (
	"github.com/chaiyawatkit/ginney"
	"upload-file/app/constants"
	"upload-file/app/env"

	_ "upload-file/app/docs"
	_healthcheck "upload-file/app/layers/deliveries/http/health_check"

	_amazonHttp "upload-file/app/layers/deliveries/http/amazon_s3"
	_amazonRepo "upload-file/app/layers/repositories/amazon_s3"
	_amazonUseCase "upload-file/app/layers/usecases/amazon_s3"

	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pseidemann/finish"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Trainning Golang MicroService
// @version 0.1.0
// @description Trainning Golang MicroService.
// @schemes http https
// @BasePath /

func main() {
	env.Init()
	//boredom.InitL(env.LogLevel)
	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())

	_healthcheck.NewEndpointHTTPHandler(ginEngine)
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ginEngine.Use(ginney.LogWithCorrelationIdMiddleware(gin.DefaultWriter, constants.NoLoggedRoutes))
	ginEngine.Use(ginney.MicroServiceCorrelationIdMiddleware())
	ginEngine.Use(ginney.FromGinContextToContextMiddleware())

	//dbConn := database.ConnectDB()
	//defer func() {
	//	_ = dbConn.Close()
	//}()
	//database.DBMigration()

	amazonRepo := _amazonRepo.Init()
	amazonUseCase := _amazonUseCase.Init(amazonRepo)
	_amazonHttp.NewEndpointHttpHandler(ginEngine, amazonUseCase)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: ginEngine,
	}

	timeOut, err := strconv.Atoi(os.Getenv("GRACEFUL_TIMEOUT"))
	if err != nil {
		timeOut = 30 // second
	}

	graceful := &finish.Finisher{Timeout: time.Duration(timeOut) * time.Second}
	graceful.Add(srv)

	go func() {
		err := srv.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	graceful.Wait()
}
