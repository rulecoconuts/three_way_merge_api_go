import assert from "assert";
import { MergeAction, MergeSource } from "./merge_presentation";

interface FileMergeActionSuccessViewProps {
    original: string | null,
    v1: string | null,
    v2: string | null,
    action: MergeAction
}

export default function FileMergeActionSuccessView({ original, v1, v2, action }: FileMergeActionSuccessViewProps) {
    assert.ok(original != null || v1 != null || v2 != null); // At list one version of the line must be non-null
    let content: string | null = null;

    switch (action.source) {
        case MergeSource.ORIGINAL:
            content = original!;
            break;
        case MergeSource.V1:
            content = v1!;
            break;
        case MergeSource.V2:
            content = v2!;
            break;
    }

    if (content == null) return (<></>);

    return (<div className="bg-white max-w-[980px] text-wrap break-words">{content}</div>);
}