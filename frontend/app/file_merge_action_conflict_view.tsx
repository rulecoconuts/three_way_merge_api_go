import assert from "assert";
import { MergeAction, MergeConflict, MergeConflictReason, MergeResolution, MergeSource, ResolvedContent } from "./merge_presentation";
import { Merger } from "./merger";
import { useState } from "react";
import LabelledBox from "./labelled_box";
import LineRemovedMessage from "./line_removed_message";
import SlimConflictResButton from "./slim_conflict_res_button";

interface FileMergeActionConflictViewProps {
    original: string | null,
    v1: string | null,
    v2: string | null,
    action: MergeAction
}

export default function FileMergeActionConflictView({ original, v1, v2, action }: FileMergeActionConflictViewProps) {
    assert.ok(original != null || v1 != null || v2 != null); // At list one version of the line must be non-null
    let [content, setContent] = useState<ResolvedContent>((new Merger()).resolveContent(original, v1, v2, action));

    function resolve(resolution: MergeResolution | null) {
        action.resolution = resolution;
        setContent((new Merger()).resolveContent(original, v1, v2, action));
    }

    let v1Content: string | null = content.v1;
    let v2Content: string | null = content.v2;
    let shouldIncludeMargins = action.resolution == null;

    return (
        <div className={`${shouldIncludeMargins && "my-5"} bg-white max-w-[980px] text-wrap break-words`}>
            {action.resolution == null && (
                <>
                    <div className="flex flex-row items-start space-x-1 mb-1">
                        {/* Resolution buttons row */}
                        <SlimConflictResButton onClick={() => resolve(MergeResolution.V1)}>
                            Keep Version #1
                        </SlimConflictResButton>
                        <SlimConflictResButton onClick={() => resolve(MergeResolution.V2)}>
                            Keep Version #2
                        </SlimConflictResButton>
                        {
                            content.conflict.reason == MergeConflictReason.BOTH_CHANGED &&
                            (<SlimConflictResButton onClick={() => resolve(MergeResolution.KEEP_BOTH)}>
                                Keep Both
                            </SlimConflictResButton>)}
                    </div>
                    <div className="border-solid border-[#D71414] rounded-md border-2 ">
                        <LabelledBox label="Version #1" labelBorder="rounded-tl" >{v1Content == null ? (<LineRemovedMessage />) : v1Content}</LabelledBox>
                        <LabelledBox label="Version #2" >{v2Content == null ? (<LineRemovedMessage />) : v2Content}</LabelledBox>
                    </div>
                </>
            )}

            {/* Resolved Content if a resolution has been made */}
            {action.resolution != null && (<div className="whitespace-pre-line">{content.toString()}</div>)}
        </div>
    );
}