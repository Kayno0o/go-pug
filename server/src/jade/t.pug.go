// Code generated by "jade.go"; DO NOT EDIT.

package jade

import (
	"io"
	"strconv"
)

const (
	t__0 = `<!DOCTYPE html><html><head><link rel="stylesheet" href="./public/css/main.css"/><title>`
	t__1 = `</title></head><body><h1>Jade - template engine`
	t__2 = `</h1><div class="container text-accent">`
	t__3 = `<p>      Jade/Pug is a terse and simple templating language with a <strong>focus</strong> on performance and powerful features.</p><div class="class-test">      ok y</div><div class="i-ph-house"></div><p>bonsoir jsais pas</p></div></body></html>`
	t__4 = `<pre>Precompile jade templates to `
	t__5 = ` code.</pre>`
	t__6 = `<p>You are amazing... </p>`
	t__7 = `<p>Get on it !!</p>`
)

func Index(values map[string]string, wr io.Writer) {
	buffer := &WriterAsBuffer{wr}

	buffer.WriteString(t__0)
	WriteEscString(values["title"], buffer)
	buffer.WriteString(t__1)

	{
		var (
			golang = "Go"
		)

		buffer.WriteString(t__4)
		WriteEscString(golang, buffer)
		buffer.WriteString(t__5)
	}

	buffer.WriteString(t__2)

	val, _ := strconv.ParseBool(values["jade"])
	if val {
		buffer.WriteString(t__6)

	} else {
		buffer.WriteString(t__7)

	}
	buffer.WriteString(t__3)

}