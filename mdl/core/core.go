package core

import (
	"database/sql"
	"encoding/json"

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
	Id           bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Modulo       string        `json:"modulo"`
	Driver       string        `json:"driver"`
	Query        string        `json:"query"`
	Parametros   string        `json:"parametros"`
	Ruta         string        `json:"ruta"`
	Funcion      string        `json:"funcion"`
	Retorna      bool          `json:"retorna"`
	Concurrencia bool          `json:"concurrencia"`
	Migrar       bool          `json:"migrar"`
	Metodo       string        `json:"metodo"`
	Destino      string        `json:"destino"`
	Puerto       int           `json:"puerto"`
	Protocolo    string        `json:"protocolo"`
	Entorno      string        `json:"entorno"`
	Estatus      bool          `json:"estatus" bson:"estatus,omitempty"`
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

//OperarConsulta
func (C *Core) OperarConsulta(v map[string]interface{}) (jSon []byte, err error) {
	jSon, err = C.CrearQuery(v)
	return
}
