package ent

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
// message PortalLegal {
//	uint32 portalId = 1;
//	uint32 timestamp = 2;
//
//	string privacyPolicy = 3;
//	string termsOfService = 4;
//	string copyright = 5;
//	string onlineTrainingAgreement = 6;
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
// message PortalMetadata {
//	uint32 portalId = 1;
//	uint32 timestamp = 2;
//
//	Kind kind = 3;
//	enum Kind {
//		KIND_DEVELOPMENT = 0;
//		KIND_INTERNAL = 1;
//		KIND_CONTENT_PARTNER = 2;
//		KIND_DISTRIBUTION_PARTNER = 3;
//		KIND_CUSTOMER = 4;
//		KIND_COMPLISPACE = 5;
//		KIND_TEAM = 6;
//	}
//
//	Lifecycle lifecycle = 4;
//	enum Lifecycle {
//		LIFECYCLE_TRIAL = 0;
//		LIFECYCLE_ONBOARD = 1;
//		LIFECYCLE_LIVE = 2;
//		LIFECYCLE_EXPIRED_TRIAL = 3;
//		LIFECYCLE_SUSPENDED = 4;
//		LIFECYCLE_CANCELLED = 5;
//		LIFECYCLE_TEST = 6;
//		LIFECYCLE_SAMPLE = 7;
//		LIFECYCLE_TEMPLATE = 8;
//	}
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
//
// message PortalMilestones {
//	uint32 portalId = 1;
//
//	uint32 goliveDate = 2;
//	uint32 conversionDate = 3;
//	uint32 expiryDate = 4;
//	uint32 cancelledDate = 5;
// }
