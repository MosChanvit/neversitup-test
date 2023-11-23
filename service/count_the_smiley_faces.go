package service

import "strings"

func CountSmileys(arr []string) int {
	count := 0

	for _, face := range arr {
		if isValidSmiley(face) {
			count++
		}
	}

	return count
}

func isValidSmiley(face string) bool {
	//eyes, nose, and mouth
	validEyes := ":;"
	validNose := "-~"
	validMouth := ")D"

	// Check  length
	if len(face) < 2 || len(face) > 3 {
		return false
	}

	// Check eye
	if strings.Index(validEyes, string(face[0])) == -1 {
		return false
	}

	// Check  mouth
	if strings.Index(validMouth, string(face[len(face)-1])) == -1 {
		return false
	}

	// check  nose
	if len(face) == 3 && strings.Index(validNose, string(face[1])) == -1 {
		return false
	}

	return true
}
