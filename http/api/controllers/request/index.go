package request

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/GoldenOwlAsia/go-golang-api/internal/i18n"
)

func (r RequestController) Index(c *gin.Context) {
	requests, err := r.requestService.GetAll()
	if err != nil {
		r.helper.Respond(c, i18n.Translate("en", "request.index.fail"), err, 0)
		return
	}

	r.helper.Respond(c, i18n.Translate("en", "request.index.success"), requests, http.StatusOK)
}
