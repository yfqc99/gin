package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ArticleModel struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

var articleList = []ArticleModel{
	{"Go", "文章0"},
	{"python", "文章1"},
	{"javaScript", "文章2"},
}

func _getList(ctx *gin.Context) {
	// 搜索，分页
	//ctx.JSON(200, articleList)
	/*[
		{
		"title": "Go",
		"content": "文章0"
		},
		{
		"title": "python",
		"content": "文章1"
		},
		{
		"title": "javaScript",
		"content": "文章2"
		}
	]*/
	ctx.JSON(200, Response{
		http.StatusOK,
		articleList,
		"获取成功",
	})
	/*{
		"code": 200,
		"data": [
	{
	"title": "Go",
	"content": "文章0"
	},
	{
	"title": "python",
	"content": "文章1"
	},
	{
	"title": "javaScript",
	"content": "文章2"
	}
	],
	"msg": "获取成功"
	}*/
}

func _getDetail(ctx *gin.Context) {
	// 获取查询参数中的id
	id, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id)
	ctx.JSON(200, Response{
		http.StatusOK,
		articleList[id],
		"获取成功",
	})
	// 0
	/*{
		"code": 200,
		"data": {
		"title": "Go",
			"content": "文章0"
	},
		"msg": "获取成功"
	}*/
	// 1
	/*{
		"code": 200,
		"data": {
		"title": "python",
			"content": "文章1"
	},
		"msg": "获取成功"
	}*/
}

func _create(ctx *gin.Context) {
	b, _ := ctx.GetRawData()
	var article ArticleModel
	json.Unmarshal(b, &article)
	articleList = append(articleList, article)
	ctx.JSON(200, Response{
		http.StatusOK,
		articleList,
		"添加成功",
	})
	/*{
		"code": 200,
		"data": [
	{
	"title": "Go",
	"content": "文章0"
	},
	{
	"title": "python",
	"content": "文章1"
	},
	{
	"title": "javaScript",
	"content": "文章2"
	},
	{
	"title": "西游记",
	"content": "文章3"
	}
	],
	"msg": "添加成功"
	}*/
}

func _update(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id)
	b, _ := ctx.GetRawData()
	var article ArticleModel
	json.Unmarshal(b, &article)
	articleList[id] = article
	ctx.JSON(200, Response{
		http.StatusOK,
		articleList,
		"修改成功",
	})
	//2
	/*{
		"code": 200,
		"data": [
	{
	"title": "Go",
	"content": "文章0"
	},
	{
	"title": "python",
	"content": "文章1"
	},
	{
	"title": "红楼梦",
	"content": "文章4"
	}
	],
	"msg": "修改成功"
	}*/
}

func _delete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	fmt.Println(id)
	articleList = append(articleList[:id], articleList[id+1:]...)
	ctx.JSON(200, Response{
		http.StatusOK,
		articleList,
		"删除成功",
	})
	// 2
	/*{
		"code": 200,
		"data": [
	{
	"title": "Go",
	"content": "文章0"
	},
	{
	"title": "python",
	"content": "文章1"
	},
	{
	"title": "西游记",
	"content": "文章3"
	}
	],
	"msg": "删除成功"
	}*/
}

func main() {
	// restful 资源定位和资源操作的风格
	// get 从服务器获取资源
	// post 在服务器建立新资源
	// put 在服务器更新资源(客户端提供完整资源数据)
	// patch 更新资源(客户端提供修改的资源数据)
	// delete 在服务器删除资源
	router := gin.Default()
	router.GET("/articles", _getList)       //列表
	router.GET("/articles/:id", _getDetail) //详情
	router.POST("/articles", _create)       // 添加
	router.PUT("/articles/:id", _update)    // 编辑
	router.DELETE("/articles/:id", _delete) // 删除
	router.Run(":8090")
}
