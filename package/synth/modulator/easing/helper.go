package easing

// NormalisePosition returns the normalized [0-1] position for currPos, between startPos and endPos
func NormalisePosition(startPos, endPos, currPos uint) float64 {
	return float64(currPos-startPos) / float64(endPos-startPos)
}

// DenormalizePosition returns the value of currVal in the startVal and endVal range
func DenormalizePosition(currVal float64, startVal, endVal uint) uint {
	return uint(currVal*float64(endVal-startVal) + float64(startVal))
}
