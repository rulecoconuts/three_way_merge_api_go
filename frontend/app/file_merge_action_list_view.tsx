import { useEffect, useState } from "react";
import { MergeAction, MergeReason } from "./merge_presentation";
import FileMergeActionSuccessView from "./file_merge_action_success_view";
import FileMergeActionConflictView from "./file_merge_action_conflict_view";

interface FileMergeActionListViewProps {
    textLines: [string | null, string | null, string | null][];
    actions: Array<MergeAction>
}



export default function FileMergeActionListView({ textLines, actions }: FileMergeActionListViewProps) {

    return (
        <div className="bg-white max-w-[980px] px-3 py-3">
            {...actions.map((action) => {

                let [oLine, v1Line, v2Line] = textLines[action.line];

                if (action.reason != MergeReason.CONFLICT) {
                    return (<FileMergeActionSuccessView key={action.line} original={oLine} v1={v1Line} v2={v2Line} action={action} ></FileMergeActionSuccessView>);
                }

                return (<FileMergeActionConflictView key={action.line} original={oLine} v1={v1Line} v2={v2Line} action={action} ></FileMergeActionConflictView>);
            })}
        </div>
    );
}