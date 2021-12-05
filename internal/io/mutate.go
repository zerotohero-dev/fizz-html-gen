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
	"fizz/internal/conf"
	"fmt"
	"github.com/gomarkdown/markdown"
	"io/ioutil"
	"regexp"
	"strings"
)

func modifyDocs(filePath, fileContent string) string {
	np := navPath(filePath)

	fmt.Println("navPath", np)

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
		fmt.Println("############ will inject document", filePath)
		firstPreIndex := strings.Index(fileContent, "<pre>")
		lastPreIndex := strings.LastIndex(fileContent, "</pre>")
		newSource = fileContent[:firstPreIndex] + `<div class="preformatted">` + "\n" + navStr + "\n<span></span>" +
			fileContent[firstPreIndex:lastPreIndex] + "</div>" + fileContent[lastPreIndex+6:]

		newSource = strings.Replace(newSource, "<document>", `<div class="nav">`, 1)
		newSource = strings.Replace(newSource, "</document>", "</div>", 1)
	} else {
		fmt.Println("############## will not inject document")
	}

	redundantHeadBlanksRegExp := regexp.MustCompile(`(?s)</style>.*</head>`)
	if redundantHeadBlanksRegExp.MatchString(newSource) {
		newSource = redundantHeadBlanksRegExp.ReplaceAllString(newSource, "</style>\n\n</head>")
	}

	return newSource
}

const scriptToInject = `<script type="text/javascript">
window.addEventListener('DOMContentLoaded', (event) => {
  me=document.querySelector("a.btn-cmt");
  if (!me) {return;}
  if (window.localStorage.getItem('comments-removed') === 'true') {
    document.getElementsByTagName("pre")[0].innerHTML = document.getElementsByTagName(
      "pre")[0].innerHTML.replaceAll(/<span class="c1">.*<\/span>/g, "")
        .replaceAll(/<span class="cm">.*<\/span>/g, "").replaceAll(/(\s*\n)/g, "\n")
	}
});

const doToggle = (evt) => {
  const ls = window.localStorage;
  const removed = () => ls.getItem("comments-removed")==="true";
  const setUnremoved = () => ls.setItem("comments-removed", "false");
  const setRemoved = () => ls.setItem("comments-removed", "true");
  const pre = () => document.getElementsByTagName("pre")[0];

  if(removed()) {
  	setUnremoved(); 
  	window.location.reload(); 
  	return false;
  }

  setRemoved();

  pre().innerHTML = pre().innerHTML
    .replaceAll(/<span class="c1">.*<\/span>/g, "")
    .replaceAll(/<span class="cm">.*<\/span>/g, "")
    .replaceAll(/(\s*\n)/g, "\n");

  return false;
}
</script>`

func commonMutations(fileContent string) string {
	htmlCommentRegExp := regexp.MustCompile(`(?s)<!--.*-->`)
	fileContent = htmlCommentRegExp.ReplaceAllString(fileContent, scriptToInject)

	inlineCssRegExp := regexp.MustCompile(`(?s)<style type="text/css">.*</style>`)
	fileContent = inlineCssRegExp.ReplaceAllString(fileContent, conf.InlineStyleToInject)

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
		fileContent = metaInjectionRegExp.ReplaceAllString(
			fileContent,
			`<meta http-equiv="content-type" content="text/html; charset=utf-8">`+
				`<meta name="viewport" content="width=device-width"><link rel="stylesheet" href="https://use.typekit.net/mqh5bnl.css">`,
		)
	}

	return fileContent
}
