package migracion

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/ipostelcore/ipostel/util"
)

type Migrar struct {
	Ruta      string
	StrCols   []string
	StrValues []string
	Tabla     string
}

//Leer procesar archivos separados por puntos y comas csv
func (m *Migrar) Leer(db *sql.DB, codigo string, doc string) bool {

	codigomd5 := strings.Split(codigo, "|")

	//insertar := ""
	archivo, err := os.Open(doc)
	util.Error(err)
	scan := bufio.NewScanner(archivo)
	i := 0
	for scan.Scan() {
		linea := strings.Split(scan.Text(), ";")
		l := len(linea)
		if l > 0 {
			m.StrCols = nil
			m.StrValues = nil
			m.picarCampos(linea[1])
			m.Tabla = linea[0]
			//fmt.Println(m.crearTabla())
			_, err := db.Exec(m.crearTabla())
			if err != nil {
				fmt.Println("ERR. AL PROCESAR ARCHIVO TXT ", m.Tabla, err.Error())
				//return false
			}

			i++
		}

	}

	fmt.Println("procesando ", i, doc, codigomd5)
	/**
	_, err = db.Exec(//)
	if err != nil {
		fmt.Println("ERR. AL PROCESAR ARCHIVO TXT ", doc, err.Error())
		return false
	}**/
	return true
}

func (m *Migrar) picarCampos(campos string) {

	splitCampos := strings.Split(campos, ",")
	for i := 0; i < len(splitCampos); i++ {
		valor := strings.Split(splitCampos[i], ".")
		campo := m.seleccionarTipo(valor[0])

		//fmt.Printf("campos: %s, tipo: %s \n", campo, valor[1])
		m.StrCols = append(m.StrCols, valor[1]+" "+campo)
		m.StrValues = append(m.StrValues, valor[1])
	}

}

func (m *Migrar) seleccionarTipo(prefijo string) string {
	var cadena string

	switch prefijo {
	case "INT":
		cadena = "INTEGER"
		break
	case "DBL":
		cadena = "NUMERIC (15,3)"
		break
	case "STR":
		cadena = "VARCHAR (256)"
		break
	case "VAR":
		cadena = "VARCHAR (256)"
		break
	case "FLO":
		cadena = "VARCHAR (256)"
		break
	case "TXT":
		cadena = "TEXT"
		break
	case "BIN":
		cadena = "BOOLEAN"
		break
	case "BIT":
		cadena = "BOOLEAN"
		break
	case "TMS":
		cadena = "TIMESTAMP"
		break
	default:
		cadena = "VARCHAR (256)"
		break

	}
	return cadena
}

func (m *Migrar) crearTabla() string {
	var tabla string
	var coma, valores string
	tabla = `CREATE TABLE IF NOT EXISTS ` + m.Tabla + ` ( m_oid serial PRIMARY KEY, `
	for i := 0; i < len(m.StrCols); i++ {
		if i > 0 {
			coma = ","
		}
		valores += coma + m.StrCols[i]
	}
	tabla += valores + ")"
	return tabla
}

func isertar(tabla string, campos string, valores string) string {
	var fila string
	fila = `INSERT INTO  ` + tabla + ` (` + campos + `) VALUES (` + valores + `)`
	return fila
}
