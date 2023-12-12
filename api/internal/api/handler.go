package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

    // Database connection setup
    connStr := "host=db port=5432 user=your_user password=your_password dbname=your_db_name sslmode=disable"

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal(err)
    }

    apiHandler := &APIHandler{DB: db}

	// Register the API handlers

    //hello world
    router.HandleFunc(("/"), func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World!"))
    }).Methods("GET")

    router.HandleFunc(("/init"), func(w http.ResponseWriter, r *http.Request) {
        db.Exec(`
CREATE TABLE verification_token
(
  identifier TEXT NOT NULL,
  expires TIMESTAMPTZ NOT NULL,
  token TEXT NOT NULL,

  PRIMARY KEY (identifier, token)
);

CREATE TABLE accounts
(
  id SERIAL,
  "userId" INTEGER NOT NULL,
  type VARCHAR(255) NOT NULL,
  provider VARCHAR(255) NOT NULL,
  "providerAccountId" VARCHAR(255) NOT NULL,
  refresh_token TEXT,
  access_token TEXT,
  expires_at BIGINT,
  id_token TEXT,
  scope TEXT,
  session_state TEXT,
  token_type TEXT,

  PRIMARY KEY (id)
);

CREATE TABLE sessions
(
  id SERIAL,
  "userId" INTEGER NOT NULL,
  expires TIMESTAMPTZ NOT NULL,
  "sessionToken" VARCHAR(255) NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE users
(
  id SERIAL,
  name VARCHAR(255),
  email VARCHAR(255),
  "emailVerified" TIMESTAMPTZ,
  image TEXT,

  PRIMARY KEY (id)
);
        `)

        w.Write([]byte("Init"))
    }).Methods("GET")



    // Auth handlers
    router.HandleFunc(("/auth/verification-token"), apiHandler.CreateVerificationTokenHandler).Methods("POST")
    router.HandleFunc(("/auth/verification-token"), apiHandler.UseVerificationTokenHandler).Methods("PUT")
    router.HandleFunc(("/auth/user"), apiHandler.CreateUserHandler).Methods("POST")
    router.HandleFunc(("/auth/user/{id}"), apiHandler.GetUserHandler).Methods("GET")
    router.HandleFunc(("/auth/user/email/{email}"), apiHandler.GetUserByEmailHandler).Methods("GET")
    router.HandleFunc(("/auth/user/account/{id}"), apiHandler.GetUserByAccountHandler).Methods("GET")
    router.HandleFunc(("/auth/user"), apiHandler.UpdateUserHandler).Methods("PUT")
    router.HandleFunc(("/auth/user/link"), apiHandler.LinkAccountHandler).Methods("POST")
    router.HandleFunc(("/auth/session"), apiHandler.CreateSessionHandler).Methods("POST")
    router.HandleFunc(("/auth/session/{id}"), apiHandler.GetSessionAndUserHandler).Methods("GET")
    router.HandleFunc(("/auth/session"), apiHandler.UpdateSessionHandler).Methods("PUT")
    router.HandleFunc(("/auth/user/link"), apiHandler.UnlinkAccountHandler).Methods("PUT")
    router.HandleFunc(("/auth/session/{id}"), apiHandler.DeleteSessionHandler).Methods("DELETE")




	return router
}
