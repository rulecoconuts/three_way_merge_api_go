'use client';
import Image from "next/image";
import AttachButton from "./attachbutton";
import Button from "@mui/material/Button";
import { useEffect, useState } from "react";
import TextFilePreview from "./text-file-preview";
import TextBlobView from "./text-blob-view";
import { Merger, MergeResult } from "./merger";
import Modal from "@mui/material/Modal";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import { MergePresentation } from "./merge_presentation";
import FileMergeActionListView from "./file_merge_action_list_view";
import MergePresentationView from "./merge_presentation_view";

export default function Home() {
  const [merged, setMerged] = useState<MergePresentation | null>(null);
  const [original, setOriginal] = useState<File | null>(null);
  const [v1, setV1] = useState<File | null>(null);
  const [v2, setV2] = useState<File | null>(null);
  const [error, setError] = useState<string | null>(null);

  function merge() {
    let currentErr: string[] = [];
    if (original == null) {
      currentErr.push("Original file is required");
    }

    if (v1 == null) {
      currentErr.push("Version #1 file is required");
    }

    if (v2 == null) {
      currentErr.push("Version #2 file is required");
    }

    if (currentErr.length > 0) {
      setError(currentErr.join(", "));
      return;
    }

    setMerged(null);

    (new Merger()).getMergeActions(original!, v1!, v2!)
      .then((presentation) => {
        setMerged(presentation);
      });
  }

  // Clear merged if any of the files change
  useEffect(() => {
    setMerged(null);
  }, [original, v1, v2]);

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
            <AttachButton label="Changed Version #1" onFileAttached={setV1}></AttachButton>
            <AttachButton label="Changed Version #2" onFileAttached={setV2}></AttachButton>
          </div>
        </div>
        <div className="flex flex-row justify-center mt-10">
          <Button variant="contained" color="primary" sx={{ px: 15 }} onClick={merge}>Merge</Button>
        </div>
        {merged == null ? (<></>) : (
          <div className="mt-2"> <MergePresentationView presentation={merged} original={original!} v1={v1!} v2={v2!} /> </div>
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
