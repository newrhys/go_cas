package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wave-admin/global"
	"wave-admin/model/system"
	"wave-admin/service"
	"wave-admin/utils"
)

var apiService = service.ServiceGroupApp.SystemServiceGroup.ApiService
var recordService = service.ServiceGroupApp.SystemServiceGroup.RecordService

func Record() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的PATH
		path, id := utils.GetRequestPath(c.Request.URL.Path)
		var body []byte
		var userId uint64
		if c.Request.Method != http.MethodGet {
			if c.Request.Method == http.MethodDelete {
				m := make(map[string]int)
				m["id"] = id
				body, _ = json.Marshal(&m)
			} else {
				var err error
				body, err = ioutil.ReadAll(c.Request.Body)
				if err != nil {
					global.GnLog.Error("read body from request error:", zap.Error(err))
				} else {
					c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
				}
			}
		} else {
			m := make(map[string]string)
			query := c.Request.URL.RawQuery
			query, _ = url.QueryUnescape(query)
			split := strings.Split(query, "&")
			for _, v := range split {
				kv := strings.Split(v, "=")
				if len(kv) == 2 {
					m[kv[0]] = kv[1]
				}
			}
			body, _ = json.Marshal(&m)
		}
		userId = utils.GetUserID(c)

		api, err := apiService.GetApiByPath(path, c.Request.Method)
		if err != nil {
			global.GnLog.Error("read path from request error:", zap.Error(err))
		}

		record := system.SysRecord{
			Description:  api.Description,
			Method:       c.Request.Method,
			Path:         path,
			Ip:           c.ClientIP(),
			Body:         string(body),
			UserID:       userId,
		}

		writer := responseBodyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		latency := time.Since(now)
		record.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		record.Status = c.Writer.Status()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := recordService.AddRecord(record); err != nil {
			global.GnLog.Error("add operation record error:", zap.Error(err))
		}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
