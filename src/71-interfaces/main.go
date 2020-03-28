package main
import (
	"71-interfaces/animales"
)

func main(){
	var p animales.Perro
	var g animales.Gato
	p.Nombre = "Firulais"
	g.Nombre = "mancha"
	// AdoptarPerro(p)
	// AdoptarGato(g)
	AdoptarMascota(p)
	AdoptarMascota(g)
}

// func AdoptarPerro(p animales.Perro){
// 	p.Comunicarse()
// }

// func AdoptarGato(g animales.Gato){
// 	g.Comunicarse()
// }

func AdoptarMascota(m animales.Mascota){
	m.Comunicarse()
}