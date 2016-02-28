package mysql

import (
	"fmt"
	"testing"
	"time"

	. "github.com/freeusd/solebtc/Godeps/_workspace/src/github.com/smartystreets/goconvey/convey"
	"github.com/freeusd/solebtc/errors"
	"github.com/freeusd/solebtc/models"
)

func TestDeductUserBalanceBy(t *testing.T) {
	Convey("Given empty mysql storage", t, func() {
		s := prepareDatabaseForTesting()

		Convey("When deduct user balance with commited transaction", func() {
			tx := s.db.MustBegin()
			tx.Commit()
			err := deductUserBalanceBy(tx, 0, 0)
			Convey("Error should be unknown", func() {
				So(err.ErrCode, ShouldEqual, errors.ErrCodeUnknown)
			})

			Reset(func() { tx.Rollback() })
		})

		Convey("When deduct user balance affecting 0 row", func() {
			tx := s.db.MustBegin()
			err := deductUserBalanceBy(tx, 0, 0)
			Convey("Error should be insufficient balance", func() {
				So(err.ErrCode, ShouldEqual, errors.ErrCodeInsufficientBalance)
			})

			Reset(func() { tx.Rollback() })
		})
	})
}

func TestInsertWithdrawal(t *testing.T) {
	Convey("Given empty mysql storage", t, func() {
		s := prepareDatabaseForTesting()

		Convey("When insert withdrawal with commited transaction", func() {
			tx := s.db.MustBegin()
			tx.Commit()
			err := insertWithdrawal(tx, 0, "", 0)
			Convey("Error should be unknown", func() {
				So(err.ErrCode, ShouldEqual, errors.ErrCodeUnknown)
			})

			Reset(func() { tx.Rollback() })
		})
	})
}

func TestCreateWithdrawal(t *testing.T) {
	Convey("Given empty mysql storage", t, func() {
		s := prepareDatabaseForTesting()
		s.CreateUser(models.User{Email: "e", BitcoinAddress: "b"})
		s.CreateRewardIncome(models.Income{UserID: 1, Income: 10}, time.Now())

		Convey("When create withdrawal", func() {
			err := s.CreateWithdrawal(models.Withdrawal{
				UserID:         1,
				BitcoinAddress: "b",
				Amount:         5,
			})

			Convey("Error should be nil", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestGetWithdrawalsSince(t *testing.T) {
	Convey("Given mysql storage", t, func() {
		s := prepareDatabaseForTesting()
		s.db.MustExec("INSERT INTO `users` (email, bitcoin_address, balance) VALUES(?, ?, ?);", "e", "b", 8388607)
		s.CreateWithdrawal(models.Withdrawal{UserID: 1, BitcoinAddress: "b", Amount: 1})
		s.CreateWithdrawal(models.Withdrawal{UserID: 1, BitcoinAddress: "b", Amount: 2})
		s.CreateWithdrawal(models.Withdrawal{UserID: 1, BitcoinAddress: "b", Amount: 3})

		Convey("When get withdrawals since now", func() {
			result, _ := s.GetWithdrawalsSince(1, time.Now().AddDate(0, 0, -1), 2)

			Convey("Withdrawals should equal", func() {
				So(result, func(actual interface{}, expected ...interface{}) string {
					withdrawals := actual.([]models.Withdrawal)
					if len(withdrawals) == 2 &&
						withdrawals[0].Amount == 1 &&
						withdrawals[1].Amount == 2 {
						return ""
					}
					return fmt.Sprintf("Withdrawals %v is not expected", withdrawals)
				})
			})
		})
	})
}

func TestGetWithdrawalsUntil(t *testing.T) {
	Convey("Given mysql storage", t, func() {
		s := prepareDatabaseForTesting()
		s.db.MustExec("INSERT INTO `users` (email, bitcoin_address, balance) VALUES(?, ?, ?);", "e", "b", 8388607)
		s.CreateWithdrawal(models.Withdrawal{UserID: 1, BitcoinAddress: "b", Amount: 1})
		s.CreateWithdrawal(models.Withdrawal{UserID: 1, BitcoinAddress: "b", Amount: 2})
		s.CreateWithdrawal(models.Withdrawal{UserID: 1, BitcoinAddress: "b", Amount: 3})

		Convey("When get withdrawals until now", func() {
			result, _ := s.GetWithdrawalsUntil(1, time.Now().AddDate(0, 0, 1), 2)

			Convey("Withdrawals should equal", func() {
				So(result, func(actual interface{}, expected ...interface{}) string {
					withdrawals := actual.([]models.Withdrawal)
					if len(withdrawals) == 2 &&
						withdrawals[0].Amount == 3 &&
						withdrawals[1].Amount == 2 {
						return ""
					}
					return fmt.Sprintf("Withdrawals %v is not expected", withdrawals)
				})
			})
		})
	})
}
