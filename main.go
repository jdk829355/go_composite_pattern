package main

import (
	"fmt"
)

// component
// 파일과 디렉토리에 들어갈 공통된 인터페이스
type Component interface {
	Search(string)
}

// leaf
// 파일 구조체
type File struct {
	Name string
}

// search 메서드를 구현했으므로 component 타입을 만족한다.
func (f File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.Name)
}

// composite
type Directory struct {
	Components []Component
	Name       string
}

// component 구현
func (d Directory) Search(keyword string) {
	for _, comp := range d.Components {
		comp.Search(keyword)
	}
}

// 복합 객체를 관리할 메서드
// component 추가
func (d *Directory) Add(c Component) {
	d.Components = append(d.Components, c)
}

// component 삭제
func (d *Directory) Remove(c Component) {

	for i, item := range d.Components {
		if item == c {
			d.Components = append(d.Components[:i], d.Components[i+1:]...)
		}
	}
}

func main() {
	// 디렉토리 선언
	var home, flutter, bin Directory

	// 파일 정의
	abc_py := File{"abc.py"}
	hello_python_py := File{"hello_python.py"}
	dart_sdk := File{"dart"}
	README_md := File{"README.md"}

	// home 디렉토리에 디렉토리 및 파일 삽입
	home.Add(flutter)
	home.Add(abc_py)
	home.Add(hello_python_py)

	// flutter 디렉토리에 component 삽입
	flutter.Add(bin)
	flutter.Add(README_md)

	// bin 디렉토리에 component 삽입
	bin.Add(dart_sdk)

	// 파일과 디렉토리에 대해 Search 메서드 호출
	README_md.Search("Google")
	/*
		Searching for keyword Google in file README.md
		Searching for keyword print('abc') in file abc.py
	*/
	home.Search("print('abc')")
	//Searching for keyword print('abc') in file hello_python.py
}
