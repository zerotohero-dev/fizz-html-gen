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
F="warm-up/014-longest-common-prefix"

D="$D/$F"
mkdir -p "$D/trie"

NAMES=("$F/doc.go"
  "$F/impl-001.go" "$F/impl-002.go"
  "$F/impl-003.go" "$F/impl-004.go" "$F/impl-005.go"
  "$F/utils-math.go" "$F/utils-const.go"
  "$F/trie/node.go" "$F/trie/trie.go"
  "$F/main.go")

"$FIZZ_HTML_GEN_ROOT/hack/colorize/pygmentize.sh" "${NAMES[@]}"
