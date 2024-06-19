package main

import (
	"errors"
	"io"
	"mime/multipart"
	"strings"
)

// Perform 3 way merge on indices of original, a, and b
func Merge(original []string, a []string, b []string) ArrayMergeResult {
	mergedFromA := make([]int, 0, 1)
	mergedFromB := make([]int, 0, 1)
	originalLength := len(original)
	aLength := len(a)
	bLength := len(b)
	maxLength := max(originalLength, bLength, aLength)

	for i := 0; i < maxLength; i++ {
		if i < originalLength && i < bLength && i < aLength {
			// index is valid for original, a, and b version
			if original[i] == a[i] && original[i] == b[i] {
				// No change was made to line. Merge from both a and b
				mergedFromA = append(mergedFromA, i)
				mergedFromB = append(mergedFromB, i)
			} else if original[i] == a[1] {
				// The line only changed in b
				mergedFromB = append(mergedFromB, i)
			} else if original[i] == b[i] {
				// The line was only changed in a
				mergedFromA = append(mergedFromA, i)
			} else if a[i] == b[i] {
				// The same change was made to the line in a and b
				mergedFromA = append(mergedFromA, i)
				mergedFromB = append(mergedFromB, i)
			}
		} else if i < originalLength && i < bLength {
			// Index is invalid for a. Merge from b
			mergedFromB = append(mergedFromB, i)
		} else if i < originalLength && i < aLength {
			// Index is invalid for b. Merge from a
			mergedFromA = append(mergedFromA, i)
		} else if i < bLength && i < aLength {
			// Index is invalid for original.
			if a[i] == b[i] {
				// Same change was made in a and b
				mergedFromA = append(mergedFromA, i)
				mergedFromB = append(mergedFromB, i)
			}
		} else if i < bLength {
			// Index is only valid for b
			mergedFromB = append(mergedFromB, i)
		} else if i < aLength {
			// Index is only valid for a
			mergedFromA = append(mergedFromA, i)
		}
	}

	return ArrayMergeResult{original: original, a: a, b: b, mergedFromA: mergedFromA, mergedFromB: mergedFromB}
}

func readLine(file multipart.File) (*string, error) {
	line := ""

	buffer := make([]byte, 1)
	_, err := file.Read(buffer)
	for err == nil {
		c := string(buffer)

		line += c
		if c == "\n" {
			break
		}
		_, err = file.Read(buffer)
	}

	if err == io.EOF && len(line) > 0 {
		return &line, nil
	} else if err != nil {
		return nil, err
	}

	n := len(line)

	if n > 0 && line[n-1:] != "\n" {
		line += "\n"
	}

	return &line, nil
}

// 3 way merge uploaded files line by line into a writer stream
func MergeFromMultipartFileIntoWriter(original multipart.File, a multipart.File, b multipart.File, writer io.Writer) error {
	for {
		oLine, oErr := readLine(original)
		aLine, aErr := readLine(a)
		bLine, bErr := readLine(b)
		aConflictMarker := strings.Repeat("<", 11) + " Change from version a"
		bConflictMarker := strings.Repeat(">", 11) + " Change from version b"

		// return unexpected errors
		if oErr != nil && oErr != io.EOF {
			return oErr
		}

		if aErr != nil && aErr != io.EOF {
			return aErr
		}

		if bErr != nil && bErr != io.EOF {
			return bErr
		}

		if bErr == nil && oErr == nil && aErr == nil {
			// All files were read properly
			if *oLine == *aLine && *oLine == *bLine {
				// No changes were made
				writer.Write([]byte(*oLine))
			} else if *oLine == *aLine {
				// Only b was changed
				writer.Write([]byte(*bLine))
			} else if *oLine == *bLine {
				// Only a was changed
				writer.Write([]byte(*aLine))
			} else if *aLine == *bLine {
				// The same change was made in a and b
				writer.Write([]byte(*aLine))
			} else {
				// Conflict
				writer.Write([]byte(aConflictMarker + "\n" + *aLine + bConflictMarker + "\n" + *bLine))
			}
		} else if bErr == nil && oErr == nil {
			// Only consider change in b
			writer.Write([]byte(*bLine))
		} else if aErr == nil && oErr == nil {
			// Only consider change in a
			writer.Write([]byte(*aLine))
		} else if aErr == nil && bErr == nil {
			// Both versions are valid and have gone farther than the original in terms of lines
			if *aLine == *bLine {
				writer.Write([]byte(*aLine))
			} else {
				// Conflict
				writer.Write([]byte(aConflictMarker + "\n" + *aLine + bConflictMarker + "\n" + *bLine))
			}
		} else if aErr == nil {
			// Only a is valid
			writer.Write([]byte(*aLine))
		} else if bErr == nil {
			// Only b is valid
			writer.Write([]byte(*bLine))
		} else if aErr != nil && bErr != nil && oErr != nil {
			break
		}
	}

	return nil
}

// Decides what kind of merge is possible, if one is possible, or if it will cause a conflict.
// At least one of the 3 lines must be non-nil.
// This function is meant to provide the standards for deciding how to merge 3 lines of input using a 3-way merge.
func GetMergeAction(line int, original *string, v1 *string, v2 *string) (*MergeAction, error) {
	if original == nil && v1 == nil && v2 == nil {
		return nil, errors.New("missing line from original, v1, and v2 in 3-way merge")
	}

	var action MergeAction

	if original != nil && v1 != nil && v2 != nil {
		if *original == *v1 && *original == *v2 {
			// No change
			action = MergeAction{Reason: NO_CHANGE_MERGE_REASON, Source: ORIGINAL_MERGE_SOURCE, Line: line}
		} else if *original == *v1 {
			// Only V2 changed
			action = MergeAction{Reason: ONLY_ONE_CHANGE_MERGE_REASON, Source: V2_MERGE_SOURCE, Line: line}
		} else if *original == *v2 {
			// Only V1 changed
			action = MergeAction{Reason: ONLY_ONE_CHANGE_MERGE_REASON, Source: V1_MERGE_SOURCE, Line: line}
		} else if *v1 == *v2 {
			// The same change was made
			action = MergeAction{Reason: SAME_CHANGE_MERGE_REASON, Source: V1_MERGE_SOURCE, Line: line}
		} else {
			action = MergeAction{Reason: CONFLICT_MERGE_REASON, Line: line}
		}
	} else if original != nil && v1 != nil {
		// Only V2 is invalid

		if *original == *v1 {
			// This means that the line, which exists in the original, was removed in V2
			action = MergeAction{Reason: LINE_REMOVED_IN_ONE_MERGE_REASON, Source: V2_MERGE_SOURCE, Line: line}
		} else {
			// A change was made to this line in V1, and simultaneously removed in V2. We cannot decided which is right
			action = MergeAction{Reason: CONFLICT_MERGE_REASON, Line: line}
		}
	} else if original != nil && v2 != nil {
		// Only V1 is invalid

		if *original == *v2 {
			// This means that the line, which exists in the original, was removed in V1
			action = MergeAction{Reason: LINE_REMOVED_IN_ONE_MERGE_REASON, Source: V1_MERGE_SOURCE, Line: line}
		} else {
			// A change was made to this line in V2, and simultaneously removed in V1. We cannot decided which is right
			action = MergeAction{Reason: CONFLICT_MERGE_REASON, Line: line}
		}
	} else if v1 != nil && v2 != nil {
		// V1 and V2 have surpassed the lines in original
		if *v1 == *v2 {
			action = MergeAction{Reason: SAME_CHANGE_MERGE_REASON, Source: V1_MERGE_SOURCE, Line: line}
		} else {
			action = MergeAction{Reason: CONFLICT_MERGE_REASON, Line: line}
		}
	} else if v1 != nil {
		// Only V1 is valid. It is the correct choice by default
		action = MergeAction{Reason: ONLY_ONE_VALID_MERGE_REASON, Source: V1_MERGE_SOURCE, Line: line}
	} else if v2 != nil {
		// Only V2 is valid. It is the correct choice by default
		action = MergeAction{Reason: ONLY_ONE_VALID_MERGE_REASON, Source: V2_MERGE_SOURCE, Line: line}
	} else {
		// The line exists only in the original. We will not keep it
		action = MergeAction{Reason: LINE_REEMOVED_IN_BOTH_MERGE_REASON, Source: V1_MERGE_SOURCE, Line: line}
	}

	return &action, nil
}

// Get the actions needed to merge 3 streams of file input from a multipart request
func GetMergeActionsFromMultipartFiles(original multipart.File, v1 multipart.File, v2 multipart.File) ([]*MergeAction, error) {
	actions := make([]*MergeAction, 0, 10)
	var line int = 0
	for {
		oLine, oErr := readLine(original)
		v1Line, v1Err := readLine(v1)
		v2Line, v2Err := readLine(v2)

		// return unexpected errors
		if oErr != nil && oErr != io.EOF {
			return nil, oErr
		}

		if v1Err != nil && v1Err != io.EOF {
			return nil, v1Err
		}

		if v2Err != nil && v2Err != io.EOF {
			return nil, v2Err
		}

		if v1Err != nil && v2Err != nil && oErr != nil {
			break
		}

		action, err := GetMergeAction(line, oLine, v1Line, v2Line)

		if err != nil {
			return nil, err
		}

		actions = append(actions, action)
		line++
	}

	return actions, nil
}
