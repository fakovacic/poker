package poker

func CompareResults(xResult, yResult Result) Compare {
	switch {
	case xResult.Score() < yResult.Score():
		return High
	case xResult.Score() > yResult.Score():
		return Low
	case xResult.Score() == yResult.Score():
		return Equal
	default:
		return ""
	}
}
