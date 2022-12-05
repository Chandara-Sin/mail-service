package mail

import (
	"mail-service/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

type sendEmailFunc func(Mail, *gin.Context) error

func (fn sendEmailFunc) SendEmail(ml Mail, c *gin.Context) error {
	return fn(ml, c)
}

func SendEmailHandler(svc sendEmailFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ml := Mail{}
		log := logger.Unwrap(c)

		if err := c.ShouldBindJSON(&ml); err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		err := svc.SendEmail(ml, c)
		if err != nil {
			log.Error(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusCreated, ml)
	}
}
