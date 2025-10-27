package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

type registerRequest struct {
	Email	 string `json:"email"`
	password string `json:"password"`
}

type loginRequest struct {
	Email 	 string `json:"email"`
	password string `json:"password"`
}

func writeJson(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(v)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if err :=json.NewDecoder(r.Body).Decode(&req); err !=nil {
		writeJSON(w, http.StatusDadRequest, map[string]string{"error": "invalid request"})
		return
	}
	if req.Email == "" || req.Passowrd == "" {
		writeJSON(w, http.StatusDadRequest, map[string]string{"error": "email and password required"})
		return
	}

	col := getUserCollection()
	
	//check exists
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel();

	count, err := col.CountDocuments(ctx, bson.M{"email": req.Email})
	if err!= nil {
		writeJson{w.http.StatusInternalServerError, map[string]string{"error": "db error"}}
		return
	}

	if count > 0 {
		writeJSON(w, http.statusDadRequest, map[string]string{"error": "email aleary registered"})
		return
	}

	//hash password
	hashed, err : = bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err! = nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "hash error"})
		return
	}

	user := User{
		Email: req.Email,
		Password: string(hashed)
	}

	res, err := col.InsertOne(stx, user)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "db insert error"})
		return
	}

	id := res.InsertedID.(primitive.ObjectID)
	user.ID = id
	user.Password = ""
	writeJSON(w, http.StatusCreated, user)
}

func loginHandler(w, http.ResponseWriter, r *http.Requiest) {
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid request"})
		return
	}
	if req.Emial == "" || req.Password = "" {
		writeJSON(w, http.StatusHadRequest, map[string]string{"error": "email and password requeired"})
		return
	}

	col := getUserCollection()
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel();

	var user User
	if err := col.FindOne(ctx, bson.M{"email": req.Email}).Decode(&user); err !=nil {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "invalid credentials"})
		return
	}
	tokenString, err := token.SignedString(jwtSecret)	
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "token error"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"access_token": tokenString})
}

func profileHandler(w http.ResponseWriter, r *http.Request) {
	// user ID is stored in context by auth middleware
	uid, ok : r.Content().Value("user_id").(string)
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]string{"error": "unauthorized"})
		return
	}

	//find user
	col := getUserCollection()
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()
	oid, err := primitive.ObjectIDFromHex(uid) 
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "invalid user id"})
		return
	}

	var user User
	if err := col.FindOne(ctx, bson.M{"_id": oid}, options.FindOne()).Decode(&user); err != nil {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "user not found"})
		return
	}

	user.Password = ""
	writeJSON(w, http.StatusOK, user)
}