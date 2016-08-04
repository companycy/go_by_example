

// By using this form of declaration, you clearly and explicitly define the hierarchy of the JSON document you are parsing. Keep in mind that the declaration doesn’t need to define all fields, but just the ones you will be using.

{
	"id": 1,
	"name": "Tit Petric",
	"address": {
		"street": "Viska cesta 49c",
		"zip": "1000",
		"city": "Ljubljana",
		"country": "Slovenia"
	}
}

type Person struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
}


// Anonymous structs
person := struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Address struct {
		City string `json:"city"`
		Country string `json:"country"`
	} `json:"address"`
}{}

err = json.Unmarshal(contents, &person)

