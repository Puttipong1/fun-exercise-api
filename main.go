package main

import (
	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
// @BasePath /
func main() {
	p, err := postgres.New()
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(p)
	e.GET("/api/v1/wallets", handler.WalletHandler)
	e.POST("/api/v1/wallets", handler.CreateWalletHandler)
	e.PUT("/api/v1/wallets", handler.UpdateWalletHandler)
	e.GET("/api/v1/users/:id/wallets", handler.FindWalletByIdHandler)
	e.DELETE("/api/v1/users/:id/wallets", handler.DeleteWalletHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
