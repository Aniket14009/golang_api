package main

import "net/http"

func test(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(200)
    w.Write([]byte("Success"))
}

func main() {
    port := "8086"
    log.Println("{INF}: Listening on port", port)
    http.HandleFunc("/test", test)

    log.Fatal(http.ListenAndServe(":"+port, nil))
}
