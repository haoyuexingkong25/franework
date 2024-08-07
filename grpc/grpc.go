package grpc

import (
	"fmt"
	"github.com/haoyuexingkong25/franework/nacos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"user/etc"
)

func Grpc(server func(s *grpc.Server)) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", etc.Conf.Service.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	server(s)

	err = nacos.ServiceRegister(etc.Conf.Service.Name, etc.Conf.Service.Port)

	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(s, healthcheck)
	if err != nil {
		log.Fatalf("failed to franework grpc ServiceRegister err: %v", err)
		return
	}
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		return
	}
}
