package main

import (
	"fmt"
	"os"

	"mobydick/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Provide one file as an argument")
		return
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// обработка данных файла, все небуквенные символы заменены на новую строку
	// все буквенные символы приведены к нижнему регистру
	data, err := functions.DataManipulation(file)
	if err != nil {
		fmt.Print(err)
		return
	}

	// подсчет встречающихся слов
	wordList := functions.WordCount(data)

	// сортировка
	functions.Sort(wordList)

	for i := 0; i < 20; i++ {
		fmt.Printf("%5d %s\n", wordList[i].Count, wordList[i].Word)
	}
}
