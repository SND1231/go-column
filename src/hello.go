package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuID, _ := uuid.NewUUID()
	fmt.Println("Hello world!!", uuID)
}
