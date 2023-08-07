package utils

import (
	"strings"

	quill "github.com/dchenk/go-render-quill"
	"github.com/k3a/html2text"
)

func QuillToHtmlAndPlainText(quillText string) (string, string, error) {
	html, err := quill.Render([]byte(quillText))
	if err != nil {
		return "", "", err
	}
	htmlText := string(html)
	html2text.SetUnixLbr(true)
	plainText := html2text.HTML2Text(htmlText)
	formattedPlainText := strings.Replace(plainText, "\n\n", "\n", -1) //html2text.HTML2Text adds empty rows, for each <p></p> block causes the character missmatch between frontend and backend
	return htmlText, formattedPlainText, nil
}
