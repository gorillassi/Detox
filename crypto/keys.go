// crypto/keys.go
package crypto

import (
	"crypto/ed25519"
	"encoding/json"
	"os"
	"path/filepath"

	"soulblog/core"
)

const KeyFile = "keys.json"

func GenerateAndSaveKeys(baseDir string) (*core.Keypair, error) {
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		return nil, err
	}
	kp := &core.Keypair{PublicKey: pub, PrivateKey: priv}
	dat, err := json.MarshalIndent(kp, "", "  ")
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(filepath.Join(baseDir, KeyFile), dat, 0600)
	if err != nil {
		return nil, err
	}
	return kp, nil
}

func LoadKeys(baseDir string) (*core.Keypair, error) {
	dat, err := os.ReadFile(filepath.Join(baseDir, KeyFile))
	if err != nil {
		return nil, err
	}
	var kp core.Keypair
	err = json.Unmarshal(dat, &kp)
	if err != nil {
		return nil, err
	}
	return &kp, nil
}
