package ent

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type AccountField struct {
	ent.Schema
}

func (AccountField) Fields() []ent.Field {
	return []ent.Field{
		field.Uint32("id").Annotations(entproto.Field(1)),
		field.Uint32("fid").Annotations(entproto.Field(2)).Comment("field ID"),
		field.String("key").Annotations(entproto.Field(3)).Comment("field name"),
		field.String("value").Annotations(entproto.Field(4)).Comment("field value"),
	}
}

func (AccountField) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("account", Account.Type).
			Ref("fields").
			Required().
			Unique().
			Annotations(entproto.Field(400)),
	}
}

func (AccountField) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(entproto.PackageName("ent")),
	}
}
