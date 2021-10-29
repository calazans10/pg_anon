package main

func main() {
	fields := "name:name,email"

	p := NewProcessor("", "", fields)
	p.Run()
}
