import assert from "assert";
import { MergeConflict, MergeConflictReason, MergePresentation, MergeSource } from "./merge_presentation";

class MergeResult {
    blob: Blob;
    milliseconds: number;
    constructor(blob: Blob, milliseconds: number) {
        this.blob = blob;
        this.milliseconds = milliseconds;
    }
}

class Merger {
    /**
     * Get the merged version obtained from performing a 3-way merge
     * @param original 
     * @param v1 
     * @param v2 
     * @returns 
     */
    async merge(original: File, v1: File, v2: File): Promise<MergeResult> {
        const formData = new FormData();
        formData.append("original_file", original);
        formData.append("v1_file", v1);
        formData.append("v2_file", v2);
        const start = Date.now();
        const response = await fetch("http://localhost:8000/merge",
            {
                method: "POST",
                body: formData,
                // headers: { 'Content-Type': 'multipart/form-data' }
            }
        );
        const end = Date.now();


        return new MergeResult(await response.blob(), end - start);
    }

    /**
     * Get a list of actions that need to be taken in order to perform a 3-way merge
     * @param original 
     * @param v1 
     * @param v2 
     * @returns 
     */
    async getMergeActions(original: File, v1: File, v2: File): Promise<MergePresentation> {
        const formData = new FormData();
        formData.append("original_file", original);
        formData.append("v1_file", v1);
        formData.append("v2_file", v2);
        const start = Date.now();
        const response = await fetch("http://localhost:8000/merge/actions",
            {
                method: "POST",
                body: formData,
                // headers: { 'Content-Type': 'multipart/form-data' }
            }
        );
        const end = Date.now();


        let presentation: MergePresentation = JSON.parse(await response.text(),);
        presentation.totalTransferTimeInNanoseconds = (end - start) * 1000 * 1000;
        return presentation;
    }

    // Deduce the kind of conflict between 2 versions of some original line in a 3-way merge
    deduceConflict(original: string | null, v1: string | null, v2: string | null): MergeConflict {
        assert.ok(original != null || v1 != null || v2 != null); // At least one version must be non-null
        assert.ok(v1 == null || v2 == null || v1 != v2); // A conflict is not possible if v1 and v2 are equal
        assert.ok((v1 == null && v2 != null) || (v1 != null && v2 == null) || (v1 != null && v2 != null)); // A conflict is not possible if the line was removed in v1 and v2

        let conflict: MergeConflict;
        if (v1 == null) {
            // Line was removed in v1 and changed in v2
            conflict = new MergeConflict(MergeConflictReason.LINE_REMOVED_FROM_ONE_AND_CHANGED_IN_OTHER, MergeSource.V2);
        } else if (v2 == null) {
            // Line was removed in v2 and changed in v1
            conflict = new MergeConflict(MergeConflictReason.LINE_REMOVED_FROM_ONE_AND_CHANGED_IN_OTHER, MergeSource.V1);
        } else {
            // Line was changed in v1 and v2
            conflict = new MergeConflict(MergeConflictReason.BOTH_CHANGED, MergeSource.V1);
        }

        return conflict;
    }
}

export { Merger, MergeResult };