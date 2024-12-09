package tokenizer

// Package tokenizer provides functions for encoding and decoding text using
// different tokenization schemes.
//
// Encoding Formats
//
// The following encoding formats are supported:
// - Cl100kBase
// - O200kBase
// - R50kBase
// - P50kBase
// - P50kEdit
// - MistralTekken
//
// Alternatively you can request a tokenizer using OpenAI's model name, the
// following OpenAI models are supported:
// - O1Preview
// - O1Mini
// - GPT4
// - GPT35Turbo
// - TextEmbeddingAda002
// - TextDavinci003
// - TextDavinci002
// - CodeDavinci002
// - CodeDavinci001
// - CodeCushman002
// - CodeCushman001
// - DavinciCodex
// - CushmanCodex
// - TextDavinci001
// - TextCurie001
// - TextBabbage001
// - TextAda001
// - Davinci
// - Curie
// - Babbage
// - Ada
// - TextSimilarityDavinci001
// - TextSimilarityCurie001
// - TextSimilarityBabbage001
// - TextSimilarityAda001
// - TextSearchDavinciDoc001
// - TextSearchCurieDoc001
// - TextSearchAdaDoc001
// - TextSearchBabbageDoc001
// - CodeSearchBabbageCode001
// - CodeSearchAdaCode001
// - TextDavinciEdit001
// - CodeDavinciEdit001
//
// Or use the Mistral name:
// - MistralNemo
// - MistralMixtral
//
// Usage Example
//
// Here is an example of how to encode a string using the `ForModel` function:
//
//	package main
//
//	import (
//		"fmt"
//		"github.com/tiktoken-go/tokenizer"
//	)
//
//	func main() {
//		enc, err := tokenizer.Get(tokenizer.Cl100kBase)
//		if err != nil {
//			panic("oh oh")
//		}
//
//		// this should print a list of token ids
//		ids, token, _ := enc.Encode("supercalifragilistic")
//		fmt.Println(ids)
//
//		// this should print the original string back
//		text, _ := enc.Decode(ids)
//		fmt.Println(text)
//}

import (
	"errors"
	"strings"

	"github.com/tiktoken-go/tokenizer/codec"
)

var (
	ErrModelNotSupported    = errors.New("model not supported")
	ErrEncodingNotSupported = errors.New("encoding not supported")
)

type Codec interface {
	GetName() string
	Encode(string) ([]uint, []string, error)
	Decode([]uint) (string, error)
}

type Model string

const (
	O1Preview                Model = "o1-preview"
	O1Mini                   Model = "o1-mini"
	GPT4o                    Model = "gpt-4o"
	GPT4                     Model = "gpt-4"
	GPT35Turbo               Model = "gpt-3.5-turbo"
	GPT35                    Model = "gpt-3.5"
	TextEmbeddingAda002      Model = "text-embedding-ada-002"
	TextDavinci003           Model = "text-davinci-003"
	TextDavinci002           Model = "text-davinci-002"
	CodeDavinci002           Model = "code-davinci-002"
	CodeDavinci001           Model = "code-davinci-001"
	CodeCushman002           Model = "code-cushman-002"
	CodeCushman001           Model = "code-cushman-001"
	DavinciCodex             Model = "davinci-codex"
	CushmanCodex             Model = "cushman-codex"
	TextDavinci001           Model = "text-davinci-001"
	TextCurie001             Model = "text-curie-001"
	TextBabbage001           Model = "text-babbage-001"
	TextAda001               Model = "text-ada-001"
	Davinci                  Model = "davinci"
	Curie                    Model = "curie"
	Babbage                  Model = "babbage"
	Ada                      Model = "ada"
	TextSimilarityDavinci001 Model = "text-similarity-davinci-001"
	TextSimilarityCurie001   Model = "text-similarity-curie-001"
	TextSimilarityBabbage001 Model = "text-similarity-babbage-001"
	TextSimilarityAda001     Model = "text-similarity-ada-001"
	TextSearchDavinciDoc001  Model = "text-search-davinci-doc-001"
	TextSearchCurieDoc001    Model = "text-search-curie-doc-001"
	TextSearchAdaDoc001      Model = "text-search-ada-doc-001"
	TextSearchBabbageDoc001  Model = "text-search-babbage-doc-001"
	CodeSearchBabbageCode001 Model = "code-search-babbage-code-001"
	CodeSearchAdaCode001     Model = "code-search-ada-code-001"
	TextDavinciEdit001       Model = "text-davinci-edit-001"
	CodeDavinciEdit001       Model = "code-davinci-edit-001"
	GPT2                     Model = "gpt2"

	// Mistral
	MistralNemo Model = "mistral_tekken"
)

type Encoding string

const (
	// OpenAI
	GPT2Enc    Encoding = "gpt2"
	R50kBase   Encoding = "r50k_base"
	P50kBase   Encoding = "p50k_base"
	P50kEdit   Encoding = "p50k_edit"
	Cl100kBase Encoding = "cl100k_base"
	O200kBase  Encoding = "o200k_base"

	// Mistral
	MistralTekken Encoding = "mistral_tekken"
)

var modelPrefixToEncoding map[Model]Encoding = map[Model]Encoding{
	"o1-":              O200kBase,
	"gpt-4o-":          O200kBase,
	"gpt-4-":           Cl100kBase,
	"gpt-3.5-turbo-":   Cl100kBase,
	"gpt-35-turbo-":    Cl100kBase,
	"ft:gpt-4":         Cl100kBase,
	"ft:gpt-3.5-turbo": Cl100kBase,
	"ft:davinci-002":   Cl100kBase,
	"ft:babbage-002":   Cl100kBase,
}

// Get returns a new instance of a Codec implementation based on the specified
// encoding format. The returned Codec instance can be used to encode (tokenize)
// and decode (reassemble) text. If the specified encoding is not supported,
// an error is returned.
func Get(encoding Encoding) (Codec, error) {
	switch encoding {
	case O200kBase:
		return codec.NewO200kBase(), nil
	case Cl100kBase:
		return codec.NewCl100kBase(), nil
	case R50kBase:
		return codec.NewR50kBase(), nil
	case P50kBase:
		return codec.NewP50kBase(), nil
	case P50kEdit:
		return codec.NewP50kEdit(), nil
	case MistralTekken:
		return codec.NewMistralTekken(), nil
	default:
		return nil, ErrEncodingNotSupported
	}
}

// ForModel returns a new instance of a Codec implementation based on the
// specified OpenAI model. If the specified model is not supported, an error
// is returned.
func ForModel(model Model) (Codec, error) {
	switch model {
	case O1Preview, O1Mini, GPT4o:
		return Get(O200kBase)

	case GPT4, GPT35, GPT35Turbo, TextEmbeddingAda002:
		return Get(Cl100kBase)

	case TextDavinci003, TextDavinci002, CodeDavinci001,
		CodeDavinci002, CodeCushman002, CodeCushman001,
		DavinciCodex, CushmanCodex:
		return Get(P50kBase)

	case TextDavinci001, TextCurie001, TextBabbage001, TextAda001, Davinci,
		Curie, Babbage, Ada, TextSimilarityDavinci001, TextSimilarityCurie001,
		TextSimilarityBabbage001, TextSimilarityAda001, TextSearchDavinciDoc001,
		TextSearchCurieDoc001, TextSearchAdaDoc001, TextSearchBabbageDoc001,
		CodeSearchBabbageCode001, CodeSearchAdaCode001:
		return Get(R50kBase)

	case TextDavinciEdit001, CodeDavinciEdit001:
		return Get(P50kEdit)

	case GPT2:
		return Get(GPT2Enc)

	// Mistral
	case MistralNemo:
		return Get(MistralTekken)
	default:
		for prefix, enc := range modelPrefixToEncoding {
			if strings.HasPrefix(string(model), string(prefix)) {
				return Get(enc)
			}
		}
		return nil, ErrModelNotSupported
	}
}
