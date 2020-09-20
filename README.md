# Combinoctocat

This is a little _work in progress_ tool to output/display all the possible Octocat combinations
that you can make on [myoctocat.com](https://myoctocat.com/build-your-octocat/). I wanted to write
this after [@jeffrafter](https://github.com/jeffrafter) nerd-sniped me.

## How to develop

I built this using Go version 1.13.4 in macOS.

1. Load [myoctocat.com](https://myoctocat.com/build-your-octocat/) in your browser and save the page as an HTML file.
1. `make && bin/combinoctocat -in PATH-TO-MYOCTOCAT-PAGE.html`
