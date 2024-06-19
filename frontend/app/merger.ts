import { MergeConflict, MergePresentation } from "./merge_presentation";

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
        console.log(formData);
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
        console.log(formData);
        const response = await fetch("http://localhost:8000/merge/actions",
            {
                method: "POST",
                body: formData,
                // headers: { 'Content-Type': 'multipart/form-data' }
            }
        );

        let presentation: MergePresentation = JSON.parse(await response.text(),);
        return presentation;
    }

    deriveConflictReason(original: string | null, v1: string | null, v2: string | null): MergeConflict {
        // TODO: Implement deriveConflictReason
        throw "Unimplemented";
    }
}

export { Merger, MergeResult };