package users

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/services"
	"bookstore_users-api/utils/errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func logError(err error) {
	fmt.Println("Error: ", err)
}

func CreateUser(c *gin.Context) {
	var user users.User
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	//fmt.Println(user)
	//if err != nil {
	//	logError(err)
	//	return
	//}
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	logError(err)
	//	return
	//}

	if err := c.ShouldBindJSON(&user); err != nil {
		logError(err)
		restErr := errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   "bad_request",
		}
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		logError(saveErr)
	}
	fmt.Println(result)
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me plz")
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me plz")
}
