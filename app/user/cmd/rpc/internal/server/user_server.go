// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"qianshi/app/user/cmd/rpc/internal/logic"
	"qianshi/app/user/cmd/rpc/internal/svc"
	"qianshi/app/user/cmd/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	__.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) EmailLogin(ctx context.Context, in *__.EmailLoginReq) (*__.LoginResp, error) {
	l := logic.NewEmailLoginLogic(ctx, s.svcCtx)
	return l.EmailLogin(in)
}

func (s *UserServer) PassLogin(ctx context.Context, in *__.PassLoginReq) (*__.LoginResp, error) {
	l := logic.NewPassLoginLogic(ctx, s.svcCtx)
	return l.PassLogin(in)
}

func (s *UserServer) UserQuery(ctx context.Context, in *__.QueryReq) (*__.UserQueryResp, error) {
	l := logic.NewUserQueryLogic(ctx, s.svcCtx)
	return l.UserQuery(in)
}

func (s *UserServer) UserHomeQuery(ctx context.Context, in *__.QueryReq) (*__.UserHomeQueryResp, error) {
	l := logic.NewUserHomeQueryLogic(ctx, s.svcCtx)
	return l.UserHomeQuery(in)
}

func (s *UserServer) UserInteractionQuery(ctx context.Context, in *__.QueryReq) (*__.UserInteractionQueryResp, error) {
	l := logic.NewUserInteractionQueryLogic(ctx, s.svcCtx)
	return l.UserInteractionQuery(in)
}

func (s *UserServer) PassChangeVerify(ctx context.Context, in *__.PassChangeVerifyReq) (*__.PassChangeVerifyResp, error) {
	l := logic.NewPassChangeVerifyLogic(ctx, s.svcCtx)
	return l.PassChangeVerify(in)
}

func (s *UserServer) PassChange(ctx context.Context, in *__.PassChangeReq) (*__.PassChangeResp, error) {
	l := logic.NewPassChangeLogic(ctx, s.svcCtx)
	return l.PassChange(in)
}