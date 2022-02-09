package proxy

import (
	"log"
	"time"
)

type ProxyCompiler struct {
	rc *RealCompiler
}

func NewProxyCompiler(rc *RealCompiler) Compiler {
	return &ProxyCompiler{rc}
}

func (p *ProxyCompiler) Compile(code []string) string {
	log.SetFlags(0)
	start := time.Now()
	log.Printf("Started compiling at %v", start)

	res := p.rc.Compile(code)

	log.Printf("Finished at %v", time.Now())
	log.Printf("It took %v", time.Since(start))

	return res
}
