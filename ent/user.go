package ent

import (
	"time"
	
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.
			Uint32("id").
			Annotations(entproto.Field(1)).
			Immutable(),
		field.Bool("isActive").
			Default(true).
			Annotations(entproto.Field(2)),
		field.Time("createdAt").
			Immutable().
			Default(time.Now).
			Annotations(entproto.Field(3)),
		field.Time("updatedAt").
			Immutable().
			Default(time.Now).
			Annotations(entproto.Field(4)),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			To("password", UserPassword.Type).
			Unique().
			Annotations(entproto.Field(400)),
		edge.
			To("accounts", Account.Type).
			Annotations(entproto.Field(401)),
		edge.
			To("emails", UserEmail.Type).
			Annotations(entproto.Field(402)),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(
			entproto.PackageName("ent"),
		),
	}
}
