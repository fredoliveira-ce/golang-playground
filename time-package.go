package main

import "fmt"
import "time"

const stretching = 1
const warming = 5
const jogging = 10
const running = 10

func startTraining() {
    fmt.Println("stretching period...")
    time.Sleep(stretching * time.Minute)

    fmt.Println("warming period...")
    time.Sleep(warming * time.Minute)

    fmt.Println("jogging period...")
    time.Sleep(jogging * time.Minute)

    fmt.Println("running period...")
    time.Sleep(running * time.Minute)

    fmt.Println("stretching period...")
    time.Sleep(stretching * time.Minute)
}