package io

import (
	"os"
	"strings"
)

const distPathPart = "/dist/"
const sourceHtmlFileSuffix = "go.html"
const entryHtmlFileSuffix = "index.html"
const sourcePathPart = "/src/" // use path.join

func needsProcessing(info os.FileInfo) bool {
	return strings.HasSuffix(info.Name(), sourceHtmlFileSuffix) ||
		strings.HasSuffix(info.Name(), entryHtmlFileSuffix)
}

func isQuestionFile(filePath string) bool {
	return strings.HasSuffix(filePath, "doc.go.html") // To constants.
}

func navPath(filePath string) string {
	return strings.Replace(
		strings.Replace(filePath, distPathPart, sourcePathPart, 1),
		"doc.go.html", "nav.html", 1,
	)
}
