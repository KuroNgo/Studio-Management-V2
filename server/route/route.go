package route

import (
	"co-studio-e-commerce/conf"
	"co-studio-e-commerce/docs"
	"co-studio-e-commerce/handler"
	"co-studio-e-commerce/middleware"
	"co-studio-e-commerce/repo"
	"co-studio-e-commerce/service"
	"co-studio-e-commerce/util"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Service struct {
	*conf.App
}

func NewService() *Service {
	s := Service{
		conf.NewApp(),
	}

	db := s.GetDB()
	repository := repo.NewRepo(db)

	userService := service.NewUser(repository)
	user := handler.NewUser(userService)

	categoryService := service.NewCategory(repository)
	category := handler.NewCategory(categoryService, userService)

	productService := service.NewProduct(repository)
	product := handler.NewProduct(productService, userService)

	// giải quyết về router
	route := s.Router

	// add middleware
	route.Use(middleware.CORSMiddleware())
	route.Use(middleware.Recover())
	route.Use(middleware.StructuredLogger(&log.Logger))
	route.Use(gzip.Gzip(gzip.DefaultCompression, gzip.WithExcludedPaths([]string{",*"})))

	route.MaxMultipartMemory = 25 << 20 // 8 MiB

	docs.SwaggerInfo.BasePath = "api/v1"
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//route automatically
	//Thực hiện tự động chuyển hướng khi chạy chương trình
	route.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusFound, "/swagger/index.html")
	})

	// option method to fix preflight request
	route.OPTIONS("/*path", middleware.OptionMessage)

	// Thêm endpoint pprof
	route.GET("/debug/pprof/*any", gin.WrapH(http.DefaultServeMux))

	v1User := route.Group("api/v1/user")
	{

		v1User.Use(
			middleware.CORSForDev(),
			middleware.RateLimiter(),
		)

		// auth

	}

	clientV1 := route.Group("api/v1/client")
	{
		clientV1.Use(
			middleware.CORSMiddleware(),
			middleware.RateLimiter(),
		)

		// auth
		clientV1.POST("/login/username", user.LoginWithUserName)
		clientV1.POST("/login/email", user.LoginWithEmail)
		clientV1.POST("/register", user.Register)
		clientV1.GET("/get-user", user.GetMeV2)
		clientV1.GET("/refresh", user.RefreshAccessToken)
		clientV1.GET("/logout", middleware.DeserializeUser(), user.LogoutUser)
		clientV1.PUT("/update", user.UpdateUser)

		// không cần đăng nhập vẫn sử dụng được
		// category
		clientV1.GET("/category/get-all", category.GetAllCategoriesForView)
		clientV1.GET("/category/get/:category_id", category.GetCategoryByIDForView)

		// product
		clientV1.GET("/product/get-all", product.GetAllProductForView)
		clientV1.GET("/product/get-name", product.GetProductByNameForView)
		clientV1.GET("/product/get-id", product.GetProductByCategoryIDForView)
		clientV1.GET("/product/get-price", product.GetProductByPriceForView)

		// image
		clientV1.GET("/image/get", util.GetUploadedFile)
		clientV1.POST("/image/upload", util.UploadAFile)
		clientV1.PUT("/images/uploads", util.UpdateFile)
		clientV1.DELETE("/image/delete", util.DeleteFile)
		clientV1.POST("/image/", util.UploadMultiFile)
	}

	// Đối với admin thì không cần phải rate limit
	// Nhưng đối với nhân viên thì khác
	// Cơ chế rate limit sẽ giới hạn đối với nhân viên
	// Nhưng số lần gửi request giới hạn sẽ được tăng lên so với người dùng user
	adminV1 := route.Group("api/v1/admin")
	{
		adminV1.Use(
			middleware.CORSForDev(),
			middleware.DeserializeUser(),
		)

		// user
		adminV1.GET("/get-all-user", user.GetAllUser)
		adminV1.GET("/get-role", user.GetUserByRole)
		adminV1.GET("/get-role-edit", user.FindUserByRole)

		// category
		adminV1.POST("/category/create", category.CreateCategory)
		adminV1.PUT("/category/update", category.UpdateCategory)
		adminV1.PATCH("/category/delete-first", category.DeleteCategoryFirst)
		adminV1.GET("/category/get-id-edit/:category_id", category.GetCategoryByIDForEdit)
		adminV1.GET("/category/get-update-date/:update_date", category.GetCategoryByUpdateDateForEdit)
		adminV1.GET("/category/get-enable/:enable", category.GetCategoryByEnableForEdit)
		adminV1.DELETE("/category/delete-second", category.DeleteCategorySecond)
		adminV1.PATCH("/category/resolve", category.ResolveCategory)

		// product
		adminV1.GET("/product/get-who_update", product.GetProductByWhoUpdateForView)
		adminV1.GET("/product/get-update_date", product.GetProductByUpdateDateForView)
		adminV1.POST("/product/create", product.CreateProduct)
		adminV1.PATCH("/product/update", product.UpdateProduct)
		adminV1.PATCH("/product/update-enable", product.UpdateEnable)
		adminV1.PATCH("/product/remove-first", product.Disable)
		adminV1.PATCH("/product/enable", product.Enable)
	}

	return &s
}
