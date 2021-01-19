package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/ipostelcore/ipostel/sys"
	"github.com/ipostelcore/ipostel/util"
	"go.mongodb.org/mongo-driver/bson"
)

//CrearQuery Creación dinamica de Consultas
func (C *Core) CrearQuery(v map[string]interface{}) (jSon []byte, err error) {
	var M util.Mensajes
	c := sys.MongoDB.Collection(sys.APICORE)
	conexion, a, xmongo := leerValores(v)
	C.ApiCore = a
	if a.Estatus != true {
		M.Msj = "Driver de conexión falló"
		M.Tipo = 1
		M.Fecha = time.Now()
		jSon, err = json.Marshal(M)

		estatus := false

		result, err := c.UpdateOne(
			sys.Contexto,
			bson.M{"funcion": a.Funcion},
			bson.D{
				{"$set", bson.D{{"estatus", estatus}}},
			},
		)

		fmt.Printf("Documento Actualizado %v\n", result.ModifiedCount)
		return jSon, err
	}
	if a.Coleccion != "" {

		jSon, err = C.CrearNOSQL(C.ApiCore.Coleccion, C.ApiCore.Query, xmongo)
		return
	}

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
				if C.ApiCore.Retorna == false && C.ApiCore.Migrar == true {
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
	var estatus bool
	var M util.Mensajes

	c := sys.MongoDB.Collection(sys.APICORE)

	lista := make([]map[string]interface{}, 0)
	//fmt.Println(consulta)
	rs, e := conexion.Query(consulta)
	if e != nil {
		M.Msj = "Select fallo"
		M.Tipo = 1
		M.Fecha = time.Now()
		jSon, err = json.Marshal(M)
		return
	}

	cols, err := rs.Columns()
	if err != nil {
		estatus = false
		_, err = c.UpdateOne(
			sys.Contexto,
			bson.M{"funcion": C.ApiCore.Funcion},
			bson.M{"$set": bson.M{"estatus": estatus}})
		return
	}
	fmt.Println(cols)
	colvals := make([]interface{}, len(cols))

	for rs.Next() {
		colassoc := make(map[string]interface{}, len(cols))
		for i := range colvals {
			colvals[i] = new(interface{})
		}
		if err := rs.Scan(colvals...); err != nil {
			estatus := false
			result, _ := c.UpdateOne(
				sys.Contexto,
				bson.M{"funcion": C.ApiCore.Funcion},
				bson.D{
					{"$set", bson.D{{"estatus", estatus}}},
				},
			)

			fmt.Printf("Documento Actualizado %v\n", result.ModifiedCount)
		}
		for i, col := range cols {
			contenido := *colvals[i].(*interface{})
			evalreflect := reflect.ValueOf(contenido)
			//fmt.Println("Evaluacion Reflect ", evalreflect.Kind())
			switch evalreflect.Kind() {
			case reflect.Slice:
				valorstr := fmt.Sprintf("%s", contenido)
				colassoc[col] = util.Utf8_decode(strings.Trim(valorstr, " "))
				break
			case reflect.String:
				colassoc[col] = evalreflect.String()
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
		estatus = true

	}
	if C.ApiCore.Migrar == true {
		go C.IUDQueryBash(C.ApiCore.Destino, lista, "", conexion)
	}

	jSon, err = json.Marshal(lista)

	fmt.Println("Finalizando _ID ", C.ApiCore.ID, estatus)
	_, err = c.UpdateOne(
		sys.Contexto,
		bson.M{"funcion": C.ApiCore.Funcion},
		bson.D{
			{"$set", bson.D{{"estatus", estatus}}},
		},
	)

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
			}

			j++
		}
		j = 0
		valores += ")"
		campos += ")"
		insert += campos + " VALUES " + valores + ";"
		campos = ""
		valores = ""
		db := sys.SQLTODO["PGODB"].DB
		_, err = db.Exec(insert)
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
	}
	return

}

func leerValores(v map[string]interface{}) (db *sql.DB, a ApiCore, mgo *mongo.Database) {

	ApiCoreAux := retornaValores(v)
	c := sys.MongoDB.Collection(sys.APICORE)
	estatus := false
	err := c.FindOne(sys.Contexto, bson.M{"funcion": ApiCoreAux.Funcion}).Decode(&a)
	if err != nil {
		fmt.Println("Error creando Query en Mongodb "+"funcion: "+ApiCoreAux.Funcion, err.Error())

	}
	if sys.SQLTODO[a.Driver].Estatus != false {
		drv := sys.SQLTODO[a.Driver]
		db = drv.DB
		estatus = true
	}
	if a.Coleccion != "" {
		drv := sys.NOSQLTODO[a.Driver]
		mgo = drv.DB
		estatus = true
	}
	a.Parametros = ApiCoreAux.Parametros
	a.Migrar = ApiCoreAux.Migrar
	a.Metodo = ApiCoreAux.Metodo
	a.Destino = ApiCoreAux.Destino
	a.Retorna = ApiCoreAux.Retorna
	a.Estatus = estatus
	a.Funcion = ApiCoreAux.Funcion
	fmt.Println("Driver seleccionado: ", a.Funcion, a.Coleccion)
	return
}

func retornaValores(v map[string]interface{}) (a ApiCore) {
	for k, vs := range v {
		switch k {
		case "coleccion":
			a.Coleccion = vs.(string)
			break
		case "funcion":
			a.Funcion = vs.(string)
			break
		case "parametros":
			a.Parametros = vs.(string)
			break
		case "metodo":
			a.Metodo = vs.(string)
			break
		case "migrar":
			a.Migrar = vs.(bool)
			break
		case "destino":
			a.Destino = vs.(string)
			break
		case "retorna":
			a.Retorna = vs.(bool)
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
