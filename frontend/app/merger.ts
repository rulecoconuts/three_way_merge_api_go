class MergeResult {
    blob: Blob;
    milliseconds: number;
    constructor(blob: Blob, milliseconds: number) {
        this.blob = blob;
        this.milliseconds = milliseconds;
    }
}

class Merger {
    async merge(original: File, a: File, b: File): Promise<MergeResult> {
        const formData = new FormData();
        formData.append("original_file", original);
        formData.append("a_file", a);
        formData.append("b_file", b);
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
}

export { Merger, MergeResult };