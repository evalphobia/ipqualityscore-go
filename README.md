ipqualityscore-go
----

[![License: MIT][401]][402] [![GoDoc][101]][102] [![Release][103]][104] [![Build Status][201]][202] [![Coveralls Coverage][203]][204] [![Codecov Coverage][205]][206]
[![Go Report Card][301]][302] [![Code Climate][303]][304] [![BCH compliance][305]][306] [![CodeFactor][307]][308] [![codebeat][309]][310] [![Scrutinizer Code Quality][311]][312] [![FOSSA Status][403]][404]


<!-- Basic -->

[101]: https://godoc.org/github.com/evalphobia/ipqualityscore-go?status.svg
[102]: https://godoc.org/github.com/evalphobia/ipqualityscore-go
[103]: https://img.shields.io/github/release/evalphobia/ipqualityscore-go.svg
[104]: https://github.com/evalphobia/ipqualityscore-go/releases/latest
[105]: https://img.shields.io/github/downloads/evalphobia/ipqualityscore-go/total.svg?maxAge=1800
[106]: https://github.com/evalphobia/ipqualityscore-go/releases
[107]: https://img.shields.io/github/stars/evalphobia/ipqualityscore-go.svg
[108]: https://github.com/evalphobia/ipqualityscore-go/stargazers


<!-- Testing -->

[201]: https://github.com/evalphobia/ipqualityscore-go/workflows/test/badge.svg
[202]: https://github.com/evalphobia/ipqualityscore-go/actions?query=workflow%3Atest
[203]: https://coveralls.io/repos/evalphobia/ipqualityscore-go/badge.svg?branch=master&service=github
[204]: https://coveralls.io/github/evalphobia/ipqualityscore-go?branch=master
[205]: https://codecov.io/gh/evalphobia/ipqualityscore-go/branch/master/graph/badge.svg
[206]: https://codecov.io/gh/evalphobia/ipqualityscore-go


<!-- Code Quality -->

[301]: https://goreportcard.com/badge/github.com/evalphobia/ipqualityscore-go
[302]: https://goreportcard.com/report/github.com/evalphobia/ipqualityscore-go
[303]: https://codeclimate.com/github/evalphobia/ipqualityscore-go/badges/gpa.svg
[304]: https://codeclimate.com/github/evalphobia/ipqualityscore-go
[305]: https://bettercodehub.com/edge/badge/evalphobia/ipqualityscore-go?branch=master
[306]: https://bettercodehub.com/
[307]: https://www.codefactor.io/repository/github/evalphobia/ipqualityscore-go/badge
[308]: https://www.codefactor.io/repository/github/evalphobia/ipqualityscore-go
[309]: https://codebeat.co/badges/142f5ca7-da37-474f-9264-f708ade08b5c
[310]: https://codebeat.co/projects/github-com-evalphobia-ipqualityscore-go-master
[311]: https://scrutinizer-ci.com/g/evalphobia/ipqualityscore-go/badges/quality-score.png?b=master
[312]: https://scrutinizer-ci.com/g/evalphobia/ipqualityscore-go/?branch=master

<!-- License -->
[401]: https://img.shields.io/badge/License-MIT-blue.svg
[402]: LICENSE.md
[403]: https://app.fossa.com/api/projects/git%2Bgithub.com%2Fevalphobia%2Fipqualityscore-go.svg?type=shield
[404]: https://app.fossa.com/projects/git%2Bgithub.com%2Fevalphobia%2Fipqualityscore-go?ref=badge_shield


Unofficial golang library for [IPQualityScore](https://www.ipqualityscore.com/).


# Quick Usage for API

```go
package main

import (
	"fmt"

	"github.com/evalphobia/ipqualityscore-go/config"
	"github.com/evalphobia/ipqualityscore-go/ipqs"
)

func main() {
	conf := config.Config{
		// you can set auth values to config directly, otherwise used from environment variables.
		APIKey:   "<your API Key>",
		Debug:    false,
	}

	svc, err := ipqs.New(conf)
	if err != nil {
		panic(err)
	}

	// execute API
	resp, err := svc.IPReputation("8.8.8.8")
	if err != nil {
		panic(err)
	}
	if resp.HasError() {
		panic(fmt.Errorf("code=[%d] message=[%s]", resp.ErrData.StatusCode, resp.ErrData.Message))
	}

	// just print response in json format
	b, _ := json.Marshal(resp)
	fmt.Printf("%s", string(b))
}
```

see example dir for more examples, and see [official API document](https://www.ipqualityscore.com/documentation/proxy-detection/overview) for more details (especially request/response).


# Environment variables

| Name | Description |
|:--|:--|
| `IPQS_APIKEY` | [IPQualityScore API Key](https://www.ipqualityscore.com/documentation/overview). |
