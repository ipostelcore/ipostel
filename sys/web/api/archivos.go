package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ipostelcore/ipostel/mdl/migracion"
	"github.com/ipostelcore/ipostel/sys"
	"github.com/ipostelcore/ipostel/util"
)

//SubirArchivo Permite procesar archivos al sistema
func (a *API) SubirArchivos(w http.ResponseWriter, r *http.Request) {
	Cabecera(w, r)
	var traza util.Traza
	var M util.Mensajes

	ip := strings.Split(r.RemoteAddr, ":")
	traza.IP = ip[0]
	traza.Time = time.Now()
	traza.Usuario = UsuarioConectado.Login

	er := r.ParseMultipartForm(32 << 20)
	if er != nil {
		fmt.Println(er)
		return
	}
	m := r.MultipartForm
	files := m.File["input-folder-2"]
	codigo := r.FormValue("txtFileID")
	directorio := "./public_web/www/temp/migracion/"
	errr := os.Mkdir(directorio, 0777)
	if errr != nil {
		fmt.Println("El directorio ya existe!")
	}
	cadena := ""
	for i := range files {
		file, errf := files[i].Open()
		defer file.Close()
		if errf != nil {
			fmt.Println(errf)
			return
		}
		out, er := os.Create(directorio + files[i].Filename)
		defer out.Close()
		if er != nil {
			fmt.Println(er.Error())
			return
		}
		_, err := io.Copy(out, file) // file not files[i] !
		if err != nil {
			fmt.Println(err)
			return
		}
		cadena += files[i].Filename + ";"
		//ProcesarTxt(files[i].Filename, codigo)
		fmt.Println(codigo)

	} // Fin de archivos
	//
	M.Msj = "Proceso exitoso"

	j, _ := json.Marshal(M)
	w.WriteHeader(http.StatusOK)
	w.Write(j)

}

//ProcesarTxt Proceso de archivo
func ProcesarTxt(doc string, codigo string) {
	var procesar migracion.Migrar
	procesar.Ruta = "./public_web/www/temp/migracion/" + doc
	db := sys.SQLTODO["PGODB"].DB
	procesar.Leer(db, codigo, doc)

}
