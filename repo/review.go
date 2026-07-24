package repo

type Review struct{
	ID          int     `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type ReviewRepo interface{
	List() ([]*Review, error)
}

// property
type reviewRepo struct{
	reviewList []*Review
}

//constractor or Constractor function 
func NewReviewRepo() ReviewRepo{
	repo := &reviewRepo{}

	generateInitialReview(repo)
	return repo
}

func (r *reviewRepo) List()([]*Review, error){
	return r.reviewList, nil
}

func generateInitialReview (r *reviewRepo){
	rev1 := &Review{
		ID: 1,
		Name: "Rakib",
		Description: "Hello, I'm Rakib. Your Company product is very good.",
	}
	rev2 := &Review{
		ID: 2,
		Name: "Rakib Hasan",
		Description: "Hello, I'm Rakib Hasan. Your Company product is very good.",
	}
	rev3 := &Review{
		ID: 2,
		Name: "MD Rakib Hasan",
		Description: "Hello, I'm MD Rakib Hasan. Your Company product is very good.",
	}
	r.reviewList = append(r.reviewList, rev1, rev2, rev3)
}