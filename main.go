package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web"))
	http.Handle("/", fs)

	addr := ":8116"
	fmt.Printf("👓 👓 VisionOptic - Eyewear & Eye Exam Booking Server đang chạy tại http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
