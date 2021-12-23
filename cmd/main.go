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

package main

import (
	"github.com/zerotohero-dev/fizz-html-gen/internal/conf"
	"github.com/zerotohero-dev/fizz-html-gen/internal/io"
	"log"
	"path/filepath"
)

func main() {
	err := filepath.Walk(conf.ProjectRoot(), io.Walk)
	if err != nil {
		log.Println(err)
		panic("Failed to mutate some of the generated HTMLs.")
	}
}
