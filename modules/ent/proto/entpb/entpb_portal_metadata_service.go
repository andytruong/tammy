// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	strconv "strconv"
	strings "strings"
	ent "tammy/ent"
	portal "tammy/ent/portal"
	portalmetadata "tammy/ent/portalmetadata"
)

// PortalMetadataService implements PortalMetadataServiceServer
type PortalMetadataService struct {
	client *ent.Client
	UnimplementedPortalMetadataServiceServer
}

// NewPortalMetadataService returns a new PortalMetadataService
func NewPortalMetadataService(client *ent.Client) *PortalMetadataService {
	return &PortalMetadataService{
		client: client,
	}
}

func toProtoPortalMetadata_Kind(e portalmetadata.Kind) PortalMetadata_Kind {
	if v, ok := PortalMetadata_Kind_value[strings.ToUpper(string(e))]; ok {
		return PortalMetadata_Kind(v)
	}
	return PortalMetadata_Kind(0)
}

func toEntPortalMetadata_Kind(e PortalMetadata_Kind) portalmetadata.Kind {
	if v, ok := PortalMetadata_Kind_name[int32(e)]; ok {
		entVal := map[string]string{
			"KIND_DEVELOPMENT":          "KIND_DEVELOPMENT",
			"KIND_INTERNAL":             "KIND_INTERNAL",
			"KIND_CONTENT_PARTNER":      "KIND_CONTENT_PARTNER",
			"KIND_DISTRIBUTION_PARTNER": "KIND_DISTRIBUTION_PARTNER",
			"KIND_CUSTOMER":             "KIND_CUSTOMER",
			"KIND_TEAM":                 "KIND_TEAM",
			"KIND_BIG_PARTNER":          "KIND_BIG_PARTNER",
		}[v]
		return portalmetadata.Kind(entVal)
	}
	return ""
}

func toProtoPortalMetadata_Lifecycle(e portalmetadata.Lifecycle) PortalMetadata_Lifecycle {
	if v, ok := PortalMetadata_Lifecycle_value[strings.ToUpper(string(e))]; ok {
		return PortalMetadata_Lifecycle(v)
	}
	return PortalMetadata_Lifecycle(0)
}

func toEntPortalMetadata_Lifecycle(e PortalMetadata_Lifecycle) portalmetadata.Lifecycle {
	if v, ok := PortalMetadata_Lifecycle_name[int32(e)]; ok {
		entVal := map[string]string{
			"LIFECYCLE_TRIAL":         "LIFECYCLE_TRIAL",
			"LIFECYCLE_ONBOARD":       "LIFECYCLE_ONBOARD",
			"LIFECYCLE_LIVE":          "LIFECYCLE_LIVE",
			"LIFECYCLE_EXPIRED_TRIAL": "LIFECYCLE_EXPIRED_TRIAL",
			"LIFECYCLE_SUSPENDED":     "LIFECYCLE_SUSPENDED",
			"LIFECYCLE_CANCELLED":     "LIFECYCLE_CANCELLED",
			"LIFECYCLE_TEST":          "LIFECYCLE_TEST",
			"LIFECYCLE_SAMPLE":        "LIFECYCLE_SAMPLE",
			"LIFECYCLE_TEMPLATE":      "LIFECYCLE_TEMPLATE",
		}[v]
		return portalmetadata.Lifecycle(entVal)
	}
	return ""
}

// toProtoPortalMetadata transforms the ent type to the pb type
func toProtoPortalMetadata(e *ent.PortalMetadata) (*PortalMetadata, error) {
	v := &PortalMetadata{}
	createdAt := timestamppb.New(e.CreatedAt)
	v.CreatedAt = createdAt
	id := int64(e.ID)
	v.Id = id
	kind := toProtoPortalMetadata_Kind(e.Kind)
	v.Kind = kind
	lifecycle := toProtoPortalMetadata_Lifecycle(e.Lifecycle)
	v.Lifecycle = lifecycle
	updatedAt := timestamppb.New(e.UpdatedAt)
	v.UpdatedAt = updatedAt
	if edg := e.Edges.Portal; edg != nil {
		id := int64(edg.ID)
		v.Portal = &Portal{
			Id: id,
		}
	}
	return v, nil
}

// toProtoPortalMetadataList transforms a list of ent type to a list of pb type
func toProtoPortalMetadataList(e []*ent.PortalMetadata) ([]*PortalMetadata, error) {
	var pbList []*PortalMetadata
	for _, entEntity := range e {
		pbEntity, err := toProtoPortalMetadata(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements PortalMetadataServiceServer.Create
func (svc *PortalMetadataService) Create(ctx context.Context, req *CreatePortalMetadataRequest) (*PortalMetadata, error) {
	portalmetadata := req.GetPortalMetadata()
	m, err := svc.createBuilder(portalmetadata)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoPortalMetadata(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements PortalMetadataServiceServer.Get
func (svc *PortalMetadataService) Get(ctx context.Context, req *GetPortalMetadataRequest) (*PortalMetadata, error) {
	var (
		err error
		get *ent.PortalMetadata
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetPortalMetadataRequest_VIEW_UNSPECIFIED, GetPortalMetadataRequest_BASIC:
		get, err = svc.client.PortalMetadata.Get(ctx, id)
	case GetPortalMetadataRequest_WITH_EDGE_IDS:
		get, err = svc.client.PortalMetadata.Query().
			Where(portalmetadata.ID(id)).
			WithPortal(func(query *ent.PortalQuery) {
				query.Select(portal.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoPortalMetadata(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements PortalMetadataServiceServer.Update
func (svc *PortalMetadataService) Update(ctx context.Context, req *UpdatePortalMetadataRequest) (*PortalMetadata, error) {
	portalmetadata := req.GetPortalMetadata()
	portalmetadataID := int(portalmetadata.GetId())
	m := svc.client.PortalMetadata.UpdateOneID(portalmetadataID)
	portalmetadataKind := toEntPortalMetadata_Kind(portalmetadata.GetKind())
	m.SetKind(portalmetadataKind)
	portalmetadataLifecycle := toEntPortalMetadata_Lifecycle(portalmetadata.GetLifecycle())
	m.SetLifecycle(portalmetadataLifecycle)
	if portalmetadata.GetPortal() != nil {
		portalmetadataPortal := int(portalmetadata.GetPortal().GetId())
		m.SetPortalID(portalmetadataPortal)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoPortalMetadata(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements PortalMetadataServiceServer.Delete
func (svc *PortalMetadataService) Delete(ctx context.Context, req *DeletePortalMetadataRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.PortalMetadata.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements PortalMetadataServiceServer.List
func (svc *PortalMetadataService) List(ctx context.Context, req *ListPortalMetadataRequest) (*ListPortalMetadataResponse, error) {
	var (
		err      error
		entList  []*ent.PortalMetadata
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.PortalMetadata.Query().
		Order(ent.Desc(portalmetadata.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(portalmetadata.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListPortalMetadataRequest_VIEW_UNSPECIFIED, ListPortalMetadataRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListPortalMetadataRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithPortal(func(query *ent.PortalQuery) {
				query.Select(portal.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoPortalMetadataList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListPortalMetadataResponse{
			PortalMetadataList: protoList,
			NextPageToken:      nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements PortalMetadataServiceServer.BatchCreate
func (svc *PortalMetadataService) BatchCreate(ctx context.Context, req *BatchCreatePortalMetadataSliceRequest) (*BatchCreatePortalMetadataSliceResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.PortalMetadataCreate, len(requests))
	for i, req := range requests {
		portalmetadata := req.GetPortalMetadata()
		var err error
		bulk[i], err = svc.createBuilder(portalmetadata)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.PortalMetadata.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoPortalMetadataList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreatePortalMetadataSliceResponse{
			PortalMetadataSlice: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *PortalMetadataService) createBuilder(portalmetadata *PortalMetadata) (*ent.PortalMetadataCreate, error) {
	m := svc.client.PortalMetadata.Create()
	portalmetadataCreatedAt := runtime.ExtractTime(portalmetadata.GetCreatedAt())
	m.SetCreatedAt(portalmetadataCreatedAt)
	portalmetadataKind := toEntPortalMetadata_Kind(portalmetadata.GetKind())
	m.SetKind(portalmetadataKind)
	portalmetadataLifecycle := toEntPortalMetadata_Lifecycle(portalmetadata.GetLifecycle())
	m.SetLifecycle(portalmetadataLifecycle)
	portalmetadataUpdatedAt := runtime.ExtractTime(portalmetadata.GetUpdatedAt())
	m.SetUpdatedAt(portalmetadataUpdatedAt)
	if portalmetadata.GetPortal() != nil {
		portalmetadataPortal := int(portalmetadata.GetPortal().GetId())
		m.SetPortalID(portalmetadataPortal)
	}
	return m, nil
}
