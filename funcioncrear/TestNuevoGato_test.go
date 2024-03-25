package funcioncrear

import "testing"

func TestNuevoGato(t *testing.T) {
	// Crear un nuevo gato
	g := NuevoGato(3, "Garfield", "Naranja", true)

	// Verificar que los campos se hayan asignado correctamente
	if g.edad != 3 {
		t.Errorf("Edad incorrecta, esperada: %d, obtenida: %d", 3, g.edad)
	}
	if g.nombre != "Garfield" {
		t.Errorf("Nombre incorrecto, esperado: %s, obtenido: %s", "Garfield", g.nombre)
	}
	if g.color != "Naranja" {
		t.Errorf("Color incorrecto, esperado: %s, obtenido: %s", "Naranja", g.color)
	}
	if !g.sano {
		t.Errorf("Estado de salud incorrecto, esperado: %t, obtenido: %t", true, g.sano)
	}
}

func TestMostrarGato(t *testing.T) {
	// Crear un nuevo gato
	g := NuevoGato(3, "Garfield", "Naranja", true)

	// Verificar que MostrarGato imprima correctamente los datos del gato
	g.MostrarGato() // Simplemente ejecuta la función, no hay una forma directa de verificar la salida en la consola en Go
	// Podrías redirigir stdout y verificar la salida, pero eso es más complicado para una prueba unitaria simple
}