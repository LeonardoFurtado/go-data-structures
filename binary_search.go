package main
import "fmt"

func BinarySearch(arr []int, target int, min int, max int) int{
    if min > max{
        return -1
    }
    middle := min + (max - min) / 2
    if arr[middle] == target {
        return middle
    } else if arr[middle] > target{
        return BinarySearch(arr, target, min, middle - 1)
    }
    return BinarySearch(arr, target, middle + 1, max)
    
}

func main() {
    meuArr := []int{1,3,5,7,9,15,17,19,21}
    result := BinarySearch(meuArr, 19, 0, 8)
    fmt.Printf("Result: %d", result)
}
