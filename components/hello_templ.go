// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.663
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func hello() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col lg:gap-20 gap-10 mx-10 lg:mx-auto\"><div class=\"flex lg:text-8xl text-5xl font-bold\"><h1 class=\"lg:ml-20\">Привет, <br>я Кристина </h1><div class=\"scale-x-[-1]\"><p class=\"animate-wiggle origin-bottom-right\">👋</p></div></div><div class=\"flex justify-end\"><h2 class=\"lg:mr-20 lg:w-2/3 lg:text-4xl text-lg font-bold\">Репетитор английского языка с 2014 года.<br>Я выбрала путь преподавания, чтобы вдохновлять и поддерживать людей в их стремлении к самосовершенствованию. <br>Изучение новых языков открывает перед вами весь мир и дарит возможность открыть новые горизонты личностного роста.<br>Моя миссия — помочь вам стать увереннее и счастливее!</h2></div><div class=\"flex justify-center\"><img id=\"hero\" class=\"rounded-lg shadow-2xl md:animate-zoom md:timeline-view\" src=\"assets/images/kris.webp\" width=\"1920\" height=\"1297\" alt=\"Моё фото\"></div><div class=\"lg:text-2xl text-lg flex flex-col gap-10 xl:w-2/3 mx-auto\"><p>Я закончила Минский Государственный Лингвистический университет, в 2023 году успешно прошла курс английского с преподавателем из Британии. Прохожу мастер-классы у топовых преподавателей и продолжаю прокачивать свои знания каждый день. Мои ученики успешно устраиваются на работу в ведущие международные компании, переезжают за границу и получают образование в престижных университетах.</p><p>Со мной скучные уроки остаются в прошлом! Я помогу вам преодолеть языковой барьер раз и навсегда. Мои занятия адаптированы для людей с различным уровнем владения английским и наполнены живым общением и интересными заданиями, далекими от стереотипов о нудных уроках. Используя коммуникативный подход, мы не просто учимся — мы общаемся. Я становлюсь не только вашим преподавателем, но и надежным другом, наставником в мире английского языкa.</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = examples().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>Индивидуальный подход  — это то, за что меня ценят студенты. Отталкиваясь от ваших запросов, я строю наши занятия. Многочисленное раскрытие скобочек, перевод бесконечных абзацев и заучивание текстов остаётся в прошлом, мы фокусируемся на реальном общении и практическом применении языка. Я использую современные интерактивные инструменты (Miro, WordWall, Notion и др.), а также дополняю занятия актуальными материалами, такими как Youtube видео, статьи из известных изданий, видео из соцсетей и тд. </p><p>Домашние задания выполняются учениками самостоятельно, с моей проверкой предварительно до занятия, а сам разбор ошибок или возникших вопросов происходит непосредственно на уроке, что обеспечивает максимальную эффективность обучения. Моя поддержка доступна 24/7, и я всегда готова помочь вам в достижении ваших целей.</p><p>В качестве знакомства я провожу бесплатную получасовую консультацию, где мы сможем узнать друг друга поближе. Я задам пару вопросов для определения целей и уровня языка, а также продемонстрирую как проходит наша работа.</p></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
