package schema

import (
	"time"

	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type PortalLegal struct {
	ent.Schema
}

func (PortalLegal) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Annotations(entproto.Field(1)),
		field.
			Time("createdAt").
			Annotations(entproto.Field(2)).
			Immutable(),
		field.Time("updatedAt").Annotations(entproto.Field(3)),
		field.String("privacyPolicy").Annotations(entproto.Field(4)),
		field.String("termOfService").Annotations(entproto.Field(5)),
		field.String("copyright").Annotations(entproto.Field(6)),
		field.String("onlineTrainingAgreement").Annotations(entproto.Field(7)),
	}
}

func (PortalLegal) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("portal", Portal.Type).
			Ref("legal").
			Required().
			Unique().
			Annotations(entproto.Field(400)),
	}
}

func (PortalLegal) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(
			entproto.PackageName(GO_PACKAGE),
		),
		entproto.Service(entproto.Methods(entproto.MethodAll)),
	}
}

type PortalMetadata struct {
	ent.Schema
}

func (PortalMetadata) Fields() []ent.Field {
	return []ent.Field{
		field.
			Int("id").
			Annotations(entproto.Field(1)),
		field.
			Time("createdAt").
			Immutable().
			Default(time.Now).
			Annotations(entproto.Field(2)),
		field.
			Time("updatedAt").
			Immutable().
			Default(time.Now).
			Annotations(entproto.Field(3)),
		field.
			Enum("kind").
			Values(
				"KIND_DEVELOPMENT",
				"KIND_INTERNAL",
				"KIND_CONTENT_PARTNER",
				"KIND_DISTRIBUTION_PARTNER",
				"KIND_CUSTOMER",
				"KIND_TEAM",
				"KIND_BIG_PARTNER",
			).
			Default("KIND_DEVELOPMENT").
			Annotations(
				entproto.Field(4),
				entproto.Enum(map[string]int32{
					"KIND_DEVELOPMENT":          0,
					"KIND_INTERNAL":             1,
					"KIND_CONTENT_PARTNER":      2,
					"KIND_DISTRIBUTION_PARTNER": 3,
					"KIND_CUSTOMER":             4,
					"KIND_TEAM":                 5,
					"KIND_BIG_PARTNER":          6,
				}),
			),
		field.
			Enum("lifecycle").
			Default("LIFECYCLE_TRIAL").
			Values(
				"LIFECYCLE_TRIAL",
				"LIFECYCLE_ONBOARD",
				"LIFECYCLE_LIVE",
				"LIFECYCLE_EXPIRED_TRIAL",
				"LIFECYCLE_SUSPENDED",
				"LIFECYCLE_CANCELLED",
				"LIFECYCLE_TEST",
				"LIFECYCLE_SAMPLE",
				"LIFECYCLE_TEMPLATE",
			).
			Annotations(
				entproto.Field(5),
				entproto.Enum(map[string]int32{
					"LIFECYCLE_TRIAL":         0,
					"LIFECYCLE_ONBOARD":       1,
					"LIFECYCLE_LIVE":          2,
					"LIFECYCLE_EXPIRED_TRIAL": 3,
					"LIFECYCLE_SUSPENDED":     4,
					"LIFECYCLE_CANCELLED":     5,
					"LIFECYCLE_TEST":          6,
					"LIFECYCLE_SAMPLE":        7,
					"LIFECYCLE_TEMPLATE":      8,
				}),
			),
	}
}

func (PortalMetadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.
			From("portal", Portal.Type).
			Ref("metadata").
			Required().
			Unique().
			Annotations(entproto.Field(400)),
	}
}

func (PortalMetadata) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(
			entproto.PackageName(GO_PACKAGE),
		),
		entproto.Service(
			entproto.Methods(entproto.MethodAll),
		),
	}
}

// message PortalMetadata {
//	uint32 portalId = 1;
//	uint32 timestamp = 2;
//
//	Kind kind = 3;
//	Lifecycle lifecycle = 4;
//
//	Edges edges = 5;
//	message Edges {
//		uint32 userId = 1;
//		uint32 partnerId = 2;
//	}
//
//	CreationChannel channel = 6;
//	enum CreationChannel {
//		DEVELOPER = 0;
//		INTERNAL = 1;
//		REFERRAL_PARTNER = 2;
//		DISTRIBUTION_PARTNER = 3;
//		SDR_ACCOUNT_EXEC = 4;
//		EXISTING_CUSTOMER = 5;
//		DIRECT_OR_INBOUND = 6;
//		PLATFORM_PARTNER = 7;
//		PORTAL_LAUNCHER = 8;
//		CONTENT_PARTNER = 9;
//	}
// }

// syntax = "proto3";
//
// package tammy.protobuf;
// option go_package = "./pkg/pb;pb";
//
// message PortalInformation {
//	uint32 portalId = 1;
//	uint32 timestamp = 2;
//
//	string name = 3;
//	string companyLocation = 4;
// }
//
// message PortalCreationConfig {
//	uint32 portalId = 1;
//	uint32 timestamp = 2;
//
//	PortalCreationConfig portalCreationConfig = 3;
//	message PortalCreationConfig {
//		string tagline = 1;
//		string taglineSecondary = 2;
//	}
// }
//
// message PortalAuthConfig {
//	uint32 portalId = 1;
//	uint32 timestamp = 2;
//
//	bool loginDisabled = 3;
//	bool registerDisabled = 4;
//	bool allowMasquerading = 5;
//	string loginTagline = 6;
//	string loginTaglineSecondary = 7;
//	string resetPasswordUrl = 8;
// }
//
// message PortalSignupConfig {
//	uint32 portalId = 1;
//	uint32 timestamp = 2;
//
//	bool enableContentPackageSignupFlow = 3;
//	bool enableLearnerSignupFlow = 4;
//	bool enableSimplifiedTermAgreement = 5;
//	bool noFocusTerms = 6;
//	bool disablePackageOverviewInSignupApp = 7;
//	bool enableCompanySizeStepper = 8;
//	string tagline = 9;
//	string taglineSecondary = 10;
// }
//
//
// message PortalMilestones {
//	uint32 portalId = 1;
//
//	uint32 goliveDate = 2;
//	uint32 conversionDate = 3;
//	uint32 expiryDate = 4;
//	uint32 cancelledDate = 5;
// }
