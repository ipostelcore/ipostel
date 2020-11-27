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
	ApiCore
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
	Destino      string `json:"destino"`
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
	C.ApiCore = a

	valores := strings.Split(C.ApiCore.Parametros, ",")

	consulta := C.ApiCore.Query
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
				if C.ApiCore.Retorna == "false" && C.ApiCore.Migrar == "true" {
					var M util.Mensajes
					M.Msj = "Proceso finalizado"
					M.Tipo = 1
					go C.Select(v, consulta, conexion)
					jSon, err = json.Marshal(M)
				} else {
					jSon, err = C.Select(v, consulta, conexion)
				}

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
			switch evalreflect.Kind() {
			case reflect.Slice:
				valorstr := fmt.Sprintf("%s", contenido)
				colassoc[col] = util.Utf8_decode(strings.Trim(valorstr, " "))
				break
			case reflect.Float32:
				colassoc[col] = evalreflect.Float()
				break
			case reflect.Float64:
				colassoc[col] = evalreflect.Float()
				break
			case reflect.Int32:
				colassoc[col] = evalreflect.Int()
				break
			case reflect.Int64:
				colassoc[col] = evalreflect.Int()
				break
			}

		}
		lista = append(lista, colassoc)

	}
	if C.ApiCore.Migrar == "true" {
		go C.IUDQueryBash(C.ApiCore.Destino, lista, "", conexion)
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
func (C *Core) IUDQueryBash(tabla string, lista []map[string]interface{}, consulta string, conexion *sql.DB) (jSon []byte, err error) {
	var M util.Mensajes
	var campos, valores string
	i := 0
	j := 0

	for _, valor := range lista {
		insert := "INSERT INTO " + tabla + " "
		campos = "("
		valores = "("
		i++

		for c, v := range valor {
			comax := ","
			if j == 0 {
				comax = ""
			}
			campos += comax + c

			evalreflect := reflect.ValueOf(v)

			switch evalreflect.Kind() {
			case reflect.String:
				valorstr := fmt.Sprintf("%s", v)
				valores += comax + "'" + util.Utf8_decode(strings.Trim(valorstr, " ")) + "'"
				break
			case reflect.Slice:
				valorstr := fmt.Sprintf("%s", v)
				valores += comax + "'" + strings.Trim(valorstr, " ") + "'"
				break
			case reflect.Float32:
				f := evalreflect.Float()
				valores += comax + strconv.FormatFloat(f, 'f', 2, 64)
				break
			case reflect.Float64:
				f := evalreflect.Float()
				valores += comax + strconv.FormatFloat(f, 'f', 2, 64)
				break
			case reflect.Int32:
				n := evalreflect.Int()
				valores += comax + strconv.FormatInt(n, 10)
				break
			case reflect.Int64:
				n := evalreflect.Int()
				valores += comax + strconv.FormatInt(n, 10)
				break
				// default:
				// 	valores += comax + "''"
				// 	fmt.Println(c+" TIPO ", " ESPACIO ")
				// 	break
			}

			j++
		}
		j = 0
		valores += ")"
		campos += ")"
		insert += campos + " VALUES " + valores + ";"
		campos = ""
		valores = ""

		_, err = sys.PuntoPostalPostgres.Exec(insert)
		M.Fecha = time.Now()
		if err != nil {
			fmt.Println("----> ", err.Error())
			M.Msj = "Erro ejecutando consulta: " + err.Error()
			M.Tipo = 0
			jSon, err = json.Marshal(M)
		} else {
			M.Msj = "Proceso Exitoso"
			M.Tipo = 1
			jSon, err = json.Marshal(M)
		}
		//fmt.Println(insert)
	}
	return

}

func leerValores(v map[string]interface{}) (db *sql.DB, a ApiCore) {

	ApiCoreAux := retornaValores(v)
	c := sys.MGOSession.DB(sys.CBASE).C(sys.APICORE)
	//fmt.Println("Api ", ApiCoreAux.Ruta)
	err := c.Find(bson.M{"ruta": ApiCoreAux.Ruta}).One(&a)
	if err != nil {
		fmt.Println("Error creando Query en Mongodb ", err.Error())
	}
	switch a.Driver {
	case "puntopostal":
		db = sys.SqlServerPuntoPostal
		break
	case "ipostel":
		db = sys.PuntoPostalPostgres
		break
	case "tracking":
		db = sys.SqlServerTracking
		break
	}
	a.Parametros = ApiCoreAux.Parametros
	a.Migrar = ApiCoreAux.Migrar
	a.Metodo = ApiCoreAux.Metodo
	a.Destino = ApiCoreAux.Destino
	a.Retorna = ApiCoreAux.Retorna
	fmt.Println("Driver seleccionado: ", a.Driver)
	return
}

func retornaValores(v map[string]interface{}) (a ApiCore) {
	for k, vs := range v {
		switch k {
		case "ruta":
			a.Ruta = vs.(string)
			break
		case "parametros":
			a.Parametros = vs.(string)
			break
		case "metodo":
			a.Metodo = vs.(string)
			break
		case "migrar":
			a.Migrar = vs.(string)
			break
		case "destino":
			a.Destino = vs.(string)
			break
		case "retorna":
			a.Retorna = vs.(string)
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
