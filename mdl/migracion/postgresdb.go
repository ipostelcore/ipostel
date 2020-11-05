package migracion

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/informaticaipsfa/tunel/util"
)

type Migrar struct{}

//Leer procesar archivos separados por puntos y comas csv
func (m *Migrar) Leer(db *sql.DB, codigo string, doc string) bool {
	var coma, concepto, tipo string
	var cantidad int
	codigomd5 := strings.Split(codigo, "|")

	insertar := ""
	archivo, err := os.Open(doc)
	util.Error(err)
	scan := bufio.NewScanner(archivo)
	i := 0
	for scan.Scan() {
		linea := strings.Split(scan.Text(), ";")
		l := len(linea)
		if i == 0 {
			concepto = linea[0]
			tipo = linea[4]
		} else { //Leyendo la primera linea
			cantidad++
			if l > 2 {
				if cantidad > 1 {
					coma = ","
				} else {
					coma = ""
				}

				cedula, _ := strconv.Atoi(strings.Split(linea[0], ".")[0])
				familiar, _ := strconv.Atoi(strings.Split(linea[1], ".")[0])
				monto := linea[2]
				insertar += coma + "('" + strconv.Itoa(cedula) + "','" + strconv.Itoa(familiar)
				insertar += "','" + codigomd5[0] + "','" + concepto + "'," + monto + ", " + tipo + ", Now(), '" + codigomd5[1] + "' )"

			} else { //DE LO CONTRARIO

				if cantidad > 1 {
					coma = ","
				} else {
					coma = ""
				}

				cedula, _ := strconv.Atoi(strings.Split(linea[0], ".")[0])
				monto := linea[1]
				insertar += coma + "('" + strconv.Itoa(cedula) + "','','" + codigomd5[0]
				insertar += "','" + concepto + "'," + monto + ", " + tipo + ", Now(), '" + codigomd5[1] + "')"
				//fmt.Println("Linea # ", i, cedula, "|", concepto)
			}

		}
		i++

	}

	fmt.Println("procesando ", i, doc)
	_, err = db.Exec(insertar)
	//fmt.Println(insertar)
	if err != nil {
		fmt.Println("ERR. AL PROCESAR ARCHIVO TXT ", doc, err.Error())
		return false
	}
	return true
}

func seleccionarTipo(prefijo string) string {
	var cadena string

	switch prefijo {
	case "int":
		cadena = "INTEGER"
		break
	case "dbl":
		cadena = "NUMERIC (15,3)"
		break
	case "str":
		cadena = "VARCHAR (256)"
		break
	case "txt":
		cadena = "TEXT"
		break
	case "fch":
		cadena = "TIMESTAMP"
		break
	default:
		cadena = "VARCHAR (256)"
		break

	}
	return cadena
}

func crearTabla(nombre string, campos string) string {
	var tabla string
	tabla = `CREATE TABLE IF NO EXISTS ` + nombre + ` ( m_oid serial PRIMARY KEY, )`
	return tabla
}

func isertar(tabla string, campos string) string {
	var fila string
	fila = `INSERT INTO  ` + tabla + ` (` + campos + `) VALUES `
	return fila
}
