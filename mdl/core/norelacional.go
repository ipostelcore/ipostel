package core

import (
	"encoding/json"
	"fmt"

	"github.com/ipostelcore/ipostel/sys"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//InsertNOSQL Insert, Generador de Consultas
func (C *Core) InsertNOSQL(v map[string]interface{}, conexion *mgo.Session) (jSon []byte, err error) {

	c := conexion.DB(sys.CBASE).C(sys.APICORE)
	err = c.Insert(v)
	if err != nil {
		fmt.Println("Error creando Query en Mongodb ", err.Error())
	}
	return

}

//Listar Insert, Generador de Consultas
func (C *Core) Listar(conexion *mgo.Session) (jSon []byte, err error) {

	c := conexion.DB(sys.CBASE).C(sys.APICORE)
	var lstApi []ApiCore

	err = c.Find(bson.M{}).All(&lstApi)
	jSon, _ = json.Marshal(lstApi)
	return

}
