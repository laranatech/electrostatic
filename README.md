# Electrostatic — ssr and ssg

Simple app to generate and serve static content.More info in
[this article (ru)](https://kucheriavyi.ru/go/kak-ya-napisal-blog-na-go/).

## Installation

0. install the package `go install larana.tech/go/electrostatic@latest`
0. initialize the project `electrostatic -m init -r /path/to/source`
0. edit `template.html` and `config.toml` as you wish
0. You are ready to go!

## Usage

This app supports 2 modes:

### SSMG (SSR) mode (recomended for development)

Just run it with `electrostatic -m serve -r /path/to/source`

### SSG mode

Build it with `electrostatic -m export -r /paht/to/source`.
Your static site will be written in `./dist` directory.
Copy it to your static server and that's all.

#### Note about routing

Links will be created in `/posts/2026-02-27-post-title` format, not
`/posts/2026-02-27-post-title.html` (the `.html` will be missing).

If you want them to work, you'll need to config your server config.

> Check out this issue: https://github.com/laranatech/electrostatic/issues/18

## Credits

[Evgenii Kucheriavyi](https://t.me/ekucheriavyi)

Support on [boosty](https://boosty.to/kucheriavyi) or
[patreon](https://patreon.com/kucheriavyi).

