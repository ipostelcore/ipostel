//configuraciones del sistema
package sys

import (
	"context"
	"database/sql"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//MongoDBConexion Conexion a Mongo DB
func MongoDBConexion(mapa map[string]CadenaDeConexion) {
	Contexto, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	c := mapa["mongodb"]
	defer cancel()
	MongoDB, _ = MGConexion.Conectar(Contexto, c)
}

//CMongoDB Conexion a Mongo DB
func CMongoDB(c CadenaDeConexion) (DB *mongo.Database, err error) {
	Contexto, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	DB, err = MGConexion.ConectarDinamicamente(Contexto, c)
	return
}

//CPostgres Funcion para conectarnos al driver de postgresSQL
func CPostgres(c CadenaDeConexion) (DB *sql.DB, err error) {
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	DB, err = sql.Open("postgres", cadena)
	if err != nil {
		color.Red("... Host: "+c.Host+" Base De Datos: ( "+c.Basedatos+" ) Error... ", err.Error())
	} else {
		err = DB.Ping()
		if err != nil {
			color.Red("... Host: "+c.Host+" Base De Datos: ( "+c.Basedatos+" ) Error...", err.Error())
			color.Red(err.Error())
			return
		}

		color.Green("... Host: " + c.Host + " Base De Datos: ( " + c.Basedatos + " )  OK... ")
	}
	return
}

//CSQLServer Funcion de Conexion a SQL SERVER
func CSQLServer(c CadenaDeConexion) (DB *sql.DB, err error) {
	DB, err = sql.Open("odbc", "server="+c.Host+";database="+c.Basedatos+";DSN=tracking;Uid="+c.Usuario+";Pwd="+c.Clave)
	if err != nil {
		color.Red("... Host: "+c.Host+" Base De Datos: ( "+c.Basedatos+" ) Error...", err.Error())
	} else {
		err = DB.Ping()
		if err != nil {
			color.Red("... Host: "+c.Host+" Base De Datos: ( "+c.Basedatos+" ) Error...", err.Error())
			color.Red(err.Error())
			return
		}
		color.Green("... Host: " + c.Host + " Base De Datos: ( " + c.Basedatos + " )  OK...")
	}
	return
}
