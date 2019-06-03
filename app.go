package main
import (
	"log"
	"fmt"
	"net/http"
	"github.com/gorilla/handlers"
	"Food/router"
	"Food/config"
)
func main(){
var c = config.Configs()
port := c.Port
fmt.Println("Starting Server at ", port)
log.Fatal(http.ListenAndServe(port, handlers.CORS()(router.FoodRouter())))
}


