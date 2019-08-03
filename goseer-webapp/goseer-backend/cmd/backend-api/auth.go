package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	payload["exp"] = "2019-08-03"

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

	encodedHeader := strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(header)), "=")
	encodedPayload := strings.TrimRight(base64.URLEncoding.EncodeToString([]byte(payload)), "=")

	sigValue := encodedHeader + "." + encodedPayload

	h.Write([]byte(sigValue))

	signature := strings.TrimRight(base64.URLEncoding.EncodeToString(h.Sum(nil)), "=")

	return &JwtAuthToken{encodedHeader, encodedPayload, signature}
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

	return fmt.Sprintf("%x", cipher), nil
}

func decSecret(sec string) (string, error) {
	goSecret := os.Getenv("GOSEERSEC")

	secretKey, _ := hex.DecodeString(sec)

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

	ivSize := aesgcm.NonceSize()
	iv, ciphertext := secretKey[:ivSize], secretKey[ivSize:]

	key, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s", string(key)), nil
}
