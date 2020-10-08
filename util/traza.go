package util

import (
	"time"

	"gopkg.in/mgo.v2"
)

//Traza Historico del Usuario
type Traza struct {
	Usuario   string    `json:"usuario" bson:"usuario"`
	Time      time.Time `json:"tiempo" bson:"tiempo"`
	Log       string    `json:"log" bson:"log"`
	Documento string    `json:"documento" bson:"documento"`
	IP        string    `json:"ip" bson:"ip"`
	Base      string    `json:"base" bson:"base"`
	Coleccion string    `json:"coleccion" bson:"coleccion"`
}

//Crear Trazabilidad
func (t *Traza) Crear(sys *mgo.Session) (err error) {
	c := sys.DB(t.Base).C(t.Coleccion)
	err = c.Insert(t)
	return
}
