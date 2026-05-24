package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api-gateway/docs"
	"api-gateway/internal/handler"
	"pb"
)

// @title API de Gerenciamento de Filmes
// @version 1.0
// @description API para gerenciamento de filmes
// @host localhost:8080
// @BasePath /
func main() {
	// 1. Configurar o Cliente gRPC para se comunicar com o Movie Service
	grpcTarget := os.Getenv("GRPC_MOVIE_SERVICE_TARGET")
	if grpcTarget == "" {
		grpcTarget = "localhost:50051"
	}

	conn, err := grpc.Dial(grpcTarget, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar no gRPC: %v", err)
	}
	defer conn.Close()

	grpcClient := pb.NewMovieServiceClient(conn)

	// 2. Configurar o Router Gin e os Handlers
	router := gin.Default()

	movieHandler := handler.NewMovieHandler(grpcClient)

	// 3. Configurar Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/movies/", movieHandler.ListMovies)
	router.GET("/movies/:id", movieHandler.GetMovie)
	router.POST("/movies/", movieHandler.CreateMovie)
	router.DELETE("/movies/:id", movieHandler.DeleteMovie)

	// 4. Iniciar o Servidor HTTP
	port := ":8080"
	log.Printf("API Gateway rodando na porta %s...", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("Falha ao rodar API Gateway: %v", err)
	}
}
