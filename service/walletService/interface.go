package walletService

type WalletService interface {
	GetAll() ([]WalletResponse, error)
	GetById(id int) (GetByIdResponse, error)
}
