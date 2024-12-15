package dao

import "time"

type BaseDao struct {
	CreatedAt time.Time `json:"created_at"`
}
