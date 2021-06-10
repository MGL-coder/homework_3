package main

import (
	"fmt"
	"github.com/MGL-coder/homework_3/atoi"
	"github.com/MGL-coder/homework_3/fibonacci"
	"github.com/MGL-coder/homework_3/itoa"
	"github.com/MGL-coder/homework_3/reverse"
	"github.com/MGL-coder/homework_3/runeByIndex"
	"github.com/MGL-coder/homework_3/sortImports"
	"log"
)

func main() {
	fmt.Println()

	// 1. Article
	fmt.Println("1. Article has been read :)")

	// 2. Atoi(n int) string
	str := "-4252"
	n, err := atoi.Atoi(str)
	if err != nil {
		log.Fatalf("converting string to integer error: %s\n", err)
	}
	fmt.Printf("2. Atoi(\"%s\") = %v\n", str, n)

	// 3. Itoa(s string) (int, error)
	n = -231
	fmt.Printf("3. Itoa(%v) = \"%s\"\n", n, itoa.Itoa(n))

	// 4. Reverse(s string) string
	str = "Алфавит! Alphabet!"
	fmt.Printf("4. Reverse(%s) = %s\n", str, reverse.Reverse(str))

	// 5. SortImports(path string) err
	path := "importSortTestFile.go"
	err = sortImports.Sort(path)
	if err != nil {
		log.Fatalf("could not sort imports of the .go file: %s\n", err)
	}
	fmt.Printf("5. Imported packages list in the file \"%s\" were sorted\n", path)

	// 6. Fibonacci() func() int
	fmt.Print("6. Fibonacci sequence: ")
	generator := fibonacci.Fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(generator(), " ")
	}
	fmt.Println()

	// 7. runeByIndex(s *string, i *int) (rune, error)
	str = "Антарктида"
	i := 9
	ch, err := runeByIndex.RuneByIndex(&str, &i)
	if err != nil {
		log.Fatalf("could not get rune by index: %s", err)
	}
	fmt.Printf("7. RuneByIndex(%s, %v) = %c\n", str, i, ch)
}
