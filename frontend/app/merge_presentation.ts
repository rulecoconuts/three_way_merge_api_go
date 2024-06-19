enum MergeReason {
    NO_CHANGE = "NO_CHANGE",
    ONLY_ONE_CHANGE = "ONLY_ONE_CHANGE",
    LINE_REMOVED_IN_ONE = "LINE_REMOVED_IN_ONE",
    LINE_REEMOVED_IN_BOTH = "LINE_REMOVED_IN_BOTH",
    SAME_CHANGE = "SAME_CHANGE",
    ONLY_ONE_VALID = "ONLY_ONE_VALID",
    CONFLICT = "CONFLICT"
}

enum MergeSource {
    ORIGINAL = "ORIGINAL",
    V1 = "V1", // From the first changed version of the original file
    V2 = "V2" // From the second changed version of the original file
}

enum MergeConflictReason {
    BOTH_CHANGED = "BOTH_CHANGED",
    LINE_REMOVED_FROM_ONE_AND_CHANGED_IN_OTHER = "LINE_REMOVED_FROM_ONE_AND_CHANGED_IN_OTHER"
}

class MergeConflict {
    reason: MergeConflictReason;

    // Source will always point to an existing line if it is not null
    source: MergeSource | null;

    constructor(
        reason: MergeConflictReason,
        source: MergeSource
    ) {
        this.reason = reason;
        this.source = source;
    }
}


/**
 * MergeAction details a line that a 3-way merge action was performed on.
 *If the action was not a conflict, it will include a reason for the merge and the final source that the merged line was taken from.
 * If the action was a conflict, source will be nil, and reason will be "CONFLICT"
 */
class MergeAction {
    line: number;
    reason: MergeReason;
    // Source will always point to an existing line if it is not null
    source: MergeSource | null;

    constructor(
        line: number,
        reason: MergeReason,
        source: MergeSource
    ) {
        this.line = line;
        this.reason = reason;
        this.source = source;
    }

}
class MergePresentation {
    actions: Array<MergeAction>;
    timeSpentInMicroseconds: number;
    totalTransferTimeInNanoseconds: number = 0;

    constructor(actions: Array<MergeAction>, timeSpentInMicroseconds: number, totalTransferTimeInNanoseconds: number = 0) {
        this.actions = actions;
        this.timeSpentInMicroseconds = timeSpentInMicroseconds;
        this.totalTransferTimeInNanoseconds = totalTransferTimeInNanoseconds;
    }
}

export { MergePresentation, MergeAction, MergeSource, MergeReason, MergeConflict, MergeConflictReason };