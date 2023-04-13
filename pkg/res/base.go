package res

import (
	"demo04/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrResponse struct {
	Code int64  `json:"status_code"`
	Msg  string `json:"status_msg"`
}

func SendErrorResponse(c *gin.Context, err error) {
	errno := errno.ConvertErr(err)
	c.JSON(http.StatusOK, ErrResponse{
		Code: errno.ErrorCode,
		Msg:  errno.ErrorMsg,
	})
}

func SendCommonResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func SendAutoResponse(c *gin.Context, err error, data interface{}) {
	if err != nil {
		SendErrorResponse(c, err)
		return
	}
	SendCommonResponse(c, data)

}