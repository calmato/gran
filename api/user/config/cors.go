package config

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
	wrapper "github.com/rs/cors/wrapper/gin"
)

func SetCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		options := cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: false,
			MaxAge:           1440, // 60m * 24h
			Debug:            false,
		}

		wrapper.New(options)
		ctx.Next()
	}
}
