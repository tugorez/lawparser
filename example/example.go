package main

import (
	"fmt"
	"github.com/tugorez/lawparser"
	"io/ioutil"
)

func main() {
	src, err := ioutil.ReadFile("bj.txt")
	p, err := lawparser.Parse(
		"Reglamento de tránsito de Benito Juarez", //nombre
		"México",        //pais
		"Quintana Roo",  //estado
		"Benito Juárez", //municipio
		"estatal",       //orden federal,estatal o municipal
		"reglamento",    //ley o reglamento o ...
		"tránsito",      //a que topico de interes comun pertenece
		"http://cancun.gob.mx/transparencia/files/2011/12/Reg_Transitoparael-Municipiode-BenitoJuarez_QRoo_30dic14.pdf", //de donde viene este doc
		src,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(p.Json()))
}
