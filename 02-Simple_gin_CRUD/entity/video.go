package entity

//not compatible with reflect.StructTag.Get: bad syntax for struct tag valuego-vet
//use json:"field" without space between json: and "field"
type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
