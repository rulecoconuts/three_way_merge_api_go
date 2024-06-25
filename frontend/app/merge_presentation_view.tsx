import { useEffect, useState } from "react";
import FileMergeActionListView from "./file_merge_action_list_view";
import { MergePresentation } from "./merge_presentation";
import Button from "@mui/material/Button";
import { Download } from "@mui/icons-material";
import { Merger } from "./merger";
var pretty = require('pretty-time');


interface MergePresentationViewProps {
    original: File,
    v1: File,
    v2: File,
    presentation: MergePresentation
}

async function getFilesText(original: File,
    v1: File,
    v2: File): Promise<[string, string, string]> {
    return [await original.text(), await v1.text(), await v2.text()];
}

export default function MergePresentationView({ original, v1, v2, presentation }: MergePresentationViewProps) {
    const [textLines, setTextLines] = useState<[string | null, string | null, string | null][]>([]);
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

                setTextLines(lineMatrix);
                setReady(true);
            });
    }, []);

    if (!isReady) return (<></>);

    function generateMergedFile() {
        // fileName consists the name of the file
        const text = (new Merger()).resolveActionsIntoText(textLines, presentation.actions);
        const element = document.createElement('a');
        const file = new Blob([text], { type: 'text/plain' });
        element.href = URL.createObjectURL(file);
        element.download = "merged.txt";
        document.body.appendChild(element);
        element.click();
        document.body.removeChild(element);
    }


    return (
        <div>
            <div className="flex flex-row items-center mb-1">
                <h2 className="font-medium">Merged</h2>
                <p className="text-gray-700 mx-1 text-sm">(in {pretty(presentation.timeSpentInMicroseconds * 1000)} | Total response time: {pretty(presentation.totalTransferTimeInNanoseconds)})</p>
                <Button role={undefined} variant="contained" color="primary" onClick={generateMergedFile}>
                    <Download />
                </Button>
            </div>
            <FileMergeActionListView textLines={textLines} actions={presentation.actions}></FileMergeActionListView>
        </div>
    );
}