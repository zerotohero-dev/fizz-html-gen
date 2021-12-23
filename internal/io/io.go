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

import "os"

func Walk(filePath string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !needsProcessing(info) {
		return nil
	}

	return process(filePath)
}
