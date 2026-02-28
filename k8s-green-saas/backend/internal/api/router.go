package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/k8s-green/backend/internal/api/handlers"
	"github.com/k8s-green/backend/internal/api/middleware"
	"github.com/k8s-green/backend/internal/config"
	"github.com/k8s-green/backend/internal/repository"
)

func NewRouter(cfg *config.Config, db *repository.Postgres, rdb *repository.Redis) *gin.Engine {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	h := handlers.NewHandler(db, rdb, cfg)

	// Middlewares globaux
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS(cfg))

	// Health / readiness (sans auth)
	r.GET("/health", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"status": "ok"}) })
	r.GET("/ready", func(c *gin.Context) {
		if err := db.Pool.Ping(c.Request.Context()); err != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "db unreachable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ready"})
	})

	// API v1 — authentifié + tenant
	v1 := r.Group("/v1",
		middleware.Auth(cfg),
		middleware.Tenant(),
		middleware.RateLimit(rdb),
		middleware.Audit(db),
	)

	// Clusters
	clusters := v1.Group("/clusters")
	{
		clusters.GET("",          h.ListClusters)
		clusters.POST("",         h.CreateCluster)
		clusters.GET("/:id/score",           h.GetScore)
		clusters.GET("/:id/waste",           h.GetWaste)
		clusters.GET("/:id/carbon",          h.GetCarbon)
		clusters.GET("/:id/savings",         h.GetSavings)
		clusters.GET("/:id/recommendations", h.GetRecommendations)
		clusters.POST("/:id/simulate",       h.Simulate)
		clusters.GET("/:id/history",         h.GetScoreHistory)
	}

	// Rapports
	v1.GET("/reports", h.GetReports)

	// WebSocket live score (auth seulement, pas de rate limit)
	ws := r.Group("/v1/ws",
		middleware.Auth(cfg),
		middleware.Tenant(),
	)
	ws.GET("/clusters/:id/live", h.LiveWebSocket)

	// Admin (superadmin uniquement)
	admin := v1.Group("/admin", middleware.RequireRole("superadmin"))
	{
		admin.GET("/tenants",     h.ListTenants)
		admin.POST("/tenants",    h.CreateTenant)
		admin.GET("/tenants/:id", h.GetTenant)
		admin.GET("/users",       h.ListUsers)
	}

	return r
}
