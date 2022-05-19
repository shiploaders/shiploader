# shiploader
## _Not your average ship loader_

[![.github/workflows/ci.yml](https://github.com/Marcus1911/shiploader/actions/workflows/ci.yml/badge.svg)](https://github.com/Marcus1911/shiploader/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/Marcus1911/shiploader/branch/main/graph/badge.svg?token=DDYNT8YXJM)](https://codecov.io/gh/Marcus1911/shiploader)
![GitHub branch checks state](https://img.shields.io/github/checks-status/Marcus1911/shiploader/main)


Discribe components.

- Type some Markdown on the left
- See HTML in the right
- ✨Magic ✨

## Releases

- release-name2 or v 2.xx
- release-name1 or v 1.xx


## Dev workflow

### Build
- `go build -o shiploader main.go`
- `chmod +x shiploader`

### Run
- Create your desired apps state in a config file (`config.yaml`)
- Run `./shiploader generate --config config.yaml`

### Tests
- `go test ./... -coverprofile=coverage.out`
- To view coverage in a html report: `go tool cover -html=coverage.out` 

## FAQ

Include FAQs as :

- faq 1
- faq 2
- faq 3
- [Breakdance](https://breakdance.github.io/breakdance/) - HTML

## Issues


| Desc. | Issue link |
| ------ | ------ |
| bug-fix-0998 | [path/path/CONTRIBUTING.md][PlDb] |
| bug-fix-8897 | [path/path/CONTRIBUTING.md][PlGh] |
| refac-22 | [path/path/CONTRIBUTING.md][PlGd] |
| issue-101 | [path/path/CONTRIBUTING.md][PlOd] |



## License (if necessary)

something

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>
