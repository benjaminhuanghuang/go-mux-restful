package repository

type reso struct{}

func NewFirestorePostRepository() PostRepository {
	return &repo{}
}
