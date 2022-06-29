package request

import "wave-admin/model/common/request"

type RecordList struct {
	Method       	string        `json:"method"`			// 请求方法
	Path         	string        `json:"path"`				// 请求路径
	Status       	int           `json:"status"`			// 请求状态
	request.PageInfo
}
