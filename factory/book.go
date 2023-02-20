package factory

import (
	"encoding/json"
	"fmt"
)

type IBook interface {
	setTitle(string)
	setAuthor(string)
	setPage(uint)
	setPrice(float64)
	getTitle() string
	getAuthor() string
	getPage() uint
	getPrice() float64
	print()
}

type Book struct {
	Title  string
	Author string
	Page   uint
	Price  float64
}

type Comics struct {
	PostCard IPostCard
	Book
}

type IPostCard interface {
	setPostCard(isPostCard bool)
	getPostCard() bool
}

type PostCard struct {
	IsPostCard bool
}

func (b *Comics) setTitle(title string)         { b.Title = title }
func (b *Comics) setAuthor(author string)       { b.Author = author }
func (b *Comics) setPage(page uint)             { b.Page = page }
func (b *Comics) setPrice(price float64)        { b.Price = price }
func (p *PostCard) setPostCard(isPostCard bool) { p.IsPostCard = isPostCard }

func (b *Comics) getTitle() string    { return b.Title }
func (b *Comics) getAuthor() string   { return b.Author }
func (b *Comics) getPage() uint       { return b.Page }
func (b *Comics) getPrice() float64   { return b.Price }
func (p *PostCard) getPostCard() bool { return p.IsPostCard }

func (b *Comics) print() {
	obj, _ := json.MarshalIndent(&b, "", "	")
	fmt.Println(string(obj))
}

func newComics() IBook {
	return &Comics{
		PostCard: &PostCard{
			IsPostCard: true,
		},
		Book: Book{
			Title:  "86",
			Author: "Asato Asato",
			Page:   300,
			Price:  325,
		},
	}
}

type Textbook struct {
	Book
}

func (b *Textbook) setTitle(title string)   { b.Title = title }
func (b *Textbook) setAuthor(author string) { b.Author = author }
func (b *Textbook) setPage(page uint)       { b.Page = page }
func (b *Textbook) setPrice(price float64)  { b.Price = price }

func (b *Textbook) getTitle() string  { return b.Title }
func (b *Textbook) getAuthor() string { return b.Author }
func (b *Textbook) getPage() uint     { return b.Page }
func (b *Textbook) getPrice() float64 { return b.Price }

func (b *Textbook) print() {
	obj, _ := json.MarshalIndent(&b, "", "	")
	fmt.Println(string(obj))
}

func newTextbook() IBook {
	return &Textbook{
		Book: Book{
			Title:  "Linear Algebra",
			Author: "John Doe",
			Page:   1000,
			Price:  19.99,
		},
	}
}

func newBook(bookType string) (IBook, error) {
	switch bookType {
	case "comics":
		return newTextbook(), nil
	case "textbook":
		return newComics(), nil
	}
	return nil, fmt.Errorf("unknown book type")
}

func Output() {
	// Create comics
	comics, err := newBook("comics")
	if err != nil {
		panic(err)
	}

	// Create textbook
	textbook, err := newBook("textbook")
	if err != nil {
		panic(err)
	}

	comics.print()
	textbook.print()
}
