package db

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/4Sao1Sad/yuemnoi-activity/internal/config"
	"github.com/4Sao1Sad/yuemnoi-activity/internal/handler"
	"github.com/4Sao1Sad/yuemnoi-activity/internal/repository"
	activity "github.com/4Sao1Sad/yuemnoi-activity/proto/activity"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

var MongoClient *mongo.Client

// InitDB initializes a connection to the MongoDB database
func InitDB(cfg *config.Config) (*mongo.Database, error) {
	// Build the MongoDB connection URI

	// Connect to MongoDB
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(cfg.Mongo.Url).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("failed to ping MongoDB: %v", err)
		return nil, err
	}

	log.Println("MongoDB connection established")
	MongoClient = client // Store the client globally if needed

	// Return the MongoDB database instance
	return client.Database(cfg.Db.Database), nil
}

func ServerInit(cfg *config.Config, db *mongo.Database) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer func() {
		listen.Close()
	}()

	fmt.Printf("Go gRPC server on port %v!", cfg.Port)
	grpcServer := grpc.NewServer()

	// Initialize the ActivityLog repository using MongoDB
	activityLogRepo := repository.NewActivityLogRepository(db)
	activityLogServer := handler.NewActivityLogGRPC(activityLogRepo)

	// Register the ActivityLogService
	activity.RegisterActivityLogServiceServer(grpcServer, activityLogServer)

	// Enable gRPC reflection for tools like grpcurl
	reflection.Register(grpcServer)

	err = grpcServer.Serve(listen)
	if err != nil {
		return fmt.Errorf("error to serve: %v", err.Error())
	}

	return nil
}
