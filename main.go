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
        // Intercambiar el elemento mínimo encontrado con el primer elemento no ordenado
        arr[i], arr[minIndex] = arr[minIndex], arr[i]
    }
}

func main() {
    // Ejemplo de uso
    arr := []int{64, 25, 12, 22, 11}
    fmt.Println("Array antes de ordenar:", arr)
    selectionSort(arr)
    fmt.Println("Array después de ordenar:", arr)
}
