package service

import (
	"log"
	"os"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupCors(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND")},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func SetupSecurity(r *gin.Engine) {
	r.SetTrustedProxies([]string{os.Getenv("FRONTEND")})
	r.Use(secure.New(secure.Config{
		// AllowedHosts:          []string{os.Getenv("FRONTEND"), os.Getenv("SSL_HOST")},
		// SSLRedirect:           true,
		SSLHost:               os.Getenv("SSL_HOST"),
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             false,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'; style-src 'self' 'unsafe-inline';",
		IENoOpen:              true,
		ReferrerPolicy:        "strict-origin-when-cross-origin",
		// SSLProxyHeaders:       map[string]string{"X-Forwarded-Proto": "https"},
	}))
}

func SetupLogger(r *gin.Engine) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))
}

func SetupRLimit() {
	var rLimit syscall.Rlimit
	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Fatalf("Error Getting Rlimit %v", err)
	}
	rLimit.Max = 25600
	rLimit.Cur = 25600

	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Fatalf("Error Setting Rlimit %v", err)
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		log.Fatalf("Error Getting Rlimit %v", err)
	}
}
