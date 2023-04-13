package errno

const (

	// Success
	SuccessCode = 0
	SuccessMsg  = "success"

	// Error
	UnexpectedTypeErrorCode    = 10001 // 未知错误
	ParamErrorCode             = 10002 // 参数错误
	ErrorAuthCheckTokenFail    = 10003 // 鉴权失败
	ErrorAuthCheckTokenTimeout = 10004 // 鉴权超时
	ErrorNotFriend             = 10005 // 不是好友
	ErrorAlreadyFriend         = 10006 // 已是好友

	// WebSocket
	WebsocketSuccessMessage = 50001
	WebsocketSuccess        = 50002
	WebsocketEnd            = 50003
	WebsocketOnlineReply    = 50004
	WebsocketOfflineReply   = 50005
	WebsocketLimit          = 50006
)
