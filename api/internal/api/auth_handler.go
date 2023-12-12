package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example.com/internal/models"
	"example.com/internal/repository"
	"github.com/gorilla/mux"
)

type APIHandler struct {
    DB *sql.DB
}

// reference: https://github.com/nextauthjs/next-auth/blob/main/packages/adapter-pg/src/index.ts

func (api *APIHandler) CreateVerificationTokenHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var token models.VerificationToken
    err := json.NewDecoder(r.Body).Decode(&token)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call the repository function
    newToken, err := repository.CreateVerificationToken(api.DB, token)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the new token
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newToken)
}

func (api *APIHandler) UseVerificationTokenHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var token models.VerificationToken
    err := json.NewDecoder(r.Body).Decode(&token)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call the repository function
    usedToken, err := repository.UseVerificationToken(api.DB, token.Identifier, token.Token)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the used token
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(usedToken)
}

func (api *APIHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call the repository function
    newUser, err := repository.CreateUser(api.DB, user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the new user
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newUser)
}

func (api *APIHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
    // Get ID from query parameters
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        http.Error(w, "ID is required", http.StatusBadRequest)
        return
    }

    // Convert id to integer
    userID, err := strconv.Atoi(id)
    if err != nil {
        http.Error(w, "Invalid ID", http.StatusBadRequest)
        return
    }

    // Call the repository function
    user, err := repository.GetUser(api.DB, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
        return
    }
    if user == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Respond with the user data
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}



func (api *APIHandler) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    email := vars["email"]
    if email == "" {
        http.Error(w, "Email is required", http.StatusBadRequest)
        return
    }

    // Call the repository function
    user, err := repository.GetUserByEmail(api.DB, email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the user data
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (api *APIHandler) GetUserByAccountHandler(w http.ResponseWriter, r *http.Request) {
    // Get provider and providerAccountId from query parameters
    provider := r.URL.Query().Get("provider")
    providerAccountId := r.URL.Query().Get("providerAccountId")
    if provider == "" || providerAccountId == "" {
        http.Error(w, "Provider and providerAccountId are required", http.StatusBadRequest)
        return
    }

    // Call the repository function
    user, err := repository.GetUserByAccount(api.DB, providerAccountId, provider)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if user == nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Respond with the user data
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (api *APIHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        log.Println("error UpdateUserHandler", err)
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call the repository function
    updatedUser, err := repository.UpdateUser(api.DB, user)
    if err != nil {
        log.Println("error UpdateUserHandler 2", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the updated user
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedUser)
}

func (api *APIHandler) LinkAccountHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var account models.Account
    err := json.NewDecoder(r.Body).Decode(&account)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call the repository function
    newAccount, err := repository.LinkAccount(api.DB, account)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the new account
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newAccount)
}

func (api *APIHandler) CreateSessionHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var session models.Session
    err := json.NewDecoder(r.Body).Decode(&session)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call the repository function
    newSession, err := repository.CreateSession(api.DB, session)
    if err != nil {
        log.Println("error CreateSessionHandler", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the new session
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(newSession)
}



func (api *APIHandler) GetSessionAndUserHandler(w http.ResponseWriter, r *http.Request) {
    // Get sessionToken from query parameters
    vars := mux.Vars(r)
    sessionToken := vars["id"]

    if sessionToken == "" {

        http.Error(w, "Session token is required", http.StatusBadRequest)
        return
    }

    // Call the repository function
    session, user, err := repository.GetSessionAndUser(api.DB, sessionToken)
    if err != nil {
        log.Println("error GetSessionAndUserHandler", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the session and user data
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(struct {
        Session models.Session `json:"session"`
        User    models.User    `json:"user"`
    }{session, user})
        
}

func (api *APIHandler) UpdateSessionHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var session models.Session
    err := json.NewDecoder(r.Body).Decode(&session)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call the repository function
    updatedSession, err := repository.UpdateSession(api.DB, session)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the updated session
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(updatedSession)
}

func (api *APIHandler) UnlinkAccountHandler(w http.ResponseWriter, r *http.Request) {
    // Parse request body
    var account models.Account
    err := json.NewDecoder(r.Body).Decode(&account)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call the repository function
    err = repository.UnlinkAccount(api.DB, account)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK)
}

func (api *APIHandler) DeleteSessionHandler(w http.ResponseWriter, r *http.Request) {
    // Get sessionToken from query parameters
    vars := mux.Vars(r)
    sessionToken := vars["id"]

    if sessionToken == "" {
        http.Error(w, "Session token is required", http.StatusBadRequest)
        return
    }

    // Call the repository function
    err := repository.DeleteSession(api.DB, sessionToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.WriteHeader(http.StatusOK)
}
