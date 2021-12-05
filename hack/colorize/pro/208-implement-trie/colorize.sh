#!/usr/bin/env zsh

# \
# \\,
#  \\\,^,.,,.                     Zero to Hero
# ,;7~((\))`;;,,               <zerotohero.dev>
#  ,(@') ;)`))\;;',    stay up to date, be curious: learn
#   )  . ),((  ))\;,
#  /;`,,/7),)) )) )\,,
# (& )`   (,((,((;( ))\,

D="$FIZZ_HTML_GEN_ROOT/dist"
F="pro/208-implement-trie"

D="$D/$F"
mkdir -p "$D/trie"

NAMES=("$F/doc.go" "$F/trie/node.go" "$F/trie/trie.go" "$F/main.go")

"$FIZZ_HTML_GEN_ROOT/hack/colorize/pygmentize.sh" "${NAMES[@]}"
