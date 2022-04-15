package main

// Given x amount, implement a function to schedule withdrawals from the mpesa spread across multiple days.
// The maximum amount that can be withdrawn from a single day is KES 300,000 and the maximum amount per transaction is
// KES 150,000.
//
// Consider that the minimum amount that can be withdrawn is KES 100, if the amount is less than this,
// it should return the balance.

// Example:
// Input: 600_000
// Output: 0, [{150_000, 0}, {150_000, 0}, 150_000, 1}, {150_000, 1}]

// Example 2:
// Input: 300_099
// Output: 99, [{150_000, 0}, {150_000, 0}]

const (
	maxTxAmount = 150_000
	minTxAmount = 100
)

type schedule struct {
	Amount float64
	Day    int
}

func mpesaSplit(schedules []schedule, amount float64, day int, increment bool) (float64, []schedule) {
	if amount < minTxAmount {
		return amount, schedules
	}

	if amount > maxTxAmount {
		amount = amount - maxTxAmount
		schedules = append(schedules, schedule{Amount: maxTxAmount, Day: day})
	} else {
		schedules = append(schedules, schedule{Amount: amount, Day: day})
		amount = 0
	}

	if increment {
		day++
	}

	increment = !increment
	return mpesaSplit(schedules, amount, day, increment)
}
