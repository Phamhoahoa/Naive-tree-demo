package dao
import (
"Food/config"
"database/sql"
_"github.com/go-sql-driver/mysql"
"fmt"
"log"

)
var c = config.Configs()
var Connect = fmt.Sprintf("%s:%s%s/%s", c.UserName,c.Password,c.DbAddress, c.DbName)
func OpenDbConnection() *sql.DB{
db, err:=sql.Open("mysql", Connect)
if err != nil{
	log.Fatal(err)
}
return db 
}