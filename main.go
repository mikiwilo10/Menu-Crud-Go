package main

import (
	"Menu_Crud_Persona/model"
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var id int

//creacion de un mapa
var users = make(map[int]model.Persona)

func crearUsuario() {

	//limpia la consola
	clearConsola()

	fmt.Print("Ingresa Nombre Usuario: ")
	nombre := LlamarFunc()
	fmt.Print("Ingresa Apellido Usuario: ")
	apellido := LlamarFunc()
	fmt.Print("Ingresa Edad Usuario: ")
	//convertir un string a entero
	edad, err := strconv.Atoi(LlamarFunc())
	if err != nil {
		panic("Error Edad")
	}
	id++
	per := model.Persona{
		IdPersona: id,
		Nombre:    nombre,
		Apellido:  apellido,
		Edad:      edad,
	}
	//almacena una Personsa en el mapa
	users[id] = per
	//fmt.Println("Persona Creada", users)

}

func listarUsuario() {

	clearConsola()
	fmt.Println("Listar Usuario")
	for _, pers := range users {
		fmt.Println(pers.IdPersona, "->", pers.Nombre, pers.Apellido)
	}
}

func actualizarUsuario() {
	fmt.Println("Actualizar Usuario")
}

func eliminarUsuario() {
	clearConsola()
	fmt.Println("Ingresa el Id de la Persona a Eliminar")
	id, _ := strconv.Atoi(LlamarFunc())

	if _, per := users[id]; per {
		delete(users, id)
	}

	fmt.Println("Eliminar Usuario\n")
}

//Funcion que limpia la consola
func clearConsola() {
	/*Dependiendo del S.O elegir el comando indicado
	En este caso Windows
	Ejecutar cualquier sentencia de comando en Go
	*/
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func LlamarFunc() string {
	read := bufio.NewReader(os.Stdin)
	//obetenemos todo lo que tipemos por el teclado
	// '\n' indica hasta que presionesmo enter

	if valor, err := read.ReadString('\n'); err != nil {
		panic("Ingrese los valores Establecidos")
	} else {
		//eliminar caracteres de un string
		return strings.TrimSpace(valor)
	}

}

func main() {

	var valor string

	for {
		fmt.Println("a) Crear")
		fmt.Println("b) Listar")
		fmt.Println("c) Eliminar")
		fmt.Println("d) Actualizar ")
		fmt.Println("q) salir ")

		fmt.Print("Elija una Opcion:")
		valor = LlamarFunc()

		if valor == "salir" || valor == "q" {
			break
		}
		switch valor {
		case "a", "Crear":
			crearUsuario()
		case "b", "Listar":
			listarUsuario()
		case "c", "Eliminar":
			eliminarUsuario()
		case "d", "Actualizar":
			actualizarUsuario()
		default:
			fmt.Println("Opcion no valida")
		}
	}
	fmt.Println("Adios")
}
