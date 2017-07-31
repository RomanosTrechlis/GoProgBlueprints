package trace

import "io"

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}
