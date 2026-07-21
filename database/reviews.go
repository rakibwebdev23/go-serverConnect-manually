package database

type Review struct{
	ID          int     `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

var reviews []Review

// GET All Procuts database Listing
func GetReviewList() []Review {
	return reviews
}


func init (){
	review1 := Review{
		ID: 1,
		Name: "Rakib",
		Description: "Hello, I'm Rakib. Your Company product is very good.",
	}
	review2 := Review{
		ID: 2,
		Name: "Rakib Hasan",
		Description: "Hello, I'm Rakib Hasan. Your Company product is very good.",
	}
	review3 := Review{
		ID: 2,
		Name: "MD Rakib Hasan",
		Description: "Hello, I'm MD Rakib Hasan. Your Company product is very good.",
	}
	reviews = append(reviews, review1, review2, review3)
}