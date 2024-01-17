package routes

import (
	"kredit_plus/src/app/handlers"
	"kredit_plus/src/middlewares"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register-admin", handlers.RegisterAdmin)
	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)

	AccountMiddlewareroute := r.Group("/change-password")
	AccountMiddlewareroute.Use(middlewares.PublicMiddleware())
	AccountMiddlewareroute.PATCH("", handlers.UpdatePassword)

	r.GET("/phones", handlers.GetAllPhone)
	r.GET("/phones/:id", handlers.GetPhoneById)
	r.GET("/phones/:id/specs-comments", handlers.GetSpecCommentByPhoneId)

	phonesMiddlewareroute := r.Group("/phones")
	phonesMiddlewareroute.Use(middlewares.AdminMiddleware())
	phonesMiddlewareroute.POST("/", handlers.CreatePhone)
	phonesMiddlewareroute.PATCH("/:id", handlers.UpdatePhone)
	phonesMiddlewareroute.DELETE("/:id", handlers.DeletePhone)

	r.GET("/brands", handlers.GetAllBrand)
	r.GET("/brands/:id", handlers.GetBrandById)
	r.GET("/brands/:id/phones", handlers.GetPhonesByBrandId)

	merkMiddlewareroute := r.Group("/brands")
	merkMiddlewareroute.Use(middlewares.AdminMiddleware())
	merkMiddlewareroute.POST("/", handlers.CreateBrand)
	merkMiddlewareroute.PATCH("/:id", handlers.UpdateBrand)
	merkMiddlewareroute.DELETE("/:id", handlers.DeleteBrand)

	r.GET("/specs", handlers.GetAllSpec)
	r.GET("/specs/:id", handlers.GetSpecById)

	specMiddlewareroute := r.Group("/specs")
	specMiddlewareroute.Use(middlewares.AdminMiddleware())
	specMiddlewareroute.POST("/", handlers.CreateSpec)
	specMiddlewareroute.PATCH("/:id", handlers.UpdateSpec)
	specMiddlewareroute.DELETE("/:id", handlers.DeleteSpec)

	r.GET("/news", handlers.GetAllNews)
	r.GET("/news/:id", handlers.GetNewsById)
	r.GET("/news/:id/comments", handlers.GetCommentByNewsId)

	newsMiddlewareroute := r.Group("/news")
	newsMiddlewareroute.Use(middlewares.AdminMiddleware())
	newsMiddlewareroute.POST("/", handlers.CreateNews)
	newsMiddlewareroute.PATCH("/:id", handlers.UpdateNews)
	newsMiddlewareroute.DELETE("/:id", handlers.DeleteNews)

	r.GET("/comments-phone", handlers.GetAllCommentPhone)
	// r.GET("/comments-phone/:id", handlers.GetCommentPhoneById)

	commentphoneMiddlewareroute := r.Group("/comments-phone")
	commentphoneMiddlewareroute.Use(middlewares.PublicMiddleware())
	commentphoneMiddlewareroute.POST("/", handlers.CreateCommentPhone)
	commentphoneMiddlewareroute.DELETE("/:id", handlers.DeleteCommentPhone)

	r.GET("/comments-news", handlers.GetAllCommentNews)
	// r.GET("/comments-news/:id", handlers.GetCommentNewsById)

	commentnewsMiddlewareroute := r.Group("/comments-news")
	commentnewsMiddlewareroute.Use(middlewares.PublicMiddleware())
	commentnewsMiddlewareroute.POST("/", handlers.CreateCommentNews)
	commentnewsMiddlewareroute.DELETE("/:id", handlers.DeleteCommentNews)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
