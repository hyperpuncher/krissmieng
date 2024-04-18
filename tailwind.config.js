/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["./**/*.html", "./**/*.templ", "./**/*.go"],
	theme: {
		extend: {
			keyframes: {
				wiggle: {
					"0%, 50%": { transform: "rotate(-10deg)" },
					"25%, 75%": { transform: "rotate(10deg)" },
				},
				zoom: {
					"0%": { transform: "scale(1)" },
					"50%": { transform: "scale(1.1)" },
				},
				"slide-chat": {
					to: {
						opacity: "1",
						transform: "translateX(0)",
					},
				},
				scroll: {
					to: { transform: "translateX(-100%)" },
				},
			},
			animation: {
				wiggle: "wiggle 1s ease-in-out",
				zoom: "scale linear",
				"slide-chat": "slide-chat ease forwards",
				scroll: "scroll linear forwards",
			},
		},
	},
	plugins: [require("daisyui")],
	daisyui: {
		themes: ["cmyk"],
	},
};
