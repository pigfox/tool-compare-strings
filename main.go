package main

import "log"

func main() {
	s1 := string1()
	s2 := string2()
	l1 := len(s1)
	l2 := len(s2)
	if l1 != l2 {
		log.Println("len of s1", l1, "len of s2", l2)
		log.Fatal("Exit")
	} else {
		log.Println("both strings are", l1, "characters longs")
	}
	err := 0
	for i := 0; i < l1; i++ {
		if s1[i] != s2[i] {
			log.Println("s1=", s1[i], "s2=", s2[i], "@ char", i)
			err++
		}
	}
	if err != 0 {
		log.Println("num errors", err)
		log.Fatal("Exit")
	}

	log.Println("Strings are equal")
}
