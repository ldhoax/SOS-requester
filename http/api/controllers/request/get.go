package request

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ldhoax/SOS-requester/internal/i18n"
)

func (r RequestController) Get(c *gin.Context) {
	id := c.Param("id")

	request, err := r.requestService.Get(id)
	if err != nil {
		r.helper.Respond(c, i18n.Translate("en", "request.get.fail"), err, 0)
		return
	}

	r.helper.Respond(c, i18n.Translate("en", "request.get.success"), request, http.StatusOK)
}
