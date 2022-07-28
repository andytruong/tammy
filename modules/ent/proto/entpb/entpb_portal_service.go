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
	ent "tammy/ent"
	account "tammy/ent/account"
	portal "tammy/ent/portal"
	portallegal "tammy/ent/portallegal"
	portalmetadata "tammy/ent/portalmetadata"
)

// PortalService implements PortalServiceServer
type PortalService struct {
	client *ent.Client
	UnimplementedPortalServiceServer
}

// NewPortalService returns a new PortalService
func NewPortalService(client *ent.Client) *PortalService {
	return &PortalService{
		client: client,
	}
}

// toProtoPortal transforms the ent type to the pb type
func toProtoPortal(e *ent.Portal) (*Portal, error) {
	v := &Portal{}
	createdAt := timestamppb.New(e.CreatedAt)
	v.CreatedAt = createdAt
	id := int64(e.ID)
	v.Id = id
	isActive := e.IsActive
	v.IsActive = isActive
	slug := e.Slug
	v.Slug = slug
	updatedAt := timestamppb.New(e.UpdatedAt)
	v.UpdatedAt = updatedAt
	for _, edg := range e.Edges.Legal {
		id := int64(edg.ID)
		v.Legal = append(v.Legal, &PortalLegal{
			Id: id,
		})
	}
	for _, edg := range e.Edges.Members {
		id := int64(edg.ID)
		v.Members = append(v.Members, &Account{
			Id: id,
		})
	}
	for _, edg := range e.Edges.Metadata {
		id := int64(edg.ID)
		v.Metadata = append(v.Metadata, &PortalMetadata{
			Id: id,
		})
	}
	return v, nil
}

// toProtoPortalList transforms a list of ent type to a list of pb type
func toProtoPortalList(e []*ent.Portal) ([]*Portal, error) {
	var pbList []*Portal
	for _, entEntity := range e {
		pbEntity, err := toProtoPortal(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements PortalServiceServer.Create
func (svc *PortalService) Create(ctx context.Context, req *CreatePortalRequest) (*Portal, error) {
	portal := req.GetPortal()
	m, err := svc.createBuilder(portal)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoPortal(res)
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

// Get implements PortalServiceServer.Get
func (svc *PortalService) Get(ctx context.Context, req *GetPortalRequest) (*Portal, error) {
	var (
		err error
		get *ent.Portal
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetPortalRequest_VIEW_UNSPECIFIED, GetPortalRequest_BASIC:
		get, err = svc.client.Portal.Get(ctx, id)
	case GetPortalRequest_WITH_EDGE_IDS:
		get, err = svc.client.Portal.Query().
			Where(portal.ID(id)).
			WithLegal(func(query *ent.PortalLegalQuery) {
				query.Select(portallegal.FieldID)
			}).
			WithMembers(func(query *ent.AccountQuery) {
				query.Select(account.FieldID)
			}).
			WithMetadata(func(query *ent.PortalMetadataQuery) {
				query.Select(portalmetadata.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoPortal(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements PortalServiceServer.Update
func (svc *PortalService) Update(ctx context.Context, req *UpdatePortalRequest) (*Portal, error) {
	portal := req.GetPortal()
	portalID := int(portal.GetId())
	m := svc.client.Portal.UpdateOneID(portalID)
	portalIsActive := portal.GetIsActive()
	m.SetIsActive(portalIsActive)
	for _, item := range portal.GetLegal() {
		legal := int(item.GetId())
		m.AddLegalIDs(legal)
	}
	for _, item := range portal.GetMembers() {
		members := int(item.GetId())
		m.AddMemberIDs(members)
	}
	for _, item := range portal.GetMetadata() {
		metadata := int(item.GetId())
		m.AddMetadatumIDs(metadata)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoPortal(res)
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

// Delete implements PortalServiceServer.Delete
func (svc *PortalService) Delete(ctx context.Context, req *DeletePortalRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Portal.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements PortalServiceServer.List
func (svc *PortalService) List(ctx context.Context, req *ListPortalRequest) (*ListPortalResponse, error) {
	var (
		err      error
		entList  []*ent.Portal
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Portal.Query().
		Order(ent.Desc(portal.FieldID)).
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
			Where(portal.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListPortalRequest_VIEW_UNSPECIFIED, ListPortalRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListPortalRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithLegal(func(query *ent.PortalLegalQuery) {
				query.Select(portallegal.FieldID)
			}).
			WithMembers(func(query *ent.AccountQuery) {
				query.Select(account.FieldID)
			}).
			WithMetadata(func(query *ent.PortalMetadataQuery) {
				query.Select(portalmetadata.FieldID)
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
		protoList, err := toProtoPortalList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListPortalResponse{
			PortalList:    protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements PortalServiceServer.BatchCreate
func (svc *PortalService) BatchCreate(ctx context.Context, req *BatchCreatePortalsRequest) (*BatchCreatePortalsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.PortalCreate, len(requests))
	for i, req := range requests {
		portal := req.GetPortal()
		var err error
		bulk[i], err = svc.createBuilder(portal)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Portal.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoPortalList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreatePortalsResponse{
			Portals: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *PortalService) createBuilder(portal *Portal) (*ent.PortalCreate, error) {
	m := svc.client.Portal.Create()
	portalCreatedAt := runtime.ExtractTime(portal.GetCreatedAt())
	m.SetCreatedAt(portalCreatedAt)
	portalIsActive := portal.GetIsActive()
	m.SetIsActive(portalIsActive)
	portalSlug := portal.GetSlug()
	m.SetSlug(portalSlug)
	portalUpdatedAt := runtime.ExtractTime(portal.GetUpdatedAt())
	m.SetUpdatedAt(portalUpdatedAt)
	for _, item := range portal.GetLegal() {
		legal := int(item.GetId())
		m.AddLegalIDs(legal)
	}
	for _, item := range portal.GetMembers() {
		members := int(item.GetId())
		m.AddMemberIDs(members)
	}
	for _, item := range portal.GetMetadata() {
		metadata := int(item.GetId())
		m.AddMetadatumIDs(metadata)
	}
	return m, nil
}
