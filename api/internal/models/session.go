package models

import "time"

type Session struct {
    ID          int       `json:"id"`
    UserID      int       `json:"userId"`
    Expires     time.Time `json:"expires"`
    SessionToken string    `json:"sessionToken"`
}