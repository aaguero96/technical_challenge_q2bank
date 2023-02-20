package walletService

type WalletService interface {
	GetAll() ([]WalletResponse, error)
}
