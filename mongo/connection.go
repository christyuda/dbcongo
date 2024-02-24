package mongo

import (
    "context"
    "fmt"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// DBInfo menyimpan informasi tentang koneksi database
type DBInfo struct {
    DBString string
}

// ConnectToDB berfungsi untuk membuat koneksi ke MongoDB
func ConnectToDB(info DBInfo) (*mongo.Database, error) {
    // Mendapatkan string koneksi dari environment variable
    uri := info.DBString
    if uri == "" {
        return nil, fmt.Errorf("DBString tidak ditemukan dalam environment variable")
    }

    // Set up client options
    clientOptions := options.Client().ApplyURI(uri)
    clientOptions.SetUseNewUrlParser(true)
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

    // Mengembalikan objek database
    return client.Database("default"), nil // Ubah sesuai nama database Anda
}
