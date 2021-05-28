package main

import (
	"encoding/hex"
	"fmt"
	"log"
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
	var tmp float32 = 0
	var res byte
	for i := 0; i < 256; i++ {
		if buf[i] >= tmp {
			tmp = buf[i]
			res = byte(i)
		}
	}
	return res
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
	str1 := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	res, err1 := hex.DecodeString(str1)
	if err1 != nil {
		log.Fatal(err1)
	}

	var answ string
	var buf []float32
	for i := 0; i < 256; i++ {
		answ = SingleXor(byte(i), res)
		buf = append(buf, ScoreStr(answ))
	}

	fmt.Println(SingleXor(BestScore(buf), res))

}
