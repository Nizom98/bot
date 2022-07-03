package product

var allProducts = []*Product {
	 {
		Title: "one",
	 },
	 {
		Title: "two",
	 },
	 {
		Title: "three",
	 },
	 {
		Title: "four",
	 },
}

type Service struct {

}

func NewService() *Service {
	return &Service{}
}

func (srv *Service) List() []*Product {
	return allProducts
}