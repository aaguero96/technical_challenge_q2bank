package registerTypeService

type RegisterTypeService interface {
	GetAll() ([]RegiterTypeResponse, error)
}
