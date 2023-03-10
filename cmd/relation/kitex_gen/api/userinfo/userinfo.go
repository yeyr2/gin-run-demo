// Code generated by Kitex v0.4.4. DO NOT EDIT.

package userinfo

import (
	"context"
	api "douSheng/cmd/relation/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return userInfoServiceInfo
}

var userInfoServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserInfo"
	handlerType := (*api.UserInfo)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register": kitex.NewMethodInfo(registerHandler, newUserInfoRegisterArgs, newUserInfoRegisterResult, false),
		"Login":    kitex.NewMethodInfo(loginHandler, newUserInfoLoginArgs, newUserInfoLoginResult, false),
		"UserInfo": kitex.NewMethodInfo(userInfoHandler, newUserInfoUserInfoArgs, newUserInfoUserInfoResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.UserInfoRegisterArgs)
	realResult := result.(*api.UserInfoRegisterResult)
	success, err := handler.(api.UserInfo).Register(ctx, realArg.Username, realArg.Password)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserInfoRegisterArgs() interface{} {
	return api.NewUserInfoRegisterArgs()
}

func newUserInfoRegisterResult() interface{} {
	return api.NewUserInfoRegisterResult()
}

func loginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.UserInfoLoginArgs)
	realResult := result.(*api.UserInfoLoginResult)
	success, err := handler.(api.UserInfo).Login(ctx, realArg.Username, realArg.Password)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserInfoLoginArgs() interface{} {
	return api.NewUserInfoLoginArgs()
}

func newUserInfoLoginResult() interface{} {
	return api.NewUserInfoLoginResult()
}

func userInfoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.UserInfoUserInfoArgs)
	realResult := result.(*api.UserInfoUserInfoResult)
	success, err := handler.(api.UserInfo).UserInfo(ctx, realArg.Token, realArg.UserId)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newUserInfoUserInfoArgs() interface{} {
	return api.NewUserInfoUserInfoArgs()
}

func newUserInfoUserInfoResult() interface{} {
	return api.NewUserInfoUserInfoResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, username string, password string) (r *api.UserResponse, err error) {
	var _args api.UserInfoRegisterArgs
	_args.Username = username
	_args.Password = password
	var _result api.UserInfoRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Login(ctx context.Context, username string, password string) (r *api.UserResponse, err error) {
	var _args api.UserInfoLoginArgs
	_args.Username = username
	_args.Password = password
	var _result api.UserInfoLoginResult
	if err = p.c.Call(ctx, "Login", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserInfo(ctx context.Context, token string, userId int64) (r *api.UserResponse, err error) {
	var _args api.UserInfoUserInfoArgs
	_args.Token = token
	_args.UserId = userId
	var _result api.UserInfoUserInfoResult
	if err = p.c.Call(ctx, "UserInfo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
