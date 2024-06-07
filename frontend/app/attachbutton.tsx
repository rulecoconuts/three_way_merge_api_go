'use client';
import Button from "@mui/material/Button";
import { styled } from '@mui/material/styles';
import CloudUploadIcon from "@mui/icons-material/CloudUpload";
import { useState, ChangeEventHandler } from "react";
import Image from "next/image";
import TextFilePreview from "./text-file-preview";

// const VisuallyHiddenInput = styled('input')({
//     clip: 'rect(0 0 0 0)',
//     clipPath: 'inset(50%)',
//     height: 1,
//     overflow: 'hidden',
//     position: 'absolute',
//     bottom: 0,
//     left: 0,
//     whiteSpace: 'nowrap',
//     width: 1,
// });

interface AttachButtonProps {
    label: string;
    onFileAttached?: (file: File) => void;
}

export default function AttachButton({ label, onFileAttached }: AttachButtonProps) {
    const [file, setFile] = useState<File | null>(null);
    const [name, setName] = useState<string | null>(null);

    let handleAttachment: ChangeEventHandler<HTMLInputElement> = (event) => {
        if (event.target.files == null) return;
        let f = event.target.files[0];
        onFileAttached?.call(null, f);
        setFile(f);
        setName(f.name);
    };

    return (
        <div className="flex flex-col justify-start">
            <div className="flex flex-row flex-wrap items-start">
                <p className="pr-2">{label}</p>
                <div className="flex flex-col jutify-start items-center">
                    <Button component="label" role={undefined} variant="contained" color="primary" tabIndex={-1} startIcon={<CloudUploadIcon />}>
                        {file == null ? "Attach File" : "Change File"}
                        <input type="file" onChange={handleAttachment} hidden />
                    </Button>
                    <div className="flex flex-col justify-start items-start">
                        {name == null ? (<></>) : (<p className="pt-2 mb-0 text-sm text-gray-600">{name}</p>)}
                        {file == null ? (<></>) : (<TextFilePreview file={file} className="dynamic-text-container mt-1 w-[70px] h-[80px] max-w-[70px] max-h-[80px] bg-white px-1 py-1"></TextFilePreview>)}
                    </div>
                </div>
            </div>
        </div>
    );
}