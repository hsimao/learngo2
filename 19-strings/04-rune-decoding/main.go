package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	const text = `Galaksinin Batı Sarmal Kolu'nun bir ucunda, haritası bile çıkarılmamış ücra bir köşede, gözlerden uzak, küçük ve sarı bir güneş vardır.
    Bu güneşin yörüngesinde, kabaca yüz kırksekiz milyon kilometre uzağında, tamamıyla önemsiz ve mavi-yeşil renkli, küçük bir gezegen döner.
    Gezegenin maymun soyundan gelen canlıları öyle ilkeldir ki dijital kol saatinin hâlâ çok etkileyici bir buluş olduğunu düşünürler.`

	fmt.Println()

	word := []byte("öyle")
	fmt.Printf("%s = % [1]x\n", word)

	// var size int
	// for i := range string(word) {
	// 	if i > 0 {
	// 		size = i
	// 		break
	// 	}
	// }

	// 第一個字母轉大寫
	_, size := utf8.DecodeRune(word)
	copy(word[:size], bytes.ToUpper(word[:size]))

	fmt.Printf("%s = % [1]x\n", word)
}
