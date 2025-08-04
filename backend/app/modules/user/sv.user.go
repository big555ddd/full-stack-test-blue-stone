package user

import (
	"app/app/message"
	"app/app/model"
	userdto "app/app/modules/user/dto"
	"context"
	"errors"
	"strings"

	"github.com/uptrace/bun"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	db *bun.DB
}

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Create(ctx context.Context, req userdto.CreateUser) (*model.User, bool, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), 14)
	if err != nil {
		return nil, false, err
	}
	m := &model.User{
		UserName: req.Username,
		Email:    req.Email,
		Password: string(bytes),
	}

	_, err = s.db.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, true, errors.New(message.EmailAlreadyExists)
		}
	}

	return m, false, err

}

func (s *Service) UpdatePassword(ctx context.Context, id string, newPassword string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	if err != nil {
		return err
	}

	_, err = s.db.NewUpdate().Model(&model.User{}).Set("password = ?", string(bytes)).Where("id = ?", id).Exec(ctx)
	return err
}

func (s *Service) Get(ctx context.Context, id userdto.GetByIDUser) (*userdto.UserResponse, error) {
	m := userdto.UserResponse{}
	err := s.db.NewSelect().
		TableExpr("users AS u").
		Column("u.id", "u.username", "u.email", "u.created_at", "u.updated_at").
		Where("id = ?", id.ID).Where("deleted_at IS NULL").Scan(ctx, &m)
	return &m, err
}

func (s *Service) Delete(ctx context.Context, id userdto.GetByIDUser) error {
	ex, err := s.db.NewSelect().Table("users").Where("id = ?", id.ID).Where("deleted_at IS NULL").Exists(ctx)
	if err != nil {
		return err
	}

	if !ex {
		return errors.New("user not found")
	}

	_, err = s.db.NewDelete().Model((*model.User)(nil)).Where("id = ?", id.ID).Exec(ctx)
	return err

}

func (s *Service) ExistEmail(ctx context.Context, email string) (bool, error) {
	ex, err := s.db.NewSelect().Model(&model.User{}).Where("email = ?", email).Exists(ctx)
	return ex, err
}

func (s *Service) ExistUserName(ctx context.Context, username string) (bool, error) {
	ex, err := s.db.NewSelect().Model(&model.User{}).Where("username = ?", username).Exists(ctx)
	return ex, err
}

func (s *Service) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	m := model.User{}
	err := s.db.NewSelect().Model(&m).Where("email = ?", email).Scan(ctx)
	if err != nil {
		return nil, err
	}
	if m.ID == "" {
		return nil, errors.New(message.UserNotFound)
	}
	return &m, nil
}

func (s *Service) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	m := model.User{}
	err := s.db.NewSelect().Model(&m).Where("username = ?", username).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
