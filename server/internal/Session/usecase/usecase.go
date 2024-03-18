package usecase

import (
	"math/rand"
	sessionRep "server/internal/Session/repository"
	userRep "server/internal/User/repository"
	"server/internal/domain/entity"
	"time"
)

const sessKeyLen = 10

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

type SessionUsecaseI interface {
	Login(user *entity.User) (*entity.Cookie, error)
	Check(SessionToken string) (uint, error)
	Logout(cookie *entity.Cookie) error
	GetUserProfile(sessionToken string) (*entity.User, error)
}

type SessionUsecase struct {
	sessionRepo sessionRep.SessionRepositoryI
	userRepo    userRep.UserRepositoryI
}

func NewSessionUsecase(sessionRep sessionRep.SessionRepositoryI, userRep userRep.UserRepositoryI) *SessionUsecase {
	return &SessionUsecase{
		sessionRepo: sessionRep,
		userRepo:    userRep,
	}
}

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (ss SessionUsecase) Login(user *entity.User) (*entity.Cookie, error) {
	us, err := ss.userRepo.FindUserByUsername(user.Name)

	if err != nil {
		return nil, err
	}

	if us == nil {
		return nil, entity.ErrBadRequest
	}

	if user.Password != us.Password {
		return nil, entity.ErrBadRequest
	}

	cookie := &entity.Cookie{
		UserID:       us.ID,
		SessionToken: randStringRunes(sessKeyLen),
		MaxAge:       150 * time.Hour,
	}

	err = ss.sessionRepo.Create(cookie)
	if err != nil {
		return nil, err
	}

	return cookie, nil

}

func (ss SessionUsecase) Check(SessionToken string) (uint, error) {

	cookie, err := ss.sessionRepo.Check(SessionToken)
	if err != nil {
		return 0, err
	}
	if cookie == nil {
		return 0, nil
	}
	user, err := ss.userRepo.FindUserByID(cookie.UserID)
	if err != nil {
		return 0, err
	}

	if user == nil {
		return 0, nil
	}
	return user.ID, nil
}

func (ss SessionUsecase) Logout(cookie *entity.Cookie) error {
	return ss.sessionRepo.Delete(entity.ToDBDeleteCookie(cookie))
}

func (ss SessionUsecase) GetUserProfile(sessionToken string) (*entity.User, error) {
	cookie, err := ss.sessionRepo.Check(sessionToken)
	if err != nil {
		return nil, err
	}

	user, err := ss.userRepo.FindUserByID(cookie.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
