package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"golang.org/x/crypto/ssh"
)

var (
	outputDir = flag.String("outputDir", "./assets/secrets", `(default "./assets/secrets")`)
)

func run(name string) error {
	sshKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	publicSSHKey, err := ssh.NewPublicKey(&sshKey.PublicKey)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path.Join(*outputDir, "id_rsa.pub"), ssh.MarshalAuthorizedKey(publicSSHKey), 0600)
	if err != nil {
		return err
	}

	// Get ASN.1 DER format
	privDER := x509.MarshalPKCS1PrivateKey(sshKey)

	// pem.Block
	privBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privDER,
	}

	err = ioutil.WriteFile(path.Join(*outputDir, "id_rsa"), pem.EncodeToMemory(&privBlock), 0600)
	if err != nil {
		return err
	}

	return nil
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "usage: %s commonName\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if err := run(flag.Arg(0)); err != nil {
		panic(err)
	}
}
