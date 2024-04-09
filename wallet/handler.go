package wallet

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	store Storer
}

type Storer interface {
	Wallets() ([]Wallet, error)
	FindByType(walletType string) ([]Wallet, error)
	FindById(id int) ([]Wallet, error)
	Insert(userId int, userName, walletName, walletType string, balance float64) error
	Update(id int, walletName, walletType string, balance float64) (int, error)
	Delete(userId int) (int, error)
}

func New(db Storer) *Handler {
	return &Handler{store: db}
}

type Err struct {
	Message string `json:"message"`
}

// WalletHandler
//
//	@Summary		Get all wallets
//	@Description	Get all wallets
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Param			wallet_type  	query    string  false  "wallet type (Savings, Credit Card, Crypto Wallet)"
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [get]
//	@Failure		500	{object}	Err
func (h *Handler) WalletHandler(c echo.Context) error {
	walletType := c.QueryParam("wallet_type")
	var wallets []Wallet
	var err error
	fmt.Printf("type = %s", walletType)
	if walletType != "" {
		wallets, err = h.store.FindByType(walletType)
	} else {
		wallets, err = h.store.Wallets()
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

// FindWalletByIdHandler
//
//	@Summary		Get all wallets by user id
//	@Description	Get all wallets by user id
//	@Tags			wallet
//	@Success		200	{object}	string
//	@Router			/api/v1/users/{id}/wallets [get]
//	@Param			id  path int true  "user id"
//	@Failure		500	{object}	Err
func (h *Handler) FindWalletByIdHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	wallets, err := h.store.FindById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, wallets)
}

// CreateWalletHandler
//
//	@Summary		Create Wallet
//	@Description	Create Wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	Wallet
//	@Router			/api/v1/wallets [post]
//	@Failure		500	{object}	Err
func (h *Handler) CreateWalletHandler(c echo.Context) error {
	var wallet Wallet
	err := c.Bind(&wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	if wallet.UserID == 0 || wallet.UserName == "" || wallet.WalletName == "" || wallet.WalletType == "" || wallet.Balance == 0.0 {
		return c.JSON(http.StatusBadRequest, Err{Message: "request body is invalid!"})
	}
	err = validateWalletType(wallet.WalletType)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}
	err = h.store.Insert(wallet.UserID, wallet.UserName, wallet.WalletName, wallet.WalletType, wallet.Balance)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.String(http.StatusCreated, "create wallet success")
}

// UpdateWalletHandler
//
//	@Summary		Update Wallet
//	@Description	Update Wallet
//	@Tags			wallet
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	UpdateWallet
//	@Router			/api/v1/wallets [put]
//	@Failure		500	{object}	Err
func (h *Handler) UpdateWalletHandler(c echo.Context) error {
	var wallet UpdateWallet
	err := c.Bind(&wallet)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	if wallet.ID == 0 || wallet.WalletName == "" || wallet.WalletType == "" || wallet.Balance == 0.0 {
		return c.JSON(http.StatusBadRequest, Err{"request body is invalid"})
	}
	row, err := h.store.Update(wallet.ID, wallet.WalletName, wallet.WalletType, wallet.Balance)
	fmt.Printf("Update wallets total := %d \n", row)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.String(http.StatusOK, "update wallet success")
}

// DeleteWalletHandler
//
//	@Summary		Get all wallets by user id
//	@Description	Get all wallets by user id
//	@Tags			wallet
//	@Success		200	{object}	string
//	@Router			/api/v1/users/{id}/wallets [delete]
//	@Param			id  path int true  "user id"
//	@Failure		500	{object}	Err
func (h *Handler) DeleteWalletHandler(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	row, err := h.store.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}
	return c.JSON(http.StatusOK, fmt.Sprintf("total delete wallets %d", row))
}

func validateWalletType(walletType string) error {
	if walletType != "Savings" && walletType != "Credit Card" && walletType != "Crypto Wallet" {
		return errors.New("Wallet type is incorrect!")
	}
	return nil
}
