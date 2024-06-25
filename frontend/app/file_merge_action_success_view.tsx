import assert from "assert";
import { MergeAction, MergeSource } from "./merge_presentation";
import { Merger } from "./merger";

interface FileMergeActionSuccessViewProps {
    original: string | null,
    v1: string | null,
    v2: string | null,
    action: MergeAction
}

export default function FileMergeActionSuccessView({ original, v1, v2, action }: FileMergeActionSuccessViewProps) {
    assert.ok(original != null || v1 != null || v2 != null); // At list one version of the line must be non-null
    let content: string | null = (new Merger()).resolveContentFromSuccess(original, v1, v2, action);

    if (content == null) return (<></>);

    return (<div className="bg-white max-w-[980px] text-wrap break-words">{content}</div>);
}