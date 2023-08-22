package core

import (
	ahocorasick "github.com/petar-dambovaliev/aho-corasick"
)

func PmEvaluate(matcher ahocorasick.AhoCorasick, value string, isCapture bool) (bool, []string) {
	iter := matcher.Iter(value)

	if isCapture {
		// Not capturing so just one match is enough.
		return iter.Next() != nil, nil
	}

	var numMatches int
	var matchedValues []string
	for {
		m := iter.Next()
		if m == nil {
			break
		}

		//tx.CaptureField(numMatches, value[m.Start():m.End()])
		matchedValues = append(matchedValues, value[m.Start():m.End()])
		numMatches++
		if numMatches == 10 {
			return true, matchedValues
		}
	}

	return numMatches > 0, matchedValues
}
