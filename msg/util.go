package msg

func paginate(length, page, perPage int) (currentPage int, lastPage int, start int, end int) {
	if length < perPage {
		currentPage = 1
		lastPage = 1
		start = 1
		end = length
	} else {
		lastPage = length / perPage
		currentPage = page
		if currentPage*perPage > length {
			currentPage = lastPage
		}
		start = (currentPage - 1) * perPage
		end = start + perPage
		if end > length {
			end = length
		}
	}
	return
}

// eof
