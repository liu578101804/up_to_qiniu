package common

import (
	"github.com/liu578101804/up_to_qiniu/model"
	"github.com/plimble/ace"
	"net/http"
)

func SendErrJSON(msg string, args ...interface{})  {

	if len(args) == 0 {
		panic("缺少 *ace.C")
	}

	var c *ace.C
	var errNo = model.ErrorCode.ERROR

	if len(args) == 1 {
		theCtx, ok := args[0].(*ace.C)
		if !ok {
			panic("缺少 *ace.C")
		}
		c = theCtx
	} else if len(args) == 2 {
		theErrNo, ok := args[0].(int)
		if !ok {
			panic("errNo不正确")
		}
		errNo = theErrNo
		theCtx, ok := args[1].(*ace.C)
		if !ok {
			panic("缺少 *ace.C")
		}
		c = theCtx
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"code": 	errNo,
		"msg": 		msg,
		"data":		"",
	})

	//终止请求
	c.Abort()
}