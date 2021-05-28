package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

//ScoreStr ....
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

func main() {

	file, err := os.Open("encrypted_text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var resultStr []string
	var buf2 []float32

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str1 := scanner.Text()
		res, err1 := hex.DecodeString(str1)

		if err1 != nil {
			log.Fatal(err1)
		}

		var ans string
		var buf []float32
		for i := 0; i < 256; i++ {
			ans = SingleXor(byte(i), res)
			buf = append(buf, ScoreStr(ans))
		}
		resultStr = append(resultStr, SingleXor(BestScore(buf), res))
		buf2 = append(buf2, ScoreStr(SingleXor(BestScore(buf), res)))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(resultStr[BestScore(buf2)])
}
