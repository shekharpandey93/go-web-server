// api/api.go
package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Initialize and return a new router
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//// Define API routes
	//router.HandleFunc("/hello", service.HelloHandler).Methods("GET")
	//router.HandleFunc("/add", service.AddHandler).Methods("POST")

	return router
}

// Start the server
func StartServerMux() {
	router := NewRouter()

	router.HandleFunc("/login", loginHandler).Methods("POST")
	router.HandleFunc("/signup", signupHandler).Methods("POST")
	router.HandleFunc("/logout", logoutHandler).Methods("POST")

	// Start the server on port 8080
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func StartServerGin() {
	router := gin.Default()
	router.GET("/users", getUserHandler)
	router.GET("/user/:id", getUserHandler)
	router.GET("/user/add", addUser)
	router.Run("localhost:8081")
}
