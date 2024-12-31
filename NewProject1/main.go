package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

// rollDice generates a random number between 1 and 6.
func rollDice() int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(6) + 1
}

func diceHandler(w http.ResponseWriter, r *http.Request) {
    dice := rollDice()
    fmt.Fprintf(w, "You rolled: %d", dice)
}

func main() {
    http.HandleFunc("/roll", diceHandler)
    fmt.Println("Starting server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}
