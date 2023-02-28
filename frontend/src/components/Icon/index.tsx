import * as icons from "./icons"

interface IconProps {
    class?: string,
    name: string
}

export default function Icon(props: IconProps) {

    return <svg class={props.class} viewBox="0 0 512 512" fill="currentColor" aria-hidden="true" innerHTML={icons[props.name].default}>
    </svg>
}