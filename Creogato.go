package main

import (
	
	"gato/objeto/funcioncrear"
)

func main() {
	g := new(funcioncrear.Gato)
	
	g.Crear_gato()

	
	p:= funcioncrear.NuevoGato(3,"Garfield","Naranja",true)

	p.MostrarGato()
}