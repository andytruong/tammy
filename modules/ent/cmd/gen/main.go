package main

import (
	"log"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	schemaPath := "./schema"
	config := &gen.Config{
		// Target: "./store",
		// Schema: "../../../schema",
		// Package:     "",
		// Header:      "",
		// Storage:     nil,
		// IDType:      nil,
		// Templates:   nil,
		// Features:    nil,
		// Hooks:       nil,
		// Annotations: nil,
	}

	graph, err := entc.LoadGraph(schemaPath, config)
	if err != nil {
		log.Fatalf("entproto: failed loading ent graph: %v", err)
	}

	if err := entproto.Generate(graph); err != nil {
		log.Fatalf(err.Error())
	}
}
