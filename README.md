# Combinoctocat

This is a little _work in progress_ tool to output/display all the possible Octocat combinations
that you can make on [myoctocat.com](https://myoctocat.com/build-your-octocat/). I wanted to write
this after [@jeffrafter](https://github.com/jeffrafter) nerd-sniped me.

## How to develop

I built this using Go version 1.13.4 in macOS.

1. Load [myoctocat.com](https://myoctocat.com/build-your-octocat/) in your browser and save the page as an HTML file.
1. `make && bin/combinoctocat -in PATH-TO-MYOCTOCAT-PAGE.html`

Sample output:

```sh
% make && bin/combinoctocat -in my-octocat.html
go build -o bin/combinoctocat combinoctocat.go
Octocat customization options:

18 body choices
456 eye choices
24 face choices

962 hair choices (including none)
- 31 colors
- 31 styles

280 Octocat facial hair choices (including none)
- 31 colors
- 9 styles

18 mouth choices
37 prop choices

262,144 accessory combinations

30,280,800 outfit choices (50 tops, 33 bottoms, 37 headgears, 16 eyewears, 31 footgears)
```
