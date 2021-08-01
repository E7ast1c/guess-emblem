package emblem

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.RouterGroup, service Service) {
	res := resource{service: service}

	r.GET("/get-random", res.getRandom)
}

type resource struct {
	service Service
}

func (r resource) getRandom(c *gin.Context)  {
	 emblem, err := r.service.GetRandom(c)
	 if err != nil {
		 c.JSON(http.StatusInternalServerError, err)
		 return
	 }
	 c.JSON(http.StatusOK, emblem)
	return
}
