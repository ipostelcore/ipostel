package util

import (
	"database/sql"
	"io/ioutil"
)

type Archivo struct {
	Responsable      int
	Ruta             string
	NombreDelArchivo string
	Codificacion     string
	Cabecera         string
	Leer             bool
	Salvar           bool
	Fecha            string
	CantidadLineas   int
	Registros        int
	PostgreSQL       *sql.DB
	Canal            chan []byte
}

//Directorio operaciones sobre un directorio
type Directorio struct {
	Ruta    string
	Listado []string `json:"listado"`
}

//Listar archivos del directorio principal
func (D *Directorio) Listar(ruta string) error {
	files, err := ioutil.ReadDir(ruta)
	if err != nil {
		return err
	}

	for _, f := range files {
		println(f.Name())
		D.Listado = append(D.Listado, f.Name())
	}
	return nil
}

func (a *Archivo) Crear(cadena string) bool {
	return true
}

// func (a *Archivo) LeerPorLinea(excelFileName string, PostgreSQLPENSIONSIGESP *sql.DB) bool {
// 	var iconstante, iconcepto string
// 	var codconcepto string
// 	xlFile, err := xlsx.OpenFile(excelFileName)
//
// 	switch excelFileName[4:7] {
// 	case "inv":
// 		codconcepto = "0000000027"
// 		break
// 	case "rcp":
// 		codconcepto = "0000000061"
// 		break
// 	case "sob":
// 		codconcepto = "0000000063"
// 		break
// 	}
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	i := 0
// 	coma := ""
// 	iconstante = `INSERT INTO sno_constantepersonal (codemp,codnom,codper,codcons,moncon,montopcon) VALUES `
// 	iconcepto = `INSERT INTO sno_conceptopersonal (codemp, codnom, codper, codconc, aplcon, valcon, acuemp,
// 		acuiniemp, acupat, acuinipat, acuinipataux, acupataux, acuiniempaux, acuempaux, valconaux) VALUES `
// 	fmt.Println("Preparando indices para el insert")
// 	for _, sheet := range xlFile.Sheets {
// 		for _, row := range sheet.Rows {
// 			if i > 0 {
// 				coma = ","
// 			}
// 			cedula := CompletarCeros(row.Cells[0].String(), 0, 10)
// 			monto := row.Cells[1].String()
// 			iconstante += coma + `('0001','0001','` + cedula + `','` + codconcepto + `',` + monto + `,0)`
// 			iconcepto += coma + `('0001','0001','` + cedula + `','` + codconcepto + `',1, 0, 0, 0, 0, 0, NULL, NULL, NULL, NULL, NULL)`
// 			i++
// 		}
// 	}
// 	fmt.Println("Insertando...")
// 	_, err = PostgreSQLPENSIONSIGESP.Exec(iconstante)
// 	if err != nil {
// 		fmt.Println("Error en la inserción: ", err.Error())
// 	}
// 	_, err = PostgreSQLPENSIONSIGESP.Exec(iconcepto)
// 	if err != nil {
// 		fmt.Println("Error en la inserción ", err.Error())
// 	}
//
// 	fmt.Println("Proceso exitoso...")
// 	fmt.Println(excelFileName[4:7])
// 	return true
// }

func (a *Archivo) LeerTodo() (f []byte, err error) {
	f, err = ioutil.ReadFile(a.NombreDelArchivo)
	return
}

func (a *Archivo) EscribirLinea(linea string) bool {
	return true
}

//Cerrar Archivos
func (a *Archivo) Cerrar() bool {
	return true
}
