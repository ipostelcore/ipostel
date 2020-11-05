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
	c := mapa["sqlserver"]
	//cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host
	//cadena := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s", c.Host, c.Usuario, c.Clave, c.Puerto, c.Basedatos)
	//cadena := "sqlserver://innova:innova2020@10.16.12.210:1433?database=PuntoPostal_06&connection+timeout=30"
	//cadena := "sqlserver://sa:operainfo123*@192.168.6.185:1433?database=PuntoPostal_06"
	//cadena := "odbc:server=10.16.12.210\\SQLExpress;user id=innova;password=innova2020;database=PuntoPostal_06;app name=innova"
	//cadena := "odbc:server=10.16.12.210;user id=innova;password=innova2020;database=PuntoPostal_06;app name=MyAppName"
	//cadena := "sqlserver://sa:operainfo123*@192.168.6.185/SQLExpress?database=PuntoPostal_06&connection+timeout=30"
	//cadena := "server=192.168.6.185;user id=sa;password=operainfo123*;"
	//fmt.Println(cadena)
	//sqlcmd -S 192.168.6.185 -U sa -P operainfo123* -Q "SELECT @@version"
	SqlServerPuntoPostal, err = sql.Open("odbc", "server=10.16.12.210;DSN=innova;Uid=innova;Pwd=innova2020")

	//SqlServerPuntoPostal, err = sql.Open("mssql", cadena)

	if err != nil {
		fmt.Println("[Punto Postal:   Error...] ", SqlServerPuntoPostal.Ping())
		fmt.Println(err.Error())
	} else {
		fmt.Println("[Punto Postal: ", c.Host, "  OK...]")
		sq, err := SqlServerPuntoPostal.Query("SELECT  codofic, descripcion FROM oficinas")
		util.Error(err)
		for sq.Next() {
			var a, b string
			sq.Scan(&a, &b)
			fmt.Println(a, b)
		}
		// ctx := context.Background()
		// err = SqlServerPuntoPostal.PingContext(ctx)
		// if err != nil {
		// 	fmt.Println("Control...", err.Error())
		// }
		// var result string

		// error := SqlServerPuntoPostal.QueryRowContext(ctx, "SELECT @@version").Scan(&result)
		// if error != nil {
		// 	fmt.Println("LLLL ", error.Error())
		// }
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
