package wallet

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/labstack/echo/v4"
)

type StubWallets struct {
	wallets     []Wallet
	err         error
	rowAffected int
}

func (s StubWallets) Wallets() ([]Wallet, error) {
	return s.wallets, s.err
}

func (s StubWallets) FindById(id int) ([]Wallet, error) {
	return s.wallets, s.err
}

func (s StubWallets) FindByType(walletType string) ([]Wallet, error) {
	return s.wallets, s.err
}

func (s StubWallets) Insert(userId int, userName, walletName, walletType string, balance float64) error {
	return s.err
}
func (s StubWallets) Update(id int, walletName, walletType string, balance float64) (int, error) {
	return s.rowAffected, s.err
}

func (s StubWallets) Delete(id int) (int, error) {
	return s.rowAffected, s.err
}

func TestWallets(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/v1/wallets")
		stubWallet := StubWallets{
			nil,
			echo.ErrInternalServerError,
			0,
		}
		w := New(stubWallet)
		w.WalletHandler(c)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})
	t.Run("given user able to getting wallet should return list of wallets", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/v1/wallets")
		stubWallet := StubWallets{
			[]Wallet{
				{ID: 1, UserID: 1234, UserName: "A1", WalletName: "A1_1"},
				{ID: 2, UserID: 1234, UserName: "A1", WalletName: "A1_2"},
				{ID: 3, UserID: 5678, UserName: "A2", WalletName: "A2_1"},
				{ID: 4, UserID: 6789, UserName: "A2", WalletName: "A2_2"},
				{ID: 5, UserID: 0000, UserName: "A3", WalletName: "A3_1"},
			},
			nil,
			0,
		}
		w := New(stubWallet)
		w.WalletHandler(c)
		gotJson := rec.Body.Bytes()
		var got []Wallet
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if !reflect.DeepEqual(got, stubWallet.wallets) {
			t.Errorf("expected %v but got %v", stubWallet.wallets, got)
		}
	})
}

func TestWalletsByType(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/v1/wallets")
		stubWallet := StubWallets{
			nil,
			echo.ErrInternalServerError,
			0,
		}
		w := New(stubWallet)
		w.WalletHandler(c)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})
	t.Run("given user able to getting wallet by type should return list of wallets", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/v1/wallets")
		c.SetParamNames("wallet_type")
		c.SetParamValues("saving")
		stubWallet := StubWallets{
			[]Wallet{
				{ID: 1, UserID: 1234, UserName: "A1", WalletName: "A1_1", WalletType: "saving"},
				{ID: 2, UserID: 1234, UserName: "A1", WalletName: "A1_2", WalletType: "saving"},
			},
			nil,
			0,
		}
		w := New(stubWallet)
		w.WalletHandler(c)
		gotJson := rec.Body.Bytes()
		var got []Wallet
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if !reflect.DeepEqual(got, stubWallet.wallets) {
			t.Errorf("expected %v but got %v", stubWallet.wallets, got)
		}
	})
}

func TestFindWalletsById(t *testing.T) {
	t.Run("given unable to get wallets should return 500 and error message", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/v1/users/:id/wallets")
		stubWallet := StubWallets{
			nil,
			echo.ErrInternalServerError,
			0,
		}
		w := New(stubWallet)
		w.WalletHandler(c)
		if rec.Code != http.StatusInternalServerError {
			t.Errorf("expected status code %d but got %d", http.StatusInternalServerError, rec.Code)
		}
	})
	t.Run("given user able to getting wallet by id should return list of wallets", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/api", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/v1/users/:id/wallets")
		c.SetParamNames("id")
		c.SetParamValues("1")
		stubWallet := StubWallets{
			[]Wallet{
				{ID: 1, UserID: 1, UserName: "A1", WalletName: "A1_1"},
				{ID: 2, UserID: 1, UserName: "A1", WalletName: "A1_2"},
				{ID: 3, UserID: 1, UserName: "A1", WalletName: "A1_3"},
				{ID: 4, UserID: 1, UserName: "A1", WalletName: "A1_4"},
				{ID: 5, UserID: 1, UserName: "A1", WalletName: "A1_5"},
			},
			nil,
			0,
		}
		w := New(stubWallet)
		w.WalletHandler(c)
		gotJson := rec.Body.Bytes()
		var got []Wallet
		if err := json.Unmarshal(gotJson, &got); err != nil {
			t.Errorf("unable to unmarshal json: %v", err)
		}
		if !reflect.DeepEqual(got, stubWallet.wallets) {
			t.Errorf("expected %v but got %v", stubWallet.wallets, got)
		}
	})
}
