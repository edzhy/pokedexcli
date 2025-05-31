package main
import (
	"fmt"
	"strings"
	"bufio"
	"os"
)
var supported_elements map[string]cliCommand
func init() {
    supported_elements = map[string]cliCommand{
        "exit": {
            name:        "exit",
            description: "Exit the Pokedex",
            callback:    commandExit,
        },
        "help": {
            name:        "help",
            description: "Displays a help message",
            callback:    commandHelp,
        },
		"map": {
            name:        "map",
            description: "Displays names of locations",
            callback:    commandMap,
        },
    }
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		str := scanner.Text()
		str_slice := cleanInput(str)
		first_word := str_slice[0]	
		
		cmd, exists := supported_elements[first_word]
		
		if !exists {
			fmt.Printf("Unknown command\n")
		} else {
		err := cmd.callback()
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	}
}
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: \n")
	for ele := range supported_elements {
		fmt.Printf("%v: %v\n", supported_elements[ele].name, supported_elements[ele].description)
	}
	return nil
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
type cliCommand struct {
	name string
	description string
	callback func() error
}

func cleanInput(text string) []string {
	var res []string
	workstr := strings.Trim(strings.ToLower(text), " ")
	res = strings.Split(workstr, " ")
	return res
}