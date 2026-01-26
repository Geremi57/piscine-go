# Ascii -art

This is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII

for example\
`go run . "Hello" | cat -e`

```
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
```

to run this program you would need the **standard txt** provided in the project folder\

and read it by means of os\

`os.ReadFile(standard.txt)`

then we will read the contents and split them\

`lines := strings.Split(string(arguments), "\n")`

after initializing an empty string\

word := ""\

we want to place all the text submitted as input via `os.Args` into the variable\
	
	for i:=1; i < len(os.Args); i++{
		word += os.Args[i]
	}

we want to replace `\n` with "\n"\
Because when we receive an input like\

go run . Hello\nThere\

this is not a real newline it receives two characters "\" and "n"

it is just text not an actual line break yet this is why we ran\

`word = strings.ReplaceAll(word, `\n`, "\n"`

with `\n` this is actual text but with "\n" this is a newline character (ascii  10)\
	
what ReplaceAll does it scans the whole string and replaces `\n` with "\n"\

	parts := strings.Split(word, "\n")

Now that the newline works after replacingAll we now want to split the word when we get a "\n" this creates an empty line after the "\n"

Creating something like\
```
parts = []string{
    "Hello"
    "There,
}
```

next is the core **loop**

```
for _, part := range parts {
	for row := 0; row <= 8; row++ {
		for _, ch := range part {
			if ch < 32 || ch > 126 {
				continue
			}
			start := (int(ch) - 32) * 9
			fmt.Print(lines[start+row])
		}
		fmt.Println()
	}
}
```

here we are printing ascii text which works like this

Each character is drawn using **8 line**

Those 8 lines are stored one after another in **standard.txt**

Each character block is **9 lines tall**(8 drawing lines + 1 empty separator)

* Print row 0 of all letters
* Then row 1 of all letters
* ...
* Then row 7 of all letter

```
for _,part := range parts
```
Remeber parts are what we slpit the string with using "\n" character to create a slice of number of words excluding the "\n"\
`parts := string.Split(words, "\n")`

Each par is one line of text\

```
for row:=0; row <= 8; row++
```
This loop Prints the ascii art row by row, not letter by letter

* Row 0: H row0 + E row0 + L row0
* Row 1: H row1 + E row1 + L row1
* ...
* Row 8: H row8 + E row8 + L row8

```
for _,ch := range part
```
Now we go character by character in the word.

```
if part = "Hi"

```

* ch = 'H'
* then ch = 'i'

this loop goes for this row, print the row of Each charater in the word

Calculating where the characetr lives in the **Standard.txt** file

```
start := (int(ch) - 32) * 9
```

this is the most important line

`-32`\
Ascii font start at space(32)\

`int(ch) - 32`

gives the character number in the font\

`*9`

Each character takes 9 lines in the file\

* 8 drawing lines
* 1 empty separator line

character block start line = index x 9

`A`\
```
int('A') = 65

(65-32) * 9 = 297
```

so 'A' starts at

lines[297]

```
fmt.Print(lines[start+row])
```
This prints\
* the current row
* of the current character

if\
* row = 3
* ch = 'A'

then it prints:

lines[297+3]

which is
row 3 of the ascii art for A

`fmt.Println()`
Newline after finishing one row after all characters have printed