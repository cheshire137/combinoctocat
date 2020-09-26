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
18 mouth choices
37 prop choices
262,144 accessory combinations

962 hair choices (including none)
- 31 colors
- 31 styles

280 facial hair choices (including none)
- 31 colors
- 9 styles

30,280,800 outfit choices
- 50 tops
- 33 bottoms
- 37 headgears
- 16 eyewears
- 31 footgears

3,869,803,707,365,326,848 possible Octocats

three quintrillion eight hundred sixty-nine quadrillion eight hundred three trillion seven hundred seven billion three hundred sixty-five million three hundred twenty-six thousand eight hundred forty-eight

What would you like to do?
[g]enerate an Octocat
[q]uit
g

Octocat: Body: #5F6196, #C69848, #E1CFC5
- Face: #9DA2C9 / Nose: #52467A
- Eyes: Confused / #B0A83E
- Mouth: Underbite
- Hair: Flattop / #E245BF
- Facial hair: Goatee / #724437
- Outfit: barbarianvest / Maxi Skirt / French shoes / Snapback / MardiGra Mask
- Prop: Ruby Gem
- Accessories (10): Camera, Cape, Cowboy Buckle Belt, Dress Buckle Belt, Eyebrow Ring, Freckles, Messenger Bag, Nose Ring, Pearls, Tattoo

What would you like to do?
[g]enerate an Octocat
[q]uit
g

Octocat: Body: #9E4D9E, #B1BDC9, #DFC562
- Face: #AB806C / Nose: #453E2D
- Eyes: Chillin / #5A0630
- Mouth: Express
- Hair: Afro Tight / #2C3EC1
- Facial hair: Cowboy Mustache / #9E4D9E
- Outfit: maxidress / Cyclist Shorts / converseshoes / Superhero Crown / no eyewear
- Prop: Magnifying Glass
- Accessories (7): Bolo Tie, Bow Tie, Bracelet, Cape, Cowboy Buckle Belt, Dress Buckle Belt, Eyebrow Ring

What would you like to do?
[g]enerate an Octocat
[q]uit
q
Goodbye
```

## How to test

```sh
make test
```
