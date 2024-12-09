package codec

import "github.com/dlclark/regexp2"

func NewMistralTekken() *Codec {
	mistralTekkenVocabOnce.Do(mistralTekkenVocabInit)

	splitRegexp := regexp2.MustCompile(
		"[^\\r\\n\\p{L}\\p{N}]?[\\p{Lu}\\p{Lt}\\p{Lm}\\p{Lo}\\p{M}]*[\\p{Ll}\\p{Lm}\\p{Lo}\\p{M}]+|[^\\r\\n\\p{L}\\p{N}]?[\\p{Lu}\\p{Lt}\\p{Lm}\\p{Lo}\\p{M}]+[\\p{Ll}\\p{Lm}\\p{Lo}\\p{M}]*|\\p{N}| ?[^\\s\\p{L}\\p{N}]+[\\r\\n/]*|\\s*[\\r\\n]+|\\s+(?!\\S)|\\s+",
		regexp2.None)

	return &Codec{
		name:        "mistral_tekken",
		vocabulary:  mistralTekkenVocab,
		splitRegexp: splitRegexp,
		specialTokens: map[string]uint{
			"<unk>":              0,
			"<s>":                1,
			"</s>":               2,
			"[INST]":             3,
			"[/INST]":            4,
			"[AVAILABLE_TOOLS]":  5,
			"[/AVAILABLE_TOOLS]": 6,
			"[TOOL_RESULTS]":     7,
			"[/TOOL_RESULTS]":    8,
			"[TOOL_CALLS]":       9,
			"<pad>":              10,
			"[PREFIX]":           11,
			"[MIDDLE]":           12,
			"[SUFFIX]":           13,
		},
	}
}
