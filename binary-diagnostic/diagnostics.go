package main

import (
	"errors"
	"strconv"
	"strings"
)

func BinaryFromString(payload string) ([]string, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	bins := strings.Split(payload, "\n")
	for i := 0; i < len(bins); i++ {
		err := validateBinString(bins[i])
		if err != nil {
			return nil, err
		}
	}
	padBinStrings(bins)
	return bins, nil
}

func CalculatePowerConsumption(bins []string) int {
	gammaStr := ""
	epsilonStr := ""
	if len(bins) == 0 {
		return 0
	}
	maxLen := len(bins[0])
	for i := 0; i < maxLen; i++ {
		balance := 0
		for j := 0; j < len(bins); j++ {
			if bins[j][i] == '0' {
				balance++
			} else {
				balance--
			}
		}
		if balance > 0 {
			gammaStr += "0"
		} else {
			gammaStr += "1"
		}
	}

	for i := 0; i < len(gammaStr); i++ {
		if gammaStr[i] == '0' {
			epsilonStr += "1"
		} else {
			epsilonStr += "0"
		}
	}

	gamma, _ := strconv.ParseInt(gammaStr, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonStr, 2, 64)
	return int(gamma * epsilon)
}

func CalculateLifeSupportRating(bins []string) (int, error) {
	oxygen, err := calculateOxygenGeneratorRating(bins)
	if err != nil {
		return -1, err
	}
	co2, err := calculateCO2ScrubberRating(bins)
	if err != nil {
		return -1, err
	}
	return oxygen * co2, nil
}

func padBinStrings(bins []string) {
	maxLen := 0
	for i := 0; i < len(bins); i++ {
		if len(bins[i]) > maxLen {
			maxLen = len(bins[i])
		}
	}
	for i := 0; i < len(bins); i++ {
		padLen := maxLen - len(bins[i])
		bins[i] = strings.Repeat("0", padLen) + bins[i]
	}
}

func validateBinString(bin string) error {
	for i := 0; i < len(bin); i++ {
		if bin[i] != '0' && bin[i] != '1' {
			return errors.New("invalid binary string")
		}
	}
	return nil
}

func filterBinStrings(bins []string, position int, bit byte) []string {
	filteredBins := make([]string, 0)
	for _, bin := range bins {
		if bin[position] == bit {
			filteredBins = append(filteredBins, bin)
		}
	}
	return filteredBins
}

func calculateOxygenGeneratorRating(bins []string) (int, error) {
	if len(bins) == 0 {
		return 0, nil
	}
	maxLen := len(bins[0])
	for i := 0; i < maxLen; i++ {
		balance := 0
		for j := 0; j < len(bins); j++ {
			if bins[j][i] == '0' {
				balance++
			} else {
				balance--
			}
		}
		if balance > 0 {
			bins = filterBinStrings(bins, i, '0')
		} else {
			bins = filterBinStrings(bins, i, '1')
		}
		if len(bins) == 1 {
			binVal, _ := strconv.ParseInt(bins[0], 2, 64)
			return int(binVal), nil
		}
	}
	return -1, errors.New("did not find bit string")
}

func calculateCO2ScrubberRating(bins []string) (int, error) {
	if len(bins) == 0 {
		return 0, nil
	}
	maxLen := len(bins[0])
	for i := 0; i < maxLen; i++ {
		balance := 0
		for j := 0; j < len(bins); j++ {
			if bins[j][i] == '0' {
				balance++
			} else {
				balance--
			}
		}
		if balance > 0 {
			bins = filterBinStrings(bins, i, '1')
		} else {
			bins = filterBinStrings(bins, i, '0')
		}
		if len(bins) == 1 {
			binVal, _ := strconv.ParseInt(bins[0], 2, 64)
			return int(binVal), nil
		}
	}
	return -1, errors.New("did not find bit string")
}
