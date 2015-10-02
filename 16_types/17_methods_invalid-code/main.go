package main

type Map map[string]string

//this is valid
func (m Map) Set(key string, value string) {
	m[key] = value
}

//this is invalid
func (m map[string]string)) Set(key string, value string) {
	m[key] = value
}

func main() {

}
