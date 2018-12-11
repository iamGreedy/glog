package glog

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Glogger struct {
	ref         *log.Logger
	indentDepth int
	indent      string
}

func New(log *log.Logger, indent string) *Glogger {
	return &Glogger{
		ref:         log,
		indent:      indent,
		indentDepth: 0,
	}
}

func (s *Glogger) SetOutput(w io.Writer) {
	s.SetOutput(w)
}
func (l *Glogger) Output(calldepth int, s string) error {
	return l.ref.Output(calldepth, strings.Repeat(l.indent, l.indentDepth)+s)
}

func (s *Glogger) Printf(format string, v ...interface{}) {
	s.Output(2, fmt.Sprintf(format, v...))
}
func (s *Glogger) Print(v ...interface{}) {
	s.Output(2, fmt.Sprint(v...))
}
func (s *Glogger) Println(v ...interface{}) {
	s.Output(2, fmt.Sprintln(v...))
}

func (s *Glogger) Fatal(v ...interface{}) {
	s.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}
func (s *Glogger) Fatalf(format string, v ...interface{}) {
	s.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
func (s *Glogger) Fatalln(v ...interface{}) {
	s.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

func (s *Glogger) Panic(v ...interface{}) {
	s.Output(2, fmt.Sprint(v...))
	panic(s)
}
func (s *Glogger) Panicf(format string, v ...interface{}) {
	s.Output(2, fmt.Sprintf(format, v...))
	panic(s)
}
func (s *Glogger) Panicln(v ...interface{}) {
	s.Output(2, fmt.Sprintln(v...))
	panic(s)
}

func (s *Glogger) Flags() int {
	return s.Flags()
}
func (s *Glogger) SetFlags(flag int) {
	s.SetFlags(flag)
}
func (s *Glogger) Prefix() string {
	return s.Prefix()
}
func (s *Glogger) SetPrefix(prefix string) {
	s.SetPrefix(prefix)
}

func (s *Glogger) Indent() *Glogger {
	return &Glogger{
		ref:         s.ref,
		indent:      s.indent,
		indentDepth: s.indentDepth + 1,
	}
}
