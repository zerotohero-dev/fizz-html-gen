/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package io

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"regexp"
	"strings"
)

const scriptToInject = `<script src="https://assets.fizzbuzz.pro/script/toggle.js"></script>`
const styleToInject = `<link rel="stylesheet" href="https://assets.fizzbuzz.pro/style/fizz.css">`

func commonMutations(fileContent string) string {
	htmlCommentRegExp := regexp.MustCompile(`(?s)<!--.*-->`)
	fileContent = htmlCommentRegExp.ReplaceAllString(fileContent, "")

	blankTitleRegExp := regexp.MustCompile(`<title></title>`)
	fileContent = blankTitleRegExp.ReplaceAllString(fileContent, "<title>FizzBuzz Pro</title>")

	blankH2RegExp := regexp.MustCompile(`<h2></h2>`)
	fileContent = blankH2RegExp.ReplaceAllString(fileContent, scriptToInject)

	inlineCssRegExp := regexp.MustCompile(`(?s)<style type="text/css">.*</style>`)
	fileContent = inlineCssRegExp.ReplaceAllString(fileContent, styleToInject)

	bannerLinkRegExp := regexp.MustCompile(`(?s)&lt;zerotohero.dev&gt;`)
	fileContent = bannerLinkRegExp.ReplaceAllString(
		fileContent,
		`&lt;<a href="https://zerotohero.dev">zerotohero.dev</a>&gt;`,
	)

	htmlUrlRegExp := regexp.MustCompile(`(https://.*\.html)`)
	fileContent = htmlUrlRegExp.ReplaceAllString(
		fileContent,
		`<a href="$1">$1</a>`,
	)

	if strings.Index(fileContent, "<!doctype html>") == -1 {
		legacyDoctypeRegExp := regexp.MustCompile(`<!DOCTYPE.*\n.*>`)
		fileContent = legacyDoctypeRegExp.ReplaceAllString(
			fileContent,
			"<!doctype html>",
		)
	}

	if strings.Index(fileContent, "use.typekit.net") == -1 {
		metaInjectionRegExp := regexp.MustCompile(
			`<meta http-equiv="content-type" content="text/html; charset=utf-8">`,
		)

		fileContent = metaInjectionRegExp.ReplaceAllString(fileContent,
			`<meta http-equiv="content-type" content="text/html; charset=utf-8">`+
				`<meta name="viewport" content="width=device-width">`+
				`<link rel="stylesheet" href="https://use.typekit.net/mqh5bnl.css">`,
		)
	}

	postMetaInjectionRegExp := regexp.MustCompile(
		`(?s)<link rel="stylesheet" ` +
			`href="https://use.typekit.net/mqh5bnl.css">(.*)<style`,
	)

	fileContent = postMetaInjectionRegExp.ReplaceAllString(fileContent,
		`<link rel="stylesheet" href="https://use.typekit.net/mqh5bnl.css"><style`)

	postMetaInjectionRegExp = regexp.MustCompile(`(?s)</style>(.*)</head`)
	fileContent = postMetaInjectionRegExp.ReplaceAllString(fileContent,
		`</style></head`)

	return fileContent
}

func commonPostMutations(fileContent string) string {
	if strings.Index(fileContent, "Copyright ©") != -1 {
		return fileContent
	}

	bodyEndRegExp := regexp.MustCompile(`(?s)</body>`)
	fileContent = bodyEndRegExp.ReplaceAllString(
		fileContent,
		`<div class="copyright">Copyright © <a href="https://volkan.io">Volkan Özçelik</a>.`+
			`<br><strong>FizzBuzz Pro</strong> is a `+
			`<a href="https://zerotohero.dev/">Zero to Hero</a> initiative.</div></body>`,
	)

	return fileContent
}

func mutateQuestion(filePath, fileContent string) string {
	np := navPath(filePath)

	b, err := ioutil.ReadFile(np)
	if err != nil {
		fmt.Print(err)
	}

	navStr := string(b)
	navStr = strings.Replace(navStr, "<document>", "", 1)
	navStr = strings.Replace(navStr, "</document>", "", 1)
	navStr = string(markdown.ToHTML([]byte(navStr), nil, nil))
	navStr = "<document>" + navStr + "</document>"

	noNavRegExp := regexp.MustCompile(`(?s)<pre><span></span>`)
	hasNoNavigation := noNavRegExp.MatchString(fileContent)

	newSource := ""
	if hasNoNavigation {
		firstPreIndex := strings.Index(fileContent, "<pre>")
		lastPreIndex := strings.LastIndex(fileContent, "</pre>")
		newSource = fileContent[:firstPreIndex] + `<div class="preformatted">` + "\n" + navStr + "\n<span></span>" +
			fileContent[firstPreIndex:lastPreIndex] + "</div>" + fileContent[lastPreIndex+6:]

		newSource = strings.Replace(newSource, "<document>", `<div class="nav">`, 1)
		newSource = strings.Replace(newSource, "</document>", "</div>", 1)
	}

	redundantHeadBlanksRegExp := regexp.MustCompile(`(?s)</style>.*</head>`)
	if redundantHeadBlanksRegExp.MatchString(newSource) {
		newSource = redundantHeadBlanksRegExp.ReplaceAllString(newSource, "</style>\n\n</head>")
	}

	return newSource
}

func mutateSourceCode(filePath, fileContent string) string {
	preRegExp := regexp.MustCompile(`<pre>`)

	fileContent = preRegExp.ReplaceAllString(
		fileContent,
		`<pre><a class="btn-cmt" onclick="doToggle()">[Toggle Comments]</a><br>`,
	)

	return fileContent
}
