package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/PacktPublishing/Echo-Essentials/chapter7/bindings"
	"github.com/PacktPublishing/Echo-Essentials/chapter7/models"
	"github.com/PacktPublishing/Echo-Essentials/chapter7/renderings"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

// Login - Login Handler will take a username and password from the request
// hash the password, verify it matches in the database and respond with a token
func Login(c echo.Context) error {
	resp := renderings.LoginResponse{}
	lr := new(bindings.LoginRequest)

	if err := c.Bind(lr); err != nil {
		resp.Success = false
		resp.Message = "Unable to bind request for login"
		return c.JSON(http.StatusBadRequest, resp)
	}

	if err := c.Validate(lr); err != nil {
		resp.Success = false
		resp.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resp)
	} // get DB from context
	db := c.Get(models.DBContextKey).(*sql.DB)
	// get user by username from models
	user, err := models.GetUserByUsername(db, lr.Username)
	if err != nil {
		resp.Success = false
		resp.Message = "Username or Password incorrect"
		return c.JSON(http.StatusUnauthorized, resp)
	}

	if err := bcrypt.CompareHashAndPassword(
		user.PasswordHash, []byte(lr.Password)); err != nil {
		resp.Success = false
		resp.Message = "Username or Password incorrect"
		return c.JSON(http.StatusUnauthorized, resp)
	} // need to make a token, successful login
	signingKey := c.Get(models.SigningContextKey).([]byte)

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		Issuer:    "service",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	if err != nil {
		resp.Success = false
		resp.Message = "Server Error"
		return c.JSON(http.StatusInternalServerError, resp)
	}

	resp.Token = ss

	return c.JSON(http.StatusOK, resp)
}
