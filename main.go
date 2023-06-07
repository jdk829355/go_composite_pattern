package main

import (
	"fmt"
)

// component
// 파일과 디렉토리에 들어갈 공통된 인터페이스
type component interface {
	search(string)
}

// leaf
// 파일 구조체
type file struct {
	name string
}

// search 메서드를 구현했으므로 component 타입을 만족한다.
func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

// composite
type directory struct {
	components []component
	name       string
}

// component 구현
func (d *directory) search(keyword string) {
	for _, comp := range d.components {
		comp.search(keyword)
	}
}

// 복합 객체를 관리할 메서드
// component 추가
func (d *directory) add(c component) {
	f.components = append(f.components, c)
}

// component 삭제
func (d *directory) remove(c component) {

	for i, item := range d.components {
		if item == c {
			d.components = append(f.components[:i], f.components[i+1:]...)
		}
	}
}

func main() {

}
