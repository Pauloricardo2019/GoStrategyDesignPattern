package main

import "fmt"

func getGmail(value string) {
	fmt.Println("Gmail: ", value)
}

func getOutlook(value string) {
	fmt.Println("Outlook: ", value)
}

var getPerson = map[string]map[string]func(value string){
	"email": {
		"gmail":   getGmail,
		"outlook": getOutlook,
	},
}

func main() {

	by := "email"
	provider := "outlook"
	value := "test@email.com"

	fn, ok := getPerson[by][provider]
	if !ok {
		fmt.Println("this method don't exits")
		return
	}

	fn(value)

	//switch by {
	//case "email":
	//	fmt.Println("email")
	//case "document":
	//	fmt.Println("document")
	//case "name":
	//	fmt.Println("name")
	//default:
	//	fmt.Println("this method don't exists")
	//	return
	//}

}
