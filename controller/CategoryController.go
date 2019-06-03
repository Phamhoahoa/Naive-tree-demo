package controller
import (
	"net/http"
	"github.com/gorilla/mux"
	"Food/model"
	"strconv"
	"fmt"
	"Food/dao"
	"encoding/json"
	"io/ioutil"
	"log"
	"bytes"
)

func CreateCategories(w http.ResponseWriter, r *http.Request){
	var category model.Category
	var res model.Respone
	err := json.NewDecoder(r.Body).Decode(&category)

	buf, _:= ioutil.ReadAll(r.Body)
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	log.Printf("BODY: %q", rdr1)

	fmt.Println(category)
	defer r.Body.Close()
	err = dao.CreateCategory(&category)
	if err != nil{
		fmt.Println(err)
		res.Err = "Create Fail!"
		json.NewEncoder(w).Encode(res) 
		return
	}
	fmt.Println("Insert successfull!")
	res.Success = "Insert successfull"
	json.NewEncoder(w).Encode(res)

}
func UpdateCategories(w http.ResponseWriter, r *http.Request){
	var res model.Respone
	var category model.Category
	vars := mux.Vars(r)
	id, _:= strconv.Atoi(vars["id"])
	decoder := json.NewDecoder(r.Body)
	category.Id = id
	err := decoder.Decode(&category)
	err = dao.UpdateCategory(id, &category)

	defer r.Body.Close()
	if err != nil{
		fmt.Println(err)
		res.Err = "update Fail!"
		json.NewEncoder(w).Encode(res) 
		return 
	}
	fmt.Println("Update success!")
	res.Success = "Update success"
	json.NewEncoder(w).Encode(res)

}
func DeleteCategories(w http.ResponseWriter, r *http.Request){
	var res model.Respone
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	err := dao.DeleteCategory(id)
	if err != nil{
		fmt.Println(err)
		res.Err = "Delete Fail!"
		json.NewEncoder(w).Encode(res) 
		return
	}
	fmt.Println("Category is deleted!")
	res.Success  = "Category is deleted"
	json.NewEncoder(w).Encode(res)

}
func GetSubCategories(w http.ResponseWriter, r *http.Request){
	var subCategory[] model.Category
	params := mux.Vars(r)
	var res model.Respone
	id,_ := strconv.Atoi(params["id"])
	subCategory = dao.GetSubCategory(id)
	i := len(subCategory)
	fmt.Println(i)
	if i != 0 {
	json.NewEncoder(w).Encode(subCategory)
	}else{
		res.Err = "No SubCategory!"
		json.NewEncoder(w).Encode(res) 
	}
}


