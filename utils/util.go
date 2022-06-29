package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
	"wave-admin/global"
	"wave-admin/model/system/request"

	"github.com/gin-gonic/gin"
)

func GetRequestPath(path string) (obj string, id int) {
	//log.Println("path=", path)
	paths := strings.Split(path, "/")
	//log.Println("paths=", paths)
	len := len(paths)
	//log.Println("len=", len)
	if len > 5 {
		obj = "/" + strings.Join(paths[3:(len-1)], "/") + "/:id"
		pid, _ := strconv.Atoi(paths[(len - 1)])
		id = pid
	} else {
		obj = "/" + strings.Join(paths[3:], "/")
	}

	return obj, id
}

// 获得随机变量
func RandomString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

// 获得图片URL
func TransformImageUrl(path string) (url string) {
	if global.GnConfig.System.OssType == "local" {
		url = "http://localhost:" + global.GnConfig.System.ServerPort + path
	}

	return url
}

func GetRoleId(ctx *gin.Context) uint64 {
	if claims, exists := ctx.Get("claims"); !exists {
		global.GnLog.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件")
		return 0
	} else {
		waitUse := claims.(*request.LoginClaims)
		return waitUse.RoleId
	}
}

// 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(ctx *gin.Context) uint64 {
	if claims, exists := ctx.Get("claims"); !exists {

		fmt.Printf("claims: %v\n", claims)
		global.GnLog.Error("从Gin的Context中获取从jwt解析出来的用户ID失败, 请检查路由是否使用jwt中间件")
		return 0
	} else {
		fmt.Printf("claims: %v\n", claims)

		waitUse := claims.(*request.LoginClaims)
		return waitUse.ID
	}
}

func ToMap(in interface{}, tagName string, ignore string) (map[string]interface{}, error) {
	myMap := make(map[string]interface{})
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("类型错误:%v", v)
	}
	t := v.Type()
	//遍历结构体字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		get := field.Tag.Get(tagName)
		if get != "" && get != ignore {
			myMap[get] = v.Field(i).Interface()
		}
	}
	return myMap, nil
}
