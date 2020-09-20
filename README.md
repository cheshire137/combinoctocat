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
Octocat body choices (18):
- Body: #282828, #494949, #99C7C9
- Body: #282828, #99C7C9, #494949
- Body: #52467A, #9DA2C9, #D6BCA5
- Body: #6B438C, #B7D5A5, #E8D575
- Body: #7F5797, #A69CB1, #EEBCB8
- Body: #9E4D9E, #B1BDC9, #DFC562
- Body: #941D1D, #CE8D8D, #D6D58B
- Body: #B0591A, #F4E1A0, #7A9E98
- Body: #874E23, #E6BD9F, #B7E7EA
- Body: #695131, #FDE5C6, #BD7586
- Body: #6D6D63, #E2BD43, #7A9E98
- Body: #AD943D, #F4E1A0, #7A9E98
- Body: #84872F, #CDCF8A, #876D5D
- Body: #557542, #B7D5A5, #D8DAF0
- Body: #457A6D, #B98DBD, #A1D4C8
- Body: #46828E, #494959, #7C6983
- Body: #4F6AB0, #89B7CC, #694E61
- Body: #5F6196, #C69848, #E1CFC5

456 Octocat eye choices

Octocat face choices (24):
- Face: #FFC5AF / Nose: #A34F3E
- Face: #D9B6A5 / Nose: #814B41
- Face: #c8ab9f / Nose: #73382C
- Face: #AB806C / Nose: #453E2D
- Face: #8E5740 / Nose: #453E2D
- Face: #F7EFCD / Nose: #A37D76
- Face: #E0DBB8 / Nose: #A37D76
- Face: #A09C7D / Nose: #7E5E6F
- Face: #F6F4B7 / Nose: #A39866
- Face: #FFEBBB / Nose: #A37C40
- Face: #CEAF82 / Nose: #A34F3E
- Face: #A88C5B / Nose: #63504C
- Face: #947F74 / Nose: #6D5145
- Face: #967D62 / Nose: #6A372D
- Face: #7C6C54 / Nose: #3D2517
- Face: #785B4F / Nose: #453E2D
- Face: #89B7CC / Nose: #446660
- Face: #99C7C9 / Nose: #494949
- Face: #CDCF8A / Nose: #84872F
- Face: #B7D5A5 / Nose: #457A6D
- Face: #7A9E98 / Nose: #4C335B
- Face: #9DA2C9 / Nose: #52467A
- Face: #B98DBD / Nose: #694E61
- Face: #A69CB1 / Nose: #495677

961 Octocat hair choices

279 Octocat facial hair choices

Octocat mouth choices (18):
Neutral, Canines, Underbite, LipstickSmile, Goofy, Happy Closed, Happy Open, Sad, Confused, Excited Open, Excited Closed, Express, Chillin, Angry, Heart Eyes, Starstruck, Zany, Winking

Octocat prop choices (37):
Baguette, Panda, Hockey Stick, Football, Githawk, Semmle, Kimono, Aussie, Club Mate Bottle, Currywurst, Doner Kebap, Sponsor Heart, Black Hole, Lasso, Laptop, Tea Cup Mug, Smartphone, Snowboard, Newspaper, Kryptonite, Wizard Wand, Treasure, Sword, Ruby Gem, Microphone, Paintbrush, FishingRod, Game Controller, Prosthetic Hand, Books, Skateboard, Soccer Ball, Magnifying Glass, Beachball, Ship It Squirrel, Guitar, Boba Cup

Octocat accessory choices (18):
Freckles, Scarf, Cowboy Buckle Belt, Dress Buckle Belt, Bow Tie, Bolo Tie, Ties Oxford, Messenger Bag, Cape, Camera, Jewlery Gold Chain, Pearls, Watch, Bracelet, Earrings, Nose Ring, Eyebrow Ring, Tattoo

30,280,800 outfit choices (50 tops, 33 bottoms, 37 headgears, 16 eyewears, 31 footgears)
```
