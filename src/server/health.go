package main

import (
    context "golang.org/x/net/context"
    google_pb "github.com/golang/protobuf/ptypes/empty"
    health "test/pb/health"
)

type HealthServer struct {}

func (s HealthServer) Check(ctx context.Context, req *google_pb.Empty) (*health.HealthCheckResponse, error) {
    return &health.HealthCheckResponse{health.HealthCheckResponse_SERVING},nil
}
