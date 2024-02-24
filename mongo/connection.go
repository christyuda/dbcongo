package mongo

import (
    "context"
    "fmt"
    "os"
    "strings"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB berfungsi untuk membuat koneksi ke MongoDB
func ConnectDB() (*mongo.Database, error) {
    // Mendapatkan string koneksi dari environment variable
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        return nil, fmt.Errorf("MONGO_URI tidak ditemukan dalam environment variable")
    }

    // Parse URI
    clientOptions, err := options.ParseURI(uri)
    if err != nil {
        return nil, fmt.Errorf("gagal memparse URI: %v", err)
    }

    // Set useUnifiedTopology menjadi true
    clientOptions.SetUseUnifiedTopology(true)

    // Membuat klien MongoDB
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        return nil, fmt.Errorf("gagal membuat koneksi ke MongoDB: %v", err)
    }

    // Konteks dengan timeout
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Koneksi ke server MongoDB
    err = client.Connect(ctx)
    if err != nil {
        return nil, fmt.Errorf("gagal terhubung ke server MongoDB: %v", err)
    }

    // Mendapatkan nama database dari URI
    dbName := getDatabaseName(uri)

    // Mengembalikan objek database
    return client.Database(dbName), nil
}

// getDatabaseName berfungsi untuk mendapatkan nama database dari URI
func getDatabaseName(uri string) string {
    parts := strings.Split(uri, "/")
    dbNameWithParams := parts[len(parts)-1]
    dbName := strings.Split(dbNameWithParams, "?")[0]
    return dbName
}
