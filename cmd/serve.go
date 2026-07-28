package cmd

import (
	"ecommerce/config"
	"ecommerce/infrastructure/db"
	"ecommerce/repo"
	"ecommerce/rest"
	"ecommerce/rest/handlers/products"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"os"
)

func Serve() {
	cnf := config.GetConfig()

	dbConnection, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbConnection, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)
	productRepo := repo.NewProductRepo(dbConnection)
	userRepo := repo.NewUserRepo(dbConnection)
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
