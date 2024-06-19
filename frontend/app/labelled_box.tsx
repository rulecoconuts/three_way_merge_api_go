import React from "react";

interface LabelledBoxProps {
    labelBorder?: string,
    border?: string,
    label: string,
    children: React.ReactNode
}

export default function LabelledBox({ labelBorder, border, label, children }: LabelledBoxProps) {
    const defaultBorder = "";
    const defaultLabelBorder = "";
    const defaultLabelBackgroundColor = "bg-[#414B5D]";
    const defaultLabelTextColor = "text-white"

    const finalBorder = border ?? defaultBorder;
    const finalLabelBorder = labelBorder ?? defaultLabelBorder;
    const finalLabelBackgroundColor = defaultLabelBackgroundColor;
    const finalLabelTextColor = defaultLabelTextColor;

    return (
        <div>
            {/*Label */}
            <div className={`${finalLabelBorder} ${finalLabelBackgroundColor} ${finalLabelTextColor} w-fit`}>{label}</div>

            <div className={`${finalBorder} min-h-[10px] py-1`}>
                {children}
            </div>
        </div>
    );
}