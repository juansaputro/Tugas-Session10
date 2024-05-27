package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

func main() {
	dialer := &net.Dialer{
		Timeout: 5 * time.Second,
	}
	conn, err := tls.DialWithDialer(dialer, "tcp", "www.example.com:443", &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		fmt.Println("Failed to connect:", err)
		return
	}
	defer conn.Close()

	state := conn.ConnectionState()
	fmt.Println("TLS Version:", tlsVersion(state.Version))
	fmt.Println("CipherSuite:", tls.CipherSuiteName(state.CipherSuite))

	for _, cert := range state.PeerCertificates {
		fmt.Println("Issuer Organization:", cert.Issuer.Organization)
		break
	}
}

func tlsVersion(version uint16) string {
	switch version {
	case tls.VersionTLS13:
		return "TLS 1.3"
	case tls.VersionTLS12:
		return "TLS 1.2"
	case tls.VersionTLS11:
		return "TLS 1.1"
	case tls.VersionTLS10:
		return "TLS 1.0"
	default:
		return "Unknown"
	}
}
