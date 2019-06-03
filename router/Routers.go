package router
import (
	"github.com/gorilla/mux"
	"net/http"
	"Food/controller"
	"fmt"
)
func FoodRouter() http.Handler{
	router := mux.NewRouter()
	fmt.Println("ran")
	router.HandleFunc("/api/home/unit/insert", controller.InsertUnits).Methods("POST")
	router.HandleFunc("/api/home/unit/get/{id}", controller.GetUnits).Methods("GET")
	router.HandleFunc("/api/home/unit/getAllUnit", controller.GetAllUnits).Methods("GET")
	router.HandleFunc("/api/home/unit/update/{id}", controller.UpdateUnits).Methods("PUT")
	router.HandleFunc("/api/home/unit/delete/{id}", controller.DeleteUnits).Methods("DELETE")
	

	router.HandleFunc("/api/home/category/insert", controller.CreateCategories).Methods("POST")
	router.HandleFunc("/api/home/category/getSub/{id}", controller.GetSubCategories).Methods("GET")
	router.HandleFunc("/api/home/category/getAllUnit", controller.GetAllUnits).Methods("GET")
	router.HandleFunc("/api/home/category/update/{id}", controller.UpdateCategories).Methods("PUT")
	router.HandleFunc("/api/home/category/delete/{id}", controller.DeleteCategories).Methods("DELETE")
	return router
}