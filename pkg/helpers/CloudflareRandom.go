package helpers

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
)

const drandUrl = "https://drand.cloudflare.com/public/latest"

type cloudflareDrandResp struct {
	Round             int    `json:"round"`
	Signature         string `json:"signature"`
	PreviousSignature string `json:"previous_signature"`
	Randomness        string `json:"randomness"`
}

func GetCloudflareRand() (*rand.Rand, error) {
	resp, err := http.Get(drandUrl)
	if err != nil {
		return nil, err
	}

	respBytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	cfResp := new(cloudflareDrandResp)
	err = json.Unmarshal(respBytes, cfResp)
	if err != nil {
		return nil, err
	}

	randomnessInt, err := hex.DecodeString(cfResp.Randomness)
	if err != nil {
		return nil, err
	}

	return rand.New(rand.NewSource(int64(binary.BigEndian.Uint64(randomnessInt)))), nil

}
