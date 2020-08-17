package routers

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/mesment/mblog/middleware/jwt"
	"github.com/mesment/mblog/pkg/logger"
	"github.com/mesment/mblog/pkg/setting"
	"github.com/mesment/mblog/routers/api"
	"github.com/mesment/mblog/routers/api/v1"
	"time"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	// Get raw zap logger
	zlog := logger.NewLogger().Raw()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(zlog, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(zlog, true))

	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	gin.SetMode(setting.RunMode)

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")

	// api/v1 都需要进行身份验证
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}