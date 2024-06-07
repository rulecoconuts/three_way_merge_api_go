import { useEffect, useState } from "react";
import { styleText } from "util";

interface TextFilePreviewProps {
    file: File;
    className?: string
}
export default function TextFilePreview({ file, className }: TextFilePreviewProps) {
    const [text, setText] = useState<string>("");
    useEffect(() => {
        file.slice(0, 600).text()
            .then((content) => {
                setText(content);
            })
    }, file == null ? [] : [file]);

    return (
        <div className={`shadow-md dynamic-text-container ${className}`} style={{ whiteSpace: "pre-wrap" }} >
            {text}
        </div >
    );
}