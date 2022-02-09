package proxy

import "time"

type Compiler interface {
	Compile(code []string) string
}

type RealCompiler struct {
}

func (rc *RealCompiler) Compile(code []string) string {
	buf := ""
	for _, v := range code {
		// lets imagine there's something more complex going on
		time.Sleep(time.Millisecond * 20)
		buf += v + "\n"
	}

	return buf
}
