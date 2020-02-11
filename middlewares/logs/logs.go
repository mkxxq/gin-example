package logs

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"ginexample/utils/log"
)

const (
	requestDataMaxByteNumber  = 10 * 1024
	responseDataMaxByteNumber = 10 * 1024
)

type logInfo struct {
	StatusCode   int       `json:"statusCode"`
	ClientIp     string    `json:"clientIp"`
	Method       string    `json:"method"`
	Path         string    `json:"path"`
	RequestData  string    `json:"requestData"`
	ResponseData string    `json:"responseData"`
	RequestTime  time.Time `json:"requestTime"`
	Email        string    `json:"email"`
	Latency      string    `json:"latency"`
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	if n, err := w.body.Write(b); err != nil {
		log.Println("custom response write error:", err)
		return n, err
	}
	return w.ResponseWriter.Write(b)
}
func CustomeLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logMiddlewareStart := time.Now()
		c.Set("reqStart", logMiddlewareStart)
		logMessage := &logInfo{}
		cc := c.Copy()
		logMessage.Method = cc.Request.Method
		logMessage.ClientIp = cc.ClientIP()
		logMessage.Path = cc.Request.URL.RequestURI()
		if logMessage.Method == http.MethodPost || logMessage.Method == http.MethodPut {
			if cc.Request.ContentLength < requestDataMaxByteNumber {
				bodyBytes, _ := ioutil.ReadAll(cc.Request.Body)
				// After r.Body data is read, there is no data in the body. You need to write the read data to r.Body again.
				cc.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
				logMessage.RequestData = fmt.Sprint(string(bodyBytes))
			}
		}
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		logMessage.StatusCode = c.Writer.Status()

		if len(blw.body.Bytes()) < responseDataMaxByteNumber {
			logMessage.ResponseData = blw.body.String()
		}
		duration := c.GetDuration("reqStart")
		logMessage.Latency = fmt.Sprintf("%v", duration)
		logMessage.RequestTime = logMiddlewareStart
		if logMessage.StatusCode%100 == 5 {
			log.Errorf("%d | %s | %s | %s | %s | %s | %s", logMessage.StatusCode, logMessage.Latency, logMessage.ClientIp, logMessage.Method, logMessage.Path, logMessage.RequestData, logMessage.ResponseData)
		} else {
			log.Infof("%d | %s | %s | %s | %s | %s | %s", logMessage.StatusCode, logMessage.Latency, logMessage.ClientIp, logMessage.Method, logMessage.Path, logMessage.RequestData, logMessage.ResponseData)
		}
	}
}
