import RemoveCircleOutlineIcon from '@mui/icons-material/RemoveCircleOutline';
export default function LineRemovedMessage() {
    return (
        <div className="flex flex-row space-x-2">
            <RemoveCircleOutlineIcon sx={{
                color: "#D71414"
            }} />
            <span>Line was removed</span>
        </div>
    );
}