package appErrors

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Res struct {
	Message string `json:"message"`
	Code    int    `json:"status"`
}

var (
	ErrInvalidToken     = errors.New("Token is invalid. please log in system again")
	ErrUserExists       = errors.New("User with this phone exists")
	ErrMethodNotAllowed = errors.New("Method not allowed")
	ErrUserDoesntExist  = errors.New("User doensn't exists")
	ErrWrongPassword    = errors.New("Wrong password")
	ErrInvalidData      = errors.New("Invalid data input")
)

func HandleErr(ctx *gin.Context) {
	ctx.Next()
	for _, err := range ctx.Errors {
		var res Res
		switch err.Err {
		case ErrInvalidToken:
			res = Res{ErrInvalidToken.Error(), http.StatusBadRequest}
		case ErrUserExists:
			res = Res{ErrUserExists.Error(), http.StatusConflict}
		case ErrMethodNotAllowed:
			res = Res{ErrMethodNotAllowed.Error(), http.StatusMethodNotAllowed}
		case ErrWrongPassword:
			res = Res{ErrWrongPassword.Error(), http.StatusBadRequest}
		case ErrUserDoesntExist:
			res = Res{ErrUserDoesntExist.Error(), http.StatusBadRequest}
		case ErrInvalidData:
			res = Res{ErrInvalidData.Error(), http.StatusBadRequest}
		default:
			res = Res{"", http.StatusInternalServerError}
		}
		log.Printf("Registred error: %d|%s", res.Code, err.Error())
		ctx.AbortWithStatusJSON(res.Code, res.Message)
	}
}
