package main

import (
    "fmt"
    "log"

    "path/to/mongo-connection/mongo"
)

func main() {
    // Terhubung ke database
    db, err := mongo.ConnectDB()
    if err != nil {
        log.Fatalf("gagal terhubung ke database: %v", err)
    }

    // Melakukan sesuatu dengan objek database
    fmt.Println("Terhubung ke database!")

    // db adalah objek database yang bisa digunakan untuk operasi MongoDB lainnya
}
