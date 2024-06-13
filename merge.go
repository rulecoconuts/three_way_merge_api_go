package main

import (
	"io"
	"mime/multipart"
	"strings"
)

type MergeResult struct {
	a           []string
	b           []string
	original    []string
	mergedFromA []int
	mergedFromB []int
}

type MergeSummary struct {
	aConflictPositions []int
	bConflictPositions []int
}

// Perform 3 way merge on indices of original, a, and b
func Merge(original []string, a []string, b []string) MergeResult {
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

	return MergeResult{original: original, a: a, b: b, mergedFromA: mergedFromA, mergedFromB: mergedFromB}
}

func readLine(file multipart.File) (string, error) {
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
		return line, nil
	} else if err != nil {
		return "", err
	}

	n := len(line)

	if n > 0 && line[n-1:] != "\n" {
		line += "\n"
	}

	return line, nil
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
			if oLine == aLine && oLine == bLine {
				// No changes were made
				writer.Write([]byte(oLine))
			} else if oLine == aLine {
				// Only b was changed
				writer.Write([]byte(bLine))
			} else if oLine == bLine {
				// Only a was changed
				writer.Write([]byte(aLine))
			} else if aLine == bLine {
				// The same change was made in a and b
				writer.Write([]byte(aLine))
			} else {
				// Conflict
				writer.Write([]byte(aConflictMarker + "\n" + aLine + bConflictMarker + "\n" + bLine))
			}
		} else if bErr == nil && oErr == nil {
			// Only consider change in b
			writer.Write([]byte(bLine))
		} else if aErr == nil && oErr == nil {
			// Only consider change in a
			writer.Write([]byte(aLine))
		} else if aErr == nil && bErr == nil {
			// Both versions are valid and have gone farther than the original in terms of lines
			if aLine == bLine {
				writer.Write([]byte(aLine))
			} else {
				// Conflict
				writer.Write([]byte(aConflictMarker + "\n" + aLine + bConflictMarker + "\n" + bLine))
			}
		} else if aErr == nil {
			// Only a is valid
			writer.Write([]byte(aLine))
		} else if bErr == nil {
			// Only b is valid
			writer.Write([]byte(bLine))
		} else if aErr != nil && bErr != nil && oErr != nil {
			break
		}
	}

	return nil
}
