package main

// si yo a mi user.go le cambiara el nombre, esto no funciona
// no puedo importar archivos .go, siempre paquetes

import (
	"fmt"
	"time"
	us "./user" 
)

// creamos estructura
/*type usuario struct {
	Id			int
	Nombre		string
	FechaAlta	time.Time
	Status 		bool
}*/

type paquito struct {
	us.Usuario //hago referencia a usuario
}


type humano interface{
	pensar()
}

type animal interface{
	respirar()
	comer()
	EsCarnivoro() bool
	sexo() string
}

/* Genero Humano */
type hombre struct{
	edad		int
	altura 		float32
	peso		float32
	pensando	bool
}

// detecta directamente la interface con los metodos implementados
func (h *hombre) pensar() { h.pensando=true }

func HumanoPensando(hu humano){
	hu.pensar()
	fmt.Println("Entra")
}

func main(){
	/*User := new(usuario)
	User.Id = 10
	User.Nombre = "Nahuel"

	fmt.Println(User)*/

	user:=new(paquito)
	user.AltaUsuario(1, "Paquito", time.Now(), true)
	fmt.Println(user)
	// user.Usuario.AltaUsuario(3, "Paquito2", time.Now(), true)
	fmt.Println(user.Usuario)

	pija := new(hombre)
	HumanoPensando(pija)
}