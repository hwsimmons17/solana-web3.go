package base58

import (
	"encoding/json"
	"log"
	"solana"
	"testing"

	"github.com/mr-tron/base58"
)

func TestParsePubkey(t *testing.T) {
	invalidBase58 := "_($#!@#$)"
	if _, err := ParsePubkey(invalidBase58); err == nil {
		log.Fatal("Expected error for invalid base58 string")
	}

	invalidLength := "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchz"
	if _, err := ParsePubkey(invalidLength); err == nil {
		log.Fatal("Expected error for invalid pubkey length")
	}

	valid := "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY"
	if _, err := ParsePubkey(valid); err != nil {
		log.Fatal("Expected no error for valid pubkey", err)
	}
}

func TestParsePubkeyBytes(t *testing.T) {
	bytes, _ := base58.Decode("5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY")
	key, _ := ParsePubkeyBytes(bytes)
	if key.String() != "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY" {
		log.Fatal("Expected pubkey to match")
	}
}

func TestPubkeyString(t *testing.T) {
	expected := "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY"
	key, _ := ParsePubkey(expected)
	if key.String() != expected {
		log.Fatal("Expected pubkey string to match")
	}
}

func TestPubkeyBytes(t *testing.T) {
	str := "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY"
	key, _ := ParsePubkey(str)
	bytes := key.Bytes()
	if len(bytes) != 32 {
		log.Fatal("Expected pubkey bytes to be 32 bytes")
	}
}

func TestIsOnCurve(t *testing.T) {
	strOnCurve := "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY"
	keyOnCurve, _ := ParsePubkey(strOnCurve)
	if !keyOnCurve.IsOnCurve() {
		log.Fatal("Expected pubkey to be on curve")
	}

	strNotOnCurve := "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrQ"
	keyNotOnCurve, _ := ParsePubkey(strNotOnCurve)
	if keyNotOnCurve.IsOnCurve() {
		log.Fatal("Expected pubkey to not be on curve")
	}
}

func TestMarshalJSON(t *testing.T) {
	type strct struct {
		Pubkey solana.Pubkey `json:"pubkey"`
	}

	str := "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY"
	key, _ := ParsePubkey(str)
	s := strct{Pubkey: key}
	json, err := json.Marshal(s)
	if err != nil {
		log.Fatal("Expected no error marshalling pubkey")
	}

	expected := `{"pubkey":"5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY"}`
	if string(json) != expected {
		log.Fatal("Expected pubkey json to match")
	}
}

func TestUnmarshalJSON(t *testing.T) {
	jsonStr := `{"pubkey":"5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY"}`
	type strct struct {
		Pubkey Pubkey `json:"pubkey"`
	}
	var result strct
	if err := json.Unmarshal([]byte(jsonStr), &result); err != nil {
		log.Fatal("Expected no error unmarshalling pubkey")
	}
	if result.Pubkey.String() != "5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejchzrY" {
		log.Fatal("Expected pubkey to match")
	}

	errJsonStr := `{"pubkey":"5oNDL3swdJJF1g9DzJiZ4ynHXgszjAEpUkxVYejc"}`
	if err := json.Unmarshal([]byte(errJsonStr), &result); err == nil {
		log.Fatal("Expected error unmarshalling pubkey")
	}
}
