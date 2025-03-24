package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func main() {
	router := gin.Default()
	//router.POST("/upload", func(context *gin.Context) {
	//	// 通过form-data 上传文件
	//	file, _ := context.FormFile("file")
	//	//fmt.Println(file.Filename)
	//	// 限制用户上传文件的大小，单位：字节
	//	//fmt.Println(file.Size)
	//	/*	金融报告.docx
	//		19460*/
	//	// 保存上传的文件,./表示当前项目下
	//	dst := "./document/" + file.Filename
	//	context.SaveUploadedFile(file, dst)
	//	context.JSON(200, "上传成功")
	//})

	router.POST("/upload", func(context *gin.Context) {
		// 上传多个文件时，只能获取到第一个 参数是上传的表单中对应的参数名
		file, _ := context.FormFile("file")
		readFile, _ := file.Open() // 返回文件对象(只读),直接读取文件所有内容
		// 获取文件对象内的数据
		//data, _ := io.ReadAll(readFile)
		// 可以对文件内容进行处理，例如判断是否合法
		dst := "./document/" + file.Filename
		writeFile, _ := os.Create(dst) // 创建指定的文件对象
		io.Copy(writeFile, readFile)
		defer writeFile.Close()
		defer readFile.Close()
		context.JSON(200, "上传成功")
	})

	router.POST("/uploads", func(context *gin.Context) {
		// 上传多个文件
		form, _ := context.MultipartForm()
		//File  map[string][]*FileHeader
		// 切片 参数是上传的表单中对应的参数名
		files := form.File["upload[]"]
		for _, file := range files {
			context.SaveUploadedFile(file, "./document/"+file.Filename)
		}
		context.JSON(200, "上传成功")
	})

	router.Run(":8090")
}

/*func (c *Context) FormFile(name string) (*multipart.FileHeader, error) {
	if c.Request.MultipartForm == nil {
		if err := c.Request.ParseMultipartForm(c.engine.MaxMultipartMemory); err != nil {
			return nil, err
		}
	}
	f, fh, err := c.Request.FormFile(name)
	if err != nil {
		return nil, err
	}
	f.Close()
	return fh, err
}*/
/*func (c *Context) MultipartForm() (*multipart.Form, error) {
	err := c.Request.ParseMultipartForm(c.engine.MaxMultipartMemory)
	return c.Request.MultipartForm, err
}*/
