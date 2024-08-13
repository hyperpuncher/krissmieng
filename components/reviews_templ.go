// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.707
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

var reviews_text = []string{
	"Я начала лучше понимать устную речь, даже если не знаю все слова, начала понимать из контекста. Плюс не очень сильно (тк мы не долго занимаемся) увеличился словарный запас. Я поняла разницу между quite и quiet. Благоприятна ли атмосфера во время занятия? Да, мне нравится, что мы много разговариваем на разные темы Больше всего мне нравятся тексты и на словарь упражнения, не люблю грамматику Немного проще стало общаться на работе",

	"Я начала разговаривать, времена разделились из одной каши в логическую структуру. Атмосфера-супер. Мне нравятся тесты после каждого раздела, особенно, упражнения со строением предложений. Определенно, изучение языка стало приятным моментом. Вижу улучшения даже в изучении польского, думаю, это имеет связь",

	"Комфортная атмосфера, уже занимаюсь 2 года. Вижу хороший прогресс, новые слова в речи при общении и на работе. Начал работать на Америку и увеличил заработок",

	"В первую очередь из важного - мне менее страшно говорить на английском. В моем топе предполагаемых причин провалов в использовании иностранного просто страх открыть рот и использовать его - на первом месте. Поэтому понижение планки страха и напряжения - это главное улучшения. Думаю, это какой- то психологический перенос - мол на занятии комфортно и безопасно говорить, и это состояние переносится в быт как-то.",

	"Атмосфера и настроение занятий однозначно очень располагающие, комфортные, позитивные. У меня нет напряжение и волнения перед занятием, как это обычно бывает и что мне очень не нравится (не хочется энергию и ресурс тратить на дополнительные переживания). Это реально супер важно, так как не отталкивает от занятий и даже в самый сложный день нет мысли мол \"блин, я так устала, никакого ресурса не осталось, а тут еще этот напряжный английский, пожалуй сольюсь\".",

	"Если подытожить: я со своей стороны не сильно вкладываюсь, не супер много ресурса в язык направляю. Но даже при этом ты умудряешься успеть что-то заложить за занятие. Какой же был бы результат, если бы я сама регулярно дополнительно над этим работала Я очень благодарна (и благодарна Саше за рекомендацию). С тобой интересно, спокойно, ободряюще и очень полезно. Тебя приятно слушать (произношение, подача), легко коммуницировать, и занятия улучшают настрой",
}

func reviews() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col justify-center mx-1 text-lg xl:mx-auto xl:w-2/3\"><h2 class=\"mb-5 text-3xl font-bold text-center\">Отзывы</h2>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for i, review := range reviews_text {
			if i%2 == 0 {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"ml-20 opacity-0 translate-x-1/2 animate-slide-chat timeline-view chat chat-end drop-shadow-md\"><div class=\"max-w-4xl chat-bubble chat-bubble-accent\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var2 string
				templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(review)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/reviews.templ`, Line: 24, Col: 14}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			} else {
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"mr-20 opacity-0 -translate-x-1/2 animate-slide-chat timeline-view chat chat-start drop-shadow-md\"><div class=\"max-w-4xl chat-bubble chat-bubble-secondary\">")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				var templ_7745c5c3_Var3 string
				templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(review)
				if templ_7745c5c3_Err != nil {
					return templ.Error{Err: templ_7745c5c3_Err, FileName: `components/reviews.templ`, Line: 30, Col: 14}
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
				_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
				if templ_7745c5c3_Err != nil {
					return templ_7745c5c3_Err
				}
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
