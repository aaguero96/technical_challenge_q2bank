package walletService

type WalletService interface {
	GetAll() ([]WalletResponse, error)
	GetById(id int) (GetByIdResponse, error)
	AddAmount(walletID int, increaseAmount float64) (AddAmountResponse, error)
}
