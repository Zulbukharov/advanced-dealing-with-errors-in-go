package stacktrace

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

const maxStacktraceDepth = 32

type Frame uintptr

func (f Frame) pc() uintptr {
	return uintptr(f) - 1
}

func (f Frame) String() string {
	fun := runtime.FuncForPC(f.pc())
	if fun == nil {
		return ""
	}

	funcName := fun.Name()
	if v := strings.LastIndex(funcName, "/"); v != -1 {
		funcName = funcName[v+1:]
	}

	fileStr, line := fun.FileLine(f.pc())

	if n := strings.LastIndex(fileStr, "/"); n != -1 {
		if m := strings.LastIndex(fileStr[:n], "/"); m != -1 {
			fileStr = fileStr[m+1:]
		} else {
			fileStr = fileStr[n:]
		}
	}

	return fmt.Sprintf("%s\n%s:%d", funcName, fileStr, line)
}

type StackTrace []Frame

func (s StackTrace) String() string {
	b := strings.Builder{}
	for _, l := range s {
		b.WriteString(l.String())
		b.WriteByte('\n')
	}
	log.Println(b.String())
	// Реализуй меня.
	return b.String()
}

// Trace возвращает стектрейс глубиной не более maxStacktraceDepth.
// Возвращаемый стектрейс начинается с того места, где была вызвана Trace.
func Trace() StackTrace {
	pc := make([]uintptr, maxStacktraceDepth)
	n := runtime.Callers(2, pc)
	st := make(StackTrace, 0, n)
	log.Println(n, len(st))

	for _, c := range pc[:n] {
		st = append(st, Frame(c))
	}
	return st
}
