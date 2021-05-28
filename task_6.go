package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math/bits"
	"strings"
)

// ScoreStr ...
func ScoreStr(InputStr string) float32 {
	characterFreq := map[byte]float32{
		'a': 8.167, 'b': 1.492, 'c': 2.782, 'd': 4.253,
		'e': 12.702, 'f': 2.228, 'g': 2.015, 'h': 6.094,
		'i': 6.094, 'j': 0.153, 'k': 0.772, 'l': 4.025,
		'm': 2.406, 'n': 6.749, 'o': 7.507, 'p': 1.929,
		'q': 0.095, 'r': 5.987, 's': 6.327, 't': 9.056,
		'u': 2.758, 'v': 0.978, 'w': 2.360, 'x': 0.150,
		'y': 1.974, 'z': 0.074, ' ': 13.000,
	}

	var scores float32
	for i := 0; i < len(InputStr); i++ {
		scores += characterFreq[InputStr[i]]
	}
	return scores
}

// BestScore ...
func BestScore(buf []float32) byte {
	var tp float32
	tp = 0
	var temp byte
	for i := 0; i < 256; i++ {
		if buf[i] >= tp {
			tp = buf[i]
			temp = byte(i)
		}
	}
	return temp
}

// SingleXor ...
func SingleXor(symbol byte, InputStr []byte) string {

	var result string
	for i := 0; i < len(InputStr); i++ {
		result += string(symbol ^ InputStr[i])
	}

	return result
}

// HammingDistance ...
func HammingDistance(str1, str2 []byte) int {
	if len(str1) != len(str2) {
		panic("Different lengths")
	}

	var res int
	for i := range str1 {
		res += bits.OnesCount8(str1[i] ^ str2[i])
	}

	return res
}

// GetLenKey ...
func GetLenKey(text string) int {
	temp := text
	KeyLength := 0
	Max := 0
	for i := 1; i < len(text)/2; i++ {
		temp = string(temp[len(text)-1]) + temp[:len(text)-1]
		Dist := 0
		for j := 0; j < len(text); j++ {
			if temp[j] == text[j] {
				Dist = Dist + 1
			}
		}
		if Dist > Max {
			Max = Dist
			KeyLength = i
		}
	}
	return KeyLength
}

//FindSingleXor ...
func FindSingleXor(InputStr []byte) byte {
	var ans string
	var buf []float32
	for i := 0; i < 256; i++ {
		ans = SingleXor(byte(i), InputStr)
		buf = append(buf, ScoreStr(ans))
	}

	return BestScore(buf)
}

// GetKeyText ...
func GetKeyText(KeyLength int, text string) string {

	LenBlock := (len(text)) / KeyLength
	Block := make([][]byte, KeyLength)

	for i := 0; i < KeyLength; i++ {
		Block[i] = make([]byte, LenBlock)
	}

	for i := 0; i < KeyLength; i++ {
		for j := 0; j < LenBlock; j++ {
			Block[i][j] = text[j*KeyLength+i]
		}
	}

	var key string
	for k := 0; k < KeyLength; k++ {
		key += string(FindSingleXor(Block[k]))
	}

	return key
}

func main() {

	file, err := ioutil.ReadFile("task6text.txt")
	if err != nil {
		log.Fatal(err)
	}
	text := strings.Split(string(file), "\n")

	var textfile string
	for i := 0; i < len(text); i++ {
		textfile += strings.TrimSpace(text[i])
	}

	byteFileText, err1 := base64.StdEncoding.DecodeString(textfile)
	if err1 != nil {
		log.Fatal(err1)
	}

	KeyValue := GetKeyText(GetLenKey(string(byteFileText)), string(byteFileText))

	var result []byte
	for i := 0; i < len(byteFileText); i++ {
		result = append(result, KeyValue[i%len(KeyValue)]^byteFileText[i])
	}

	fmt.Println(string(result))
}
