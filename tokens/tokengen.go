package tokens

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/Mdromi/go-ecom-yt.git/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// SignedDetails represents the claims of a JWT token
type SignedDetails struct {
	Email      string `json:"email"`
	First_Name string `json:"first_name"`
	Last_Name  string `json:"last_name"`
	Uid        string `json:"uid"`
	jwt.StandardClaims
}

var UserData *mongo.Collection = database.UserData(database.Client, "Users")

var err = godotenv.Load()
var SECRET_KEY = os.Getenv("SECRECT_KEY")

// TokenGenerator generates a signed token and a signed refresh token
func TokenGenerator(email string, firstName string, lastName string, uid string) (signedToken string, signedRefreshToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_Name: firstName,
		Last_Name:  lastName,
		Uid:        uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", "", err
	}
	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panicln(err)
		return
	}
	return token, refreshtoken, err
}

// ValidateToken validates a signed token and returns the claims
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "the token is already expired"
		return
	}

	return claims, msg
}

// UpdateAllTokens updates the signed token and signed refresh token for a user
func UpdateAllTokens(signedToken string, signedRefreshToken string, userID string) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

	var updateObj primitive.D

	updateObj = append(updateObj, bson.E{Key: "token", Value: signedToken})
	updateObj = append(updateObj, bson.E{Key: "refresh_token", Value: signedRefreshToken})

	updatedAt, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	updateObj = append(updateObj, bson.E{Key: "updated_at", Value: updatedAt})

	upsert := true
	filter := bson.M{"user_id": userID}
	opt := options.UpdateOptions{
		Upsert: &upsert,
	}

	_, err := UserData.UpdateOne(ctx, filter, bson.D{
		{Key: "$set", Value: updateObj},
	}, &opt)

	defer cancel()
	if err != nil {
		log.Panic(err)
		return
	}
}
