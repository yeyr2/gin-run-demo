// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messagefunc

import (
	"context"
	api "douSheng/cmd/message/kitex_gen/api"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return messageFuncServiceInfo
}

var messageFuncServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "MessageFunc"
	handlerType := (*api.MessageFunc)(nil)
	methods := map[string]kitex.MethodInfo{
		"MessageAction": kitex.NewMethodInfo(messageActionHandler, newMessageFuncMessageActionArgs, newMessageFuncMessageActionResult, false),
		"MessageChat":   kitex.NewMethodInfo(messageChatHandler, newMessageFuncMessageChatArgs, newMessageFuncMessageChatResult, false),
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

func messageActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.MessageFuncMessageActionArgs)
	realResult := result.(*api.MessageFuncMessageActionResult)
	success, err := handler.(api.MessageFunc).MessageAction(ctx, realArg.Token, realArg.Content, realArg.ToUserId, realArg.ActionType)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageFuncMessageActionArgs() interface{} {
	return api.NewMessageFuncMessageActionArgs()
}

func newMessageFuncMessageActionResult() interface{} {
	return api.NewMessageFuncMessageActionResult()
}

func messageChatHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.MessageFuncMessageChatArgs)
	realResult := result.(*api.MessageFuncMessageChatResult)
	success, err := handler.(api.MessageFunc).MessageChat(ctx, realArg.Token, realArg.PreMsgTime, realArg.ToUserId)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newMessageFuncMessageChatArgs() interface{} {
	return api.NewMessageFuncMessageChatArgs()
}

func newMessageFuncMessageChatResult() interface{} {
	return api.NewMessageFuncMessageChatResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) MessageAction(ctx context.Context, token string, content string, toUserId int64, actionType int32) (r *api.MessageResponse, err error) {
	var _args api.MessageFuncMessageActionArgs
	_args.Token = token
	_args.Content = content
	_args.ToUserId = toUserId
	_args.ActionType = actionType
	var _result api.MessageFuncMessageActionResult
	if err = p.c.Call(ctx, "MessageAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) MessageChat(ctx context.Context, token string, preMsgTime int64, toUserId int64) (r *api.MessageResponse, err error) {
	var _args api.MessageFuncMessageChatArgs
	_args.Token = token
	_args.PreMsgTime = preMsgTime
	_args.ToUserId = toUserId
	var _result api.MessageFuncMessageChatResult
	if err = p.c.Call(ctx, "MessageChat", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
