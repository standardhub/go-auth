package main

import "time"

var (
	jwtSecretKey = []byte("replace-with-a-secrure-secret")

	mongoURI      = "mongodb://localhost:27017"
	mongoDatabase = "goauth"

	tokenExpiry = time.Hour * 24
)