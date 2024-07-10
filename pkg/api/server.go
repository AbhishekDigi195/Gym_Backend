package api

import (
	"gym/pkg/api/handlers"
	middleware "gym/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func NewServer(r *gin.Engine,handler *handlers.Handler) *gin.Engine{
	r.Use(middleware.CorsMiddleware)
    v1 := r.Group("/api/v0")
    {
        v1.POST("/login", handler.Login)
        // Apply the AuthMiddleware to all routes below this line
        v1.Use(middleware.AuthMiddleware())
        {
            v1.POST("/send_otp", handler.SendOtp)
            v1.POST("/check_otp", handler.CheckOTP)
            v1.POST("/register_admin", handler.RegisterAdmin)
            v1.POST("/register", handler.RegisterUser)
            v1.POST("/bulkregister", handler.BulkRegisterUser)
            v1.POST("/verify", handler.VerifyUser)
            v1.POST("/purchase", handler.PurchaseMembership)
            v1.POST("/activate", handler.ActivateMembership)
            v1.POST("/access", handler.AccessControl)
            v1.POST("/update_status", handler.UpdateMembershipStatus)
            v1.POST("/check_status", handler.CheckMembershipStatus)
            v1.POST("/renew", handler.RenewMembership)
        }
    }
    
    return r
}