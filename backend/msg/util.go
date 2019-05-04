package msg

import (
	"fmt"
	"math"
)

func paginate(length, page, perPage int) (current int, last int, start int, end int) {
	fmt.Printf("length: %d, page: %d, perPage: %d\n", length, page, perPage)
	if length < perPage {
		current = 1
		last = 1
		start = 0
		end = length
	} else {
		last = int(math.Ceil(float64(length) / float64(perPage)))
		current = page
		if current*perPage > length {
			current = last
		}
		start = (current - 1) * perPage
		end = start + perPage
		if end > length {
			end = length
		}
	}
	fmt.Printf("current: %d, last: %d, start: %d, end: %d\n", current, last, start, end)
	return
}

// eof
