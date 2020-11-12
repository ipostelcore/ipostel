//configuraciones del sistema
package sys

import (
	"database/sql"
	"fmt"

	mgo "gopkg.in/mgo.v2"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	"github.com/ipostelcore/ipostel/util"
	_ "github.com/lib/pq"
)

//MongoDBConexion Conexion a Mongo DB
func MongoDBConexion(mapa map[string]CadenaDeConexion) {
	c := mapa["mongodb"]
	MGOSession, Error = mgo.Dial(c.Host + ":27017")
	fmt.Println("Cargando Conexión Con MongoDB...")
	util.Error(Error)
}

//ConexionPuntoPostal Funcion de Conexion a Postgres
func ConexionPuntoPostal(mapa map[string]CadenaDeConexion) (err error) {
	c := mapa["puntopostal"]
	SqlServerPuntoPostal, err = sql.Open("odbc", "server=192.168.6.185;DSN=ipostel;Uid=ipostel;Pwd=Za63qj2p##")
	if err != nil {
		fmt.Println("[Punto Postal:   Error...] ", SqlServerPuntoPostal.Ping())
		fmt.Println(err.Error())
	} else {
		fmt.Println("[Punto Postal: ", c.Host, "  OK...]")
		sq, err := SqlServerPuntoPostal.Query("SELECT TOP 2 codofic, descripcion FROM oficinas")
		util.Error(err)
		for sq.Next() {
			var a, b string
			sq.Scan(&a, &b)
			fmt.Println(a, b)
		}
		fmt.Println("Controlando la situación")
	}
	return
}

//ConexionPuntoPostalPostgres Funcion de Conexion a Postgres
func ConexionPuntoPostalPostgres(mapa map[string]CadenaDeConexion) {
	c := mapa["postgres"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PuntoPostalPostgres, _ = sql.Open("postgres", cadena)
	if PuntoPostalPostgres.Ping() != nil {
		fmt.Println("[Punto Postal Postgres:   Error...] ", PuntoPostalPostgres.Ping())
	} else {
		fmt.Println("[Punto Postal Postgres: ", c.Host, "  OK...]")
	}

}

//ConexionPuntoPostalIpostel Funcion de Conexion a Postgres
func ConexionPuntoPostalIpostel(mapa map[string]CadenaDeConexion) {
	c := mapa["ipostel"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PuntoPostalIpostel, _ = sql.Open("postgres", cadena)
	if PuntoPostalIpostel.Ping() != nil {
		fmt.Println("[Punto Postal Ipostel:   Error...] ", PuntoPostalIpostel.Ping())
	} else {
		fmt.Println("[Punto Postal Ipostel: ", c.Host, "  OK...]")
	}

}

//ConexionTARJETA
func ConexionPENSIONSIGESP(mapa map[string]CadenaDeConexion) {
	c := mapa["pensiones"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLPENSIONSIGESP, _ = sql.Open("postgres", cadena)
	if PostgreSQLPENSIONSIGESP.Ping() != nil {
		fmt.Println("[Pensiones SIGESP: Error...] ", PostgreSQLPENSIONSIGESP.Ping())
	} else {
		fmt.Println("[Pensiones SIGESP: OK...]")
	}
	return
}

//ConexionMYSQL
func ConexionMYSQL(mapa map[string]CadenaDeConexion) {
	c := mapa["mysql"]
	cadena := c.Usuario + ":" + c.Clave + "@tcp(" + c.Host + ":3306)/sssifanb"
	MysqlFullText, _ = sql.Open("mysql", cadena)
	if MysqlFullText.Ping() != nil {
		fmt.Println("[mysql FULLTEXT: Error...] ", MysqlFullText.Ping())
	} else {
		fmt.Println("[mysql FULLTEXT: OK...]")
	}
	return
}
