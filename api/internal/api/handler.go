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
