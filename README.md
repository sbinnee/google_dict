# Dictionary project
It parses json output either from [`sdcv`] and a 3rd party Google dictionary API
[dictionary.dev] and append the searched word to a file.

Toy project to learn Golang. I mainly use [`sdcv`] for cli dictionary tool. But
if I couldn't find a word, then I tend to open a browser, go to Google, type
dictionary, and type in the word, which is a very annoying process. This simple
Go program does that in my terminal with two words.

I don't like relying on 3rd party API but I'm going to just use it until I find
some other solutions.

I use a bash script to wrap [`sdcv`] and `parse_dict` and save the word I searched
in a file so that I can memorize more words.

## Program structure
1. Parse json from [`sdcv`]

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

## Installation
1. Install `sdcv`

2. Download dictionaries

    Go to https://web.archive.org/web/20200702000038/http://download.huzheng.org/ and
    download dictionaries to use.

3. Place the dictionaries in a folder and set `$STARDICT_DATA_DIR`


## Usage
```bash
# sdcv
sdcv --non-interactive --utf8-output --color --json-output "$word" \
    | COLUMNS="${$(tput cols):-72}" parse_dict

# google dict
curl --silent "https://api.dictionaryapi.dev/api/v2/entries/${LANG:-en_US}/$word}" \
    | COLUMNS=$COLUMNS parse_dict -google
```




## TODO
- [ ] Let user configure the history file path through a flag
- [ ] Multiple words found e.g.) try "unpair" to `sdcv`
- [ ] No word found

<!-- Links -->
[`sdcv`]: https://dushistov.github.io/sdcv/
[dictionary.dev]: https://dictionaryapi.dev/
