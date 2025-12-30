# Electrostatic — ssr and ssg

Simple app to generate and serve static content.More info in
[this article (ru)](https://kucheriavyi.ru/go/kak-ya-napisal-blog-na-go/).

## Installation

0. install the package `go install github.com/laranatech/electrostatic@latest`
0. initialize the project `electrostatic -m init -r /path/to/source`
0. edit `template.html` and `meta.json` as you wish
0. You are ready to go!

## Usage

This app supports 2 modes:

### SSMG (SSR) mode (recomended for development)

Just run it with `electrostatic -m serve -r /path/to/source`

### SSG mode

Build it with `electrostatic -m export -r /paht/to/source`.
Your static site will be written in `./dist` directory.
Copy it to your static server and that's all.

## Credits

[Evgenii Kucheriavyi](https://t.me/frontend_director)

Support on [boosty](https://boosty.to/kucheriavyi) or
[patreon](https://patreon.com/kucheriavyi).

