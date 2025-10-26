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

