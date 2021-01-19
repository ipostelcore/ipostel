package webscraper

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

//CodigoFuente permite crear archivos ts, js
type CodigoFuente struct {
	Nombre      string   `json:"nombre"`
	Goblal      string   `json:"global"`
	Constructor string   `json:"constructor"`
	Obtener     string   `json:"obtener"`
	Cargar      string   `json:"cargar"`
	Crear       string   `json:"crear"`
	Items       []string `json:"Items"`
	TypeScript  string   `json:"ts"`
	JavaScript  string   `json:"js"`
}

//WebScraping Herramienta que permite leer codigo fuente html para obtener id
type WebScraping struct {
	Ruta    string
	Archivo io.Reader
	Items   []string `json:"Items"`
	Codigo  []CodigoFuente
}

//Leer Archivos
func (W *WebScraping) Leer() (err error) {
	W.Archivo, err = os.Open(W.Ruta)
	W.obtener()
	W.generar()
	return
}

//Obtener datos del archivo
func (W *WebScraping) obtener() error {
	doc, err := goquery.NewDocumentFromReader(W.Archivo)

	doc.Find("form").Each(func(i int, s *goquery.Selection) {
		var cf CodigoFuente

		cf.Nombre, _ = s.Attr("id")
		cf.Goblal, cf.Constructor, cf.Obtener, cf.Cargar = W.elemento("input", s)

		W.Codigo = append(W.Codigo, cf)
	})

	return err
}

func (W *WebScraping) elemento(element string, doc *goquery.Selection) (cgbl string, ccon string, cobt string, ccarg string) {
	doc.Find(element).Each(func(i int, s *goquery.Selection) {
		valor, _ := s.Attr("id")
		a := strings.ToLower(valor)
		W.Items = append(W.Items, a)
		cgbl += "\t" + a[3:] + " : " + seleccionarTipo(a[:3]) + ";\n"
		ccon += "\t\tthis." + a[3:] + " = \"\";\n"
		cobt += "\tthis." + a[3:] + " = $(\"#" + valor + "\").val();\n"
		ccarg += "\t$(\"#" + valor + "\").val(this." + a[3:] + ");\n"
	})
	return
}

func seleccionarTipo(prefijo string) (cadena string) {
	switch prefijo {
	case "cmb":
		cadena = "string"
		break
	case "txt":
		cadena = "string"
		break
	case "num":
		cadena = "number"
		break
	case "chk":
		cadena = "number"
		break
	case "rdb":
		cadena = "boolean"
		break
	case "tms":
		cadena = "date"
		break
	default:
		cadena = "string"
		break
	}
	return cadena
}

//generar datos del archivo
func (W *WebScraping) generar() {

	for _, v := range W.Codigo {
		fmt.Println(codigoClase(v))
	}
	return
}
