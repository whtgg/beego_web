package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"strconv"
	"time"
	"log"
	"mime/multipart"
	"fmt"
	"io"
	"path"
)

type UploadController struct {
	BaseController
	ImagesFilePath string
	FilesFilePath  string
}

const (
	PATH_IMAGE = "./static/upload/images"
)

func (ctx *UploadController) Prepare() {
	year, month, _ := time.Now().Date()
	ctx.ImagesFilePath = PATH_IMAGE + "/" + strconv.Itoa(year) + month.String() + "/"
	if err := os.MkdirAll(ctx.ImagesFilePath, 0777); err != nil {
		beego.Error(err)
	}
}

//图片上传
func (c *UploadController) UploadImg() {
	files, err := c.GetFiles("file")
	ch := make(chan string, len(files))
	if err != nil {
		log.Fatal("$v", err)
	}
	list := make([]string, 0)
	go c.uploads(files, ch)
	for v := range ch {
		list = append(list, v)
	}
	if len(list) == 0 {
		c.ajaxMsgWang("上传失败", list, MSG_ERR)
	}
	c.ajaxMsgWang("上传成功", list, MSG_OK)
}

func (c *UploadController) uploads(files []*multipart.FileHeader, ch chan<- string) {
	for k, _ := range files {
		fileSuffix := path.Ext(files[k].Filename)
		file, err := files[k].Open()
		if err != nil {
			fmt.Println("Err for Open", err)
		}
		defer file.Close()
		newname := strconv.FormatInt(time.Now().UnixNano(), 10) + strconv.Itoa(k) + fileSuffix
		path_url := c.ImagesFilePath + newname
		dst, _ := os.Create(path_url)
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			log.Fatal("$v", err)
		} else {
			ch <- path_url[1:]
		}
	}
	close(ch)
}
