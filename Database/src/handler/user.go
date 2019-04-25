package handler

import (
	"../model"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.User{ID: bson.NewObjectId()}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	if u.FirstName == "" || u.LastName == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "No first name or last name"}
	}

	// Save user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("cw_user").C("users").Insert(u); err != nil {
		return
	}

	//Json
	first_name_s := "\"first_name\""
	last_name_s := "\"last_name\""
	user_id_s := "\"user_id\""

	responsemessage := join("{ ", first_name_s, " : ", u.FirstName, ", ", last_name_s, " : ", u.LastName, "}")

	return c.JSON(http.StatusCreated, responsemessage)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.User)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("cw_user").C("users").
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return c.JSON(http.StatusCreated, u)
	}

	//-----
	// JWT
	//-----

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response
	u.Token, err = token.SignedString([]byte(Key))
	if err != nil {
		return err
	}

	u.Password = "" // Don't send password

	//Json
	is_aut_s := "\"is_auth\""
	token_s := "\"token\""
	first_name_s := "\"first_name\""
	last_name_s := "\"last_name\""
	user_id_s := "\"user_id\""

	responsemessage := join("{ ", is_aut_s," : true, ", token_s, " : ", u.Token,
		", ", first_name_s, " : ", u.FirstName, ", ", last_name_s, " : ", u.LastName, "}")
	fmt.Println(responsemessage)
	return c.JSON(http.StatusOK, responsemessage)
}

func userIDFromToken(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims["id"].(string)
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
