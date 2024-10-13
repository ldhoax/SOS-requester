package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ldhoax/SOS-requester/pkg/erru"
	"github.com/ldhoax/SOS-requester/pkg/sentry"
	"github.com/ldhoax/SOS-requester/utils"
	"github.com/gin-gonic/gin"
)

/*
Don’t have to repeat yourself every time you respond to user, instead you can use some helper functions.
*/
func (s Helper) Respond(c *gin.Context, msg string, data interface{}, status int) {
	var respData interface{}
	switch v := data.(type) {
	case nil:
	case erru.ErrArgument:
		status = http.StatusBadRequest
		respData = ErrorResponse{ErrorMessage: v.Unwrap().Error()}
		fmt.Println(v.Unwrap())
	case error:
		if http.StatusText(status) == "" {
			status = http.StatusInternalServerError
		} else {
			respData = ErrorResponse{ErrorMessage: v.Error()}
		}
		sentry.Log(v)
	default:
		respData = data
	}
	if status >= 300 {
		c.JSON(status, utils.GetRespError(msg, respData))
	} else {
		c.JSON(status, utils.GetRespSuccess(msg, respData))
	}

}

// it does not read to the memory, instead it will read it to the given 'v' interface.
func (s Helper) Decode(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}
