package models

type Video struct {
	Id 			int		`json:"id"`
	Src			string	`json:"src"`
	Title 		string	`json:"title"`
	Artist 		string	`json:"artist"`
	Sifat		bool	`json:"sifat"`
	Deskripsi	string 	`json:"deskripsi"`
}