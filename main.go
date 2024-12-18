package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func encryptWithOffset(input string) string {
	var encrypted strings.Builder
	words := strings.Fields(input) // 按空格分割成单词

	// 遍历每个单词
	for _, word := range words {
		// 对每个字母根据位置从右边开始加密
		encryptedWord := encryptWordFromRight(word)
		encrypted.WriteString(encryptedWord + " ")
	}

	return strings.TrimSpace(encrypted.String())
}

func encryptWordFromRight(word string) string {
	var encryptedWord strings.Builder
	wordLength := len(word)

	for i, ch := range word {
		if unicode.IsLetter(ch) {
			// 计算偏移量：从右边开始 (从1开始)
			offset := wordLength - i
			// 处理大写字母和小写字母
			if unicode.IsLower(ch) {
				encryptedWord.WriteRune(shiftLetter(ch, offset))
			} else if unicode.IsUpper(ch) {
				encryptedWord.WriteRune(shiftLetter(ch, offset))
			}
		} else {
			encryptedWord.WriteRune(ch)
		}
	}
	return encryptedWord.String()
}

// 偏移字母（字母环绕）
func shiftLetter(ch rune, offset int) rune {
	var base rune
	if unicode.IsLower(ch) {
		base = 'a'
	} else {
		base = 'A'
	}

	return base + (ch-base+rune(offset))%26
}

func decryptWithOffset(input string) string {
	var decrypted strings.Builder
	words := strings.Fields(input) // 按空格分割成单词

	for _, word := range words {
		decryptedWord := decryptWordFromRight(word)
		decrypted.WriteString(decryptedWord + " ")
	}

	return strings.TrimSpace(decrypted.String())
}

func decryptWordFromRight(word string) string {
	var decryptedWord strings.Builder
	wordLength := len(word)

	for i, ch := range word {
		if unicode.IsLetter(ch) {
			offset := wordLength - i
			if unicode.IsLower(ch) {
				decryptedWord.WriteRune(reverseShiftLetter(ch, offset))
			} else if unicode.IsUpper(ch) {
				decryptedWord.WriteRune(reverseShiftLetter(ch, offset))
			}
		} else {
			decryptedWord.WriteRune(ch)
		}
	}
	return decryptedWord.String()
}

func reverseShiftLetter(ch rune, offset int) rune {
	var base rune
	if unicode.IsLower(ch) {
		base = 'a'
	} else {
		base = 'A'
	}

	return base + (ch-base-rune(offset)+26)%26
}

func main() {
	fmt.Println("1+ 2-：")
	var operationType int
	fmt.Scanln(&operationType)

	// 获取要加解密的文本，可以包含多个单词（空格分隔）
	fmt.Println("请输入+-：")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	// 根据操作类型加解密
	var result string
	if operationType == 1 {
		// 加密
		result = encryptWithOffset(text)
		fmt.Printf("+：%s\n", result)
	} else if operationType == 2 {
		// 解密
		result = decryptWithOffset(text)
		fmt.Printf("-：%s\n", result)
	} else {
		fmt.Println("无效的操作类型。请输入1或2。")
	}
}
