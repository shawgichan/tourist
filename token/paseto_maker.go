package token

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMaker struct {
	Paseto       *paseto.V2
	SymmetricKey []byte
}

func NewPastoMaker(SymmetricKey string) (Maker, error) {

	if len(SymmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size : must be exactly %d char ", chacha20poly1305.KeySize)
	}

	maker := &PasetoMaker{
		Paseto:       paseto.NewV2(),
		SymmetricKey: []byte(SymmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMaker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		log.Println("tesitng errr")
		return "", err
	}
	token, err := maker.Paseto.Encrypt(maker.SymmetricKey, payload, nil)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (maker *PasetoMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.Paseto.Decrypt(token, maker.SymmetricKey, payload, nil)
	if err != nil {
		return nil, errors.New("error is while verify")
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}
