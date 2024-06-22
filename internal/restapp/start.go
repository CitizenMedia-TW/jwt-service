package restapp

import (
	"encoding/json"
	"jwt-service/internal/config"
	"net/http"
)

type RestServer struct {
	config config.Config
}

func New(cnf config.Config) *RestServer {
	return &RestServer{
		config: cnf,
	}
}

func (s *RestServer) Routes() http.Handler {
	// Declare a new router
	mux := http.NewServeMux()

	mux.Handle("/", s.middlewares(http.HandlerFunc(s.home)))
	mux.Handle("/generate", s.middlewares(http.HandlerFunc(s.GenerateToken)))
	mux.Handle("/verify", s.middlewares(http.HandlerFunc(s.VerifyToken)))

	return mux
}

// Define the middlewares for the application
func (s *RestServer) middlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		// log.Print("Executing middlewareTwo again")
	})
}

// home is the handler for the home page
func (s *RestServer) home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a HandleFunc #1"))
}

func (s *RestServer) httpError(w *http.ResponseWriter, err string, statusCode int) {
	(*w).Header().Set("Content-Type", "application/json")
	res := struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}{Message: "Failed", Error: err}
	json.NewEncoder(*w).Encode(res)
	return
}
