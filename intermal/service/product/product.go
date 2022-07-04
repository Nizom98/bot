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

func List() []*Product {
	return allProducts
}
