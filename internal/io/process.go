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

import "io/ioutil"

func process(filePath string) error {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	fileContent := commonMutations(string(b))

	if isQuestionFile(filePath) {
		fileContent = mutateQuestion(filePath, fileContent)
	} else {
		fileContent = mutateSourceCode(filePath, fileContent)
	}

	fileContent = commonPostMutations(fileContent)

	err = ioutil.WriteFile(filePath, []byte(fileContent), 0644)
	if err != nil {
		return err
	}

	return nil
}
