package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserEmail struct {
	ent.Schema
}

func (UserEmail) Fields() []ent.Field {
	return []ent.Field{
		field.
			Int("id").
			Annotations(entproto.Field(1)),
		field.
			String("value").
			Immutable().
			Unique().
			Annotations(entproto.Field(2)),
		field.Time("createdAt").
			Immutable().
			Annotations(entproto.Field(3)),
		field.Bool("isVerified").
			Immutable().
			Annotations(entproto.Field(4)),
	}
}

func (UserEmail) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("user", User.Type).
			Ref("emails").
			Unique().
			Annotations(entproto.Field(400)),
	}
}

func (UserEmail) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(
			entproto.PackageName(GO_PACKAGE),
		),
		entproto.Service(
			entproto.Methods(entproto.MethodAll),
		),
	}
}
