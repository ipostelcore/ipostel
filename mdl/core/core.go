package core

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID           primitive.ObjectID `json:"id" bson:"_id"`
	Modulo       string             `json:"modulo"`
	Driver       string             `json:"driver"`
	Coleccion    string             `json:"coleccion"`
	Query        string             `json:"query"`
	Parametros   string             `json:"parametros"`
	Ruta         string             `json:"ruta"`
	Funcion      string             `json:"funcion"`
	Retorna      bool               `json:"retorna"`
	Concurrencia bool               `json:"concurrencia"`
	Migrar       bool               `json:"migrar"`
	Metodo       string             `json:"metodo"`
	Destino      string             `json:"destino"`
	Puerto       int                `json:"puerto"`
	Protocolo    string             `json:"protocolo"`
	Entorno      string             `json:"entorno"`
	Estatus      bool               `json:"estatus" bson:"estatus,omitempty"`
}

//Object Objeto para reflexiones
type Object map[string]interface{}

//OperarConsulta Control de resultados
func (C *Core) OperarConsulta(v map[string]interface{}) (jSon []byte, err error) {

	jSon, err = C.CrearQuery(v)
	return
}
