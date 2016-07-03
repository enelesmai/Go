// Crear un programa que declare dos funciones anónimas. Una que cuente del 0 al 200 y otra que cuente de 200 a 0.
// Mostrar cada número con un identificador único para cada goroutine.
// Crear goroutines a partir de estas funciones.
// No permitir que main regrese hasta que ambas goroutines hayan terminado de correr.

package main

// Agregar imports.
import (
		"runtime"
		"fmt"
		"sync"
		"bytes"
		"strconv"
		)

func init() {
	// Alojar un procesador para que el scheduler lo use.
	runtime.GOMAXPROCS(1)
}

func increment(wg *sync.WaitGroup){
	// Declarar un loop que cuente del 0 al 200 e imprimir cada valor
	var buffer bytes.Buffer
	for i := 1; i <= 200; i++ {
		buffer.WriteString("increment: " + strconv.Itoa(i) + "\t")
	}
	fmt.Println(buffer.String())

	// Decrementar la cuenta del wait group.
	defer wg.Done()
}

func decrement(wg *sync.WaitGroup){
	// Declarar un loop que cuente del 200 al 0 e imprimir cada valor
	var buffer bytes.Buffer
	for i := 200; i > 0; i-- {
		buffer.WriteString("decrement: " + strconv.Itoa(i) + "\t")
	}
	fmt.Println(buffer.String())

	// Decrementar la cuenta del wait group.
	defer wg.Done()
}

func main() {

	// Declarar el wait group e iniciar el contador en 2.
	var wg sync.WaitGroup
    wg.Add(2)

	// Declarar una función anónima y crear una goroutine.
	go increment(&wg)

	// Declarar una función anónima y crear una goroutine.
	go decrement(&wg)

	// Esperar a que las goroutines terminen.
	wg.Wait()
	fmt.Printf("hello, the routines have finished\n")
}
