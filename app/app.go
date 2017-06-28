package main

import (
	"log"

	"github.com/blevesearch/bleve"
)

func main() {
	// open a new index
	// mapping := bleve.NewIndexMapping()
	// index, err := bleve.New("example.bleve", mapping)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// data := struct {
	// 	Name string
	// }{
	// 	Name: "text",
	// }

	// // index some data
	// index.Index("id", data)

	// // search for some text
	// query := bleve.NewMatchQuery("text")
	// search := bleve.NewSearchRequest(query)
	// searchResults, err := index.Search(search)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(searchResults)mapping := bleve.NewIndexMapping()
	// index, err := bleve.New("example.bleve", mapping)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// data := struct {
	// 	Name string
	// }{
	// 	Name: "text",
	// }

	// // index some data
	// index.Index("id", data)

	// // search for some text
	// query := bleve.NewMatchQuery("text")
	// search := bleve.NewSearchRequest(query)
	// searchResults, err := index.Search(search)
	// if err != nil {
	// 	fmt.Println(err)log.Println(s
	// 	return
	// }
	// fmt.Println(searchResults)

	// type message struct {
	// 	Id   string
	// 	From string
	// 	Body string
	// }

	// ms0 := message{
	// 	Id:   "example0",
	// 	From: "marty.schoch@gmail.com",
	// 	Body: "bleve indexing is easy",
	// }

	// ms1 := message{
	// 	Id:   "example1",
	// 	From: "marty.schoch@gmail.com",
	// 	Body: "bleve indexing is easy",
	// }
	// mapping := bleve.NewIndexMapping()
	// index, err := bleve.New("example.bleve", mapping)
	// if err != nil {
	// 	panic(err)
	// }
	// index.Index(ms0.Id, ms0)
	// index.Index(ms1.Id, ms1)

	// index.Close()

	indexN, _ := bleve.Open("example.bleve")
	query := bleve.NewQueryStringQuery("easy")
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"Id", "From", "Body"}
	searchResult, _ := indexN.Search(searchRequest)

	log.Println(searchResult)
	log.Println(searchResult.Hits[0].Fields["From"])

	// msg, _ := indexN.Document("example0")
	// log.Println(msg.GoString())
	indexN.Close()
}
