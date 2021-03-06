package v1

import (
	"time"

	"github.com/gorilla/websocket"

	"github.com/solefaucet/sole-server/models"
)

// dependencies
type (
	// user
	dependencyGetUserByID         func(int64) (models.User, error)
	dependencyGetUserByEmail      func(string) (models.User, error)
	dependencyCreateUser          func(models.User) error
	dependencyUpdateUserStatus    func(int64, string) error
	dependencyGetReferees         func(userID int64, limit, offset int64) ([]models.User, error)
	dependencyGetNumberOfReferees func(userID int64) (int64, error)

	// auth token
	dependencyCreateAuthToken func(models.AuthToken) error
	dependencyDeleteAuthToken func(string) error

	// session
	dependencyUpsertSession     func(models.Session) error
	dependencyGetSessionByToken func(string) (models.Session, error)

	// email
	dependencySendEmail func(recipients []string, subject string, html string) error

	// total reward
	dependencyGetLatestTotalReward func() models.TotalReward
	dependencyIncrementTotalReward func(time.Time, int64) error

	// reward rate
	dependencyGetRewardRatesByType func(string) []models.RewardRate

	// system config
	dependencyGetSystemConfig func() models.Config

	// income
	dependencyCreateRewardIncome          func(models.Income, time.Time) error
	dependencyCreateSuperrewardsIncome    func(income models.Income, transactionID, offerID string) error
	dependencyCreateKiwiwallIncome        func(income models.Income, transactionID, offerID string) error
	dependencyCreateAdscendMediaIncome    func(income models.Income, transactionID, offerID string) error
	dependencyCreateAdgateMediaIncome     func(income models.Income, transactionID, offerID string) error
	dependencyCreateOffertoroIncome       func(income models.Income, transactionID, offerID string) error
	dependencyCreatePersonalyIncome       func(income models.Income, offerID string) error
	dependencyCreateClixwallIncome        func(income models.Income, offerID string) error
	dependencyCreatePtcwallIncome         func(income models.Income) error
	dependencyGetRewardIncomes            func(userID int64, limit, offset int64) ([]models.Income, error)
	dependencyGetNumberOfRewardIncomes    func(userID int64) (int64, error)
	dependencyGetOfferwallIncomes         func(userID int64, limit, offset int64) ([]models.Income, error)
	dependencyGetNumberOfOfferwallIncomes func(userID int64) (int64, error)
	dependencyInsertIncome                func(interface{}) // cache for broadcasting
	dependencyChargebackIncome            func(incomeID int64) error

	// websocket
	dependencyPutConn          func(*websocket.Conn)
	dependencyBroadcast        func([]byte)
	dependencyGetUsersOnline   func() int
	dependencyGetLatestIncomes func() []interface{}

	// withdrawals
	dependencyGetWithdrawals         func(userID int64, limit, offset int64) ([]models.Withdrawal, error)
	dependencyGetNumberOfWithdrawals func(userID int64) (int64, error)
	dependencyConstructTxURL         func(tx string) string

	// validation
	dependencyValidateAddress func(string) (bool, error)

	// captcha
	dependencyRegisterCaptcha func() (string, error)
	dependencyGetCaptchaID    func() string

	// superrewards
	dependencyGetNumberOfSuperrewardsOffers func(transactionID string, userID int64) (int64, error)

	// clixwall
	dependencyGetNumberOfClixwallOffers func(offerID string, userID int64) (int64, error)

	// personaly
	dependencyGetNumberOfPersonalyOffers func(offerID string, userID int64) (int64, error)

	// kiwiwall
	dependencyGetNumberOfKiwiwallOffers func(transactionID string, userID int64) (int64, error)

	// adscend media
	dependencyGetAdscendMediaOffer func(transactionID string, userID int64) (*models.AdscendMedia, error)

	// adgate media
	dependencyGetNumberOfAdgateMediaOffers func(transactionID string, userID int64) (int64, error)

	// offertoro
	dependencyGetNumberOfOffertoroOffers func(transactionID string, userID int64) (int64, error)
)
