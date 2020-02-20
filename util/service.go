package util

type Service struct {
	names []string
}

func NewService() *Service {
	return &Service{
		names: []string{"Egor", "Roman", "Ivan", "Svetlana"},
	}
}
