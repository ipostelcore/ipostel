//configuraciones del sistema
package sys

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

//MongoDBConexion Conexion a Mongo DB
func MongoDBConexion(mapa map[string]CadenaDeConexion) {
	Contexto, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	MongoDB = MGConexion.Conectar(Contexto, mapa)
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

//ConexionPuntoPostal Funcion de Conexion a SQL SERVER
func ConexionPuntoPostal(mapa map[string]CadenaDeConexion) (err error) {
	c := mapa["puntopostal"]

	SqlServerPuntoPostal, err = sql.Open("odbc", "server="+c.Host+";database="+c.Basedatos+";DSN=ipostel;Uid="+c.Usuario+";Pwd="+c.Clave)
	if err != nil {
		fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + "  Error...] ")
	} else {
		//"SELECT TOP 2 codofic, descripcion FROM oficinas"
		_, err := SqlServerPuntoPostal.Query("SELECT TOP 2 CLIENTE, DESCRIPCION FROM CLIENTES")

		if err != nil {
			fmt.Println("[Host: "+c.Host+" Base De Datos: "+c.Basedatos+"  Error...] ", err.Error())
			return err
		} else {
			fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + " OK...]")
			// for sq.Next() {
			// 	var a string
			// 	errx := sq.Scan(&a)
			// 	if errx != nil {
			// 		fmt.Println("Error cargando fila ", errx.Error())
			// 	}
			// 	fmt.Println(a)
			// }
			// fmt.Println("Controlando la situación")
		}

	}
	return
}

//ConexionTracking Funcion de Conexion a SQL SERVER
func ConexionTracking(mapa map[string]CadenaDeConexion) (err error) {
	c := mapa["tracking"]

	SqlServerTracking, err = sql.Open("odbc", "server="+c.Host+";database="+c.Basedatos+";DSN=tracking;Uid="+c.Usuario+";Pwd="+c.Clave)
	if err != nil {
		fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + "  Error...] ")
	} else {
		//"SELECT TOP 2 codofic, descripcion FROM oficinas"
		_, err := SqlServerTracking.Query("SELECT TOP 1 of_Id, of_Codigo FROM MA_OFICINA")

		if err != nil {
			fmt.Println("[Host: "+c.Host+" Base De Datos: "+c.Basedatos+"  Error...] ", err.Error())
			return err
		} else {
			fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + " OK...]")
			// for sq.Next() {
			// 	var a, b string
			// 	sq.Scan(&a, &b)
			// 	fmt.Println(a, b)
			// }
			// fmt.Println("Controlando la situación")
		}

	}
	return
}

//ConexionMaestros Funcion de Conexion a SQL SERVER
func ConexionMaestros(mapa map[string]CadenaDeConexion) (err error) {
	c := mapa["maestros"]

	SqlServerMaestros, err = sql.Open("odbc", "server="+c.Host+";database="+c.Basedatos+";DSN=maestros;Uid="+c.Usuario+";Pwd="+c.Clave)
	if err != nil {
		fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + "  Error...] ")
	} else {
		_, err := SqlServerMaestros.Query("SELECT TOP 1 * FROM MA_Servicio")

		if err != nil {
			fmt.Println("[Host: "+c.Host+" Base De Datos: "+c.Basedatos+"  Error...] ", err.Error())
			return err
		} else {
			fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + " OK...]")

		}
	}
	return
}

//ConexionPuntoPostalPostgres Funcion de Conexion a Postgres
func ConexionPuntoPostalPostgres(mapa map[string]CadenaDeConexion) (err error) {
	c := mapa["postgres"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PuntoPostalPostgres, err = sql.Open("postgres", cadena)
	if err != nil {
		fmt.Println("[Host: "+c.Host+" Base De Datos: "+c.Basedatos+" Error...] ", err.Error())
	} else {
		fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + " OK...]")
	}
	return
}

//ConexionPENSIONSIGESP Control de datos
func ConexionPENSIONSIGESP(mapa map[string]CadenaDeConexion) (err error) {
	c := mapa["pensiones"]
	cadena := "user=" + c.Usuario + " dbname=" + c.Basedatos + " password=" + c.Clave + " host=" + c.Host + " sslmode=disable"
	PostgreSQLPENSIONSIGESP, _ = sql.Open("postgres", cadena)
	if err != nil {
		fmt.Println("[Host: "+c.Host+" Base De Datos: "+c.Basedatos+" Error...] ", err.Error())
	} else {
		fmt.Println("[Host: " + c.Host + " Base De Datos: " + c.Basedatos + " OK...]")
	}
	return
}

//ConexionMYSQL Control de datos MyISAM
func ConexionMYSQL(mapa map[string]CadenaDeConexion) (err error) {
	c := mapa["mysql"]
	cadena := c.Usuario + ":" + c.Clave + "@tcp(" + c.Host + ":3306)/sssifanb"
	MysqlFullText, err = sql.Open("mysql", cadena)
	if MysqlFullText.Ping() != nil {
		fmt.Println("[MySQL Host: "+c.Host+" Base De Datos: "+c.Basedatos+"  Error...] ", err.Error())
	} else {
		fmt.Println("[MySQL Host: " + c.Host + " Base De Datos: " + c.Basedatos + "  OK...]")
	}
	return
}
