package main

import (
	"log"

	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	schemaPath := "./ent"
	config := &gen.Config{
		Target: "./pkg",
		// Schema:      "",
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
		log.Fatalf("entproto: failed generating protos: %s", err)
	}
}
