const styles = {
    primary: "border-primary-active bg-primary-disabled text-primary-main",
    secondary: "border-secondary-active bg-secondary-disabled text-secondary-main",
    error: "border-error-active bg-error-disabled text-error-main",
    warning: "border-warning-active bg-warning-disabled text-warning-main",
    info: "border-info-active bg-info-disabled text-info-main",
    success: "border-success-active bg-success-disabled text-success-main",
    neutral: "border-neutral-active bg-neutral-disabled text-neutral-main",
}

interface BlockProps {
    color?:
    "primary" |
    "secondary" |
    "error" |
    "warning" |
    "info" |
    "success" |
    "neutral",
}

export default function Block(props: BlockProps) {
    return <div class={"border-2 rounded-lg px-4 py-3 " + styles[props.color || "neutral"]}>
        primary: "text-primary-main hover:text-primary-active",
        secondary: "text-secondary-main hover:text-secondary-active",
        error: "text-error-main hover:text-error-active",
        warning: "text-warning-main hover:text-warning-active",
        info: "text-info-main hover:text-info-active",
        success: "text-success-main hover:text-success-active",
        neutral: "text-neutral-main hover:text-neutral-active",
    </div>
}