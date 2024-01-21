/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./**/*.{templ,js,html}"],
  theme: {
    container: {
      center: true,
      screens: {
        sm: "800px",
        md: "1024px",
        lg: "1280px",
        xl: "1540px",
        "2xl": "2560px",
      }
    },
    extend: {},
  },
  plugins: [],
}
