package login

// IService ...
type IService interface {
}

// Service ...
type Service struct{}

// NewService ...
func NewService() IService {
	return &Service{}
}
