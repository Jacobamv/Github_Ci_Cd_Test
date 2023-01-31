package main
import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
	"encoding/json"
)
 
type Response struct{
	Message  string `json:"message"`
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "123",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
 
func main() {
      
    router := mux.NewRouter()
    router.HandleFunc("/test", TestHandler)
    http.Handle("/",router)
 
    fmt.Println("Server is listening...")
    http.ListenAndServe(":8181", nil)
}