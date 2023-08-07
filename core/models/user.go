package models

import (
	"github.com/google/uuid"
	"github.com/pepusz/go_redirect/gateways/jwt"
)

type User struct {
	AuthID   string `json:"auth_id"`
	UserID   *int   `json:"user_id"`
	UserType string `json:"user_type"`
}

func (User) IsEntity() {}

func GetCoreUserFromToken(serviceJWTToken string) (*User, error) {
	jwtGW := jwt.Gateway{}
	valid, claims, err := jwtGW.ValidateToken(serviceJWTToken)
	if err != nil || !valid {
		return nil, err
	}
	return &User{
		AuthID:   claims["authID"].(string),
		UserID:   getIntValue(claims, "userID"),
		UserType: claims["userType"].(string),
	}, nil
}

func getIntValue(claims map[string]interface{}, name string) *int {
	if claims[name] == nil {
		return nil
	}
	val := int(claims[name].(float64))
	return &val
}
func getUUIDValue(claims map[string]interface{}, name string) *uuid.UUID {
	if claims[name] == nil {
		return nil
	}
	uuidVal, err := uuid.Parse(claims[name].(string))
	if err != nil || uuidVal == uuid.Nil {
		return nil
	}
	return &uuidVal
}
