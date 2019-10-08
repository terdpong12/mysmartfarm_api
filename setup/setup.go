package setup

import (
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/rebelstech/ralunar-buyer-api/controllers"
)

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
}

func StartGin() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type, Ralunar-Platform, Ralunar-App-Version,Ralunar-Language,Ralunar-Dev-Model",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	v1_0 := router.Group("/v1.0")
	{
		//USER AUTHORIZATION
		v1_0.POST("/oauth/access_token", controllers.AuthenticationController)
		v1_0.POST("/oauth/refresh_token", controllers.RefreshToken)

		//HOME
		v1_0.GET("/home/brands", controllers.HomeBrands)
		v1_0.GET("/home/banners", controllers.HomeBanners)
		v1_0.GET("/home/featured_categories", controllers.HomeCategories)
		v1_0.GET("/home/featured_brands", controllers.HomeFeaturedBrands)
		v1_0.GET("/home/hot_brands", controllers.HomeHotBrands)
		v1_0.GET("/home/flash_sales", controllers.HomeFlashSale)
		v1_0.GET("/home/flash_sales/page/:page", controllers.HomeFlashSaleSeemore)
		v1_0.GET("/home/hot_products", controllers.HomeHotProducts)
		v1_0.GET("/home/hot_products/page/:page", controllers.HomeHotProductsSeemore)

		//GET CATEGORIES
		v1_0.GET("/categories/web", controllers.GetCategoriesWeb)
		v1_0.POST("/categories/level/:id", controllers.GetCategoriesFromLevelId)
		v1_0.POST("/categories/nested/:id", controllers.GetNestedCategories)
		v1_0.POST("/categories_web/nested/:id", controllers.GetNestedCategoriesWeb)

		//GET PRODUCT BY FILTERS
		v1_0.POST("/products/page/:page", controllers.GetProducts)
		v1_0.GET("/product/:slug", controllers.GetProduct)

		//PRODUCT COLLECTION SEE MORE
		v1_0.POST("/product_collection/:slug/page/:page", controllers.ProductCollection)

		//LOGISTIC
		v1_0.GET("/logistics/productId/:product_id", controllers.Logistics)

		//SEARCH
		v1_0.POST("/search_all", controllers.SearchAll)
		v1_0.POST("/search_recent", controllers.GetUserSearchRecentController)
		v1_0.POST("/search_recent/add", controllers.AddUserSearchRecentController)
		v1_0.POST("/search_recent/delete", controllers.DeleteUserSearchRecentController)

		//USER PROFILE
		v1_0.GET("/user/:uid/profile_ordered/status/:status/page/:page", controllers.GetUserProfileOrdered)
		v1_0.GET("/user/:uid/profile", controllers.GetUserProfile)
		v1_0.GET("/user/:uid/profile_settings", controllers.GetUserSettings)
		v1_0.PUT("/user/:uid/profile/edit", controllers.EditUserProfile)

		//USER LIKES
		v1_0.POST("/user/:uid/like", controllers.UserLike)
		v1_0.GET("/user/:uid/profile_liked/page/:page", controllers.GetUserProfileLiked)
		v1_0.GET("/user_likes/product/:product_id/page/:page", controllers.ListUserLikeProduct)

		//USER ACTIVITIES
		v1_0.POST("/user/:uid/change_password", controllers.ChangePassword)

		//USER SHIPPING ADDRESSES
		v1_0.GET("/user/:uid/shipping_address", controllers.GetUserShippingAddress)
		v1_0.POST("/user/:uid/shipping_address/add", controllers.AddUserShippingAddress)
		v1_0.PUT("/user/:uid/shipping_address/:id/edit", controllers.UpdateUserShippingAddress)
		v1_0.DELETE("/user/:uid/shipping_address/:id/delete", controllers.DeleteUserShippingAddress)

		//USER FORGOT PASSWORD
		v1_0.POST("/recover_pass", controllers.RecoverPassword)
		v1_0.POST("/reset_password/:key", controllers.ResetPassword)

		//CART
		v1_0.POST("/user/:uid/cart_product/add", controllers.AddUserCartProduct)
		v1_0.POST("/user/:uid/cart", controllers.GetUserCart)
		v1_0.DELETE("/user/:uid/cart", controllers.ClearUserCart)
		v1_0.DELETE("/user/:uid/cart_product_variant/:variant_id/delete", controllers.RemoveUserCartProduct)
		v1_0.PUT("/user/:uid/cart_product/quantity", controllers.UpdateUserCartProductQuantity)

		//ORDER
		v1_0.POST("/check_promotion_code", controllers.CheckPromotionCode)
		v1_0.POST("/user/:uid/order/checkout", controllers.Checkout)
		v1_0.POST("/user/:uid/cancel_order", controllers.UserCancelOrder)
		v1_0.GET("/order/detail/:order_no", controllers.OrderDetailController)

		//RATING
		v1_0.POST("/rating/orderId/:order_item_id", controllers.CreateRating)
		v1_0.GET("/rating/productId/:product_id/stars/:stars/pageId/:page_id", controllers.GetRating)
		v1_0.POST("/rating/productId/:product_id/product_image/upload", controllers.CreateRatingImage)

		//2C2P
		v1_0.POST("/2c2p/callback", controllers.CallBack2C2P)

		//THAILAND DATA
		v1_0.GET("/provinces", controllers.GetProvinces)
		v1_0.GET("/amphurs/:province_id", controllers.GetAmphures)
		v1_0.GET("/zipcodes/:amphur_id", controllers.GetZipcodes)

		//RESOURCES
		v1_0.GET("/resources", controllers.Resources)

		//HEALTH CHECK
		v1_0.GET("/healthcheck", controllers.HealthCheck)
		v1_0.GET("/redis_health_check", controllers.RedisHealthCheck)

		//BLOGS
		v1_0.GET("/blogs/all", controllers.GetBlogs)
		v1_0.GET("/blog/:slugblog", controllers.GetBlogSlug)
		v1_0.POST("/blogs/create", controllers.CreateBlogs)

		//PAYMENT SERVICE 2C2P
		v1_0.POST("/paymentService2c2p", controllers.PaymentService2c2p)

		//BANK TRANSFER
		v1_0.POST("/bankTransfer", controllers.CreateBankTransfer)
		v1_0.POST("/bankTransfer_image", controllers.AddBankTransferImage)

		//LALAMOVE
		v1_0.POST("/lalamove/quotation", controllers.LalamoveQuotation)
		v1_0.POST("/lalamove/orders", controllers.LalamoveOrderCreate)
		v1_0.PUT("/lalamove/orders", controllers.LalamoveOrderCancelling)
		// v1_0.GET("/lalamove/order/:lalamove_order_id", controllers.LalamoveDetailOrder)
	}
	return router
}
