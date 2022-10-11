// Models
package models

// type Book struct {
// 	Id     int    `json:"id"`
// 	Title  string `json:"title"`
// 	Author string `json:"author"`
// 	Desc   string `josn:"desc"`
// }
// new models for postgresql db
type Book struct {
	Id     int    `json:"id" gorm: "primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}
