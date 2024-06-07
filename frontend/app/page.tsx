'use client';
import Image from "next/image";
import AttachButton from "./attachbutton";
import Button from "@mui/material/Button";
import { useState } from "react";
import TextFilePreview from "./text-file-preview";
import TextBlobView from "./text-blob-view";
import { Merger, MergeResult } from "./merger";
import Modal from "@mui/material/Modal";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";

export default function Home() {
  const [merged, setMerged] = useState<MergeResult | null>(null);
  const [original, setOriginal] = useState<File | null>(null);
  const [a, setA] = useState<File | null>(null);
  const [b, setB] = useState<File | null>(null);
  const [error, setError] = useState<string | null>(null);

  function merge() {
    let currentErr: string[] = [];
    if (original == null) {
      currentErr.push("Original file is required");
    }

    if (a == null) {
      currentErr.push("Version #1 file is required");
    }

    if (b == null) {
      currentErr.push("Version #2 file is required");
    }

    if (currentErr.length > 0) {
      setError(currentErr.join(", "));
      return;
    }

    setMerged(null);

    (new Merger()).merge(original!, a!, b!)
      .then((blob) => {
        setMerged(blob);
      });
  }

  function closeError() {
    setError(null);
  }


  return (
    <main className="flex min-h-screen flex-col items-center justify-start p-24">
      <div className="z-10 w-full max-w-5xl items-center justify-center font-mono text-sm lg:flex">
        <h1 className="font-semibold text-xl text-center">3 Way Merge Improvement Project</h1>

      </div>
      <p className="pt-5">Please run and submit test results to help improve this tool</p>
      <div>
        <div className="flex flex-row pt-10 items-center justify-center w-auto">
          <div className="flex flex-row flex-wrap space-x-3 items-start">
            <AttachButton label="Original Version" onFileAttached={setOriginal}></AttachButton>
            <AttachButton label="Changed Version #1" onFileAttached={setA}></AttachButton>
            <AttachButton label="Changed Version #2" onFileAttached={setB}></AttachButton>
          </div>
        </div>
        <div className="flex flex-row justify-center mt-10">
          <Button variant="contained" color="primary" sx={{ px: 15 }} onClick={merge}>Merge</Button>
        </div>
        {merged == null ? (<></>) : (
          <div>
            <div className="flex flex-row items-center">
              <h2 className="font-medium">Merged</h2>
              <p className="text-gray-700 ml-1 text-sm">(in {merged.milliseconds} ms)</p>
            </div>
            <TextBlobView blob={merged.blob} className="mt-10 bg-white px-3 py-3 max-w-[980px] text-wrap break-words"></TextBlobView>
          </div>
        )}
        <Modal open={(error?.length ?? 0) > 0} onClose={closeError} aria-labelledby="modal-modal-title"
          aria-describedby="modal-modal-description">
          <Box sx={{
            position: 'absolute' as 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
            width: 400,
            bgcolor: 'background.paper',
            border: '2px solid #000',
            boxShadow: 24,
            p: 4,
          }}>
            <Typography id="modal-modal-title" variant="h6" component="h2">
              Missing Files
            </Typography>
            <Typography id="modal-modal-description" sx={{ mt: 2 }}>
              {error}
            </Typography>
          </Box>
        </Modal>
      </div>
    </main>
  );
}
