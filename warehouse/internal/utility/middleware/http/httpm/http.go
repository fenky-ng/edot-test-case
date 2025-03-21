package httpm

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body   *bytes.Buffer
	header http.Header
}

func (w *bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

func (w *bodyLogWriter) WriteHeader(statusCode int) {
	w.header = w.ResponseWriter.Header().Clone()
	w.ResponseWriter.WriteHeader(statusCode)
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader("X-Request-ID") // check if client provided one
		if requestId == "" {
			requestId = uuid.New().String() // generate a new UUID
		}

		// store in context for use in handlers
		c.Set("RequestID", requestId)

		// add request ID to response headers
		c.Writer.Header().Set("X-Request-ID", requestId)

		// get request header
		headerBytes, _ := json.Marshal(c.Request.Header)

		// get request body
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
		}
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer, header: make(http.Header)}
		c.Writer = blw

		// log request
		if len(bodyBytes) == 0 {
			log.Printf("[REQUEST] id: %v | method: %v | uri: %v | header: %s",
				requestId, c.Request.Method, c.Request.RequestURI, string(headerBytes))
		} else {
			buffer := new(bytes.Buffer)
			json.Compact(buffer, bodyBytes)
			log.Printf("[REQUEST] id: %v | method: %v | uri: %v | header: %s | body: %s",
				requestId, c.Request.Method, c.Request.RequestURI, string(headerBytes), buffer.String())
		}

		// process request
		c.Next()

		// log response
		log.Printf("[RESPONSE] id: %v | method: %v | uri: %v | code: %v | header: %v | body: %s",
			requestId, c.Request.Method, c.Request.RequestURI, c.Writer.Status(), blw.header, blw.body.String())
	}
}
