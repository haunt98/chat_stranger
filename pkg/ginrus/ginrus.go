package ginrus

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		c.Next()

		end := time.Now()

		if len(c.Errors) != 0 {
			logrus.WithFields(logrus.Fields{
				"module": "gin",
			}).Errorf("error=%s", c.Errors.String())
		}
		logrus.WithFields(logrus.Fields{
			"module": "gin",
		}).Infof("latency=%s method=%s path=%s status=%d",
			end.Sub(start), c.Request.Method, path, c.Writer.Status())
	}
}
