/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'selector',
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  safelist: [],
  extend: {
    padding: {
      '64-important': '16rem !important',
    },
  },
  plugins: [
    function ({ addUtilities }) {
      addUtilities({
        '.py-64-important': {
          paddingTop: '16rem !important',
          paddingBottom: '16rem !important',
        },
      });
    },
  ],
}

