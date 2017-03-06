package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"

	"github.com/matryer/goblueprints/chapter1/trace"
)

// templ은 하나의 템플릿을 나타냄
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP가 HTTP 요청을 처리한다.
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}

func main() {
	var addr = flag.String("addr", ":8080", "애플리케이션의 주소이다."")
	flag.Parse() // 플래그 파싱

	r := newRoom()
	r.tracer = trace.New(os.Stdout)

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)

	// room을 가져옴
	go r.run()

	// 웹 서버 시작
	log.Println("다음 주소에서 웹 서버 시작", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}
