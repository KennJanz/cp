package main

import "fmt"

func selectionSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        minIndex := i
        for j := i + 1; j < n; j++ {
            if arr[j] < arr[minIndex] {
                minIndex = j
            }
        }
        
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Array antes de ordenar:", arr)
    selectionSort(arr)
    fmt.Println("Array después de ordenar:", arr)
}


//bsc

package main

import (
	"fmt"
)

var numeroCambiar int = 8


func main() {
	var arrays =[10]int{1,7,3,4,6,8,2,10,9,5}

  linealsearch(arrays[:])
}


func linealsearch(arrays []int){

  for i:=0;i<len(arrays);i++{
    if arrays[i] != numeroCambiar{
      fmt.Println("numero no encontrado")
    }else{
      fmt.Println("numero",numeroCambiar ,"encontrado en la posicion",arrays[i]-arrays[i]+i)
      
    }
  }
}


// bs

package main

import "fmt"

func bubbleSort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                // Intercambiar arr[j] y arr[j+1]
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }
}

func main() {
    arr := []int{64, 34, 25, 12, 22, 11, 90}
    fmt.Println("Array original:", arr)
    bubbleSort(arr)
    fmt.Println("Array ordenado:", arr)
}


//bb

package main

import "fmt"

func binarySearch(arr []int, target int) int {
    low := 0
    high := len(arr) - 1

    for low <= high {
        mid := (low + high) / 2
        guess := arr[mid]

        if guess == target {
            return mid
        } else if guess < target {
            low = mid + 1
        } else {
            high = mid - 1
        }
    }

    return -1
}

func main() {
    // Array de ejemplo
    arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
    target := 11

    // Llamada a la función de búsqueda binaria
    index := binarySearch(arr, target)

    // Imprimir resultado
    if index != -1 {
        fmt.Printf("El elemento %d se encuentra en la posición %d\n", target, index)
    } else {
        fmt.Printf("El elemento %d no se encuentra en el array\n", target)
    }
}

//oi

package main

import "fmt"

// Función para ordenar un slice usando el algoritmo de inserción
func insertionSort(arr []int) {
    n := len(arr)
    for i := 1; i < n; i++ {
        key := arr[i]
        j := i - 1
        // Mover los elementos de arr[0..i-1] que son mayores que key
        // a una posición adelante de su posición actual
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j = j - 1
        }
        arr[j+1] = key
    }
}

func main() {
    // Ejemplo de uso
    arr := []int{12, 11, 13, 5, 6}
    fmt.Println("Slice antes de ordenar:", arr)
    insertionSort(arr)
    fmt.Println("Slice después de ordenar:", arr)
}
