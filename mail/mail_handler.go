package mail

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type sendEmailFunc func(Mail) error

func (fn sendEmailFunc) SendEmail(ml Mail) error {
	return fn(ml)
}

func SendEmailHandler(svc sendEmailFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ml := Mail{}

		if err := c.ShouldBindJSON(&ml); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := svc.SendEmail(ml)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, ml)
	}
}
