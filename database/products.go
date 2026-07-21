package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"img_url"`
}

var productList []Product

// POST create product store to database
func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

// GET All Procuts database Listing
func GetList() []Product {
	return productList
}

// GET Single Product by id from database
func SingleGet(productID int) *Product {
	for _, product := range productList {
		if product.ID == productID {
			return &product
		}
	}
	return nil
}

// Update product from database
func Update(product Product) {
	for index, p := range productList {
		if p.ID == product.ID {
			productList[index] = product
		}
	}
}

// Delete product from database
func Delete(productID int) {
	var tempList []Product
	for _, p := range productList {
		if p.ID != productID {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func init() {
	product1 := Product{
		ID:          1,
		Title:       "Banana",
		Description: "A delicious yellow fruit",
		Price:       1.99,
		ImgUrl:      "https://example.com/banana.jpg",
	}
	product2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "A sweet red fruit",
		Price:       0.99,
		ImgUrl:      "https://example.com/apple.jpg",
	}
	product3 := Product{
		ID:          3,
		Title:       "Orange",
		Description: "A juicy citrus fruit",
		Price:       1.49,
		ImgUrl:      "https://example.com/orange.jpg",
	}
	product4 := Product{
		ID:          4,
		Title:       "Grapes",
		Description: "A bunch of small round fruits",
		Price:       2.99,
		ImgUrl:      "https://example.com/grapes.jpg",
	}
	product5 := Product{
		ID:          5,
		Title:       "Strawberry",
		Description: "A red heart-shaped fruit",
		Price:       3.49,
		ImgUrl:      "https://example.com/strawberry.jpg",
	}
	product6 := Product{
		ID:          6,
		Title:       "Pineapple",
		Description: "A tropical fruit with a spiky exterior",
		Price:       4.99,
		ImgUrl:      "https://example.com/pineapple.jpg",
	}
	product7 := Product{
		ID:          7,
		Title:       "Watermelon",
		Description: "A large green fruit with red flesh",
		Price:       5.99,
		ImgUrl:      "https://example.com/watermelon.jpg",
	}
	product8 := Product{
		ID:          8,
		Title:       "Mango",
		Description: "A sweet tropical fruit with orange flesh",
		Price:       2.49,
		ImgUrl:      "https://example.com/mango.jpg",
	}
	product9 := Product{
		ID:          9,
		Title:       "Blueberry",
		Description: "A small round blue fruit",
		Price:       3.99,
		ImgUrl:      "https://example.com/blueberry.jpg",
	}
	product10 := Product{
		ID:          10,
		Title:       "Kiwi",
		Description: "A small brown fruit with green flesh",
		Price:       1.49,
		ImgUrl:      "https://example.com/kiwi.jpg",
	}

	productList = append(productList, product1, product2, product3, product4, product5, product6, product7, product8, product9, product10)
}
