/*
Create a Go program that computes how long does it take for the Sunlight to reach the Earth if we know that the distance from the Sun to Earth is 149.6 million km and the speed of light  is 299,792,458 m / s

The formula used is: time = distance / speed
*/
package main
 
import "fmt"
 
func main() {
    sunToEart := 149.6
	kmToM := 1000
	oneMillion := 1000000
    distance := int(sunToEart * float64(kmToM) * float64(oneMillion))
	speedOfLight := 299_792_458 
    time := distance/speedOfLight
	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time)

 
}
