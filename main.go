package main

import (
    "log"
    "net"
    "net/http"

    "github.com/graphql-go/graphql"
    "github.com/graphql-go/handler"
    "google.golang.org/grpc"

    "paper_social/notification_service/graph"
    "paper_social/notification_service/grpc"
    "paper_social/notification_service/metrics"
    "paper_social/notification_service/queue"
)

func main() {
    go startGRPCServer()

    schema := graph.GetSchema()
    h := handler.New(&handler.Config{
        Schema: &schema,
        Pretty: true,
    })

    http.Handle("/graphql", h)
    log.Println("GraphQL server running on http://localhost:8081/graphql")

    metrics.Init()
    http.Handle("/metrics", metrics.Handler())
    log.Println("Metrics available on http://localhost:2112/metrics")

    queue.StartWorkerPool(5)

    log.Fatal(http.ListenAndServe(":8081", nil))
}

func startGRPCServer() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }
    grpcServer := grpc.NewServer()
    grpc.Register(grpcServer)
    log.Println("gRPC server running on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
