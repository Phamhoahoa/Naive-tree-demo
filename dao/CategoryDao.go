package dao 
import (
"Food/model"
"fmt"
"context"
)
func CreateCategory(category *model.Category) error{
	db := OpenDbConnection()
	statement := fmt.Sprintf("INSERT INTO categories(`name`, `code`, `parent`, `show`) VALUES('%s','%s',%d,%t)", category.Name, category.Code, category.Parent, category.Show)
	fmt.Println(statement)
	_, err := db.Exec(statement)
	query := fmt.Sprintf("SELECT LAST_INSERT_ID();")
	ctx := context.Background()
	_ = db.QueryRowContext(ctx, query).Scan(&category.Id)
	fmt.Println(category.Id)
	statement1:= fmt.Sprintf("INSERT INTO trees(parent, child, depth) VALUES (%d, %d, 0)", category.Id, category.Id)
	fmt.Println(statement1)
	statement2 := fmt.Sprintf("INSERT INTO trees(parent, child, depth) SELECT parent, %d, depth+1 FROM trees WHERE child = %d",category.Id, category.Parent)
	_, err = db.Exec(statement1)
	_, err = db.Exec(statement2)
	if err != nil{
		fmt.Println(err)
	}
	db.Close()
	return err

}
func UpdateCategory(id int, category *model.Category) error{
	db := OpenDbConnection()
	statement := fmt.Sprintf("UPDATE categories set `code` ='%s',`name` = '%s', `show`=%t where id = %d", category.Code,category.Name, category.Show, category.Id)
	_, err := db.Exec(statement)
  	if err != nil {
  	 	fmt.Println(err)
   	}
   	db.Close()
	return err
}


func DeleteCategory(id int) error{
	Db:=OpenDbConnection()
	statement := fmt.Sprintf("SELECT child FROM trees WHERE parent = %d;", id)
	rows, _ := Db.Query(statement)
	var categoryIds []int
	for rows.Next(){
		var category model.Category
		_ = rows.Scan(&category.Id)
		categoryIds = append(categoryIds,category.Id)
	}
	fmt.Println(categoryIds)
	statement1:=fmt.Sprintf("DELETE p FROM `trees` p JOIN `trees` a USING (`child`) WHERE a.`parent` = %d;", id)
	_, err :=Db.Exec(statement1)
	fmt.Println("Err:", err)
	fmt.Println(statement1)
	statement3 := fmt.Sprintf("SET FOREIGN_KEY_CHECKS=0;")
	_, err = Db.Exec(statement3)
	for i:=0;i<len(categoryIds);i++{
		
		statement2:=fmt.Sprintf("DELETE FROM categories WHERE id = %d",categoryIds[i])
		_, err = Db.Exec(statement2)
	}
	statement4 := fmt.Sprintf("SET FOREIGN_KEY_CHECKS=1;")
	_, err = Db.Exec(statement4)
	return err
}

func GetSubCategory(id int) []model.Category{
	Db := OpenDbConnection()
	var subCategories[] model.Category
	statement := fmt.Sprintf("SELECT c.* FROM categories c JOIN trees t  ON (c.id = t.child) WHERE t.parent = %d  AND t.depth = 1;", id)
	rows, _ := Db.Query(statement)
	for rows.Next(){
		var category model.Category
		_ = rows.Scan(&category.Id, &category.Name, &category.Code, &category.Parent, &category.Show)
		subCategories = append(subCategories, category)
	}
	return subCategories
}