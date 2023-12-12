package models

import "time"

type User struct {
    ID           int       `json:"id"`
    Name         string    `json:"name,omitempty"`
    Email        string    `json:"email,omitempty"`
    EmailVerified time.Time `json:"emailVerified,omitempty"`
    Image        string    `json:"image,omitempty"`
}