package core

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

func InitTrans() ut.Translator {
	trans, _ := ut.New(zh.New()).GetTranslator("zh")
	validate := binding.Validator.Engine().(*validator.Validate)
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		label := fld.Tag.Get("label")
		if label == "" {
			return fld.Name
		}
		return label
	})
	zhTranslations.RegisterDefaultTranslations(validate, trans)
	//zhTranslations.RegisterDefaultTranslations(global.GnValidate, trans)

	return trans
}