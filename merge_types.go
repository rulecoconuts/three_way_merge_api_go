package main

type ArrayMergeResult struct {
	a           []string
	b           []string
	original    []string
	mergedFromA []int
	mergedFromB []int
}

type MergeReason string
type MergeSource string

// MERGE ACTION REASONS
const (
	NO_CHANGE_MERGE_REASON             MergeReason = "NO_CHANGE"
	ONLY_ONE_CHANGE_MERGE_REASON       MergeReason = "ONLY_ONE_CHANGE"
	LINE_REMOVED_IN_ONE_MERGE_REASON   MergeReason = "LINE_REMOVED_IN_ONE"
	LINE_REEMOVED_IN_BOTH_MERGE_REASON MergeReason = "LINE_REMOVED_IN_BOTH"
	SAME_CHANGE_MERGE_REASON           MergeReason = "SAME_CHANGE"
	ONLY_ONE_VALID_MERGE_REASON        MergeReason = "ONLY_ONE_VALID"
	CONFLICT_MERGE_REASON              MergeReason = "CONFLICT"
)

// MERGE ACTION SOURCES
const (
	ORIGINAL_MERGE_SOURCE MergeSource = "ORIGINAL"
	V1_MERGE_SOURCE       MergeSource = "V1" // From the first changed version of the original file
	V2_MERGE_SOURCE       MergeSource = "V2" // From the second changed version of the original file
)

// MergeAction details a line that a 3-way merge action was performed on.
// If the action was not a conflict, it will include a reason for the merge and the final source that the merged line was taken from.
// If the action was a conflict, source will be nil, and reason will be "CONFLICT"
type MergeAction struct {
	line   int
	reason MergeReason
	source MergeSource
}
