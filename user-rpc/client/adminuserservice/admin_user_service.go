// Code generated by goctl. DO NOT EDIT.
// Source: rpc.proto

package adminuserservice

import (
	"context"
	"github.com/doubunv/rpc-client/user-rpc/userRpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AdminUserListItem            = userRpc.AdminUserListItem
	AdminUserListReq             = userRpc.AdminUserListReq
	AdminUserListResp            = userRpc.AdminUserListResp
	AdminUserLoginReq            = userRpc.AdminUserLoginReq
	AdminUserLoginResp           = userRpc.AdminUserLoginResp
	CheckUserGoogleCodeByIdReq   = userRpc.CheckUserGoogleCodeByIdReq
	CheckUserGoogleCodeByIdResp  = userRpc.CheckUserGoogleCodeByIdResp
	CreateAdminUserReq           = userRpc.CreateAdminUserReq
	CreateAdminUserResp          = userRpc.CreateAdminUserResp
	DeleteAdminUserReq           = userRpc.DeleteAdminUserReq
	DeleteAdminUserResp          = userRpc.DeleteAdminUserResp
	GetAdminDetailReq            = userRpc.GetAdminDetailReq
	GetAdminDetailResp           = userRpc.GetAdminDetailResp
	GetUserDetailByIdReq         = userRpc.GetUserDetailByIdReq
	GetUserDetailByIdResp        = userRpc.GetUserDetailByIdResp
	GetUserDetailByQueryReq      = userRpc.GetUserDetailByQueryReq
	GetUserDetailByQueryResp     = userRpc.GetUserDetailByQueryResp
	GetUserListByIdsItem         = userRpc.GetUserListByIdsItem
	GetUserListByIdsReq          = userRpc.GetUserListByIdsReq
	GetUserListByIdsResp         = userRpc.GetUserListByIdsResp
	GetUserListByPageItem        = userRpc.GetUserListByPageItem
	GetUserListByPageReq         = userRpc.GetUserListByPageReq
	GetUserListByPageResp        = userRpc.GetUserListByPageResp
	LoginByAccountReq            = userRpc.LoginByAccountReq
	LoginByAccountResp           = userRpc.LoginByAccountResp
	LoginByEmailReq              = userRpc.LoginByEmailReq
	LoginByEmailResp             = userRpc.LoginByEmailResp
	LogoutReq                    = userRpc.LogoutReq
	LogoutResp                   = userRpc.LogoutResp
	ParseTokenReq                = userRpc.ParseTokenReq
	ParseTokenResp               = userRpc.ParseTokenResp
	RegisterByAccountReq         = userRpc.RegisterByAccountReq
	RegisterByAccountResp        = userRpc.RegisterByAccountResp
	RegisterByEmailReq           = userRpc.RegisterByEmailReq
	RegisterByEmailResp          = userRpc.RegisterByEmailResp
	UpdateAdminDetailReq         = userRpc.UpdateAdminDetailReq
	UpdateAdminDetailResp        = userRpc.UpdateAdminDetailResp
	UpdateUserDetailByIdReq      = userRpc.UpdateUserDetailByIdReq
	UpdateUserDetailByIdResp     = userRpc.UpdateUserDetailByIdResp
	UpdateUserGoogleCodeByIdReq  = userRpc.UpdateUserGoogleCodeByIdReq
	UpdateUserGoogleCodeByIdResp = userRpc.UpdateUserGoogleCodeByIdResp

	AdminUserService interface {
		// 管理员板块
		AdminUserLogin(ctx context.Context, in *AdminUserLoginReq, opts ...grpc.CallOption) (*AdminUserLoginResp, error)
		CreateAdminUser(ctx context.Context, in *CreateAdminUserReq, opts ...grpc.CallOption) (*CreateAdminUserResp, error)
		DeleteAdminUser(ctx context.Context, in *DeleteAdminUserReq, opts ...grpc.CallOption) (*DeleteAdminUserResp, error)
		AdminUserList(ctx context.Context, in *AdminUserListReq, opts ...grpc.CallOption) (*AdminUserListResp, error)
		GetAdminDetail(ctx context.Context, in *GetAdminDetailReq, opts ...grpc.CallOption) (*GetAdminDetailResp, error)
		UpdateAdminDetail(ctx context.Context, in *UpdateAdminDetailReq, opts ...grpc.CallOption) (*UpdateAdminDetailResp, error)
	}

	defaultAdminUserService struct {
		cli zrpc.Client
	}
)

func NewAdminUserService(cli zrpc.Client) AdminUserService {
	return &defaultAdminUserService{
		cli: cli,
	}
}

// 管理员板块
func (m *defaultAdminUserService) AdminUserLogin(ctx context.Context, in *AdminUserLoginReq, opts ...grpc.CallOption) (*AdminUserLoginResp, error) {
	client := userRpc.NewAdminUserServiceClient(m.cli.Conn())
	return client.AdminUserLogin(ctx, in, opts...)
}

func (m *defaultAdminUserService) CreateAdminUser(ctx context.Context, in *CreateAdminUserReq, opts ...grpc.CallOption) (*CreateAdminUserResp, error) {
	client := userRpc.NewAdminUserServiceClient(m.cli.Conn())
	return client.CreateAdminUser(ctx, in, opts...)
}

func (m *defaultAdminUserService) DeleteAdminUser(ctx context.Context, in *DeleteAdminUserReq, opts ...grpc.CallOption) (*DeleteAdminUserResp, error) {
	client := userRpc.NewAdminUserServiceClient(m.cli.Conn())
	return client.DeleteAdminUser(ctx, in, opts...)
}

func (m *defaultAdminUserService) AdminUserList(ctx context.Context, in *AdminUserListReq, opts ...grpc.CallOption) (*AdminUserListResp, error) {
	client := userRpc.NewAdminUserServiceClient(m.cli.Conn())
	return client.AdminUserList(ctx, in, opts...)
}

func (m *defaultAdminUserService) GetAdminDetail(ctx context.Context, in *GetAdminDetailReq, opts ...grpc.CallOption) (*GetAdminDetailResp, error) {
	client := userRpc.NewAdminUserServiceClient(m.cli.Conn())
	return client.GetAdminDetail(ctx, in, opts...)
}

func (m *defaultAdminUserService) UpdateAdminDetail(ctx context.Context, in *UpdateAdminDetailReq, opts ...grpc.CallOption) (*UpdateAdminDetailResp, error) {
	client := userRpc.NewAdminUserServiceClient(m.cli.Conn())
	return client.UpdateAdminDetail(ctx, in, opts...)
}
