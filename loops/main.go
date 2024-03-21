package main

func main() {
	// for i := 0; i <= 10; i++ {
	// 	println(i)
	// }

	// animals := []string{"dog", "fish", "horse", "cow", "cat"}

	// for i, animal := range animals {
	// 	println(i, animal)
	// }

	// for _, animal := range animals {
	// 	println(animal)
	// }

	// animals := make(map[string]string)
	// animals["dog"] = "John"
	// animals["fish"] = "Doe"
	// for animalType, animal := range animals {
	// 	println(animalType, animal)
	// }

	firstLine := "Once upon a midnight dreary"

	for i, l := range firstLine {
		println(i, l)
	}

}
