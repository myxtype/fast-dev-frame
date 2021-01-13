package rest

import (
	"frame/models"
	"time"
)

type UserVo struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
}

func NewUserVo(v *models.User) *UserVo {
	return &UserVo{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Username:  v.Username,
	}
}
