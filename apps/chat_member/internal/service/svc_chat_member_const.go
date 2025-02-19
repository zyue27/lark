package service

import "errors"

const (
	ERROR_CODE_CHAT_MEMBER_QUERY_DB_FAILED                       int32 = 800001
	ERROR_CODE_CHAT_MEMBER_GRPC_SERVICE_FAILURE                  int32 = 800002 // 服务故障
	ERROR_CODE_CHAT_MEMBER_CHCHE_MEMBER_FAILED                   int32 = 800003
	ERROR_CODE_CHAT_MEMBER_MISS_USER_INFO                        int32 = 800004
	ERROR_CODE_CHAT_MEMBER_UPDATE_VALUE_FAILED                   int32 = 800005
	ERROR_CODE_CHAT_MEMBER_CHCHE_MEMBER_BASIC_LIST_FAILED        int32 = 800006
	ERROR_CODE_CHAT_MEMBER_UPDATE_MEMBER_CONNECTED_SERVER_FAILED int32 = 800007
	ERROR_CODE_CHAT_MEMBER_REDIS_GET_FAILED                      int32 = 800008
	ERROR_CODE_CHAT_MEMBER_REDIS_SET_FAILED                      int32 = 800009
	ERROR_CODE_CHAT_MEMBER_NON_CHAT_MEMBER                       int32 = 800010
)

const (
	ERROR_CHAT_MEMBER_QUERY_DB_FAILED                       = "查询失败"
	ERROR_CHAT_MEMBER_GRPC_SERVICE_FAILURE                  = "服务故障"
	ERROR_CHAT_MEMBER_CHCHE_MEMBER_FAILED                   = "缓存成员信息失败"
	ERROR_CHAT_MEMBER_MISS_USER_INFO                        = "缺失用户信息"
	ERROR_CHAT_MEMBER_UPDATE_VALUE_FAILED                   = "更新Value失败"
	ERROR_CHAT_MEMBER_CHCHE_MEMBER_BASIC_LIST_FAILED        = "缓存Chat成员基础信息列表JSON失败"
	ERROR_CHAT_MEMBER_UPDATE_MEMBER_CONNECTED_SERVER_FAILED = "更新 ServerID 失败"
	ERROR_CHAT_MEMBER_REDIS_GET_FAILED                      = "读取redis缓存失败"
	ERROR_CHAT_MEMBER_REDIS_SET_FAILED                      = "设置redis缓存失败"
	ERROR_CHAT_MEMBER_NON_CHAT_MEMBER                       = "非Chat成员"
)

var (
	ERR_CHAT_MEMBER_CHCHE_MEMBER_FAILED = errors.New("缓存成员信息失败")
)
