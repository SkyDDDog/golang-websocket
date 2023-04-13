package errno

var (
	// Success
	Success = NewErrNo(SuccessCode, SuccessMsg)

	UnexpectedTypeError             = NewErrNo(UnexpectedTypeErrorCode, "Unknown Error")
	ServiceInternalError            = NewErrNo(UnexpectedTypeErrorCode, "Service Internal Error")
	ParamError                      = NewErrNo(ParamErrorCode, "Parameter Error")
	ErrorAuthCheckTokenError        = NewErrNo(ErrorAuthCheckTokenFail, "AuthCheckToken Failed")
	ErrorAuthCheckTokenTimeoutError = NewErrNo(ErrorAuthCheckTokenTimeout, "AuthCheckToken Timeout")
	ErrorNotFriendError             = NewErrNo(ErrorNotFriend, "You're Not Friends")
	ErrorAlreadyFriendError         = NewErrNo(ErrorAlreadyFriend, "You're Already Friends")
)
