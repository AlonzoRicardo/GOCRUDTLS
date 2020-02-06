package main

import (
	"crypto/tls"
	// "go_auth/keygen"
	"go_auth/user"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var (
	certFile    = os.Getenv("CERT_FILE")
	keyFile     = os.Getenv("KEY_FILE")
	serviceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {
	// For loggin all requests
	logger := log.New(os.Stdout, "go_auth ", log.LstdFlags|log.Lshortfile)
	user.InitialMigration()
	r := mux.NewRouter()
	srv := newServer(r, serviceAddr)
	userHandler := user.NewHandlers(logger)
	userHandler.SetupRoutes(r)

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// uuid := keygen.GenerateKey()
	// 	w.Header().Set("Content-Type", "text/plain")
	// 	w.WriteHeader(http.StatusOK)
	// 	w.Write([]byte(uuid))
	// })

	err := srv.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}

func newServer(router *mux.Router, serverAddr string) *http.Server {
	tlsConfig := &tls.Config{
		// Causes servers to use Go's default cipher suite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have assembly implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519, // Go 1.8 only
		},

		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	srv := &http.Server{
		Addr:         serverAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      router,
	}

	return srv
}
