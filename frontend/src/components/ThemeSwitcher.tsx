import { createSignal } from "solid-js"
import Icon from "./Icon"

export default function ThemeSwitcher() {
    const [theme, setTheme] = createSignal("system" as "light" | "system" | "dark")

    const handleClick = () => {
        if (theme() === "light") {
            setTheme("system")
            return
        }

        if (theme() === "system") {
            setTheme("dark")
            return
        }

        if (theme() === "dark") {
            setTheme("light")
            return
        }
    }

    return <button class="flex items-center gap-3 text-primary-main hover:text-primary-active" onClick={handleClick}>
        <div class="text-base sm:text-sm select-none">
            {theme()}
        </div>
        {theme() === "light" && <Icon class="h-6" name="sunny" />}
        {theme() === "system" && <Icon class="h-6" name="contrast" />}
        {theme() === "dark" && <Icon class="h-6" name="moon" />}
    </button >
}