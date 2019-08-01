package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"os"
	"strings"
	"time"
)

type (
	JwtAuthToken struct {
		Header    string
		Payload   string
		Signature string
	}
)

func createToken(user *User) *JwtAuthToken {
	header := make(map[string]string)
	header["alg"] = "HS256"
	header["typ"] = "JWT"

	payload := make(map[string]string)
	payload["user"] = user.Username
	payload["exp"] = string(time.Now().Unix() + 100000)

	userSecret, err := decSecret(user.SecKey)
	if err != nil {
		panic(err)
	}

	headerStr, headErr := json.Marshal(header)
	payloadStr, payloadErr := json.Marshal(payload)

	if headErr != nil || payloadErr != nil {
		panic("There was a problem converting the header or payload!")
	}

	return generateNewToken(string(headerStr), string(payloadStr), userSecret)
}

func generateNewToken(header string, payload string, secret string) *JwtAuthToken {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)

	sigValue := strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(header)), "=") + "." + strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(payload)), "=")

	h.Write([]byte(sigValue))

	signature := sigValue + base64.StdEncoding.EncodeToString(h.Sum(nil))

	return &JwtAuthToken{header, payload, signature}
}

func encSecret(sec string) (string, error) {
	goSecret := os.Getenv("GOSEERSEC")

	if goSecret == "" {
		panic("GoSeer secret not found!")
	}

	block, err := aes.NewCipher([]byte(goSecret))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	iv := make([]byte, aesgcm.NonceSize())
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	cipher := aesgcm.Seal(iv, iv, []byte(sec), nil)

	return string(cipher), nil
}

func decSecret(sec string) (string, error) {
	goSecret := os.Getenv("GOSEERSEC")

	if goSecret == "" {
		panic("GoSeer secret not found!")
	}

	block, err := aes.NewCipher([]byte(goSecret))
	if err != nil {
		return "", err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	if len([]byte(sec)) < aesgcm.NonceSize() {
		panic("Bad secret key")
	}

	key, err := aesgcm.Open(nil, []byte(sec)[:aesgcm.NonceSize()], []byte(sec)[aesgcm.NonceSize():], nil)
	if err != nil {
		return "", err
	}

	return string(key), nil
}
