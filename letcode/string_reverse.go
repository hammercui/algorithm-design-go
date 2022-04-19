package main

import "log"

/**
字符串翻转
 */

func main()  {
	reverse()
}

var str = "hello,world"

func reverse()   {
	s := []byte(str)
	for i,j :=0,len(s)-1; i<j; i,j = i+1,j-1  {
		s[i],s[j] = s[j],s[i]
	}
	log.Println(string(s))
}

func reverse2()   {

}


