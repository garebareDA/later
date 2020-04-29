package database
import(
	"github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=name dbname=later password=pass")
  return db, err
}