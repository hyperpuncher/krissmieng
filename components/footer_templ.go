// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func footer() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex gap-5 justify-center items-center text-center\"><p class=\"text-xl font-bold\">\"Понедельник\" уже настал;) <br>Записывайся на ознакомительную консультацию </p></div><footer class=\"flex justify-between items-center p-4 footer\"><aside class=\"flex items-center\"><svg height=\"20\" width=\"20\" xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 512 512\" xml:space=\"preserve\"><path d=\"M476.886 337.609c-29.533-29.533-72.785-37.15-109.422-22.902l-53.474-53.474 145.25-145.25c32.028-32.028 32.028-83.954 0-115.981L256 203.242 52.759 0c-32.026 32.028-32.026 83.954 0 115.981l145.25 145.25-53.474 53.474c-36.636-14.248-79.889-6.631-109.422 22.902-39.841 39.843-39.841 104.668 0 144.511 39.843 39.843 104.668 39.843 144.511 0 36.303-36.303 39.512-93.344 9.656-133.334l66.721-66.721 66.721 66.721c-29.856 39.987-26.648 97.029 9.655 133.334 39.843 39.843 104.668 39.843 144.511 0 39.84-39.841 39.84-104.668-.002-144.509zM140.212 442.706c-18.109 18.111-47.577 18.111-65.686 0-18.111-18.109-18.111-47.577 0-65.686 18.109-18.111 47.577-18.111 65.686 0 18.11 18.109 18.11 47.577 0 65.686zm297.262 0c-18.109 18.111-47.577 18.111-65.686 0-18.111-18.109-18.111-47.577 0-65.686 18.109-18.111 47.577-18.111 65.686 0 18.111 18.109 18.111 47.577 0 65.686z\"></path></svg><p>by <a class=\"font-bold link link-secondary\" target=\"_blank\" href=\"https://github.com/hyperpuncher\">igor</a></p></aside><nav class=\"flex relative gap-4 md:justify-self-end md:place-self-center\"><p class=\"absolute bottom-8 right-1/3 text-3xl animate-bounce select-none\">👇</p><a href=\"https://t.me/krissmieng\" target=\"_blank\" aria-label=\"Telegram канал ссылка\"><svg class=\"hover:opacity-80\" width=\"24\" height=\"24\" xmlns=\"http://www.w3.org/2000/svg\" fill-rule=\"evenodd\" clip-rule=\"evenodd\"><path d=\"M20.699 11.069c1.947 1.015 3.3 2.81 3.3 4.977 0 1.204-.433 2.387-1.202 3.304-.032 1.075.601 2.618 1.171 3.741-1.529-.277-3.702-.886-4.686-1.49-.833.202-1.631.295-2.385.295-2.919 0-5.185-1.398-6.278-3.271 6.427-.207 9.818-4.482 10.08-7.556zm-10.798-10.07c-5.282 0-10 3.522-10 8.34 0 1.708.615 3.385 1.705 4.687.046 1.525-.851 3.713-1.66 5.307 2.168-.392 5.251-1.257 6.648-2.114 7.696 1.873 13.307-2.837 13.307-7.88 0-4.844-4.751-8.34-10-8.34zm3.048 10.74c-.212.071-.441-.044-.511-.256 0 0-1.496-.501-3.272.088l1.166 1.814c.121.179.037.422-.164.49l-.794.263-.167.028c-.149 0-.294-.064-.396-.179l-1.604-1.77c-.873.276-1.862-.042-2.192-.865-.076-.188-.114-.389-.114-.592 0-.585.327-1.183 1.038-1.533 3.559-1.747 4.128-3.696 4.128-3.696-.086-.262.11-.532.382-.532.169 0 .327.108.383.277l2.371 5.952c.071.212-.044.441-.254.511zm1.687-2.043l-.614-.261c.267-.634.293-1.371.014-2.058-.279-.688-.808-1.199-1.441-1.466l.26-.615c1.265.535 2.046 1.771 2.046 3.091 0 .454-.092.898-.265 1.309zm-1.063-.451l-.599-.254c.154-.365.169-.789.008-1.185-.16-.396-.466-.69-.83-.843l.253-.601c.518.219.952.635 1.179 1.198.229.564.207 1.165-.011 1.685z\"></path></svg></a> <a href=\"https://t.me/KrisSmi27\" target=\"_blank\" aria-label=\"Telegram ссылка\"><svg class=\"hover:opacity-80\" width=\"24px\" height=\"24px\" version=\"1.1\" xmlns=\"http://www.w3.org/2000/svg\" xmlns:xlink=\"http://www.w3.org/1999/xlink\" xml:space=\"preserve\" xmlns:serif=\"http://www.serif.com/\" style=\"fill-rule:evenodd;clip-rule:evenodd;stroke-linejoin:round;stroke-miterlimit:1.41421;\"><path id=\"telegram-1\" d=\"M18.384,22.779c0.322,0.228 0.737,0.285 1.107,0.145c0.37,-0.141 0.642,-0.457 0.724,-0.84c0.869,-4.084 2.977,-14.421 3.768,-18.136c0.06,-0.28 -0.04,-0.571 -0.26,-0.758c-0.22,-0.187 -0.525,-0.241 -0.797,-0.14c-4.193,1.552 -17.106,6.397 -22.384,8.35c-0.335,0.124 -0.553,0.446 -0.542,0.799c0.012,0.354 0.25,0.661 0.593,0.764c2.367,0.708 5.474,1.693 5.474,1.693c0,0 1.452,4.385 2.209,6.615c0.095,0.28 0.314,0.5 0.603,0.576c0.288,0.075 0.596,-0.004 0.811,-0.207c1.216,-1.148 3.096,-2.923 3.096,-2.923c0,0 3.572,2.619 5.598,4.062Zm-11.01,-8.677l1.679,5.538l0.373,-3.507c0,0 6.487,-5.851 10.185,-9.186c0.108,-0.098 0.123,-0.262 0.033,-0.377c-0.089,-0.115 -0.253,-0.142 -0.376,-0.064c-4.286,2.737 -11.894,7.596 -11.894,7.596Z\"></path></svg></a> <a href=\"https://www.instagram.com/smychnikova/\" target=\"_blank\" aria-label=\"Instagram ссылка\"><svg class=\"hover:opacity-80\" xmlns=\"http://www.w3.org/2000/svg\" x=\"0px\" y=\"0px\" width=\"24\" height=\"24\" viewBox=\"0 0 50 50\"><path d=\"M 16 3 C 8.83 3 3 8.83 3 16 L 3 34 C 3 41.17 8.83 47 16 47 L 34 47 C 41.17 47 47 41.17 47 34 L 47 16 C 47 8.83 41.17 3 34 3 L 16 3 z M 37 11 C 38.1 11 39 11.9 39 13 C 39 14.1 38.1 15 37 15 C 35.9 15 35 14.1 35 13 C 35 11.9 35.9 11 37 11 z M 25 14 C 31.07 14 36 18.93 36 25 C 36 31.07 31.07 36 25 36 C 18.93 36 14 31.07 14 25 C 14 18.93 18.93 14 25 14 z M 25 16 C 20.04 16 16 20.04 16 25 C 16 29.96 20.04 34 25 34 C 29.96 34 34 29.96 34 25 C 34 20.04 29.96 16 25 16 z\"></path></svg></a></nav></footer>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
