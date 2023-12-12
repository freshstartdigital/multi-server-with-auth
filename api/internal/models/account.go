package models

type Account struct {
    ID                int       `json:"id"`
    UserID            int       `json:"userId"`
    Type              string    `json:"type"`
    Provider          string    `json:"provider"`
    ProviderAccountID string    `json:"providerAccountId"`
    RefreshToken      string    `json:"refreshToken,omitempty"`
    AccessToken       string    `json:"accessToken,omitempty"`
    ExpiresAt         int64     `json:"expiresAt,omitempty"`
    IDToken           string    `json:"idToken,omitempty"`
    Scope             string    `json:"scope,omitempty"`
    SessionState      string    `json:"sessionState,omitempty"`
    TokenType         string    `json:"tokenType,omitempty"`
}