package io

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

func process(filePath string) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	fileContent := commonMutations(string(b))

	if isQuestionFile(filePath) {
		fmt.Println("Found question file", filePath)
		fileContent = modifyDocs(filePath, fileContent)
	} else {
		preRegExp := regexp.MustCompile(`<pre>`)
		if strings.Index(filePath, "index.html") == -1 {
			fileContent = preRegExp.ReplaceAllString(
				fileContent,
				`<pre><a class="btn-cmt" onclick="doToggle()">[Toggle Comments]</a><br>`,
			)
		}
	}

	// Common post-mutations
	if strings.Index(fileContent, "Copyright ©") == -1 {
		bodyEndRegExp := regexp.MustCompile(`(?s)</body>`)
		fileContent = bodyEndRegExp.ReplaceAllString(
			fileContent,
			`<div class="copyright">Copyright © <a href="https://volkan.io">Volkan Özçelik</a>.`+
				`<br><strong>FizzBuzz Pro</strong> is a `+
				`<a href="https://zerotohero.dev/">Zero to Hero</a> initiative.</div></body>`,
		)
	}

	// fmt.Println("will write file » ", filePath)
	err = ioutil.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
