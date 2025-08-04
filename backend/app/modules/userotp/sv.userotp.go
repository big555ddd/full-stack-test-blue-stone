package userotp

import (
	"app/app/model"
	userotpdto "app/app/modules/userotp/dto"
	"context"
	"errors"
	"time"

	"github.com/uptrace/bun"
)

type Service struct {
	db *bun.DB
}

func NewService(db *bun.DB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Create(ctx context.Context, req userotpdto.CreateUserOtp) (*model.UserOtp, error) {

	m := &model.UserOtp{
		UserID:    req.UserID,
		Otp:       req.Otp,
		ExpiresAt: time.Now().Unix() + 900,
	}

	_, err := s.db.NewInsert().Model(m).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return m, err

}

func (s *Service) Get(ctx context.Context, id string) (*model.UserOtp, error) {
	m := &model.UserOtp{}
	err := s.db.NewSelect().
		Model(m).
		Where("id = ?", id).Scan(ctx)
	return m, err
}

func (s *Service) UpdateUsed(ctx context.Context, id string) error {
	ex, err := s.db.NewSelect().Model(&model.UserOtp{}).Where("id = ?", id).Exists(ctx)
	if err != nil {
		return err
	}

	if !ex {
		return errors.New("user OTP not found")
	}

	_, err = s.db.NewUpdate().Model(&model.UserOtp{}).Set("used = true").Where("id = ?", id).Exec(ctx)
	return err
}

func (s *Service) Delete(ctx context.Context, id string) error {
	ex, err := s.db.NewSelect().Model(&model.UserOtp{}).Where("id = ?", id).Exists(ctx)
	if err != nil {
		return err
	}

	if !ex {
		return errors.New("user not found")
	}

	_, err = s.db.NewDelete().Model(&model.UserOtp{}).Where("id = ?", id).Exec(ctx)
	return err

}
