package postgres

import (
	"errors"
	"fmt"
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Wallets() ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) FindByType(walletType string) ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE wallet_type = $1", walletType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) FindById(id int) ([]wallet.Wallet, error) {
	rows, err := p.Db.Query("SELECT * FROM user_wallet WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) Insert(userId int, userName, wallet_name, walletType string, balance float64) error {
	query := `INSERT INTO user_wallet (user_id,user_name, wallet_name, wallet_type, balance)
	VALUES ($1, $2, $3, $4, $5)`
	_, err := p.Db.Exec(query, userId, userName, wallet_name, walletType, balance)
	return err
}

func (p *Postgres) Update(id int, wallet_name, walletType string, balance float64) (int, error) {
	query := `UPDATE user_wallet SET user_name = $1, wallet_name = $2, wallet_type = $3, balance = $4
	WHERE id = $5`
	res, err := p.Db.Exec(query, wallet_name, walletType, balance, id)
	if err != nil {
		return 0, err
	}
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowAffected == 0 {
		return 0, errors.New("")
	}
	return int(rowAffected), nil
}

func (p *Postgres) Delete(userId int) (int, error) {
	query := `DELETE FROM user_wallet WHERE user_id = $1`
	res, err := p.Db.Exec(query, userId)
	rowAffected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	if rowAffected == 0 {
		return 0, errors.New(fmt.Sprintf("Can't delete wallet by user id %d", rowAffected))
	}
	return int(rowAffected), nil
}
