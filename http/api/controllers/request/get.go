package request

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/GoldenOwlAsia/go-golang-api/internal/i18n"
	"strconv"
)

func (r RequestController) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		r.helper.Respond(c, i18n.Translate("en", "request.get.fail"), err, 0)
		return
	}

	request, err := r.requestService.Get(uint(id))
	if err != nil {
		r.helper.Respond(c, i18n.Translate("en", "request.get.fail"), err, 0)
		return
	}

	r.helper.Respond(c, i18n.Translate("en", "request.get.success"), request, http.StatusOK)
}
