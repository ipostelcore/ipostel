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

//Config Generacion de conexiones
type Config struct{}

//DriverSQL Establecer Driver's de conexion
type DriverSQL struct {
	Nombre   string
	Contexto context.Context
	DB       *sql.DB
	Estatus  bool
	Error    error
}

//Variables del modelo
var (
	Version                  string = "V.2.1.2"
	MySQL                    bool   = false
	SQLServerPuntoPostal     bool   = false
	SQLServerTracking        bool   = false
	SQLServerMaestros        bool   = false
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
	SQLTODO                  = make(map[string]DriverSQL)
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
	ID          string
	Driver      string
	Usuario     string
	Basedatos   string
	Clave       string
	Host        string
	Puerto      string
	SUrl        string
	Descripcion string
	Estatus     bool
}

//Conexiones 0: PostgreSQL, 1: MySQL, 2: MongoDB
var Conexiones []CadenaDeConexion

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
			ID:          valor.ID,
			Driver:      valor.Driver,
			Usuario:     valor.Usuario,
			Basedatos:   valor.Basedatos,
			Clave:       valor.Clave,
			Host:        valor.Host,
			Puerto:      valor.Puerto,
			Descripcion: valor.Descripcion,
			Estatus:     valor.Estatus,
		}
		fmt.Println("")
		fmt.Println("Conectando: ", valor.Descripcion)
		switch valor.Driver {
		case "mongodb":
			MongoDBConexion(cad)
		}
	}
	util.Error(e)
}

//ConexionesDinamicas Permite establecer multiples conexiones
func (C *Config) ConexionesDinamicas(c CadenaDeConexion) bool {

	fmt.Println("Conectando: ", c.Descripcion)
	switch c.Driver {
	case "sqlserver17":
		db, er := CSQLServer(c)
		SQLTODO[c.ID] = DriverSQL{
			Nombre:   c.Driver,
			DB:       db,
			Estatus:  true,
			Contexto: Contexto,
			Error:    er,
		}
	case "postgres13":
		db, er := CPostgres(c)
		SQLTODO[c.ID] = DriverSQL{
			Nombre:   c.Driver,
			DB:       db,
			Estatus:  true,
			Contexto: Contexto,
			Error:    er,
		}

	case "mysql":
		MySQL = true
	default:
		fmt.Println("Driver: no funciona para ", c.Driver)
	}
	return true
}
