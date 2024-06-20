import assert from "assert";
import { MergeAction, MergeConflict, MergeConflictReason, MergeSource } from "./merge_presentation";
import { Merger } from "./merger";
import { useState } from "react";
import LabelledBox from "./labelled_box";
import LineRemovedMessage from "./line_removed_message";

interface FileMergeActionConflictViewProps {
    original: string | null,
    v1: string | null,
    v2: string | null,
    action: MergeAction
}

export default function FileMergeActionConflictView({ original, v1, v2, action }: FileMergeActionConflictViewProps) {
    assert.ok(original != null || v1 != null || v2 != null); // At list one version of the line must be non-null
    let [conflict, setConflict] = useState<MergeConflict>((new Merger()).deduceConflict(original, v1, v2));

    let v1Content: string | null = null;
    let v2Content: string | null = null;

    if (conflict.reason == MergeConflictReason.BOTH_CHANGED) {
        v1Content = v1;
        v2Content = v2;
    } else {
        // Line removed from one
        switch (conflict.source) {
            case MergeSource.V1:
                v1Content = v1;
            case MergeSource.V2:
                v2Content = v2;
        }
    }



    return (
        <div className="border-solid border-[#D71414] rounded-md border-2 my-5 bg-white max-w-[980px] text-wrap break-words">
            <LabelledBox label="Version #1" labelBorder="rounded-tl" >{v1Content == null ? (<LineRemovedMessage />) : v1Content}</LabelledBox>
            <LabelledBox label="Version #2" >{v2Content == null ? (<LineRemovedMessage />) : v2Content}</LabelledBox>
        </div>
    );
}