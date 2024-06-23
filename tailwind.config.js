import { fontFamily } from "tailwindcss/defaultTheme";

module.exports = {
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  safelist: [],
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px"
      }
    },
  },
  extend: {},
  fontFamily: {
    sans: [...fontFamily.sans]
  }
};

