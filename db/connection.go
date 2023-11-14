package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// EL DSN deberia ser brindado por variables de entorno pero para facilidad de desarrollo y despliegue se guarda dentro del codigo.
var DSN = "host=localhost user=postgres password=postgres dbname= TesisDB port=5432"
var DB *gorm.DB

func DBConnection() {
	var err error
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DBConnection succeeded")
	}
}
