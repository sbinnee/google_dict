# Dictionary project
It parses json output either from [`sdcv`](https://dushistov.github.io/sdcv/)
and a 3rd party Google dictionary API
[dictionary.dev](https://dictionaryapi.dev/) and append the searched word to a
file. 

Toy project to learn Golang. I mainly use `sdcv` for cli dictionary tool. But if
I couldn't find a word, then I tend to open a browser, go to Google, type
dictionary, and type in the word, which is a very annoying process. This simple
Go program does that in my terminal with two words. 

I don't like relying on 3rd party API but I'm going to just use it until I find
some other solutions. 

I use a bash script to wrap `sdcv` and `parse_dict` and save the word I searched
in a file so that I can memorize more words. 

## Program structure
1. Parse json from `sdcv`

If `--google` flag is given, 
1. Use Google dicdtionary 3rd party API
	- https://dictionaryapi.dev/
2. Get request
3. Parse json response
	- https://github.com/buger/jsonparser
4. Display it pretty
	- https://github.com/liamg/tml
	- https://github.com/gookit/color
	- https://github.com/1dot75cm/gocolor

