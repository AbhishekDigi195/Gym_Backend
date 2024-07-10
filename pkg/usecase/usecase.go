package usecase

import (
	"gym/pkg/models"
	"time"

	"gorm.io/gorm"
)

type Usecase struct {
	db *gorm.DB
}

func NewUseCase(db *gorm.DB) *Usecase {
	return &Usecase{db}
}

func (uu *Usecase) RegisterUser(user models.User) error {
	if err := uu.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (uu *Usecase) GetUserDetailsById(userid string) (*models.User, error) {
	var user models.User
	err := uu.db.Where("id=?", userid).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uu *Usecase) PurchaseMembership(input models.Membership) error {
	if err := uu.db.Create(&input).Error; err != nil {
		return err
	}
	return nil
}

func (uu *Usecase) UpdateMembershipStatus(memid string, status string) error {
	tx := uu.db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	var membership models.Membership
	if err := tx.First(&membership,"id= ?", memid).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&membership).Update("status", status).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (uu *Usecase) RenewMembership(memid string, enddate time.Time) error {
	tx := uu.db.Begin()

	if tx.Error != nil {
		return tx.Error
	}

	var membership models.Membership
	if err := tx.First(&membership,"id= ?", memid).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Model(&membership).Update("end_date", enddate).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (uu *Usecase) VerifyUser(userid string) (bool, error) {
	var user models.User
	result := uu.db.First(&user, "id = ?", userid)
	if result.Error != nil {
		return false, result.Error
	}

	return user.Verified, nil
}

func (uu *Usecase) SetUserVerified(email string) error {
	result := uu.db.Model(&models.User{}).Where("email = ?", email).Update("verified", true)
	return result.Error
}

func (uu *Usecase) GetMembershipStatus(memberid string) (string,error) {
	var membership models.Membership
	result := uu.db.First(&membership, "id = ?", memberid)
	if result.Error != nil {
		return "", result.Error
	}

	return membership.Status, nil
}

func (uu *Usecase) Login(email string) (*models.Admins,error) {
	var user models.Admins
	err := uu.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil 
}

func (uu *Usecase) RegisterAdmin(admin models.Admins) error {
	if err := uu.db.Create(&admin).Error; err != nil {
		return err
	}
	return nil
}