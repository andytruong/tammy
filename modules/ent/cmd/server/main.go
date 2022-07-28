package main

import (
	"context"
	"fmt"
	"log"
	"net"

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
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", ":8484"))
	if err = server.Serve(listener); nil != err {
		log.Fatalf("server error: %v", err)
	}
}

func buildServer(client *ent.Client) *grpc.Server {
	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	entpb.RegisterUserServiceServer(server, entpb.NewUserService(client))
	entpb.RegisterUserPasswordServiceServer(server, entpb.NewUserPasswordService(client))
	entpb.RegisterAccountServiceServer(server, entpb.NewAccountService(client))
	entpb.RegisterAccountFieldServiceServer(server, entpb.NewAccountFieldService(client))
	entpb.RegisterPortalServiceServer(server, entpb.NewPortalService(client))

	return server
}
