/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
	theme: {
		extend: {
			keyframes: {
				wiggle: {
					"0%, 50%": { transform: "rotate(-3deg)" },
					"25%, 75%": { transform: "rotate(3deg)" },
				},
			},
			animation: {
				wiggle: "wiggle 1s ease-in-out",
			},
		},
	},
	plugins: [require("daisyui")],
	daisyui: {
		themes: ["cmyk"],
	},
};
