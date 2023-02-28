/** @type {import('tailwindcss').Config} */
module.exports = {
  future: {
    hoverOnlyWhenSupported: true,
  },
  content: ['./src/**/*.{js,jsx,ts,tsx}', './index.html'],
  theme: {
    colors: {
      primary: {
        main: 'rgb(var(--color-primary-main) / <alpha-value>)',
        active: 'rgb(var(--color-primary-active) / <alpha-value>)',
        disabled: 'rgb(var(--color-primary-disabled) / <alpha-value>)',
      },
      secondary: {
        main: 'rgb(var(--color-secondary-main) / <alpha-value>)',
        active: 'rgb(var(--color-secondary-active) / <alpha-value>)',
        disabled: 'rgb(var(--color-secondary-disabled) / <alpha-value>)',
      },
      error: {
        main: 'rgb(var(--color-error-main) / <alpha-value>)',
        active: 'rgb(var(--color-error-active) / <alpha-value>)',
        disabled: 'rgb(var(--color-error-disabled) / <alpha-value>)',
      },
      warning: {
        main: 'rgb(var(--color-warning-main) / <alpha-value>)',
        active: 'rgb(var(--color-warning-active) / <alpha-value>)',
        disabled: 'rgb(var(--color-warning-disabled) / <alpha-value>)',
      },
      info: {
        main: 'rgb(var(--color-info-main) / <alpha-value>)',
        active: 'rgb(var(--color-info-active) / <alpha-value>)',
        disabled: 'rgb(var(--color-info-disabled) / <alpha-value>)',
      },
      success: {
        main: 'rgb(var(--color-success-main) / <alpha-value>)',
        active: 'rgb(var(--color-success-active) / <alpha-value>)',
        disabled: 'rgb(var(--color-success-disabled) / <alpha-value>)',
      },
      neutral: {
        main: 'rgb(var(--color-neutral-main) / <alpha-value>)',
        active: 'rgb(var(--color-neutral-active) / <alpha-value>)',
        disabled: 'rgb(var(--color-neutral-disabled) / <alpha-value>)',
      },
      action: {
        active: 'rgb(var(--color-action-active) / <alpha-value>)',
        contrast: 'rgb(var(--color-action-contrast) / <alpha-value>)',
        disabled: 'rgb(var(--color-action-disabled) / <alpha-value>)',
        main: 'rgb(var(--color-action-main) / <alpha-value>)',
      },
    },
  },
}