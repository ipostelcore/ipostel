package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/ipostelcore/ipostel/sys"
	"github.com/ipostelcore/ipostel/util"
	"gopkg.in/mgo.v2/bson"
)

//Core Ejecucion
type Core struct {
	Sistema  string
	Consulta string
	Ruta     string
	Autor    string
}

//Oficina Describe una oficinas
type Oficina struct {
	ID            string `json:"id"`
	CodigoOficina string `json:"codigo"`
	Descripcion   string `json:"descripcion"`
}

//ApiCore Estructura de conexion
type ApiCore struct {
	Modulo       string `json:"modula"`
	Driver       string `json:"driver"`
	Query        string `json:"query"`
	Campos       string `json:"campos"`
	Parametros   string `json:"parametros"`
	Ruta         string `json:"ruta"`
	Retorna      string `json:"retorna"`
	Concurrencia string `json:"concurrencia"`
	Migrar       string `json:"migrar"`
	Metodo       string `json:"metodo"`
}

//Object Objeto para reflexiones
type Object map[string]interface{}

//Oficinas Reporte de oficinas
func (C *Core) Oficinas() (jSon []byte, err error) {
	var lst []Oficina
	s := `SELECT id, codoficina, descripcion FROM oficinas`
	sq, err := sys.PuntoPostalPostgres.Query(s)
	util.Error(err)

	for sq.Next() {
		var ofc Oficina
		var id, cod, desc sql.NullString
		sq.Scan(&id, &cod, &desc)
		ofc.ID = util.ValidarNullString(id)
		ofc.CodigoOficina = util.ValidarNullString(cod)
		ofc.Descripcion = util.ValidarNullString(desc)
		lst = append(lst, ofc)

	}
	jSon, err = json.Marshal(lst)
	return
}

//CrearQuery Creación dinamica de Consultas
func (C *Core) CrearQuery(v map[string]interface{}) (jSon []byte, err error) {
	conexion, a := leerValores(v)
	valores := strings.Split(a.Parametros, ",")

	consulta := a.Query
	cantidad := len(valores)
	for i := 0; i < cantidad; i++ {

		svalor := valores[i]
		pos := "$" + strconv.Itoa(i)
		consulta = strings.Replace(consulta, pos, svalor, -1)
	}
	tpQuery := evaluarQuery(consulta)

	for cf, vf := range tpQuery {
		if vf == true {
			switch cf {
			case "select":
				jSon, err = C.Select(v, consulta, conexion)
				break
			default:
				jSon, err = C.IUDQuery(consulta, conexion)
				break

			}
		}

	}

	return
}

//Select Crear Consultas Sql
func (C *Core) Select(v map[string]interface{}, consulta string, conexion *sql.DB) (jSon []byte, err error) {

	lista := make([]map[string]interface{}, 0)
	//fmt.Println("VIDA ", consulta)
	rs, _ := conexion.Query(consulta)
	cols, err := rs.Columns()
	if err != nil {
		panic(err)
	}
	colvals := make([]interface{}, len(cols))

	for rs.Next() {
		colassoc := make(map[string]interface{}, len(cols))
		for i := range colvals {
			colvals[i] = new(interface{})
		}
		if err := rs.Scan(colvals...); err != nil {
			panic(err)
		}
		for i, col := range cols {
			contenido := *colvals[i].(*interface{})
			evalreflect := reflect.ValueOf(contenido)

			//fmt.Println("TIPO ", evalreflect.Kind())
			switch evalreflect.Kind() {
			case reflect.Slice:
				valorstr := fmt.Sprintf("%s", contenido)
				colassoc[col] = strings.Trim(valorstr, " ")
				break
			case reflect.Float64:
				colassoc[col] = evalreflect.Float()
				break
				//WHERE dbo.ESTADISTICAS.CODOPT = '$0'
			case reflect.Int32:
				colassoc[col] = evalreflect.Int()
				break
			}

		}
		lista = append(lista, colassoc)

	}

	jSon, err = json.Marshal(lista)
	return
}

//IUDQuery Insert, Update, Delete Generador de Consultas
func (C *Core) IUDQuery(consulta string, conexion *sql.DB) (jSon []byte, err error) {
	var M util.Mensajes
	_, err = conexion.Exec(consulta)
	M.Fecha = time.Now()
	if err != nil {
		M.Msj = "Erro ejecutando consulta: " + err.Error()
		M.Tipo = 0
		jSon, err = json.Marshal(M)
	} else {
		M.Msj = "Proceso Exitoso"
		M.Tipo = 1
		jSon, err = json.Marshal(M)
	}

	return

}

//IUDQueryBash Insert, Update, Delete Generador de Consultas
func (C *Core) IUDQueryBash(campos string, lista []map[string]interface{}, consulta string, conexion *sql.DB) (jSon []byte, err error) {
	var M util.Mensajes

	for clave, valor := range lista {

	}
	_, err = conexion.Exec(consulta)
	M.Fecha = time.Now()
	if err != nil {
		M.Msj = "Erro ejecutando consulta: " + err.Error()
		M.Tipo = 0
		jSon, err = json.Marshal(M)
	} else {
		M.Msj = "Proceso Exitoso"
		M.Tipo = 1
		jSon, err = json.Marshal(M)
	}

	return

}

func leerValores(v map[string]interface{}) (db *sql.DB, a ApiCore) {
	parametro, ruta, metodo, migrar := retornaValores(v)
	c := sys.MGOSession.DB(sys.CBASE).C(sys.APICORE)
	err := c.Find(bson.M{"ruta": ruta}).One(&a)
	if err != nil {
		fmt.Println("Error creando Query en Mongodb ", err.Error())
	}
	switch a.Driver {
	case "puntopostal":
		db = sys.SqlServerPuntoPostal
		break
	case "ipostel":
		db = sys.PuntoPostalIpostel
		break
	case "tracking":
		db = sys.SqlServerTracking
		break
	}
	a.Parametros = parametro
	a.Migrar = migrar
	a.Metodo = metodo
	fmt.Println("Driver seleccionado: ", a.Driver)
	return
}

func retornaValores(v map[string]interface{}) (parametro string, ruta string, metodo string, migrar string) {
	for k, vs := range v {

		switch k {
		case "ruta":
			ruta = vs.(string)
			break
		case "parametros":
			parametro = vs.(string)
			break
		case "metodo":
			metodo = vs.(string)
			break
		case "migrar":
			migrar = vs.(string)
			break
		}
	}
	return
}

func evaluarQuery(cadena string) (s map[string]bool) {
	s = make(map[string]bool)
	s["select"] = strings.Contains(cadena, "SELECT")
	s["insert"] = strings.Contains(cadena, "INSERT")
	s["update"] = strings.Contains(cadena, "UPDATE")
	s["delete"] = strings.Contains(cadena, "DELETE")
	return s
}
