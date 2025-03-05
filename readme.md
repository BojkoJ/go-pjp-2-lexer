## Lexikální analyzátor

---

Napište program, který přečte vstup a převede jej na **posloupnost lexikálních symbolů - tokenů**. Každý token je dvojice, skládá se z typu a případně hodnoty.

Definice tokenů závisí na vás a je považována za součást řešení.

### Specifikace vstupu

Vstup může obsahovat následující symboly:

- **identifikátory** - skládají se z posloupnosti písmen a číslic začínající písmenem.
- **čísla** - tvořená posloupností desetinných číslic.
- **operátory** - symboly `'+', '-', '*' a '/'`
- **oddělovače** - symboly `'(', ')' a ';'`
- **klíčová slova** - `div` a `mod`.

Symboly mohou být odděleny posloupností mezer, tabulátorů a zlomů řádků.

Ve vstupu mohou být poznámky. Poznámkám předchází sekvence `\\` a pokračují až na konec řádku.

_Bílé mezery a poznámky nevytvářejí žádné lexikální symboly._

### Specifikace výstupu

Převede zadaný vstup na posloupnost tokenů a zapíše je na výstup. Každý token zapíše na samostatný řádek.

### Příklad:

**vstup:**

```
    -2 + (245 div 3);  // note
2 mod 3 * hello
```

**výstup:**

```
OP:-
NUM:2
OP:+
LPAR
NUM:245
DIV
NUM:3
RPAR
SEMICOLON
NUM:2
MOD
NUM:3
OP:*
ID:hello
```

### Spuštění a kompilace:

- Je třeba mít nainstalovaný **[Golang](https://go.dev/)**
- Zkompilujte kód pomocí `go build ./main.go`
- Spusťte zkompilovanou binárku `./main` (nebo `./main.exe`)
- Otestujte lexikální analyzátor
