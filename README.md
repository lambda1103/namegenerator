# namegenerator

A random name generator (for projects, servers, cluster nodes, etc ...)
implementation in Golang.

## Example of use

```go
package main

import (
    "github.com/lambda1103/namegenerator"
)

func main() {
    seed := time.Now().UTC().UnixNano()
    nameGenerator := namegenerator.NewNameGenerator(seed)

    name := nameGenerator.Generate(3, "-") //generate string of 2 adjectives + 1 noun, seperate words with "-"

    fmt.Println(name)
}
```

## License

Copyright (c) 2023 lambda1103
Copyright (c) 2018 Goomba project Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

[License-Image]: https://img.shields.io/badge/License-Apache-blue.svg
[License-URL]: http://opensource.org/licenses/Apache
[CircleCI-Image]: https://circleci.com/gh/goombaio/namegenerator.svg?style=svg
[CircleCI-URL]: https://circleci.com/gh/goombaio/namegenerator
[Coverage-Image]: https://codecov.io/gh/goombaio/namegenerator/branch/master/graph/badge.svg
[Coverage-URL]: https://codecov.io/gh/goombaio/namegenerator
[GoReportCard-Image]: https://goreportcard.com/badge/github.com/goombaio/namegenerator
[GoReportCard-URL]: https://goreportcard.com/report/github.com/goombaio/namegenerator
[CII-Image]: https://bestpractices.coreinfrastructure.org/projects/2237/badge
[CII-URL]: https://bestpractices.coreinfrastructure.org/projects/2237
[GoDoc-Image]: https://godoc.org/github.com/goombaio/namegenerator?status.svg
[GoDoc-URL]: http://godoc.org/github.com/goombaio/namegenerator
