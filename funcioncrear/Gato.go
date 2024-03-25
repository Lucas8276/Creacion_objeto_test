package funcioncrear

import "fmt"

type Gato struct {
	edad          int
	nombre, color string
	sano          bool
}
func NuevoGato(edad int, nombre,color string, sano bool) * Gato{
	return &Gato{
		edad: edad,
		nombre: nombre,
		color: color,
		sano: sano,
	}
}
func (g *Gato)MostrarGato(){
	fmt.Println("Nombre del gato",g.nombre,"\n Edad del gato ",g.edad,"\nColor del gato",g.color,"\nEsta sano",g.sano)
}
func (g *Gato) Crear_gato() {
	var sano string
	fmt.Println("Asigne los siguientes datos a su gato")
	fmt.Print("Edad: ")
	fmt.Scanln(&g.edad)
	fmt.Print("Nombre: ")
	fmt.Scanln(&g.nombre)
	fmt.Print("Color: ")
	fmt.Scanln(&g.color)
	fmt.Print("¿Está sano? (si/no): ")
	fmt.Scanln(&sano)
	g.sano = sano == "si"
}
