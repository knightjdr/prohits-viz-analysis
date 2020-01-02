package dotplot

func defineScoreConverter(d *Dotplot, numColors int) func(score float64) int {
	maxIndex := defineMaxIndex(numColors)

	if d.ScoreType == "gte" {
		return func(score float64) int {
			if score >= d.PrimaryFilter {
				return maxIndex
			}
			if score < d.PrimaryFilter && score >= d.SecondaryFilter {
				return maxIndex / 2
			}
			return maxIndex / 4
		}
	}
	return func(score float64) int {
		if score <= d.PrimaryFilter {
			return maxIndex
		}
		if score > d.PrimaryFilter && score <= d.SecondaryFilter {
			return maxIndex / 2
		}
		return maxIndex / 4
	}
}

func defineMaxIndex(numColors int) int {
	if numColors%2 == 0 {
		return numColors
	}
	return numColors - 1
}
