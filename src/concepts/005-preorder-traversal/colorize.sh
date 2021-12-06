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
F="concepts/005-preorder-traversal"

D="$D/$F"
mkdir -p "$D"

NAMES=("$F/doc.go" "$F/helpers.go" "$F/impl.go" "$F/node.go" "$F/main.go")

"$FIZZ_HTML_GEN_ROOT/hack/colorize/pygmentize.sh" "${NAMES[@]}"
