package sys

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	util "github.com/ipostelcore/ipostel/util"
	"go.mongodb.org/mongo-driver/mongo"
)

type config struct{}

//Variables del modelo
var (
	Version                  string = "V.2.1.2"
	MySQL                    bool   = false
	SQLServer                bool   = false
	Oracle                   bool   = false
	BaseDeDatos              BaseDatos
	MGConexion               *Mongo
	Contexto                 context.Context
	MongoDB                  *mongo.Database
	PostgreSQLSAMAN          *sql.DB
	PuntoPostalPostgres      *sql.DB
	PostgreSQLPENSION        *sql.DB
	PostgreSQLPENSIONSIGESP  *sql.DB
	PostgreSQLEMPLEADOSIGESP *sql.DB
	MysqlFullText            *sql.DB
	SqlServerPuntoPostal     *sql.DB
	SqlServerTracking        *sql.DB
	SqlServerMaestros        *sql.DB
	Error                    error
	HostIPPace               string = ""
	HostUrlPace              string = ""
	HostIPPension            string = ""
	HostUrlPension           string = ""
	ListadoConexiones        []string
)

//Constantes del sistema
const (
	ACTIVAR_CONEXION_REMOTA       bool   = true
	DESACTIVAR_CONEXION_REMOTA    bool   = false
	ACTIVAR_LOG_REGISTRO          bool   = true
	DESACTIVAR_LOG_REGISTRO       bool   = false
	ACTIVAR_ROLES                 bool   = true
	DESACTIVAR_ROLES              bool   = false
	ACTIVAR_LIMITE_DE_CONSULTA    bool   = true
	DESACTIVAR_LIMITE_DE_CONSULTA bool   = false
	PUERTO                        string = "8012"
	PUERTO_SSL                    string = "2286"
	CODIFCACION_DE_ARCHIVOS       string = "UTF-8"
	MAXIMO_LIMITE_DE_USUARIO      int    = 100
	MAXIMO_LIMITE_DE_CONSULTAS    int    = 10
)

//BaseDatos Estructuras
type BaseDatos struct {
	CadenaDeConexion map[string]CadenaDeConexion
}

//CadenaDeConexion Conexion de datos
type CadenaDeConexion struct {
	Driver      string
	Usuario     string
	Basedatos   string
	Clave       string
	Host        string
	Puerto      string
	StrUrl      string
	Descripcion string
}

//Conexiones 0: PostgreSQL, 1: MySQL, 2: MongoDB
var Conexiones []CadenaDeConexion

//init Inicio y control
func init() {

	Magenta := color.New(color.FgMagenta)
	BoldMagenta := Magenta.Add(color.Bold)

	fmt.Println("")
	BoldMagenta.Println("..........................................................")
	BoldMagenta.Println("...                                                       ")
	BoldMagenta.Println("...           Versión del Panel ", Version, "             ")
	BoldMagenta.Println("...      Iniciando Carga de Conexiones a Drivers          ")
	BoldMagenta.Println("..........................................................")
	BoldMagenta.Println("")
	var a util.Archivo
	a.NombreDelArchivo = "sys/config_dev.json"
	data, _ := a.LeerTodo()
	e := json.Unmarshal(data, &Conexiones)
	for _, valor := range Conexiones {
		ListadoConexiones = append(ListadoConexiones, valor.Driver)
		cad := make(map[string]CadenaDeConexion)
		cad[valor.Driver] = CadenaDeConexion{
			Driver:      valor.Driver,
			Usuario:     valor.Usuario,
			Basedatos:   valor.Basedatos,
			Clave:       valor.Clave,
			Host:        valor.Host,
			Puerto:      valor.Puerto,
			Descripcion: valor.Descripcion,
		}
		fmt.Println("")
		fmt.Println("Conectando: ", valor.Descripcion)
		switch valor.Driver {
		case "mongodb":
			MongoDBConexion(cad)
		case "puntopostal":
			ConexionPuntoPostal(cad)
		case "tracking":
			ConexionTracking(cad)
		case "maestros":
			ConexionMaestros(cad)
		case "postgres":
			ConexionPuntoPostalPostgres(cad)
		case "mysql":
			MySQL = true
			ConexionMYSQL(cad)
		}
	}
	util.Error(e)
}
