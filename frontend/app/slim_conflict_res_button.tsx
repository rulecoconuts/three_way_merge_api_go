import Button from "@mui/material/Button";

interface SlimConflictResButtonProps {
    onClick: () => void;
    children: React.ReactNode
}


export default function SlimConflictResButton({ onClick, children }: SlimConflictResButtonProps) {
    return (
        <Button role={undefined} variant="contained" color="primary" onClick={onClick}>
            {children}
        </Button>
    )
}