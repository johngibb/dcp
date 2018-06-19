package prob

// This problem was asked by Facebook.
//
// Given a array of numbers representing the stock prices of a company in
// chronological order, write a function that calculates the maximum profit you
// could have made from buying and selling that stock once. You must buy before
// you can sell it.
//
// For example, given [9, 11, 8, 5, 7, 10], you should return 5, since you could
// buy the stock at 5 dollars and sell it at 10 dollars.

func MaxProfit(prices []int) int {
	buy := prices[0]
	profit := 0
	for _, p := range prices[1:] {
		profit = max(profit, p-buy) // what if we sold now?
		buy = min(buy, p)           // what if we bought now?
	}
	return profit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
