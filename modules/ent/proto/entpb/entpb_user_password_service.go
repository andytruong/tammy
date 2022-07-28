// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strconv "strconv"
	ent "tammy/ent"
	user "tammy/ent/user"
	userpassword "tammy/ent/userpassword"
)

// UserPasswordService implements UserPasswordServiceServer
type UserPasswordService struct {
	client *ent.Client
	UnimplementedUserPasswordServiceServer
}

// NewUserPasswordService returns a new UserPasswordService
func NewUserPasswordService(client *ent.Client) *UserPasswordService {
	return &UserPasswordService{
		client: client,
	}
}

// toProtoUserPassword transforms the ent type to the pb type
func toProtoUserPassword(e *ent.UserPassword) (*UserPassword, error) {
	v := &UserPassword{}
	createdAt := e.CreatedAt
	v.CreatedAt = createdAt
	hashedValue := e.HashedValue
	v.HashedValue = hashedValue
	id := int64(e.ID)
	v.Id = id
	isActive := e.IsActive
	v.IsActive = isActive
	if edg := e.Edges.User; edg != nil {
		id := int64(edg.ID)
		v.User = &User{
			Id: id,
		}
	}
	return v, nil
}

// toProtoUserPasswordList transforms a list of ent type to a list of pb type
func toProtoUserPasswordList(e []*ent.UserPassword) ([]*UserPassword, error) {
	var pbList []*UserPassword
	for _, entEntity := range e {
		pbEntity, err := toProtoUserPassword(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements UserPasswordServiceServer.Create
func (svc *UserPasswordService) Create(ctx context.Context, req *CreateUserPasswordRequest) (*UserPassword, error) {
	userpassword := req.GetUserPassword()
	m, err := svc.createBuilder(userpassword)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoUserPassword(res)
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

// Get implements UserPasswordServiceServer.Get
func (svc *UserPasswordService) Get(ctx context.Context, req *GetUserPasswordRequest) (*UserPassword, error) {
	var (
		err error
		get *ent.UserPassword
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetUserPasswordRequest_VIEW_UNSPECIFIED, GetUserPasswordRequest_BASIC:
		get, err = svc.client.UserPassword.Get(ctx, id)
	case GetUserPasswordRequest_WITH_EDGE_IDS:
		get, err = svc.client.UserPassword.Query().
			Where(userpassword.ID(id)).
			WithUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoUserPassword(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements UserPasswordServiceServer.Update
func (svc *UserPasswordService) Update(ctx context.Context, req *UpdateUserPasswordRequest) (*UserPassword, error) {
	userpassword := req.GetUserPassword()
	userpasswordID := int(userpassword.GetId())
	m := svc.client.UserPassword.UpdateOneID(userpasswordID)
	userpasswordCreatedAt := uint32(userpassword.GetCreatedAt())
	m.SetCreatedAt(userpasswordCreatedAt)
	userpasswordHashedValue := userpassword.GetHashedValue()
	m.SetHashedValue(userpasswordHashedValue)
	userpasswordIsActive := userpassword.GetIsActive()
	m.SetIsActive(userpasswordIsActive)
	if userpassword.GetUser() != nil {
		userpasswordUser := int(userpassword.GetUser().GetId())
		m.SetUserID(userpasswordUser)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoUserPassword(res)
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

// Delete implements UserPasswordServiceServer.Delete
func (svc *UserPasswordService) Delete(ctx context.Context, req *DeleteUserPasswordRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.UserPassword.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements UserPasswordServiceServer.List
func (svc *UserPasswordService) List(ctx context.Context, req *ListUserPasswordRequest) (*ListUserPasswordResponse, error) {
	var (
		err      error
		entList  []*ent.UserPassword
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.UserPassword.Query().
		Order(ent.Desc(userpassword.FieldID)).
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
			Where(userpassword.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListUserPasswordRequest_VIEW_UNSPECIFIED, ListUserPasswordRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListUserPasswordRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
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
		protoList, err := toProtoUserPasswordList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListUserPasswordResponse{
			UserPasswordList: protoList,
			NextPageToken:    nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements UserPasswordServiceServer.BatchCreate
func (svc *UserPasswordService) BatchCreate(ctx context.Context, req *BatchCreateUserPasswordsRequest) (*BatchCreateUserPasswordsResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.UserPasswordCreate, len(requests))
	for i, req := range requests {
		userpassword := req.GetUserPassword()
		var err error
		bulk[i], err = svc.createBuilder(userpassword)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.UserPassword.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoUserPasswordList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateUserPasswordsResponse{
			UserPasswords: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *UserPasswordService) createBuilder(userpassword *UserPassword) (*ent.UserPasswordCreate, error) {
	m := svc.client.UserPassword.Create()
	userpasswordCreatedAt := uint32(userpassword.GetCreatedAt())
	m.SetCreatedAt(userpasswordCreatedAt)
	userpasswordHashedValue := userpassword.GetHashedValue()
	m.SetHashedValue(userpasswordHashedValue)
	userpasswordIsActive := userpassword.GetIsActive()
	m.SetIsActive(userpasswordIsActive)
	if userpassword.GetUser() != nil {
		userpasswordUser := int(userpassword.GetUser().GetId())
		m.SetUserID(userpasswordUser)
	}
	return m, nil
}