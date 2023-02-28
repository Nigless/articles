import { JSX } from "solid-js/jsx-runtime"

const styles = {
    primary: "text-primary-main hover:text-primary-active",
    secondary: "text-secondary-main hover:text-secondary-active",
    error: "text-error-main hover:text-error-active",
    warning: "text-warning-main hover:text-warning-active",
    info: "text-info-main hover:text-info-active",
    success: "text-success-main hover:text-success-active",
    neutral: "text-neutral-main hover:text-neutral-active",
}

interface LinkProps {
    href: string,
    children: JSX.Element,
    class?: string,
    color?:
    "primary" |
    "secondary" |
    "error" |
    "warning" |
    "info" |
    "neutral" |
    "success",
}

export default function Link(props: LinkProps) {
    return <a
        href={props.href}
        class={(props.class || "") + " transition-colors " + styles[props.color || "info"]}
    >
        {props.children}
    </a>
}