package main

import (
	"bufio"   // Balíček pro čtení vstupu řádek po řádku
	"fmt"     // Balíček pro formátovaný výstup na konzoli
	"os"      // Balíček pro práci se systémovými operacemi a vstupy/výstupy
	"strings" // Balíček pro práci s řetězci
	"unicode" // Balíček pro práci s Unicode znaky
)

// Definice typů tokenů - konstanty pro různé druhy tokenů
const (
	NUMBER    = "NUM"  // Čísla
	ID        = "ID"   // Identifikátory (např. proměnné jako hello, x, ...)
	OP        = "OP"   // Operátory (+, -, *, /)
	LPAR      = "LPAR" // (
	RPAR      = "RPAR" // )
	SEMICOLON = "SEMICOLON"
	DIV       = "DIV"     // 'div' - celočíselné dělení
	MOD       = "MOD"     // 'mod' - zbytek po dělení (modulo)
	EOF       = "EOF"     // Konec vstupu
	INVALID   = "INVALID" // Neplatný token
)

// Token reprezentuje jednu lexikální jednotku - složenou z typu a hodnoty
type Token struct {
	Type  string // Typ tokenu (např. NUMBER, ID, OP, ...)
	Value string // Hodnota tokenu (např. "42", "hello", "+", ...)
}

// Lexer je zodpovědný za konverzi vstupního řetězce na tokeny - náš lexikální analyzátor
type Lexer struct {
	input       string // Vstupní řetězec, který chceme analyzovat
	position    int    // Aktuální pozice v řetězci (index)
	currentChar rune   // Aktuální znak, který zpracováváme
}

// NewLexer vytvoří nový lexikální analyzátor pro daný vstupní řetězec
func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input, position: 0} // Vytvoříme nový lexer a nastavíme vstup a pozici na začátek

	if len(input) > 0 {
		lexer.currentChar = rune(input[0]) // Pokud má vstup aspoň jeden znak, nastavíme ho jako aktuální
	} else {
		lexer.currentChar = 0
	}

	return lexer
}

// Metoda advance posune lexer na další znak ve vstupním řetězci
func (l *Lexer) advance() {
	l.position++ // Zvýšíme pozici o 1

	if l.position < len(l.input) { // Pokud nejsme na konci vstupu
		l.currentChar = rune(l.input[l.position]) // Načteme další znak
	} else {
		l.currentChar = 0 // Jinak nastavíme konec vstupu (EOF)
	}
}

// Metoda skipWhitespace přeskočí všechny bílé znaky (mezery, tabulátory, nové řádky)
func (l *Lexer) skipWhitespace() {
	for l.currentChar != 0 && unicode.IsSpace(l.currentChar) { // Dokud (while) máme hodnotu v currentChar a je to bílý znak
		l.advance() // Posuneme se na další znak
	}
}

// Metoda skipComment přeskočí komentáře, které začínají // a končí koncem řádku
func (l *Lexer) skipComment() {
	// Přeskočíme znaky '//'
	l.advance() // Přeskočíme první '/'
	l.advance() // Přeskočíme druhý '/'

	// Pokračujeme dokud nenarazíme na konec řádku nebo konec vstupu
	for l.currentChar != 0 && l.currentChar != '\n' {
		l.advance() // Přeskakujeme všechny znaky v komentáři
	}

	// Přeskočíme i znak konce řádku, pokud na něj narazíme
	if l.currentChar == '\n' {
		l.advance()
	}
}

// Metoda readNumber načte celé číslo ze vstupu - sekvenci číslic
func (l *Lexer) readNumber() string {
	num := ""

	for l.currentChar != 0 && unicode.IsDigit(l.currentChar) { // Dokud máme číslice
		num += string(l.currentChar) // Přidáme aktuální číslici do řetězce
		l.advance()                  // Posuneme se na další znak
	}

	return num
}

// Metoda readIdentifier načte identifikátor nebo klíčové slovo - sekvenci písmen a číslic začínající písmenem
func (l *Lexer) readIdentifier() string {
	id := ""

	// Dokud (while) je máme v currentChar hodnotu jinou než 0 a je to písmeno nebo číslice
	for l.currentChar != 0 && (unicode.IsLetter(l.currentChar) || unicode.IsDigit(l.currentChar)) {
		id += string(l.currentChar) // Přidáme aktuální znak do identifikátoru
		l.advance()                 // Posuneme se na další znak
	}

	return id
}

// Metoda peek se podívá na další znak bez posunutí pozice - tzv. "kouknutí dopředu" - použito pro komentáře
func (l *Lexer) peek() rune {
	peekPos := l.position + 1 // Pozice následujícího znaku

	if peekPos >= len(l.input) { // Pokud jsme za koncem vstupu
		return 0 // Vrátíme EOF
	}

	return rune(l.input[peekPos]) // Jinak vrátíme následující znak
}

// Metoda getNextToken vrací další token ze vstupního řetězce - hlavní metoda lexikálního analyzátoru
func (l *Lexer) getNextToken() Token {

	for l.currentChar != 0 { // Dokud nejsme na konci vstupu

		// Přeskočíme bílé znaky (mezery, tabulátory, konce řádků)
		if unicode.IsSpace(l.currentChar) {
			l.skipWhitespace()
			continue
		}

		// Přeskočíme komentáře začínající "//"
		if l.currentChar == '/' && l.peek() == '/' {
			l.skipComment()
			continue
		}

		// Zpracování číslic - pokud narazíme na číslici, načteme celé číslo
		if unicode.IsDigit(l.currentChar) {
			return Token{NUMBER, l.readNumber()} // Vrátíme token s typem NUMBER a hodnotou čísla
		}

		// Zpracování identifikátorů a klíčových slov - pokud narazíme na písmeno
		if unicode.IsLetter(l.currentChar) {
			id := l.readIdentifier() // Načteme celý identifikátor

			// Kontrola, zda jde o klíčové slovo
			switch strings.ToLower(id) {
			case "div": // Klíčové slovo 'div'
				return Token{DIV, ""} // Vrátíme token DIV bez hodnoty
			case "mod": // Klíčové slovo 'mod'
				return Token{MOD, ""} // Vrátíme token MOD bez hodnoty
			default: // Běžný identifikátor
				return Token{ID, id} // Vrátíme token ID s hodnotou identifikátoru
			}
		}

		// Zpracování operátorů a oddělovačů - jeden znak
		switch l.currentChar {
		case '+', '-', '*', '/': // Operátory
			op := string(l.currentChar) // Uložíme si operátor jako string
			l.advance()
			return Token{OP, op}
		case '(': // Levá závorka
			l.advance()
			return Token{LPAR, ""}
		case ')': // Pravá závorka
			l.advance()
			return Token{RPAR, ""}
		case ';': // Středník
			l.advance()
			return Token{SEMICOLON, ""}
		default: // Neznámý znak - chyba
			// Neplatný znak - vrátíme token INVALID s hodnotou znaku
			invalidChar := string(l.currentChar)
			l.advance()
			return Token{INVALID, invalidChar}
		}
	}

	return Token{EOF, ""} // Pokud jsme došli na konec vstupu, vrátíme token EOF bez hodnoty
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var inputLines []string // Pole pro ukládání řádků vstupu

	fmt.Printf("Enter your sequence:\n")

	for scanner.Scan() {
		line := scanner.Text() // Načteme řádek
		if line == "" {        // Pokud je řádek prázdný, ukončíme čtení
			break
		}
		inputLines = append(inputLines, line) // Přidáme řádek do pole
	}

	// Spojíme všechny řádky vstupu do jednoho řetězce s konci řádků
	input := strings.Join(inputLines, "\n")

	// Vytvoříme nový lexikální analyzátor
	lexer := NewLexer(input)

	for {
		token := lexer.getNextToken() // Získáme další token

		if token.Type == EOF { // Pokud jsme na konci vstupu, ukončíme zpracování
			break
		}

		// Vypíšeme token - buď s hodnotou nebo bez ní
		if token.Value != "" {
			fmt.Printf("%s:%s\n", token.Type, token.Value)
		} else {
			fmt.Println(token.Type)
		}
	}
}
