package main

import (
	"fmt"
	"kulkasku/internal/config"
	foodHandler "kulkasku/internal/handler/food"
	recipeHandler "kulkasku/internal/handler/recipe"
	userHandler "kulkasku/internal/handler/user"
	foodRepository "kulkasku/internal/repository/food"
	recipeRepository "kulkasku/internal/repository/recipe"
	userRepository "kulkasku/internal/repository/user"
	foodService "kulkasku/internal/service/food"
	recipeService "kulkasku/internal/service/recipe"
	userService "kulkasku/internal/service/user"
	"kulkasku/pkg/internalsql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	res := gin.Default()
	validate := validator.New()
	cfg, err := config.LoadConfigDatabase()
	if err != nil {
		panic(err)
	}

	db, err := internalsql.ConnectMySQL(cfg)
	if err != nil {
		panic(err)
	}

	res.Use(gin.Logger())
	res.Use(gin.Recovery())

	res.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Connected Gin Golang",
		})
	})

	userRepository := userRepository.NewRepository(db)
	userService := userService.NewService(cfg, userRepository)
	uh := userHandler.NewHandler(res, validate, userService)
	uh.RouteList(cfg.SecretJwt)

	foodRepository := foodRepository.NewRepository(db)
	foodService := foodService.NewService(foodRepository)
	fh := foodHandler.NewHandler(res, validate, foodService)
	fh.RouteList(cfg.SecretJwt)

	recipeRepo := recipeRepository.NewRepository(db)
	recipeSvc := recipeService.NewService(foodRepository, recipeRepo)
	rh := recipeHandler.NewHandler(res, validate, recipeSvc)
	rh.RouteList(cfg.SecretJwt)

	server := fmt.Sprintf("127.0.0.1:%s", cfg.Port)
	res.Run(server)
}
