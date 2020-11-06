package main

import (
	"fmt"
	"time"
)

var mostrar int
var num_procesos int

func proceso(id int, c chan int) {
	terminar := false
	for i := uint64(0); terminar != true; i++ {
		time.Sleep(time.Millisecond * 500)
		select {
		case msg := <-c:
			if msg == id {
				fmt.Println("Eliminando...")
				terminar = true
			}
		default:
			if mostrar == 1 {
				fmt.Println(id, " : ", i)
			}
		}

	}
	fmt.Println("Proceso con id ", id, " terminado!")
	num_procesos--
}
func main() {
	var input string //pausa
	contador := 1
	c := make(chan int)
	opc := 1
	id := 0

	for opc != 4 {
		mostrar = 0
		fmt.Println("MENU PROCESOS")
		fmt.Println("1.- Crear")
		fmt.Println("2.- MOSTRAR")
		fmt.Println("3.- ELIMINAR")
		fmt.Println("4.- SALIR")
		fmt.Print("Ingresa el numero de la opcion: ")
		fmt.Scanf("%d\r", &opc)

		switch opc {
		case 1:
			go proceso(contador, c)
			fmt.Println("Proceso creado!")
			contador++
			num_procesos++
		case 2:
			mostrar = 1
			fmt.Scanln(&input)
		case 3:
			fmt.Print("Ingresa el id del proceso:")
			fmt.Scanf("%d\r", &id)
			for i := 0; i < num_procesos; i++ {
				c <- id
			}
		case 4:
			fmt.Println("Saliendo...")
		}
	}
}
