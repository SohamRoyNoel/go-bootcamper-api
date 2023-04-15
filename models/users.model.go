package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID            primitive.ObjectID `bson:"_id"`
	FIRST_NAME    *string            `json:"firstName" validate:"min=2,max=100"`
	LAST_NAME     *string            `json:"lastName" validate:"min=2,max=100"`
	EMAIL         *string            `json:"email" validate:"required,max=100"`
	PASSWORD      *[]byte            `json:"password" validate:"required,min=8,max=20"`
	PHONE         *string            `json:"phone" validate:"min=10,max=10"`
	ACCESS_TOKEN  *[]byte            `json:"accessToken"`
	REFRESH_TOKEN *[]byte            `json:"refreshToken"`
	CREATED_AT    time.Time          `json:"createdAt"`
	UPDATED_AT    time.Time          `json:"updatedAt"`
	ADDRESS       *string            `json:"address"`
	ROLE *string `json:"role"`
}
