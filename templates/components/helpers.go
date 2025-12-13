package components

import (
	"context"
	"io"

	"github.com/a-h/templ"
)

func rawHTML(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := w.Write([]byte(html))
		return err
	})
}

