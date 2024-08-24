package apis

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"init-project/Modal"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	val := r.Body
	var user Modal.User
	err := json.NewDecoder(val).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	validate := validator.New()

	// Validate the User struct
	err = validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Fprintf(w, "Validation failed on field '%s', condition: %s\n", err.Field(), err.Tag())
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {

}

func signupHandler(w http.ResponseWriter, r *http.Request) {

}

func getUserHandler(c *gin.Context) {
	id := c.Param("id")
	if id != "" {
		c.IndentedJSON(http.StatusOK, Modal.User{Id: 1, Username: "Shekhar", Email: "sp@gmail.com", Phone: "8265362777"})
		return
	} else {
		users := []Modal.User{}
		users = append(users, Modal.User{Id: 1, Username: "Shekhar", Email: "sp@gmail.com", Phone: "8265362777"})
		users = append(users, Modal.User{Id: 3, Username: "Rahul", Email: "rahul@gmail.com", Phone: "9165362777"})
		users = append(users, Modal.User{Id: 3, Username: "Ravi", Email: "ravi@gmail.com", Phone: "9265362777"})
		c.IndentedJSON(http.StatusOK, &users)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}

func addUser(c *gin.Context) {
	var user Modal.User
	if err := c.BindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Invalid JSON"})
	}
}
