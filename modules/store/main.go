package main

import (
	"context"
	"log"
	"net"
	"time"
	
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"tammy/ent"
	"tammy/ent/proto/entpb"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	
	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	
	server := buildServer(client)
	listener, err := net.Listen("tcp", "localhost:8484")
	if err != nil {
		log.Fatalf("failed net listener: %v", err)
	} else {
		log.Println("listen localhost:8484")
	}
	
	if err = server.Serve(listener); nil != err {
		log.Fatalf("server error: %v", err)
	}
}

func buildServer(client *ent.Client) *grpc.Server {
	opts := []grpc.ServerOption{
		grpc.MaxConcurrentStreams(12),
		// grpc.MaxRecvMsgSize(4),
		// Setting this to zero (default) will disable workers and spawn a new goroutine for each stream.
		grpc.NumStreamWorkers(7),
		grpc.ConnectionTimeout(12 * time.Second),
	}
	server := grpc.NewServer(opts...)
	
	entpb.RegisterUserServiceServer(server, entpb.NewUserService(client))
	entpb.RegisterUserPasswordServiceServer(server, entpb.NewUserPasswordService(client))
	entpb.RegisterAccountServiceServer(server, entpb.NewAccountService(client))
	entpb.RegisterAccountFieldServiceServer(server, entpb.NewAccountFieldService(client))
	entpb.RegisterPortalServiceServer(server, entpb.NewPortalService(client))
	
	return server
}
