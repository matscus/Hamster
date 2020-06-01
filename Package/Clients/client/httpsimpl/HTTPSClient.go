package httpsimpl

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	caCertPool = x509.NewCertPool()
)

//NewHTTPSClient - return new HTTPS client whis ca cert
func NewHTTPSClient() (*http.Client, error) {
	caCert, err := ioutil.ReadFile(os.Getenv("SERVERREM"))
	caCertPool.AppendCertsFromPEM(caCert)
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
			TLSHandshakeTimeout: time.Second * 10,
		},
		Timeout: time.Second * 1,
	}, err
}
