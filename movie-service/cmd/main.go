package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	"tech-test-movies/movie-service/internal/adapters/db"
	grpcAdapter "tech-test-movies/movie-service/internal/adapters/grpc"
	"tech-test-movies/movie-service/internal/core/domain"
	"tech-test-movies/movie-service/internal/core/services"
	"tech-test-movies/pb"
)

func main() {
	// 1. Conectar ao MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("Erro ao conectar no MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())

	dbName := "moviesdb"
	collection := client.Database(dbName).Collection("movies")

	// 2. Semear o banco de dados (apenas se estiver vazio)
	seedDatabase(collection)

	// 3. Configurar o Repositório e o Serviço
	repo := db.NewMongoRepository(collection)
	service := services.NewMovieService(repo)
	grpcServerHandler := grpcAdapter.NewGrpcServer(service)

	// 4. Iniciar o servidor gRPC
	port := ":50051"
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Falha ao escutar porta: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterMovieServiceServer(server, grpcServerHandler)

	fmt.Printf("Movie Service rodando na porta %s (gRPC)...\n", port)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Falha ao servir gRPC: %v", err)
	}
}

// seedDatabase semear o banco de dados
func seedDatabase(collection *mongo.Collection) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Println("Erro ao verificar coleção:", err)
		return
	}

	if count > 0 {
		fmt.Println("Banco de dados já contém filmes. Pulando seed.")
		return
	}

	fmt.Println("Banco vazio. Lendo movies.json para semear dados...")
	file, err := os.ReadFile("data/movies.json")
	if err != nil {
		log.Println("Aviso: arquivo data/movies.json não encontrado. Seed cancelado.")
		return
	}

	var movies []domain.Movie
	if err := json.Unmarshal(file, &movies); err != nil {
		log.Println("Erro ao fazer parse do JSON:", err)
		return
	}

	var docs []interface{}
	for _, m := range movies {
		docs = append(docs, m)
	}

	if len(docs) > 0 {
		_, err = collection.InsertMany(ctx, docs)
		if err != nil {
			log.Println("Erro ao inserir filmes:", err)
		} else {
			fmt.Printf("Seed concluído! %d filmes inseridos com sucesso.\n", len(docs))
		}
	}
}
