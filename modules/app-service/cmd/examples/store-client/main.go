package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"tammy/ent/proto/entpb"
)

// go run tammy/app/cmd/examples/store-client

func main() {
	ctx := context.Background()

	// https://pkg.go.dev/google.golang.org/grpc#DialOption
	conn, err := grpc.Dial(
		"localhost:8484",
		grpc.WithTimeout(7*time.Second),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithInsecure(),
	)

	if nil != err {
		panic("connection error: " + err.Error())
	}

	client := entpb.NewPortalServiceClient(conn)

	testCreate(ctx, client)
}

func testCreate(ctx context.Context, client entpb.PortalServiceClient) {
	portal, err := client.Create(ctx, &entpb.CreatePortalRequest{
		Portal: &entpb.Portal{
			Slug:     "qa",
			IsActive: true,
		},
	})

	if nil != err {
		panic("error: " + err.Error())
	}

	fmt.Println("portal created: ", portal)

	testGet(ctx, client, portal)
	testDelete(ctx, client, portal)
}

func testGet(ctx context.Context, client entpb.PortalServiceClient, portal *entpb.Portal) {
	portal, err := client.Get(ctx, &entpb.GetPortalRequest{Id: portal.Id})

	if nil != err {
		panic("error: " + err.Error())
	}

	fmt.Println("portal load: ", portal)
}

func testDelete(ctx context.Context, client entpb.PortalServiceClient, portal *entpb.Portal) {
	empty, err := client.Delete(ctx, &entpb.DeletePortalRequest{Id: portal.Id})
	if nil != err {
		panic("error: " + err.Error())
	}

	fmt.Println("portal delete: ", empty.String())
}
