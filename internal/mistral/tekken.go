package mistral

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Config struct {
	Pattern                 string `json:"pattern"`
	NumVocabTokens          int    `json:"num_vocab_tokens"`
	DefaultVocabSize        int    `json:"default_vocab_size"`
	DefaultNumSpecialTokens int    `json:"default_num_special_tokens"`
	Version                 string `json:"version"`
}

type VocabEntry struct {
	Rank       int    `json:"rank"`
	TokenBytes string `json:"token_bytes"`
	TokenStr   string `json:"token_str"`
}

type Multimodal struct {
	ImagePatchSize int `json:"image_patch_size"`
	MaxImageSize   int `json:"max_image_size"`
}

type TekkenJSON struct {
	Config     Config       `json:"config"`
	Vocab      []VocabEntry `json:"vocab"`
	Multimodal Multimodal   `json:"multimodal"`
}

func UnmarshalTekken(r io.Reader) (*TekkenJSON, error) {
	var t TekkenJSON

	if err := json.NewDecoder(r).Decode(&t); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}

	return &t, nil
}

func (t *TekkenJSON) ToTikToken() *bytes.Buffer {
	buf := new(bytes.Buffer)

	for _, v := range t.Vocab {
		fmt.Fprintf(buf, "%s %d\n", v.TokenBytes, v.Rank)
	}

	return buf
}
