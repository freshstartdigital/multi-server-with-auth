package repository

import (
	"database/sql"
	"log"

	"example.com/internal/models"
)


func CreateVerificationToken(db *sql.DB, verificationToken models.VerificationToken) (models.VerificationToken, error) {
	const sql = `
	INSERT INTO verification_token (identifier, expires, token) 
	VALUES ($1, $2, $3)
	`
	_, err := db.Exec(sql, verificationToken.Identifier, verificationToken.Expires, verificationToken.Token)
	if err != nil {
		return models.VerificationToken{}, err
	}
	return verificationToken, nil
}

func UseVerificationToken(db *sql.DB, identifier, token string) (models.VerificationToken, error) {
	const sql = `
	DELETE FROM verification_token
	WHERE identifier = $1 AND token = $2
	RETURNING identifier, expires, token
	`

	row := db.QueryRow(sql, identifier, token)
	var verificationToken models.VerificationToken
	err := row.Scan(&verificationToken.Identifier, &verificationToken.Expires, &verificationToken.Token)
	if err != nil {

		return models.VerificationToken{}, err
	}
	return verificationToken, nil
}

func CreateUser(db *sql.DB, user models.User) (models.User, error) {
	const sql = `
	INSERT INTO users (name, email, "emailVerified", image) 
	VALUES ($1, $2, $3, $4) 
	RETURNING id, name, email, "emailVerified", image
	`
	row := db.QueryRow(sql, user.Name, user.Email, user.EmailVerified, user.Image)
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.EmailVerified, &user.Image)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUser(db *sql.DB, id int) (*models.User, error) {
	const sql = `
	SELECT * FROM users WHERE id = $1
	`
	row := db.QueryRow(sql, id)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.EmailVerified, &user.Image)
	if err != nil {

		return nil, err
	}
	return &user, nil
}

func GetUserByEmail(db *sql.DB, email string) (*models.User, error) {
    const sqlQuery = `
    SELECT * FROM users WHERE email = $1
    `
    row := db.QueryRow(sqlQuery, email)

    var user models.User
    err := row.Scan(&user.ID, &user.Name, &user.Email, &user.EmailVerified, &user.Image)

    if err != nil {
        if err == sql.ErrNoRows {
            // No user was found, return nil user and no error
            return nil, nil
        }
        log.Println(err)
        return nil, err // An error occurred during the query execution
    }
    return &user, nil
}


func GetUserByAccount(db *sql.DB, providerAccountId, provider string) (*models.User, error) {
	const sql = `
	SELECT u.* FROM users u 
	JOIN accounts a ON u.id = a."userId"
	WHERE a.provider = $1 AND a."providerAccountId" = $2
	`
	row := db.QueryRow(sql, provider, providerAccountId)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.EmailVerified, &user.Image)
	if err != nil {

		return nil, err
	}
	return &user, nil
}

func UpdateUser(db *sql.DB, user models.User) (models.User, error) {
	const fetchSQL =  `select * from users where id = $1`
	const updateSQL = `
	UPDATE users set
	name = $2, email = $3, "emailVerified" = $4, image = $5
	where id = $1
	RETURNING name, id, email, "emailVerified", image
	`

	// Fetch the user from the database
	row := db.QueryRow(fetchSQL, user.ID)
	var existingUser models.User
	err := row.Scan(&existingUser.ID, &existingUser.Name, &existingUser.Email, &existingUser.EmailVerified, &existingUser.Image)
	if err != nil {
		return models.User{}, err
	}

	// Update the user
	row = db.QueryRow(updateSQL, user.ID, user.Name, user.Email, user.EmailVerified, user.Image)
	err = row.Scan(&user.Name, &user.ID, &user.Email, &user.EmailVerified, &user.Image)
	if err != nil {
		return models.User{}, err
	}
	return user, nil

}

func LinkAccount(db *sql.DB, account models.Account) (models.Account, error) {
	const sql = `
	insert into accounts 
	(
	  "userId", 
	  provider, 
	  type, 
	  "providerAccountId", 
	  access_token,
	  expires_at,
	  refresh_token,
	  id_token,
	  scope,
	  session_state,
	  token_type
	)
	values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
	returning
	  id,
	  "userId", 
	  provider, 
	  type, 
	  "providerAccountId", 
	  access_token,
	  expires_at,
	  refresh_token,
	  id_token,
	  scope,
	  session_state,
	  token_type
	`
	row := db.QueryRow(sql, account.UserID, account.Provider, account.Type, account.ProviderAccountID, account.AccessToken, account.ExpiresAt, account.RefreshToken, account.IDToken, account.Scope, account.SessionState, account.TokenType)
	err := row.Scan(&account.ID, &account.UserID, &account.Provider, &account.Type, &account.ProviderAccountID, &account.AccessToken, &account.ExpiresAt, &account.RefreshToken, &account.IDToken, &account.Scope, &account.SessionState, &account.TokenType)
	if err != nil {
		return models.Account{}, err
	}
	return account, nil

}

func CreateSession(db *sql.DB, session models.Session) (models.Session, error) {
	const sql = `insert into sessions ("userId", expires, "sessionToken")
	values ($1, $2, $3)
	RETURNING id, "sessionToken", "userId", expires`
	row := db.QueryRow(sql, session.UserID, session.Expires, session.SessionToken)
	err := row.Scan(&session.ID, &session.SessionToken, &session.UserID, &session.Expires)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil
}


func GetSessionAndUser(db *sql.DB, sessionToken string) (models.Session, models.User, error) {
	const session_sql = `select * from sessions where "sessionToken" = $1`
	const user_sql = `select * from users where id = $1`

	row := db.QueryRow(session_sql, sessionToken)
	var session models.Session
	err := row.Scan(&session.ID, &session.SessionToken, &session.UserID, &session.Expires)
	if err != nil {
		return models.Session{}, models.User{}, err
	}

	row = db.QueryRow(user_sql, session.UserID)
	var user models.User
	err = row.Scan(&user.ID, &user.Name, &user.Email, &user.EmailVerified, &user.Image)
	if err != nil {
		return models.Session{}, models.User{}, err
	}

	return session, user, nil
}


func UpdateSession(db *sql.DB, session models.Session) (models.Session, error) {
	const fetchSQL =  `select * from sessions where id = $1`
	const updateSQL = `
	UPDATE sessions set
	expires = $2
	where "sessionToken" = $1
	`

	// Fetch the session from the database
	row := db.QueryRow(fetchSQL, session.ID)
	var existingSession models.Session
	err := row.Scan(&existingSession.UserID, &existingSession.Expires, &existingSession.SessionToken)
	if err != nil {
		return models.Session{}, err
	}

	// Update the session
	row = db.QueryRow(updateSQL, session.SessionToken, session.Expires)
	err = row.Scan(&session.UserID, &session.Expires, &session.SessionToken)
	if err != nil {
		return models.Session{}, err
	}

	return session, nil

}

func DeleteSession(db *sql.DB, sessionToken string) error {
	const sql = `delete from sessions where "sessionToken" = $1`
	_, err := db.Exec(sql, sessionToken)
	if err != nil {
		return err
	}
	return nil
}

func UnlinkAccount(db *sql.DB, account models.Account) error {
	const sql = `delete from accounts where "providerAccountId" = $1 and provider = $2`
	_, err := db.Exec(sql, account.ProviderAccountID, account.Provider)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAccount(db *sql.DB, account models.Account) error {
	const user_sql = `delete from users where id = $1`
	const session_sql = `delete from sessions where "userId" = $1`
	const account_sql = `delete from accounts where "userId" = $1`

	_, err := db.Exec(user_sql, account.UserID)
	if err != nil {
		return err
	}

	_, err = db.Exec(session_sql, account.UserID)
	if err != nil {
		return err
	}

	_, err = db.Exec(account_sql, account.UserID)
	if err != nil {
		return err
	}

	return nil

}