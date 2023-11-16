package rest

import (
	"frame/model"
	"time"
)

type UserVo struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Username  string    `json:"username"`
}

func NewUserVo(v *model.User) *UserVo {
	return &UserVo{
		ID:        v.ID,
		CreatedAt: v.CreatedAt,
		UpdatedAt: v.UpdatedAt,
		Username:  v.Username,
	}
}
