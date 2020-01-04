package usecase

// UserUsecase - UserUsecaseインターフェース
type UserUsecase interface {
	Create() error
}

type userUsecase struct{}

// NewUserUsecase - UserUsecaseの生成
func NewUserUsecase() UserUsecase {
	return &userUsecase{}
}

func (u *userUsecase) Create() error {
	// TODO: validation check

	// TODO: create user entity

	// TODO: insert to user in firebase and firestore

	return nil
}
