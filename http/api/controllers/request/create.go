package request

import (
	"net/http"

	requestService "github.com/GoldenOwlAsia/go-golang-api/internal/request/service"
	"github.com/GoldenOwlAsia/go-golang-api/pkg/erru"
	"github.com/gin-gonic/gin"
	"github.com/GoldenOwlAsia/go-golang-api/internal/i18n"
)

func (r RequestController) Create(c *gin.Context) {
	type request struct {
		PhoneNumber    string  `form:"phone_number"`
		Email          string  `form:"email"`
		EmergencyLevel int     `form:"emergency_level"`
		Location       string  `form:"location"`
		Latitude       float64 `form:"latitude"`
		Longitude      float64 `form:"longitude"`
		ShortDescription string  `form:"short_description"`
		Description    string  `form:"description"`
	}

	var requestInput request
	if err := c.ShouldBind(&requestInput); err != nil {
		r.helper.Respond(c, i18n.Translate("en", "request.create.fail"), erru.ErrArgument{Wrapped: err}, 0)
		return
	}

	req, err := r.requestService.Create(c, requestService.CreateParams{
		PhoneNumber:    requestInput.PhoneNumber,
		Email:          requestInput.Email,
		EmergencyLevel: requestInput.EmergencyLevel,
		Location:       requestInput.Location,
		Latitude:       requestInput.Latitude,
		Longitude:      requestInput.Longitude,
		ShortDescription: requestInput.ShortDescription,
		Description:      requestInput.Description,
	})

	if err != nil {
		r.helper.Respond(c, i18n.Translate("en", "request.create.fail"), err, 0)
		return
	}
	r.helper.Respond(c, i18n.Translate("en", "request.create.success"), req, http.StatusOK)
}
