package kafka

import (
	"crypto/tls"
	"crypto/x509"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/scram"
)

func newDialer(pem, username, password string) *kafka.Dialer {
	mechanism, err := scram.Mechanism(scram.SHA512, username, password)
	if err != nil {
		panic(err)
	}
	tlsConfig := pemToTLS(pem)
	dialer := &kafka.Dialer{
		Timeout:       10 * time.Second,
		DualStack:     true,
		SASLMechanism: mechanism,
		TLS:           tlsConfig,
	}

	return dialer
}

func pemToTLS(pemCert string) *tls.Config {
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM([]byte(pemCert))

	return &tls.Config{
		RootCAs: caCertPool,
	}
}

