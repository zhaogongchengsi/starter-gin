package utils

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"net"
	"os"
	"path"
	"time"
)

type CertConfig struct {
	ValidFor            time.Duration // 365*24*time.Hour, "Duration that certificate is valid for
	host                string        // Comma-separated hostnames and IPs to generate a certificate for
	ValidFrom           string        // Creation date formatted as Jan 1 15:04:05 2011
	IsCA                bool          // whether this cert should be its own Certificate Authority
	RsaBits             int           // Size of RSA key to generate. Ignored if --ecdsa-curve is set
	EcdsaCurve          string        // ECDSA curve to use to generate a key. Valid values are P224, P256 (recommended), P384, P521
	Ed25519Key          bool          // Generate an Ed25519 key
	certificate         x509.Certificate
	priv                any
	notBefore, notAfter time.Time
}

func NewCertConifg(host string) *CertConfig {
	return &CertConfig{
		host:       host,
		Ed25519Key: false,
		EcdsaCurve: "",
		RsaBits:    2048,
		IsCA:       false,
		ValidFrom:  "",
		ValidFor:   365 * 24 * time.Hour,
	}
}

func (c *CertConfig) GetHost() string {
	if len(c.host) == 0 {
		return "localhost"
	}
	return c.host
}

func (c *CertConfig) generatePriv() (any, error) {
	ecdsaCurve := c.EcdsaCurve
	ed25519Key := c.Ed25519Key
	rsaBits := c.RsaBits

	var priv any
	var err error

	switch ecdsaCurve {
	case "":
		if ed25519Key {
			_, priv, err = ed25519.GenerateKey(rand.Reader)
		} else {
			priv, err = rsa.GenerateKey(rand.Reader, rsaBits)
		}
	case "P224":
		priv, err = ecdsa.GenerateKey(elliptic.P224(), rand.Reader)
	case "P256":
		priv, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	case "P384":
		priv, err = ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	case "P521":
		priv, err = ecdsa.GenerateKey(elliptic.P521(), rand.Reader)
	default:
		return priv, fmt.Errorf("unrecognized elliptic curve: %q", ecdsaCurve)
	}
	if err != nil {
		return priv, err
	}

	c.priv = priv

	return priv, err
}

func (c *CertConfig) validPeriod(f time.Duration) (time.Time, time.Time, error) {
	var notBefore time.Time
	if len(c.ValidFrom) == 0 {
		notBefore = time.Now()
	} else {
		notBefore, err := time.Parse("Jan 2 15:04:05 2006", c.ValidFrom)
		if err != nil {
			return time.Time{}, time.Time{}, err
		}
		notAfter := notBefore.Add(c.ValidFor)
		c.notBefore = notBefore
		c.notAfter = notAfter
		return notBefore, notAfter, nil
	}
	notAfter := notBefore.Add(c.ValidFor)
	c.notBefore = notBefore
	c.notAfter = notAfter
	return notBefore, notAfter, nil
}

func (c *CertConfig) CreateCertificateTemp() (x509.Certificate, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)

	if err != nil {
		return x509.Certificate{}, err
	}

	notBefore, notAfter, err := c.validPeriod(c.ValidFor)

	if err != nil {
		return x509.Certificate{}, err
	}

	keyUsage := x509.KeyUsageDigitalSignature
	priv, err := c.generatePriv()
	if err != nil {
		return x509.Certificate{}, err
	}

	if _, isRSA := priv.(*rsa.PrivateKey); isRSA {
		keyUsage |= x509.KeyUsageKeyEncipherment
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	ip := net.ParseIP(c.GetHost())
	if ip != nil {
		template.IPAddresses = append(template.IPAddresses, ip)
	} else {
		template.DNSNames = append(template.DNSNames, c.host)
	}

	if c.IsCA {
		template.IsCA = true
		template.KeyUsage |= x509.KeyUsageCertSign
	}

	c.certificate = template

	return template, nil
}

func publicKey(priv any) any {
	switch k := priv.(type) {
	case *rsa.PrivateKey:
		return &k.PublicKey
	case *ecdsa.PrivateKey:
		return &k.PublicKey
	case ed25519.PrivateKey:
		return k.Public().(ed25519.PublicKey)
	default:
		return nil
	}
}

func (c *CertConfig) CreateCertificate() (*pem.Block, error) {
	c.CreateCertificateTemp()
	derBytes, err := x509.CreateCertificate(rand.Reader, &c.certificate, &c.certificate, publicKey(c.priv), c.priv)
	if err != nil {
		return &pem.Block{}, err
	}
	pem := &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}
	return pem, nil
}

func (c *CertConfig) CreateKey() (*pem.Block, error) {
	privBytes, err := x509.MarshalPKCS8PrivateKey(c.priv)
	if err != nil {
		return &pem.Block{}, err
	}
	pem := &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}
	return pem, nil
}

func (c *CertConfig) Generate(filepath, cetFileName, keyFileName string) error {
	cert, err := c.CreateCertificate()
	if err != nil {
		return err
	}
	key, err := c.CreateKey()
	if err != nil {
		return err
	}

	// write cert
	certPath := path.Join(filepath, cetFileName+".pem")
	cretOut, err := os.OpenFile(certPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	if err := pem.Encode(cretOut, cert); err != nil {
		return err
	}
	cretOut.Close()

	// write ket
	ketPath := path.Join(filepath, keyFileName+".pem")
	keyOut, err := os.OpenFile(ketPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	if err := pem.Encode(keyOut, key); err != nil {
		return err
	}
	keyOut.Close()

	return nil
}
