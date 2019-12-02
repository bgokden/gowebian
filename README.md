# GoWebian

GoWebian is a go library to generate html files and WebAssembly bindings.
*GoWebian is work in progress*

It can only be used for server side html generation and building websites with Golang.
WebAssembly bindings are easy to do and doesn't require writing Javascript code.

Main target of GoWebian is splitting the visual logic and application logic.
Visual logic will be done with HTML/CSS/Javascript and application logic is done with WebAssembly.

GoWebian support page page definition and custom component definition.
Javascript event binding, dom manipulation and messaging between components.

### Running Examples

Git Clone this repo and go into the directory

```shell
$ git clone https://github.com/bgokden/gowebian.git
$ cd gowebian
```
Basic example:

```shell
$ ./buildpage.sh examples/basicpage/
$ go run serve/serve.go ./examples/basicpage/public/
````

Go to your browser http://localhost:8080
Click on the button, it will add random component to the list.
Enter text to the text field and press enter, text will be copied to the text field

MDB basic page

```shell
$ ./buildpage.sh examples/mdbpage1/
$ go run serve/serve.go ./examples/mdbpage1/public/
````

MDB grid page

```shell
$ ./buildpage.sh examples/mdbpage2/
$ go run serve/serve.go ./examples/mdbpage2/public/
````

### Current problems in WebAssembly and solutions in GoWebian:

#### Problem:
Currently go WebAssembly WASM file sizes are big. Example page sizes are reaching to 5Mb. With Gzip it can be compressed to less than 1.5Mb. Main reason is Go runtime is bundled in this executable.
There are promising projects like tinygo which has a smaller runtime, but GoWebian doesn't work with it due to depending go packages ( template/text, template/html packages and go routines ). `I don't fully use these packages so I may prefer to rewrite to fully support tinygo builds.`
Also there are proposals to embed common libraries of common languages into the browser runtimes which is a good solution.
Probably we will see both solutions in a year.

#### Current solution:

GoWebian generates a working HTML representation per page. This can be integrated into web server, user will see this page on load. While user see a functioning page, the WASM code loads. WASM has support for streaming instantiation, which means compile as the payload downloading. Depending on the network, WASM download and compile time can be faster than complex Javascript libraries.

#### Notes:

* Material Design Bootstrap page and components are added for examples, they will be probably move to another repository.
* I will a client that do tasks of `buildpage.sh` and `serve/serve.go`
