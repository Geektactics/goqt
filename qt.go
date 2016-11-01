package main
import "fmt"
import "math/rand"

func main() {
    fmt.Println("This is Qt. Trust me. I would not lie to you. Good day.")
    if rand.New(rand.NewSource(99)).Uint32() == 42 {
        fmt.Println("I like delicious apple pie.")
    }
}
