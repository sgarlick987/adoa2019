package main

import (
	"log"
	"math"
)

/**
a sevenish number is defined as a value that is a unique power of 7 or a sum of two unique powers of 7
this function is to find the sevenish number at n in the sequence (eg 1=1, 2=7, 3=8, 4=49, 5=50, 6=51 ...)
 */
func sevenishNumber(num uint64) uint64 {
	if num < 1 {
		log.Fatal("cannot get sevenish number in sequence that is less than 1")
	}

	/**
	short circuit our uniques seed values
	*/
	if num == 1 {
		return 1
	}
	if num == 2 {
		return 7
	}

	/**
	seed uniques with first 2 sevenish numbers
	set max number of sevenish we can handle with our initial 2 uniques which can cover 3 values in the sequence
	*/
	uniques := []uint64{1, 7}
	maxNum := uint64(3)

	for num > maxNum {
		/**
		if the num requested is maxNum + 1 then get the next power of 7
		which is the length of our uniques slice and return it
		*/
		if num == maxNum+1 {
			return uint64(math.Pow(7, float64(len(uniques))))
		}

		/**
		until maxNum is higher than num append the next unique and update the max number in the sequence we can handle
		which is maxNum + how many uniques are in the updated list
		*/
		uniques = append(uniques, uint64(math.Pow(7, float64(len(uniques)))))
		maxNum = maxNum + uint64(len(uniques))
	}

	/**
	num 1, 2 and any num in sequence that is a unique power of 7 are handled at this point
	we need to handle sevenish values that are sums of unique powers of 7
	first we get the last unique value in our slice which we know will be the first power of 7 we are going to add
	the next unique value we need to add we get by starting at the last unique power of 7 and subtract
	maxNum - num + 1 so for example if maxNum == num then we just subtract 1 from the end of the slice and use that
	*/
	return uniques[len(uniques)-1] + uniques[uint64(len(uniques))-1-(maxNum-num+1)]
}

func main() {
	log.Println(sevenishNumber(1))
	log.Println(sevenishNumber(2))
	log.Println(sevenishNumber(3))
	log.Println(sevenishNumber(4))
	log.Println(sevenishNumber(5))
	log.Println(sevenishNumber(6))
	log.Println(sevenishNumber(7))
	log.Println(sevenishNumber(8))
	log.Println(sevenishNumber(9))
	log.Println(sevenishNumber(10))
	log.Println(sevenishNumber(11))
	log.Println(sevenishNumber(13))
	log.Println(sevenishNumber(14))
	log.Println(sevenishNumber(15))
	log.Println(sevenishNumber(16))
	log.Println(sevenishNumber(17))
	log.Println(sevenishNumber(18))
	log.Println(sevenishNumber(19))
	log.Println(sevenishNumber(20))
	log.Println(sevenishNumber(21))
	log.Println(sevenishNumber(22))
	log.Println(sevenishNumber(23))
}
