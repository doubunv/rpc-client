// Code generated by goctl. DO NOT EDIT.
// Source: rpc.proto

package userloginservice

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

	UserLoginService interface {
		ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error)
		LoginByAccount(ctx context.Context, in *LoginByAccountReq, opts ...grpc.CallOption) (*LoginByAccountResp, error)
		LoginByEmail(ctx context.Context, in *LoginByEmailReq, opts ...grpc.CallOption) (*LoginByEmailResp, error)
		Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error)
		RegisterByAccount(ctx context.Context, in *RegisterByAccountReq, opts ...grpc.CallOption) (*RegisterByAccountResp, error)
		RegisterByEmail(ctx context.Context, in *RegisterByEmailReq, opts ...grpc.CallOption) (*RegisterByEmailResp, error)
	}

	defaultUserLoginService struct {
		cli zrpc.Client
	}
)

func NewUserLoginService(cli zrpc.Client) UserLoginService {
	return &defaultUserLoginService{
		cli: cli,
	}
}

func (m *defaultUserLoginService) ParseToken(ctx context.Context, in *ParseTokenReq, opts ...grpc.CallOption) (*ParseTokenResp, error) {
	client := userRpc.NewUserLoginServiceClient(m.cli.Conn())
	return client.ParseToken(ctx, in, opts...)
}

func (m *defaultUserLoginService) LoginByAccount(ctx context.Context, in *LoginByAccountReq, opts ...grpc.CallOption) (*LoginByAccountResp, error) {
	client := userRpc.NewUserLoginServiceClient(m.cli.Conn())
	return client.LoginByAccount(ctx, in, opts...)
}

func (m *defaultUserLoginService) LoginByEmail(ctx context.Context, in *LoginByEmailReq, opts ...grpc.CallOption) (*LoginByEmailResp, error) {
	client := userRpc.NewUserLoginServiceClient(m.cli.Conn())
	return client.LoginByEmail(ctx, in, opts...)
}

func (m *defaultUserLoginService) Logout(ctx context.Context, in *LogoutReq, opts ...grpc.CallOption) (*LogoutResp, error) {
	client := userRpc.NewUserLoginServiceClient(m.cli.Conn())
	return client.Logout(ctx, in, opts...)
}

func (m *defaultUserLoginService) RegisterByAccount(ctx context.Context, in *RegisterByAccountReq, opts ...grpc.CallOption) (*RegisterByAccountResp, error) {
	client := userRpc.NewUserLoginServiceClient(m.cli.Conn())
	return client.RegisterByAccount(ctx, in, opts...)
}

func (m *defaultUserLoginService) RegisterByEmail(ctx context.Context, in *RegisterByEmailReq, opts ...grpc.CallOption) (*RegisterByEmailResp, error) {
	client := userRpc.NewUserLoginServiceClient(m.cli.Conn())
	return client.RegisterByEmail(ctx, in, opts...)
}
