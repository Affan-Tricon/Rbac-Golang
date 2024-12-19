package middleware

import (
    "net/http"
    "strings"
)

func RBACMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        role := r.Header.Get("Role")
        if role == "" {
            http.Error(w, "Role header is missing", http.StatusBadRequest)
            return
        }

        if !isAllowed(role, r.Method) {
            http.Error(w, "Invalid request: insufficient permissions", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}


func isAllowed(role, method string) bool {
    method = strings.ToUpper(method) 

    switch role {
    case "admin":
        return true
    case "user", "guest":
        return method == http.MethodGet 
    default:
        return false 
    }
}
