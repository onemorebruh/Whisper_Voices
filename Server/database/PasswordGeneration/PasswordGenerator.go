//██████╗░░█████╗░░██████╗░██████╗░██╗░░░░░░░██╗░█████╗░██████╗░██████╗░░░░░░░░░░██████╗░███████╗███╗░░██╗███████╗██████╗░░█████╗░████████╗░█████╗░██████╗░
//██╔══██╗██╔══██╗██╔════╝██╔════╝░██║░░██╗░░██║██╔══██╗██╔══██╗██╔══██╗░░░░░░░░██╔════╝░██╔════╝████╗░██║██╔════╝██╔══██╗██╔══██╗╚══██╔══╝██╔══██╗██╔══██╗
//██████╔╝███████║╚█████╗░╚█████╗░░╚██╗████╗██╔╝██║░░██║██████╔╝██║░░██║░░░░░░░░██║░░██╗░█████╗░░██╔██╗██║█████╗░░██████╔╝███████║░░░██║░░░██║░░██║██████╔╝
//██╔═══╝░██╔══██║░╚═══██╗░╚═══██╗░░████╔═████║░██║░░██║██╔══██╗██║░░██║░░░░░░░░██║░░╚██╗██╔══╝░░██║╚████║██╔══╝░░██╔══██╗██╔══██║░░░██║░░░██║░░██║██╔══██╗
//██║░░░░░██║░░██║██████╔╝██████╔╝░░╚██╔╝░╚██╔╝░╚█████╔╝██║░░██║██████╔╝███████╗╚██████╔╝███████╗██║░╚███║███████╗██║░░██║██║░░██║░░░██║░░░╚█████╔╝██║░░██║
//╚═╝░░░░░╚═╝░░╚═╝╚═════╝░╚═════╝░░░░╚═╝░░░╚═╝░░░╚════╝░╚═╝░░╚═╝╚═════╝░╚══════╝░╚═════╝░╚══════╝╚═╝░░╚══╝╚══════╝╚═╝░░╚═╝╚═╝░░╚═╝░░░╚═╝░░░░╚════╝░╚═╝░░╚═╝

// This is a class programm uses for generating passwords and invites

package PasswordGenerator

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var Password_generator PasswordGenerator

type PasswordGenerator struct {
	dictionary [1000]string
}

func Init() PasswordGenerator {
	//read dictionary
	var dictionary [1000]string
	var i int = 0
	file, err := os.Open("dictionary.txt")

	if err != nil {
		fmt.Println("password dictionary: false")
		fmt.Println("dictionary.txt is not loaded")
		fmt.Println(err.Error())
		path, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println("current dirctory is:", path)
		dictionary[0] = "dictionary"
		dictionary[1] = "is"
		dictionary[1] = "not"
		dictionary[1] = "loaded"
	} else {
		fmt.Println("password dictionary: true")
	}
	file_scanner := bufio.NewScanner(file)

	file_scanner.Split(bufio.ScanLines)

	for file_scanner.Scan() {
		dictionary[i] = file_scanner.Text()
		i++
	}

	password_generator := PasswordGenerator{
		dictionary: dictionary,
	}
	return password_generator
}

func (password_generator *PasswordGenerator) charge(words string, max_length int) string {
	var the_biggest_element_index int
	var the_biggest_element_length int

	var words_array []string //programm need it to delete unnecessary word when length of words is bigger than max_length

	for len(words) < max_length {
		//insert random word
		rand.Seed(time.Now().UnixNano())
		new_word := password_generator.dictionary[rand.Intn(1000)]
		if rand.Int() == 1 {
			new_word = strings.ToUpper(new_word)
		}
		words_array = append(words_array, new_word)
		if len(new_word) > the_biggest_element_length { // It is here to not go through all the words after finish just to make words combination fit in 32 charachters
			the_biggest_element_length = len(new_word)
			the_biggest_element_index = len(words_array)
		}
		words = strings.Join(words_array, "")
	}
	words = strings.Join(words_array, "")

	if len(words) > max_length {
		words = strings.Join(append(words_array[:the_biggest_element_index], words_array[(the_biggest_element_index+1):]...), "")
	}

	return words
}

func (password_generator *PasswordGenerator) Create_password() string {
	var password string
	password = password_generator.charge(password, 32)
	return password
}
