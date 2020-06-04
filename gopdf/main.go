package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/russross/blackfriday"
	"github.com/signintech/gopdf"
	"golang.org/x/text/unicode/norm"
)

const textcard = `TEST
what is here?
and? next?`

// replace the paragraphs with double blank lines for the pdf renderer
func toCustomMarkdown(input string) string {
	renderer := blackfriday.HtmlRenderer(0, "", "")
	markdowned := string(blackfriday.MarkdownOptions([]byte(input), renderer, blackfriday.Options{Extensions: blackfriday.EXTENSION_HARD_LINE_BREAK}))

	markdowned = norm.NFC.String(markdowned)

	//TODO: uncomment if print error
	// markdowned = strings.Replace(markdowned, "<br>", " ", -1)
	markdowned = strings.Replace(markdowned, "</p>\n\n<p>", "<br><br>", -1)

	// align text center if text length less than 100 char
	if len(input) < 100 {
		markdowned = strings.Replace(markdowned, "<p>", "<p><center>", -1)
	}

	return markdowned
}

func replaceAmpersand(input string) string {
	s := strings.Replace(input, "&amp;", "&", -1)
	s = strings.Replace(s, "&quot;", "\"", -1)
	s = strings.Replace(s, "&#34;", "\"", -1)
	return s
}

func main() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{
		Unit:     gopdf.Unit_MM,
		PageSize: *gopdf.PageSizeA4,
	})
	pdf.AddPage()
	err := pdf.AddTTFFont("opensan", "OpenSansEmoji.ttf")
	if err != nil {
		log.Panicln(err.Error())
	}

	err = pdf.SetFont("opensan", "", 16)
	if err != nil {
		log.Panicln(err.Error())
	}

	htmlText := toCustomMarkdown("😃😋\n <3 Hello\n Test Me")
	fmt.Println(htmlText)
	lst := strings.Split(textcard, "\n")
	for _, s := range lst {
		fmt.Println(s)
	}
	// htmlText = replaceAmpersand(htmlText)
	// html.Write(10, strings.Replace(htmlText, "\n", "", -1))

	err = pdf.Cell(nil, textcard)
	if err != nil {
		fmt.Println("Loi roi", err.Error())
	}

	// var buf bytes.Buffer
	// err = pdf.Write(&buf)
	pdf.WritePdf("hello.pdf")

}
