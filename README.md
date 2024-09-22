## Ana

Anagram CLI is a simple command-line tool that generates anagrams from a list of words or a file.   
It leverages a dictionary file to find valid anagrams and outputs them in various formats (table, JSON, YAML, TOML).   

- [Install](###install)
- [Usage](###usage)
- [Dev](###dev)
- [Tests](###tests)




---


### Install
> Install the `ana` go package
```bash
go install github.com/moffgato/ana
```
> Optional but recommended. get a word list
```rb
sudo apt install wamerican

# word list for lookup will default to /usr/share/dict/words

op@xx ~/l/ana (main)> ll /usr/share/dict/
Permissions Size User Date Modified Name
.rw-r--r--  985k root 20 Jan  2022  american-english
.rw-r--r--   199 root  2 Dec  2021  README.select-wordlist
lrwxrwxrwx    30 root 22 Sep 00:22  words -> /etc/dictionaries-common/words
lrwxrwxrwx    16 root 20 Jan  2022  words.pre-dictionaries-common -> american-english
```

### Usage

```bash
ana g -h
```
```ts
Generate anagrams from word lists. You can input words directly or via a file.

Usage:
  ana generate [flags]

Aliases:
  generate, g, gen

Flags:
  -d, --dict string     Path to the word list (dictionary) (default "/usr/share/dict/american-english")
  -i, --file string     File path to read words from
  -f, --format string   Output format (table, json, yaml, toml) (default "table")
  -h, --help            help for generate
  -o, --output string   Output destination (file path or stdout)
  -p, --progress        Display a fancy progress bar
  -w, --words string    Comma-separated list of words
```
> Command completion is optional but available
```bash
ana completion -h
```
> example with fish
```
ana completion fish > ~/.config/fish/completions/ana.fish
```

```bash
Generate the autocompletion script for ana for the specified shell.
See each sub-command's help for details on how to use the generated script.

Usage:
  ana completion [command]

Available Commands:
  bash        Generate the autocompletion script for bash
  fish        Generate the autocompletion script for fish
  powershell  Generate the autocompletion script for powershell
  zsh         Generate the autocompletion script for zsh

Flags:
  -h, --help   help for completion

Use "ana completion [command] --help" for more information about a command.
```
```bash
op@xx ~/l/ana (main)> ana g -
-d  --dict         (Path to the word list (dictionary))  -i  --file                (File path to read words from)
-f  --format  (Output format (table, json, yaml, toml))  -o  --output  (Output destination (file path or stdout))
-h  --help                          (help for generate)  -w  --words              (Comma-separated list of words)
```


### Dev
> Run in dev mode, `air` reloads on file changes.
```bash
air
```
```bash
  __    _   ___
 / /\  | | | |_)
/_/--\ |_| |_| \_ v1.52.3, built with Go go1.23.0

watching .
watching cmd
watching pkg
watching pkg/printer
!exclude tmp
building...
running...
{
  "results": [
    {
      "word": "trader",
      "subsets": [
        "a",
        "ad",
        "d",
        "e",
        "r",
        "re",
        "t",
        "tad",
        "tar",
        "trade",
        "trader"
      ],
      "anagrams": [
        "retard",
        "tarred",
        "trader"
      ]
    },
```


### Test
> Find anagrams for average trader's aptitude
```
go test cmd/anagram_test.go
```

```bash
op@xx ~/l/ana (main)> go test cmd/anagram_test.go
ok      command-line-arguments  0.003s
```
