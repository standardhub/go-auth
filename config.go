package main

import "time"

var (
	jwtSecret = []byte("replace-with-a-secure-secret")

	mongoURI      = "mongodb://localhost:27017"
	mongoDatabase = "goauth"

	tokenExpiry = time.Hour * 24
)