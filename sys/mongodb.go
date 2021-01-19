package sys

import (
	"context"
	"encoding/json"

	"github.com/fatih/color"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//Mongo Database
type Mongo struct {
	DataBase *mongo.Database
	Ctx      context.Context
}

//Conectar Establecer conexión
func (m *Mongo) Conectar(Ctx context.Context, c CadenaDeConexion) (DB *mongo.Database, e error) {
	var cnf Config

	uri := "mongodb://" + c.Host + ":" + c.Puerto
	Client, err := mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := Client.Ping(Ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	DB = Client.Database(c.Basedatos)

	color.Green("... Host: " + c.Host + " Base De Datos: ( " + c.Basedatos + " )  OK...")
	rs, e := DB.Collection("drivers").Find(Ctx, bson.D{{}})
	for rs.Next(Ctx) {
		var cadena CadenaDeConexion
		e = rs.Decode(&cadena)
		if e != nil {
			color.Red("No se logro establecer la definición del driver, por favor verifique")
			return
		}
		if cadena.Estatus {
			cnf.ConexionesDinamicas(cadena)
		}
	}
	return
}

//ConectarDinamicamente Establecer conexión
func (m *Mongo) ConectarDinamicamente(Ctx context.Context, c CadenaDeConexion) (DB *mongo.Database, e error) {

	uri := "mongodb://" + c.Host + ":" + c.Puerto
	Client, err := mongo.Connect(Ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	if err := Client.Ping(Ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	DB = Client.Database(c.Basedatos)

	color.Green("... Host: " + c.Host + " Base De Datos: ( " + c.Basedatos + " )  OK...")

	return
}

//Buscar un registro
func (m *Mongo) Buscar(coleccion string, filtro interface{}) (j []byte, err error) {
	var doc interface{}
	c := m.DataBase.Collection(coleccion)
	//.Select(bson.M{"clave": false})
	err = c.FindOne(Contexto, bson.M{"cedula": ""}).Decode(&doc)
	j, _ = json.Marshal(doc)
	return
}

//Salvar documentos
func (m *Mongo) Salvar(mgo interface{}, coleccion string) (err error) {

	// c := MGOSession.DB(CBASE).C(coleccion)
	// err = c.Insert(mgo)

	return
}
