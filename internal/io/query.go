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
	"os"
	"strings"
)

const distPathPart = "/dist/"
const sourceHtmlFileSuffix = "go.html"
const questionFileSuffix = "doc.go.html"
const entryHtmlFileSuffix = "index.html"
const sourcePathPart = "/src/" // TODO: this is fragile!

func needsProcessing(info os.FileInfo) bool {
	return strings.HasSuffix(info.Name(), sourceHtmlFileSuffix) ||
		strings.HasSuffix(info.Name(), entryHtmlFileSuffix)
}

func isQuestionFile(filePath string) bool {
	return strings.HasSuffix(filePath, questionFileSuffix)
}

func navPath(filePath string) string {
	return strings.Replace(
		strings.Replace(filePath, distPathPart, sourcePathPart, 1),
		"doc.go.html", "nav.html", 1,
	)
}
