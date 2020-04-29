package database
import(
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

var(
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
)


func ConnectDB() (*gorm.DB, error) {
	connectTemplate := "host=localhost port=5432 user=%s dbname=later password=%s"
	connect := fmt.Sprintf(connectTemplate, DBUser, DBPass)
	db, err := gorm.Open("postgres", connect)
  return db, err
}