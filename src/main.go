package main

import (
	"embed"
	"io/fs"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

//go:embed public/*
var publicFS embed.FS

// 定义日期格式
const LAYOUT = "2006/01/02"

const BASE_URL = "http://127.0.0.1:8088"

func getFileName(ext string) string {
	// 获取当前日期
	now := time.Now()

	nanoid, _ := gonanoid.New()

	return strings.Join([]string{
		"./upload",
		now.Format(LAYOUT),
		nanoid + ext,
	}, "/")
}

func main() {
	app := gin.Default()

	// 上传文件目录
	app.Static("/upload", "upload")

	// 静态资源
	public, _ := fs.Sub(publicFS, "public")

	// 静态资源
	assets, _ := fs.Sub(public, "assets")
	app.StaticFS("/assets", http.FS(assets))
	app.StaticFileFS("/vite.svg", "vite.svg", http.FS(public))

	// 首页
	app.GET("/", func(ctx *gin.Context) {
		data, _ := fs.ReadFile(public, "index.html")
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", data)
	})

	// 文件上传接口
	app.POST("/api/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")

		ext := filepath.Ext(file.Filename)
		dst := getFileName(ext)

		// 上传文件至指定的完整文件路径
		ctx.SaveUploadedFile(file, dst)

		fileUrl, _ := url.JoinPath(BASE_URL, dst)

		ctx.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": gin.H{
				"fileUrl": fileUrl,
			},
			"msg": "success",
		})
	})

	// 监听并在 http://127.0.0.1:8080 上启动服务
	app.Run(":8088")
}
