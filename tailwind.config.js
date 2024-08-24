import { fontFamily } from "tailwindcss/defaultTheme";

module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  safelist: [],
  theme: {
    container: {
      center: true,
      padding: {
        DEFAULT: "1rem",
        sm: "1rem",
        md: "2rem",
      },
      screens: {
        "2xl": "1400px"
      }
    },
  },
  plugins: [
    require('daisyui'),
  ],
  daisyui: {
    styled: true,
    prefix: "daisy-",
    themes: ["bumblebee"],
  },

  extend: {},
  fontFamily: {
    sans: [...fontFamily.sans]
  }
};

