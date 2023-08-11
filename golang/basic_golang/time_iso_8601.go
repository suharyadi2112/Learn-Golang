package main

import (
	"fmt"
	"time"
)

func main() {
	// Buat instance tanggal dan waktu saat ini
	now := time.Now()

	// Format tanggal dan waktu menggunakan layout yyyy-MM-dd'T'HH:mm:ssTZD
	formattedDate := now.Format("2006-01-02T15:04:05-07:00")

	fmt.Println(formattedDate)
}