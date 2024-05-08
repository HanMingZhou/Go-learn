package main

import (
	"fmt"
	"go-learn/StructMode"
)

type circul struct {
	redius float64
}

type Book struct {
	title  string
	author string
	id     int
}

func main() {
	circulA := circul{10.1}
	fmt.Println(circulA.getAera())

	book := new(Book)
	book.storageBook()
	fmt.Println(book.FindBookById(1))

	fmt.Println(Book{"the game of throne", "john sonw", 7787})

	// 结构体指针
	var BookEnglish *Book
	BookEnglish = &Book{}
	fmt.Println("BookEnglish", BookEnglish)
	BookEnglish.title = "111"
	BookEnglish.author = "222"
	BookEnglish.id = 2
	fmt.Println("BookEnglish", BookEnglish)
	BookEnglish = &Book{"---", "2----2", 3}
	fmt.Println("BookEnglish", BookEnglish)

	var BookChinese Book = Book{"水浒传", "hhh", 111122}
	fmt.Println("BookChinese,before change", BookChinese)
	Change(&BookChinese)
	fmt.Println("BookChinese,after change", BookChinese)

	book1 := mode.Books{}
	book1.Title = "Hary Poter"
	book1.Author = "andam"
	book1.Subject = "history"
	fmt.Println("book1", book1)

}

func Change(b *Book) {
	b.id = 8812
}
func (c circul) getAera() float64 {
	return c.redius * c.redius * 3.14
}

// storage your book information
func (B *Book) storageBook() *Book {
	var Booktitle string
	var Bookauthor string
	var Bookid int

	fmt.Println("input book information:")
	fmt.Println("input book title:")
	fmt.Scanln(&Booktitle)
	fmt.Println("input book author:")
	fmt.Scanln(&Bookauthor)
	fmt.Println("input book id:")
	fmt.Scanln(&Bookid)

	B.title = Booktitle
	B.author = Bookauthor
	B.id = Bookid

	fmt.Println("input successfully")
	return B

}

// find your book information by id
func (B *Book) FindBookById(BookId int) (string, string, int) {

	fmt.Println("input the BookId you want to find out")
	fmt.Scanln(&BookId)
	if B.id == BookId {
		return B.title, B.author, B.id
	} else {
		fmt.Println("not found")
		return "", "", 0
	}

}
