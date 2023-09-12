package models

import (
	"authService/db"
	"authService/lib"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Password string             `bson:"password,omitempty"`
}

type LoggedInUser struct {
	Id       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Token 	 string             `json:"token"`
}

type AuthorizedUser struct {
	Id       string `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
}

var usersCollection *mongo.Collection

func init() {
	db := db.InitDB()
	if users, err := db.UsersCollection(); err != nil {
		panic("fail to get 'users' collection")
	} else {
		usersCollection = users
	}
}

func MakeUser(username string, email string, password string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if existingUser, _ := GetUserByEmail(email); existingUser != nil {
		return nil, errors.New("mail is already used")
	}

	hashedPsw, err := lib.Hash(password)
	if err != nil {
		return nil, errors.New("fail to hash password")
	}
	user := User{
		Username: username,
		Email:    email,
		Password: hashedPsw,
	}

	if insert, err := usersCollection.InsertOne(ctx, user); err != nil {
		return nil, err
	} else {
		id, ok := insert.InsertedID.(primitive.ObjectID)

		if !ok {
			return nil, errors.New("fail to get insertedID")
		}
		user.Id = id
		return &user, nil
	}
}

func GetUserByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := User{}
	if err := usersCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func ToLogedInUser(user *User, token string) *LoggedInUser {
	return &LoggedInUser{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
		Token:    token,
	}
}

func ToAuthorizedUser(user *User) *AuthorizedUser {
	return &AuthorizedUser{
		Id:       user.Id.Hex(),
		Username: user.Username,
		Email:    user.Email,
	}
}

func ClaimToAuthorizedUser(claim lib.Claims) AuthorizedUser {
	return AuthorizedUser{
		Id:       claim.UserID,
		Username: claim.Username,
		Email:    claim.Email,
	}
}