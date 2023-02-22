package registerTypeService

type RegisterTypeService interface {
	GetAll() ([]RegisterTypeResponse, error)
}
