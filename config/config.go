package config
import(
	"encoding/json"
	"os"
	"fmt"
)
type Config struct{
	UserName string
	Password string
	Port string 
	DbName string
	DbAddress string 
}
func Configs() Config{
	file,_ :=os.Open("app.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil{
		fmt.Println("error: ", err)
	}
	return configuration
}