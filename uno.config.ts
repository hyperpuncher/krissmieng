import { defineConfig, presetUno, presetWebFonts } from "unocss";

export default defineConfig({
	cli: {
		entry: {
			patterns: ["./**/*.html", "./**/*.templ", "./**/*.go"],
			outFile: "assets/css/uno.css",
		},
	},
	presets: [presetUno(), presetWebFonts({
		fonts: {
			sans: "Montserrat:100,200,300,400,500,600,700,800,900",
		}
	})],
});
