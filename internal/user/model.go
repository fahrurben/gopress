package user

import "time"

const (
	TYPE_USER = iota
	TYPE_ADMIN
)

type User struct {
	Id        int        `db:"id" json:"id"`
	Email     string     `db:"email" json:"email"`
	Name      string     `db:"name" json:"name"`
	Password  string     `db:"password" json:"-"`
	Type      int        `db:"type" json:"type"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at" json:"deleted_at"`
}
