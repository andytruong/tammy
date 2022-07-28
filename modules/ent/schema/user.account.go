package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Account struct {
	ent.Schema
}

func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Annotations(entproto.Field(1)),
		field.Bool("isActive").Annotations(entproto.Field(3)),
		field.Bool("isBlocked").Annotations(entproto.Field(4)),
	}
}

func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("accounts").
			Required().
			Unique().
			Annotations(entproto.Field(400)),
		edge.
			From("portal", Portal.Type).
			Ref("members").
			Annotations(entproto.Field(401)),
		edge.To("fields", AccountField.Type).
			Annotations(entproto.Field(402)),
	}
}

func (Account) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(
			entproto.PackageName(GO_PACKAGE),
		),
		entproto.Service(
			entproto.Methods(entproto.MethodAll),
		),
	}
}