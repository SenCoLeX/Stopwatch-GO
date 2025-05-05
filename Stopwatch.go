package main

import (
	"bufio" // Para leer la entrada del usuario (Enter)
	"fmt"   // Para imprimir mensajes en pantalla
	"os"    // Para acceder a la entrada estándar (teclado)
	"time"  // Para trabajar con el tiempo (reloj, duración, etc.)
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== Cronómetro en Go ===")
	fmt.Println("Presioná Enter para iniciar el cronómetro...")
	reader.ReadString('\n') // Espera que el usuario presione Enter

	start := time.Now() // Guarda el momento actual como "inicio"
	fmt.Println("Cronómetro iniciado. Presioná Enter para detener...")

	stop := make(chan bool) // Canal para indicar cuándo detener el cronómetro

	// Goroutine que actualiza el tiempo cada segundo
	go func() {
		for {
			select {
			case <-stop:
				return // Sale cuando recibe una señal de stop
			default:
				elapsed := time.Since(start) // Tiempo transcurrido desde el inicio
				fmt.Printf("\rTiempo transcurrido: %s", formatDuration(elapsed))
				time.Sleep(1 * time.Second) // Espera 1 segundo
			}
		}
	}()

	reader.ReadString('\n') // Espera otro Enter para detener el cronómetro
	stop <- true            // Envía señal para detener la goroutine

	final := time.Since(start) // Calcula el tiempo total transcurrido
	fmt.Printf("\n⏱ Tiempo total cronometrado: %s\n", formatDuration(final))
}

// Función para formatear el tiempo en formato HH:MM:SS
func formatDuration(d time.Duration) string {
	segundos := int(d.Seconds()) % 60
	minutos := int(d.Minutes()) % 60
	horas := int(d.Hours())

	return fmt.Sprintf("%02d:%02d:%02d", horas, minutos, segundos)
}
