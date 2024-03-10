package alphabet

import (
	"crypto/rand"
	"errors"
	"math/big"
)

type Alphabet struct {
	chars       []rune
	charToIndex map[rune]int
	indexToChar map[int]rune
	len         int
}

func New(chars []rune) *Alphabet {
	charToIndex := make(map[rune]int)
	indexToChar := make(map[int]rune)

	for i, c := range chars {
		charToIndex[c] = i
		indexToChar[i] = c
	}

	return &Alphabet{
		chars:       chars,
		charToIndex: charToIndex,
		indexToChar: indexToChar,
		len:         len(chars),
	}
}

func (a *Alphabet) Encrypt(text string, key string) (string, error) {
	textRunes := []rune(text)
	keyRunes := []rune(key)

	if len(textRunes) != len(keyRunes) {
		return "", errors.New("text and key must have same length")
	}

	result := make([]rune, len(textRunes))

	for i := range result {
		ti, ok := a.charToIndex[textRunes[i]]
		if !ok {
			return "", errors.New("text contains chars not within the alphabet")
		}

		ki, ok := a.charToIndex[keyRunes[i]]
		if !ok {
			return "", errors.New("key contains chars not within the alphabet")
		}

		newi := (ti + ki) % a.len
		result[i] = a.indexToChar[newi]
	}

	return string(result), nil
}

func (a *Alphabet) Decrypt(text string, key string) (string, error) {
	textRunes := []rune(text)
	keyRunes := []rune(key)

	if len(textRunes) != len(keyRunes) {
		return "", errors.New("text and key must have same length")
	}

	result := make([]rune, len(textRunes))

	for i := range result {
		ti, ok := a.charToIndex[textRunes[i]]
		if !ok {
			return "", errors.New("text contains chars not within the alphabet")
		}

		ki, ok := a.charToIndex[keyRunes[i]]
		if !ok {
			return "", errors.New("key contains chars not within the alphabet")
		}

		newi := (ti - ki) % a.len
		if newi < 0 {
			newi += a.len
		}
		result[i] = a.indexToChar[newi]
	}

	return string(result), nil
}

func (a *Alphabet) RandKey(length int) string {
	max := big.NewInt(int64(len(a.chars)))
	key := make([]rune, length)

	for i := range key {
		n, _ := rand.Int(rand.Reader, max)
		key[i] = a.indexToChar[int(n.Int64())]
	}

	return string(key)
}
