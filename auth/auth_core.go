package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/optclblast/audioservice/logger"
)

const (
	PEM_LOCATION        = "./auth//bin/rsapems/"
	PUBLIC_KEY_LOCATION = "./auth/bin/rsapems/public.pem"
	PATH_TO_BINS        = "./auth/bin/"
	ERR                 = "The system cannot find the file specified"
)

type UserAccount struct {
	Id         int
	Login      string
	Password   string
	Created_at time.Time
	PublicKey  string
}

func HandleGetRSAKey() string {
	publicKeyBytes, err := os.ReadFile(PUBLIC_KEY_LOCATION)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/HandleGetRSAKey() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return "EOF"
	}
	return string(publicKeyBytes)
}

func GenerateKeyPair(withForce bool) error {
	if _, err := os.Stat(PUBLIC_KEY_LOCATION); !errors.Is(err, os.ErrNotExist) && !withForce {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  "RSA keys exist. Nothing to do here",
		})
		return nil
	}
	if withForce {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  "Re-generating RSA keys with force!",
		})
		err := os.Remove(PEM_LOCATION + "private.pem")
		if err != nil {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
				Content:  fmt.Sprintf("Error: %s", err),
			})
			return err
		}
		err = os.Remove(PUBLIC_KEY_LOCATION)
		if err != nil {
			logger.Logger(logger.LogEntry{
				DateTime: time.Now(),
				Level:    logger.ERROR,
				Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
				Content:  fmt.Sprintf("Error: %s", err),
			})
			return err
		}
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  "RSA keys have been removed!",
		})
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}

	privateKeyFile, err := os.Create(PEM_LOCATION + "private.pem")
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}

	privateKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}

	err = pem.Encode(privateKeyFile, privateKeyPEM)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}
	privateKeyFile.Close()

	// Save public key to file
	publicKeyFile, err := os.Create(PEM_LOCATION + "public.pem")
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}

	publicKey := privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}

	publicKeyPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	err = pem.Encode(publicKeyFile, publicKeyPEM)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}
	publicKeyFile.Close()
	logger.Logger(logger.LogEntry{
		DateTime: time.Now(),
		Level:    logger.INFO,
		Location: "AUTH_CORE/GenerateKeyPair() SCOPE",
		Content:  "RSA keys successfully created",
	})

	return nil
}

func encryptData(data []byte, publicKeyFile string, outputFile string) error {
	publicKeyBytes, err := os.ReadFile(publicKeyFile)
	if err != nil {
		return err
	}

	block, _ := pem.Decode(publicKeyBytes)
	if block == nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/encryptData() SCOPE",
			Content:  "failed to decode PEM block containing public key",
		})
		return os.ErrInvalid
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "AUTH_CORE/encryptData() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey.(*rsa.PublicKey), data)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "AUTH_CORE/encryptData() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}

	err = os.WriteFile(outputFile, encryptedData, 0644)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.INFO,
			Location: "AUTH_CORE/encryptData() SCOPE",
			Content:  fmt.Sprintf("Error: %s", err),
		})
		return err
	}

	return nil
}

func decryptData(encryptedData []byte, privateKeyFile string, outputFile string) ([]byte, error) {
	privateKeyBytes, err := os.ReadFile(privateKeyFile)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/DecryptData() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	block, _ := pem.Decode(privateKeyBytes)
	if block == nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/DecryptData() SCOPE",
			Content:  "failed to decode PEM block containing private key",
		})
		return nil, errors.New("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/DecryptData() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	decryptedData, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedData)
	if err != nil {
		logger.Logger(logger.LogEntry{
			DateTime: time.Now(),
			Level:    logger.ERROR,
			Location: "AUTH_CORE/DecryptData() SCOPE",
			Content:  fmt.Sprint(err),
		})
		return nil, err
	}

	return decryptedData, nil
}
