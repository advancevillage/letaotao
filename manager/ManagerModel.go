//author: richard
package manager

type DatabaseConfig struct {
	Host 		string 		`json:"host"`
	Port 	 	string 		`json:"port"`
	User 	 	string 		`json:"user"`
	Password	string 		`json:"password"`
	Schema		string 		`json:"schema"`
	CharSet		string 		`json:"charset"`
}
