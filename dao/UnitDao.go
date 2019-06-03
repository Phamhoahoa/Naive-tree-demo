package dao
import (
"Food/model"
"fmt"
"context"
)
func InsertUnit(unit *model.Unit) error{
	Db := OpenDbConnection()
	defer Db.Close()
	statement := fmt.Sprintf("INSERT INTO units(code,name) VALUES('%s', '%s')", unit.Code,unit.Name)
	fmt.Println(statement)
	_, err := Db.Exec(statement)
	return err 
}
func GetUnit(id int) model.Unit{
	Db := OpenDbConnection()
	defer Db.Close()
	var unit model.Unit
	ctx := context.Background()
	statement := fmt.Sprintf("select * from units where id = %d", id)
	_ = Db.QueryRowContext(ctx,statement).Scan(&unit.Id,&unit.Name,&unit.Code)
	return unit
}
func GetAllUnit()[]model.Unit{
	Db:= OpenDbConnection()
	defer Db.Close()
	var units[] model.Unit
	statement := fmt.Sprintf("select * from units")
	rows,_ := Db.Query(statement)
	for rows.Next(){
		var unit model.Unit
		_ = rows.Scan(&unit.Id,&unit.Code, &unit.Name)
		units = append(units, unit)
	}
	return units 
}
func UpdateUnit(id int, unit *model.Unit) error{
	Db := OpenDbConnection()
	statement := fmt.Sprintf("update units set code='%s', name='%s' where id=%d ",unit.Code ,unit.Name , id)
	_, err:=Db.Exec(statement)
	return err
}
func DeleteUnit(id int)error{
	Db:=OpenDbConnection()
	statement := fmt.Sprintf("delete from units where id=%d",id)
	_, err :=Db.Exec(statement) 
	return err 
}
