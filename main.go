package main

import (
    "fmt"
    "net/http"
    "RBAC/middleware"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Request successful: %s %s\n", r.Method, r.URL.Path)
}

func main() {
    mux := http.NewServeMux()

    mux.Handle("/api/resource", middleware.RBACMiddleware(http.HandlerFunc(mainHandler)))

    fmt.Println("Server running on :3000")
    http.ListenAndServe(":3000", mux)
}
