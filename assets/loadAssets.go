package assets

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

var Texts struct {
	Help   string `yaml:"help"`
	Encode string `yaml:"encode"`
	Decode string `yaml:"decode"`
}

var EncryptionTable struct {
	Source         []string `yaml:"source"`
	Additional     []string `yaml:"additional"`
	UseAdditionals bool     `yaml:"useAdditionals"`
}

//go:embed texts.yaml
var textsFile []byte

//go:embed encryptionTable.yaml
var encryptionTableFile []byte

func LoadAssets() {
	yaml.Unmarshal(textsFile, &Texts)
	yaml.Unmarshal(encryptionTableFile, &EncryptionTable)
}
