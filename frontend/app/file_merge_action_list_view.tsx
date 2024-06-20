import { useEffect, useState } from "react";
import { MergeAction, MergeReason } from "./merge_presentation";
import FileMergeActionSuccessView from "./file_merge_action_success_view";
import FileMergeActionConflictView from "./file_merge_action_conflict_view";

interface FileMergeActionListViewProps {
    original: File,
    v1: File,
    v2: File,
    actions: Array<MergeAction>
}

async function getFilesText(original: File,
    v1: File,
    v2: File): Promise<[string, string, string]> {
    return [await original.text(), await v1.text(), await v2.text()];
}


export default function FileMergeActionListView({ original, v1, v2, actions }: FileMergeActionListViewProps) {
    const [textFiles, setTextFiles] = useState<[string | null, string | null, string | null][]>([]);
    const [isReady, setReady] = useState(false);

    useEffect(() => {
        // Convert files into a 2d array of lines with each position in the inner array corresponding to either original, v1, or v2
        getFilesText(original, v1, v2)
            .then(([oText, v1Text, v2Text]) => {
                let oLines = oText.split("\n");
                let v1Lines = v1Text.split("\n");
                let v2Lines = v2Text.split("\n");

                let n = Math.max(oLines.length, v1Lines.length, v2Lines.length);

                let lineMatrix: [string | null, string | null, string | null][] = Array.from({ length: n }).map((v, i) => {
                    return [i < oLines.length ? oLines[i] : null, i < v1Lines.length ? v1Lines[i] : null, i < v2Lines.length ? v2Lines[i] : null];
                })

                setTextFiles(lineMatrix);
                setReady(true);
            });
    }, []);

    if (!isReady) return (<></>);

    return (
        <div className="bg-white max-w-[980px] px-3 py-3">
            {...actions.map((action) => {

                let [oLine, v1Line, v2Line] = textFiles[action.line];

                if (action.reason != MergeReason.CONFLICT) {
                    return (<FileMergeActionSuccessView key={action.line} original={oLine} v1={v1Line} v2={v2Line} action={action} ></FileMergeActionSuccessView>);
                }

                return (<FileMergeActionConflictView key={action.line} original={oLine} v1={v1Line} v2={v2Line} action={action} ></FileMergeActionConflictView>);
            })}
        </div>
    );
}