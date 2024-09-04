/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: 'selector',
  content: ["./**/*.html", "./**/*.templ", "./**/*.go",],
  safelist: [],
  theme: {
    extend: {
      padding: {
        '52': '12rem',
        '44': '11rem',
        '24': '6rem',
      },
      screens: {
        'xs': '500px',
      }
    },
  },
  plugins: [
    function ({ addUtilities }) {
      addUtilities({
        '.pt-52-important': {
          paddingTop: '12rem !important',
        },
        '.pb-44-important': {
          paddingBottom: '11rem !important',
        },
        '.py-64-important': {
          paddingTop: '16rem !important',
          paddingBottom: '16rem !important',
        },
        '.py-24-important': {
          paddingTop: '6rem !important',
          paddingBottom: '6rem !important',
        },
        '.py-16-important': {  // Add this new utility
          paddingTop: '5rem !important',
          paddingBottom: '4rem !important',
        },
        '.hidden-important': {
          display: 'none !important',
        },
        '.block-important': {
          display: 'block !important',
        },
        '.pb-32-important': {
          paddingBottom: '8rem !important',
        },
        '.dark-autofill': {
          '.dark &:-webkit-autofill': {
            '-webkit-text-fill-color': '#ffffff',
            'box-shadow': '0 0 0px 1000px rgb(7 89 133 / var(--tw-bg-opacity)) inset',
          },
          '.dark &:-webkit-autofill:focus': {
            '-webkit-text-fill-color': '#ffffff',
            'box-shadow': '0 0 0px 1000px rgb(7 89 133 / var(--tw-bg-opacity)) inset',
          },
        },
        '.flex-important': {
          display: 'flex !important',
        },
        '.hidden-xs-flex': {
          '@media (max-width: 599px)': {
            display: 'none !important',
          },
          '@media (min-width: 600px)': {
            display: 'flex !important',
          },
        },
        '.carousel': {
          'overflow-x': 'auto',
          'scroll-snap-type': 'x mandatory',
          'scroll-behavior': 'smooth',
          '-webkit-overflow-scrolling': 'touch',
        },
        '.carousel-inner': {
          'display': 'flex',
          'flex-wrap': 'nowrap',
        },
        '.carousel-item': {
          'flex': '0 0 auto',
          'scroll-snap-align': 'center',
        },
      });
    },
  ],
}

