package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UserPassword struct {
	ent.Schema
}

func (UserPassword) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Annotations(entproto.Field(1)),
		field.String("hashedValue").Annotations(entproto.Field(3)),
		field.Bool("isActive").Annotations(entproto.Field(4)),
		field.Uint32("createdAt").Annotations(entproto.Field(5)),
	}
}

func (UserPassword) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("password").
			Unique().
			Required().
			Annotations(entproto.Field(400)),
	}
}

func (UserPassword) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(
			entproto.PackageName(GO_PACKAGE),
		),
		entproto.Service(
			entproto.Methods(entproto.MethodAll),
		),
	}
}
