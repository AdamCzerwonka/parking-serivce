package entities

import "time"

type EmailToken struct {
	Token     string    `db:"token"`
	UserId    int       `db:"user_id"`
	ValidFrom time.Time `db:"valid_from"`
	ValidTo   time.Time `db:"valid_to"`
}
