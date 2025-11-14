/** @type {import('tailwindcss').Config} */
module.exports = {
    theme: {
        extend: {
            typography: {
                DEFAULT: {
                    css: {
                        'li::marker': {
                            color: '#2b231b',
                        },
                        'ul > li::marker': {
                            color: '#2b231b',
                        },
                        'ol > li::marker': {
                            color: '#2b251b',
                            fontWeight: '600',
                        },
                        'blockquote': {
                            borderLeftColor: '#BAAE98',
                            color: "#696357",
                            borderLeftWidth: '0.25rem',
                            borderLeftStyle: 'solid',
                            paddingLeft: '1rem',
                            fontStyle: 'italic',
                        },
                        'strong': {
                            fontWeight: '700',
                        },
                        'h1, h2': {
                            fontFamily: 'Yeseva One, serif',
                            color: "#000000",
                        },
                        'h3, h4, h5, h6': {
                            fontWeight: '900',
                        },
                    },
                },
            },
        },
    },
}
