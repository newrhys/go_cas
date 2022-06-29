package system

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
	"wave-admin/global"
	"wave-admin/model/common/response"
	"wave-admin/utils"
)

type AttachApi struct{}

// @Tags Attach
// @Summary 上传文件
// @Produce application/json
// @Param file path string true "文件"
// @Success 200 {string} json "{"code":200,"data":null,"msg":"上传成功！"}"
// @Router /api/v1/attach/upload [post]
func (a *AttachApi) Upload(ctx *gin.Context)  {
	var uploadDir string
	uploadDir = global.GnConfig.Local.Path + "images/"
	//log.Println(uploadDir)

	isPath,_ := attachService.PathExists(uploadDir)
	if !isPath {
		os.MkdirAll(uploadDir,os.ModePerm)
	}

	//获取上传的源文件
	srcFile,head,err := ctx.Request.FormFile("file")
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	//创建一个新文件
	suffix := ".png"
	uploadFileName := head.Filename
	tmp := strings.Split(uploadFileName,".")
	if len(tmp) > 1{
		suffix = "."+tmp[len(tmp) -1]
	}
	fileType := ctx.Request.FormValue("filetype")
	if len(fileType) > 0{
		suffix = fileType
	}
	prefixName := fmt.Sprintf("%d%04d",time.Now().Unix(),rand.Int31())
	fileName := fmt.Sprintf("%s%s", prefixName, suffix)
	dsFile,err := os.Create(uploadDir+fileName)
	if err != nil{
		response.FailWithMessage(ctx, err.Error())
		return
	}
	//将源文件内容copy到新文件
	_,err = io.Copy(dsFile,srcFile)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}

	file, err := os.Open(uploadDir+fileName)
	if err != nil {
		response.FailWithMessage(ctx, err.Error())
		return
	}
	var img image.Image
	var imgErr error
	switch suffix {
	case ".png":
		img, imgErr = png.Decode(file)
	case ".jpg":
		img, imgErr = jpeg.Decode(file)
	case ".jpeg":
		img, imgErr = jpeg.Decode(file)
	case ".gif":
		img, imgErr = gif.Decode(file)
	default:
		imgErr = errors.New("不支持该图片格式")
	}
	if imgErr != nil {
		response.FailWithMessage(ctx, "图片格式错误："+imgErr.Error())
		return
	}

	imgWidth := img.Bounds().Max.X
	imgHeight := img.Bounds().Max.Y
	log.Println("宽：",imgWidth,"高：",imgHeight)
	file.Close()

	// resize to width 300 using Lanczos resampling
	// and preserve aspect ratio
	//var reWidth uint
	//if imgWidth > 300 {
	//	reWidth = 300
	//} else {
	//	reWidth = 0
	//}
	//m := resize.Resize(uint(imgWidth), uint(imgHeight), img, resize.Lanczos3)
	//thumbName := fmt.Sprintf("thumb_%s%s", prefixName, suffix)
	//out, err := os.Create(uploadDir+thumbName)
	//if err != nil {
	//	response.Fail(ctx, -1, err.Error())
	//	return
	//}
	//defer out.Close()
	//// write new image to file
	//jpeg.Encode(out, m, nil)

	//将新文件路径转换成url地址
	urls := []string{utils.TransformImageUrl("/"+uploadDir+fileName)}

	response.Success(ctx, urls, "上传成功！")
}