package besttimetobuy

import "log"

func maxProfit(prices []int) int {
	log.Println("Given: ", prices)
	profit := 0

	if len(prices) < 2 {
		return profit
	}

	for i := 0; i != len(prices); {
		log.Println("i:", i)
		bi := -1
		si := -1
		// Find buy day
		for ; i < len(prices)-1; i++ {
			if prices[i] < prices[i+1] {
				bi = i
				break
			}

		}
		// Find sell day
		// si = bi + 1
		for ; i < len(prices)-1; i++ {
			if prices[i] > prices[i+1] {
				si = i
				break
			}
		}
		if si == -1 && i != len(prices) && bi != -1 && prices[bi] < prices[i] {
			si = i
		}
		log.Printf("Day to buy: %d and sell: %d\n", bi, si)
		if bi == -1 || si == -1 {
			log.Println("Couldn't find anymore profit.")
			break
		}
		profit += prices[si] - prices[bi]
		log.Printf("Profit (sell %d - buy %d): %d\n", prices[si], prices[bi], profit)
	}

	log.Println("Total profit:", profit)
	return profit
}
