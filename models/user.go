package models

import "time"

type User struct {
	CreatedAt                    time.Time   `json:"created_at"`
	Email                        string      `json:"email"`
	EmailVerifiedAt              time.Time   `json:"email_verified_at"`
	EmergencyContact             interface{} `json:"emergency_contact"`
	EmergencyContactRelationship interface{} `json:"emergency_contact_relationship"`
	ID                           int         `json:"id"`
	IsDeleted                    int         `json:"is_deleted"`
	Mobile                       interface{} `json:"mobile"`
	Name                         string      `json:"name"`
	PhotoURL                     string      `json:"photo_url"`
	ReferralCode                 string      `json:"referral_code"`
	ReferralLink                 string      `json:"referral_link"`
	Role                         string      `json:"role"`
	UpdatedAt                    time.Time   `json:"updated_at"`
}
