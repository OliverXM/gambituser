package bd

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/models"
	"github.com/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println(" > Inicio de Registro del Usuario en la DB")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("SignUp > Usuario Registrado en la DB EXITOSAMENTE")
	return nil
}
