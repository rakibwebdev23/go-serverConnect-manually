package cmd

import (
	"ecommerce/config"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	middlewares := middleware.NewMiddlewares(cnf)

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUserRepo()
	reviewRepo := repo.NewReviewRepo()

	productHandler := products.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)
	reviewHandler := review.NewHandler(reviewRepo)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
		reviewHandler,
	)
	server.Start()
}

