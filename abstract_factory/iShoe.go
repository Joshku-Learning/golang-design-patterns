package abstract_factory

// iShoe 抽象鞋子產品
type IShoe interface {
	SetLogo(logo string)
	GetLogo() string
	SetSize(size int)
	GetSize() int
}

type shoe struct {
	logo string
	size int
}

// setLogo func
func (s *shoe) SetLogo(logo string) {
	s.logo = logo
}

// getLogo func
func (s *shoe) GetLogo() string {
	return s.logo
}

// setSize func
func (s *shoe) SetSize(size int) {
	s.size = size
}

// getSize func
func (s *shoe) GetSize() int {
	return s.size
}
