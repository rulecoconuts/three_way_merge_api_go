import { assert } from "console";
import { MergeAction } from "./merge_presentation";

interface FileMergeActionConflictViewProps {
    original: string | null,
    v1: string | null,
    v2: string | null,
    action: MergeAction
}

export default function FileMergeActionConflictView({ original, v1, v2, action }: FileMergeActionConflictViewProps) {
    assert(original != null || v1 != null || v2 != null); // At list one version of the line must be non-null


    return (<></>);
}