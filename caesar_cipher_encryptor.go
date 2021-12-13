package main

import "fmt"

func main () {

	result := CaesarCipherEncryptor("abc",3)
	fmt.Println(result)
}

func CaesarCipherEncryptor(str string, key int) string {
	//values from 97 -> 122
	//key = n
	var str1 string
	if key > 26 {
		key = key%26
	}
	for _,v := range str {
		if v >= 97 && v <= 122 {
			if (v+int32(key)) <= 122 {
				str1 += fmt.Sprintf("%c",v+int32(key))
			} else {
				str1 += fmt.Sprintf("%c",96+(v+int32(key))%122)
            }
		}

    }
	return str1
}

