import { useEffect, useState } from "react";

interface TextBlobViewProps {
    blob: Blob;
    className?: string;
}
export default function TextBlobView({ blob, className }: TextBlobViewProps) {
    const [text, setText] = useState<string>("");
    useEffect(() => {
        blob.text()
            .then((content) => {
                setText(content);
            })
    }, blob == null ? [] : [blob]);

    return (
        <div className={`shadow-md ${className}`} style={{ whiteSpace: "pre-wrap" }} >
            {text}
        </div >
    );
}