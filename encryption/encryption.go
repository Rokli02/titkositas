package encryption

import (
	"fmt"
	"math/rand"
	"strings"
	"titkositas/assets"
)

func Encode(text string) string {
	codeBuilder := strings.Builder{}
	for i, letter := range text {
		letterArray, hasLetter := encoderMap[letter]
		if !hasLetter {
			codeBuilder.WriteRune(letter)

			continue
		}

		codeBuilder.WriteRune(letterArray[i%len(letterArray)])
	}

	return codeBuilder.String()
}

func Decode(code string) string {
	textBuilder := strings.Builder{}

	for _, character := range code {
		letter, hasLetter := decoderMap[character]
		if !hasLetter {
			textBuilder.WriteRune(character)

			continue
		}

		textBuilder.WriteRune(letter)
	}

	return textBuilder.String()
}

func InitializeCoderMaps(useAdditionals bool) {
	if err := checkForRedundantLetters(); err != nil {
		panic(err)
	}

	sizeOfSource := len(assets.EncryptionTable.Source)

	sourceSlice := make([]string, sizeOfSource)
	targetSlice := make([]string, sizeOfSource)

	encoderMap = make(map[rune][]rune)
	decoderMap = make(map[rune]rune)

	copy(sourceSlice, assets.EncryptionTable.Source)
	copy(targetSlice, assets.EncryptionTable.Source)

	// Mindegyik source elemhez hozzárendelni egy target-et
	// Ha saját magát választaná, akkor X alkalommal újrapróbálja, ha akkor sincs eredmény, akkor veszi vagy az első, vagy az utolsó elemet
	for _, sourceLetter := range sourceSlice {
		targetIndex := findTargetIndex(sourceLetter, targetSlice)
		if targetIndex == -1 {
			panic(fmt.Sprintf("Couldn't pair '%s' letter to any of the target letters", sourceLetter))
		}

		decoderMap[[]rune((targetSlice[targetIndex]))[0]] = []rune(sourceLetter)[0]
		encoderMap[[]rune(sourceLetter)[0]] = []rune(targetSlice[targetIndex])

		targetSlice[targetIndex] = targetSlice[len(targetSlice)-1]
		targetSlice = targetSlice[:len(targetSlice)-1]
	}

	// A 'targetSlice'-ban maradt betűket hozzácsapni 'random' betűkhöz
	if useAdditionals {
		targetSlice = append(targetSlice, assets.EncryptionTable.Additional...)

		for i := len(targetSlice) - 1; i >= 0; i-- {
			nextIndex := rand.Intn(sizeOfSource)

			letters := encoderMap[[]rune(sourceSlice[nextIndex])[0]]
			letters = append(letters, []rune(targetSlice[i])[0])
			encoderMap[[]rune(sourceSlice[nextIndex])[0]] = letters

			decoderMap[[]rune(targetSlice[i])[0]] = []rune(sourceSlice[nextIndex])[0]

			targetSlice = targetSlice[:i]
		}
	}
}

func findTargetIndex(sourceLetter string, targetSlice []string) int {
	var nextIndex int
	var hasValue bool

	for tries := 0; tries < maxMixTry; tries++ {
		nextIndex = rand.Intn(len(targetSlice))

		if sourceLetter != targetSlice[nextIndex] {
			hasValue = true

			break
		}
	}

	if !hasValue {
		if sourceLetter != targetSlice[0] {
			return 0
		}

		if sourceLetter != targetSlice[len(targetSlice)-1] {
			return len(targetSlice) - 1
		}

		return -1
	}

	return nextIndex
}

func checkForRedundantLetters() error {
	var checkMap map[string]bool = make(map[string]bool)
	defer func() {
		clear(checkMap)
		checkMap = nil
	}()

	for _, letter := range assets.EncryptionTable.Source {
		if _, hasLetter := checkMap[letter]; hasLetter {
			return fmt.Errorf("letter '%s' is redundant in the 'source' list", letter)
		}

		checkMap[letter] = true
	}

	for _, letter := range assets.EncryptionTable.Additional {
		if _, hasLetter := checkMap[letter]; hasLetter {
			return fmt.Errorf("letter '%s' is redundant in the 'additional' list", letter)
		}

		checkMap[letter] = true
	}

	return nil
}

func GetDecoderMap() string {
	res := strings.Builder{}

	for key, value := range decoderMap {
		res.WriteString(fmt.Sprintf("\t[%s]: %s\n", string(key), string(value)))
	}

	return res.String()
}

func GetEncoderMap() string {
	res := strings.Builder{}

	for key, value := range encoderMap {
		res.WriteString(fmt.Sprintf("\t[%s]: [", string(key)))

		for i, letter := range value {
			if i != 0 {
				res.WriteString(", ")
			}
			res.WriteString(string(letter))
		}

		res.WriteString("]\n")
	}

	return res.String()
}
