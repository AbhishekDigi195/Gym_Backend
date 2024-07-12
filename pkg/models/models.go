package models

import (
	"math/big"
	"time"
)

type Admins struct {
	Id int `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}

type User struct {
	ID         string 
	Name       string
	Email      string `gorm:"primaryKey"`
	Contact    string
	IDProof    string
	Verified   bool
	Membership *Membership `gorm:"foreignKey:Email"`
}

type Membership struct {
	ID        string `gorm:"primaryKey"`
	Email    string `gorm:"index"`
	Type      string
	StartDate time.Time
	EndDate   time.Time
	Status    string
}

type UserRegistered struct {
	Name    string
	Email   string
	Contact string
}

type MembershipPurchasedEvent struct {
	Email          string
	MembershipId   *big.Int
	MembershipType string
	StartDate      *big.Int
	EndDate        *big.Int
}

type MembershipRenewedEvent struct {
	Email        string
	MembershipId *big.Int
	NewEndDate   *big.Int
}
