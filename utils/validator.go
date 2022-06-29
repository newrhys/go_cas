package utils

import (
	"github.com/go-playground/validator/v10"
	"strings"
	"wave-admin/global"
)

func Error(err error) string {
	var ret []string
	if validationErrors,ok := err.(validator.ValidationErrors);!ok {
		return err.Error()
	} else {
		for _, e := range validationErrors {
			//log.Println(e.Translate(global.GnTrans))
			ret = append(ret, e.Translate(global.GnTrans))
		}
	}
	//fmt.Println(ret)
	return strings.Join(ret, ";")
}
