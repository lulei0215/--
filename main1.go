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
	words := strings.Fields(input) // 

	// 遍历每个单词
	for _, word := range words {
		// 对每个字母根据位置加密
		encryptedWord := encryptWord(word)
		encrypted.WriteString(encryptedWord + " ")
	}

	return strings.TrimSpace(encrypted.String())
}

// 加密单个单词：每个字母根据其在单词中的位置进行偏移
func encryptWord(word string) string {
	var encryptedWord strings.Builder
	for i, ch := range word {
		if unicode.IsLetter(ch) {
			// 计算偏移量：字母在单词中的位置 (从1开始，i+1)
			offset := i + 1
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

	// 计算偏移后的字母
	return base + (ch-base+rune(offset))%26
}

// 解密函数：根据单词中每个字母的位置来决定偏移量
func decryptWithOffset(input string) string {
	var decrypted strings.Builder
	words := strings.Fields(input) // 按空格分割成单词

	// 遍历每个单词
	for _, word := range words {
		// 对每个字母根据位置解密
		decryptedWord := decryptWord(word)
		decrypted.WriteString(decryptedWord + " ")
	}

	return strings.TrimSpace(decrypted.String())
}

// 解密单个单词：每个字母根据其在单词中的位置进行反向偏移
func decryptWord(word string) string {
	var decryptedWord strings.Builder
	for i, ch := range word {
		if unicode.IsLetter(ch) {
			// 计算偏移量：字母在单词中的位置 (从1开始，i+1)
			offset := i + 1
			// 处理大写字母和小写字母
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

// 反向偏移字母（字母环绕）
func reverseShiftLetter(ch rune, offset int) rune {
	var base rune
	if unicode.IsLower(ch) {
		base = 'a'
	} else {
		base = 'A'
	}

	// 计算反向偏移后的字母
	return base + (ch-base-rune(offset)+26)%26
}

func main() {
	// 提示用户输入操作类型（加密或解密）
	fmt.Println("请输入操作类型 (1为加密，2为解密)：")
	var operationType int
	fmt.Scanln(&operationType)

	// 获取要加解密的文本，可以包含多个单词（空格分隔）
	fmt.Println("请输入要加解密的文本：")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	// 根据操作类型加解密
	var result string
	if operationType == 1 {
		// 加密
		result = encryptWithOffset(text)
		fmt.Printf("加密后的文本：%s\n", result)
	} else if operationType == 2 {
		// 解密
		result = decryptWithOffset(text)
		fmt.Printf("解密后的文本：%s\n", result)
	} else {
		fmt.Println("无效的操作类型。请输入1或2。")
	}
}
