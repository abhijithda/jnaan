package addbinary

import (
	"log"
	"strconv"
	"strings"
)

func addBinary(a string, b string) string {
	log.Printf("Given a: %s; b:%s", a, b)
	la := len(a)
	lb := len(b)
	var lr, le int
	if la < lb {
		lr = lb
		le = la
	} else {
		lr = la
		le = lb
	}
	lr++
	res := make([]string, lr)
	log.Println("Result length:", lr)
	log.Println("Equal length:", le)
	carry := 0
	i := 1
	for ; i <= le; i++ {
		d1, _ := strconv.Atoi(string(a[la-i]))
		d2, _ := strconv.Atoi(string(b[lb-i]))
		sum := carry + d1 + d2
		res[lr-i] = strconv.Itoa(sum % 2)
		carry = sum / 2
		log.Printf("Equal length - i: %d; Sum: %d; Carry: %d", i, sum, carry)
	}
	if la > lb {
		for ; i < lr; i++ {
			d1, _ := strconv.Atoi(string(a[la-i]))
			sum := carry + d1
			res[lr-i] = strconv.Itoa(sum % 2)
			carry = sum / 2
			log.Printf("i: %d; Sum: %d; Carry: %d", i, sum, carry)
		}
	} else if la < lb {
		for ; i < lr; i++ {
			d2, _ := strconv.Atoi(string(b[lb-i]))
			sum := carry + d2
			res[lr-i] = strconv.Itoa(sum % 2)
			carry = sum / 2
			log.Printf("i: %d; Sum: %d; Carry: %d", i, sum, carry)
		}
	}
	if carry != 0 {
		log.Println("Adding carry!", carry)
		res[0] = strconv.Itoa(carry)
	} else {
		res = res[1:len(res)]
	}
	log.Println("Result:", res)
	return strings.Join(res, "")
}
