package main

import (
    "flag"
    "fmt"
    "golang.org/x/net/context"
    "google.golang.org/grpc"
    google_pb "github.com/golang/protobuf/ptypes/empty"
    health "test/pb/health"
)

var (
    tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
    serverAddr = flag.String("server_addr", "127.0.0.1:9026", "The server address in the format of host:port")
)

func isHealthy(client health.HealthClient, req *google_pb.Empty) (*health.HealthCheckResponse, error){
    return client.Check(context.Background(), req)
}

func main() {
    var opts []grpc.DialOption
    opts = append(opts, grpc.WithInsecure())
    conn, err := grpc.Dial(*serverAddr, opts...)
    if(err != nil) {
        fmt.Println(err)
        conn.Close()
    }
    req := &google_pb.Empty{}
    client := health.NewHealthClient(conn)
    res,err := isHealthy(client, req)
    fmt.Println(res)
}

