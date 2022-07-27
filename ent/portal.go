package ent

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Portal struct {
	ent.Schema
}

func (Portal) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug"),
	}
}

func (Portal) Fields() []ent.Field {
	return []ent.Field{
		field.
			Uint32("id").
			Immutable().
			Annotations(entproto.Field(1)),
		field.
			Time("createdAt").Default(time.Now).
			Immutable().
			Annotations(entproto.Field(2)),
		field.
			Time("updatedAt").Default(time.Now).
			Immutable().
			Annotations(entproto.Field(3)),
		field.
			Bool("isActive").
			Default(false).
			Annotations(entproto.Field(4)),
		field.
			String("slug").
			Unique().
			Immutable().
			Annotations(entproto.Field(5)),
	}
}

func (Portal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("members", Account.Type).
			Annotations(entproto.Field(400)),
		edge.
			To("metadata", PortalMetadata.Type).
			Annotations(entproto.Field(401)),
		edge.
			To("legal", PortalLegal.Type).
			Annotations(entproto.Field(402)),
	}
}

func (Portal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(
			entproto.PackageName("ent"),
		),
	}
}
