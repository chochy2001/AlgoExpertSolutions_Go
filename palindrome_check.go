package main

import "fmt"

func main (){
	result := IsPalindrome("jorge")
	//hannah, oro, (es un palindromo)
	fmt.Println(result)
}

func IsPalindrome(str string)bool{
	var word1 string
	var word2 string
	for i:= 0; i<len(str);i++{
		word1 += string(str[i])
	}
	for i:= len(str)-1; i>=0;i--{
        word2 += string(str[i])
    }
	if word1 == word2{
		return true
	}
	return false
}
