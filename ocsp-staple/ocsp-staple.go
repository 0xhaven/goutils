package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/crypto/ocsp"
)

func getResponse(cert, issuer *x509.Certificate) ([]byte, error) {
	if len(cert.OCSPServer) == 0 {
		return nil, errors.New("no OCSPServer provided")
	}

	req, err := ocsp.CreateRequest(cert, issuer, nil)
	if err != nil {
		return nil, err
	}

	for _, server := range cert.OCSPServer {
		if u, err := url.Parse(server); err == nil {
			u.Path = base64.StdEncoding.EncodeToString(req)
			if resp, err := http.Get(u.String()); err == nil && resp.StatusCode == http.StatusOK {
				body, err := ioutil.ReadAll(resp.Body)
				resp.Body.Close()
				if err == nil {
					return body, nil
				}
				log.Println(err)
			} else {
				log.Println(err)
			}
		} else {
			log.Println(err)
		}
	}
	return nil, errors.New("no OCSP response")
}

func main() {
	for _, host := range os.Args[1:] {
		fmt.Println(host)
		host := net.JoinHostPort(host, "443")
		conn, err := tls.Dial("tcp", host, nil)
		peerCerts := conn.ConnectionState().PeerCertificates
		resp, err := getResponse(peerCerts[0], peerCerts[1])
		if err != nil {
			log.Println(err)
			continue
		}
		ocspResp, err := ocsp.ParseResponse(resp, peerCerts[1])
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("ProducedAt:", ocspResp.ProducedAt)
		fmt.Println("NextUpdate:", ocspResp.NextUpdate)
		fmt.Println("Delta:", ocspResp.NextUpdate.Sub(ocspResp.ProducedAt))
	}
}
