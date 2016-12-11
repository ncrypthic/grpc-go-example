package main

import (
    "flag"
    "net"
    "fmt"
    grpc "google.golang.org/grpc"
    pb "test/pb/health"
)

var (
    tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
    serverAddr = flag.String("server_addr", "127.0.0.1:9026", "The server address in the format of host:port")
    port       = flag.Int("port", 9026, "The server port")
)

func newServer() pb.HealthServer {
    return HealthServer{}
}

func main() {
    flag.Parse()
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
	fmt.Println("failed to listen: %v", err)
    }
    var opts []grpc.ServerOption
    grpcServer := grpc.NewServer(opts...)
    pb.RegisterHealthServer(grpcServer, newServer())
    fmt.Printf("Serving on %d\n", *port)
    grpcServer.Serve(lis)
}

