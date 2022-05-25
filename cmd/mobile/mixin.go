package mobile

import (
	"context"
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	bot "github.com/MixinNetwork/bot-api-go-client"
	"time"
)

type MixinKeychain struct {
	ClientId   string `json:"client_id"`
	SessionId  string `json:"session_id"`
	PinToken   string `json:"pin_token"`
	PrivateKey string `json:"private_key"`
}

//TODO how to store secrets
const (
	clientId   = "e119e713-35ef-4f4e-819c-836b29a55914"
	sessionId  = "a7e121da-d546-46dd-9e24-380069ec5796"
	privateKey = "5mkvZPhR993d4jbS5WQTrZ2cr9iq21Dcg_0bLUYuopPN-r0xmGWUb8VSMG2HQm9Gt0Wn7mNE_kwDsUWRBIydGQ"
)

func CreateNetworkUserWithPIN(pin string) (string, error) {
	ctx := context.Background()

	user, userPrivateKeyBase64, err := createNetworkUser(ctx)
	if err != nil {
		return "", err
	}

	err = setupPin(ctx, pin, user, userPrivateKeyBase64)
	if err != nil {
		return "", err
	}

	mixinKeychain := MixinKeychain{ClientId: user.UserId, SessionId: user.SessionId, PinToken: user.PINTokenBase64, PrivateKey: userPrivateKeyBase64}
	data, _ := json.Marshal(mixinKeychain)
	return string(data), nil
}

func createNetworkUser(ctx context.Context) (*bot.User, string, error) {
	publicKey, userPrivateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, "", err
	}
	sessionSecret := base64.RawURLEncoding.EncodeToString(publicKey[:])
	userPrivateKeyBase64 := base64.RawURLEncoding.EncodeToString(userPrivateKey[:])

	user, err := bot.CreateUser(ctx, sessionSecret, "Vlow User", clientId, sessionId, privateKey)
	if err != nil {
		return nil, "", err
	}

	return user, userPrivateKeyBase64, nil
}

func setupPin(ctx context.Context, pin string, user *bot.User, userSessionKey string) error {
	encryptedPIN, err := bot.EncryptEd25519PIN(ctx, pin, user.PINTokenBase64, user.SessionId, userSessionKey, uint64(time.Now().UnixNano()))
	if err != nil {
		return err
	}
	err = bot.UpdatePin(ctx, "", encryptedPIN, user.UserId, user.SessionId, userSessionKey)
	if err != nil {
		return err
	}
	return nil
}

func EncryptPin(pin string, userPinToken string, userSessionId string, userPrivateKey string) (string, error) {
	encryptedPIN, err := bot.EncryptEd25519PIN(nil, pin, userPinToken, userSessionId, userPrivateKey, uint64(time.Now().UnixNano()))
	if err != nil {
		return "", err
	}
	return encryptedPIN, nil
}
