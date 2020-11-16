package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

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

type ApiCore struct {
	Modulo     string `json:"modula"`
	Driver     string `json:"driver"`
	Query      string `json:"query"`
	Campos     string `json:"campos"`
	Parametros string `json:"parametros"`
	Ruta       string `json:"ruta"`
	Retorna    string `json:"retorna"`
}

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

	lista := make([]map[string]interface{}, 0)
	c, a := leerValores(v)
	valores := strings.Split(a.Parametros, ",")
	consulta := a.Query
	cantidad := len(valores)

	for i := 0; i < cantidad; i++ {

		svalor := valores[i]
		pos := "$" + strconv.Itoa(i)
		consulta = strings.Replace(a.Query, pos, svalor, -1)
	}

	fmt.Println(consulta)
	rs, _ := c.Query(consulta)
	cols, err := rs.Columns()
	if err != nil {
		panic(err)
	}
	colvals := make([]interface{}, len(cols))

	for rs.Next() {
		colassoc := make(map[string]interface{}, len(cols))
		for i, _ := range colvals {
			colvals[i] = new(interface{})
		}
		if err := rs.Scan(colvals...); err != nil {
			panic(err)
		}
		for i, col := range cols {
			contenido := *colvals[i].(*interface{})
			colassoc[col] = fmt.Sprintf("%s", contenido)
		}
		lista = append(lista, colassoc)

	}
	jSon, err = json.Marshal(lista)
	return
}

func leerValores(v map[string]interface{}) (db *sql.DB, a ApiCore) {

	parametro, ruta := retornaValores(v)

	c := sys.MGOSession.DB(sys.CBASE).C(sys.APICORE)
	err := c.Find(bson.M{"ruta": ruta}).One(&a)
	if err != nil {
		fmt.Println("Error creando Query en Mongodb ", err.Error())
	}
	switch a.Driver {
	case "puntopostal":
		db = sys.SqlServerPuntoPostal
		break
	case "tracking":
		db = sys.SqlServerTracking
		break
	}
	a.Parametros = parametro
	fmt.Println("Driver seleccionado: ", a.Driver)
	return
}
func retornaValores(v map[string]interface{}) (parametro string, ruta string) {
	for k, vs := range v {

		switch k {
		case "ruta":
			ruta = vs.(string)
			break
		case "parametros":
			parametro = vs.(string)
			break
		case "metodo":

		}
	}
}
