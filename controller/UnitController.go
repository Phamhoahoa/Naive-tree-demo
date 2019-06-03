package controller
import (
"net/http"
"Food/model"
"fmt"
"encoding/json"
"Food/dao"
"github.com/gorilla/mux"
"strconv"
)
// func setupResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func InsertUnits(w http.ResponseWriter, r *http.Request){
	var unit model.Unit
	var res model.Respone 
	decoder := json.NewDecoder(r.Body)
	err:= decoder.Decode(&unit) 
	if err != nil {
		fmt.Println("Error 1: inser_decode")
		return 
	}
	defer r.Body.Close()
	
	err = dao.InsertUnit(&unit)
	fmt.Println(err)
		if err != nil {
			fmt.Println("Error 2 : insert_dao")
			res.Err = "Insert Fail!"
			json.NewEncoder(w).Encode(res) 
			return
		}
		fmt.Println("insert_successful!")
		res.Success = "Insert Success"
		json.NewEncoder(w).Encode(res)

}
func GetUnits(w http.ResponseWriter, r *http.Request){
	params :=mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	unit := dao.GetUnit(id)
	json.NewEncoder(w).Encode(unit)

}
func GetAllUnits(w http.ResponseWriter, r *http.Request){
	var units[] model.Unit
	units = dao.GetAllUnit()
	json.NewEncoder(w).Encode(units)
}
func UpdateUnits(w http.ResponseWriter, r *http.Request){
	var unit model.Unit
	var res model.Respone
	vars := mux.Vars(r)
	id,_:=strconv.Atoi(vars["id"])
	decoder := json.NewDecoder(r.Body)
	unit.Id = id 
	err := decoder.Decode(&unit)
	if err != nil {
		return
	}
	defer r.Body.Close()
	err = dao.UpdateUnit(id, &unit)
	if err != nil {
		fmt.Println("update_err")
		res.Err = "Update Fail!"
		json.NewEncoder(w).Encode(res) 
		return
	}
	fmt.Println("update_success")
	res.Success = "update Success!"
	json.NewEncoder(w).Encode(res)
}
func DeleteUnits(w http.ResponseWriter, r * http.Request){
	var res model.Respone
	params := mux.Vars(r)
	id,_ := strconv.Atoi(params["id"])
	err := dao.DeleteUnit(id)
	if err != nil{
		fmt.Println(err)
		fmt.Println("Error 1 : delete_dao")
		res.Err = "Delete Fail!"
		json.NewEncoder(w).Encode(res) 
		return
	}
	fmt.Println("Deleted")
	res.Success = "Deleted!"
	json.NewEncoder(w).Encode(res) 


}