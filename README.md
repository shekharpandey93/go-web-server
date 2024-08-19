# go-web-server
Small init project to use Mux and Gin web service.
Mux web service run on :8080 port and Gin web service rin on :8081
1. Mux web service
2. Gin web service
   
# import package 
```
  GIN: go get github.com/gin-gonic/gin
  Mux: go get -u github.com/gorilla/mux
```

# sample code mux
```
router := mux.NewRouter()
router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	w.WriteHeader(http.StatusOK)
}).Methods("GET")
```
# sample code Gin
```
router := gin.Default()
router.GET("/:id", getUserHandler)
func getHandler(c *gin.Context) {
   //Params get
   id := c.Param("id")
   c.IndentedJSON(http.StatusOK, {})
}
```
