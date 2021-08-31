package numbersequence

func findSequence(numbers []int, sequence []int) bool {

	sequenceLength := len(sequence)

	for numbersIter := 0; numbersIter < (len(numbers)); numbersIter++ {
		for sequenceIter := 0; sequenceIter < sequenceLength; {
			if sequence[sequenceIter] == numbers[numbersIter] {

				// move sequence needle by one on match
				sequenceIter++
				if numbersIter < (len(numbers) - 1) {
					// move numbers needle by one on match
					numbersIter++
				}

				if sequenceIter == sequenceLength {
					// on sequence length match return true
					return true
				}
			} else {
				// break inner loop when sequence elems don't match
				break
			}
		}
	}
	return false
}
