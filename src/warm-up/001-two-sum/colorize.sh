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
F="warm-up/001-two-sum"

D="$D/$F"
mkdir -p "$D"

NAMES=("$F/doc.go" "$F/utils.go"
  "$F/impl-001.go" "$F/impl-002.go" "$F/impl-003.go"
  "$F/main.go")

"$FIZZ_HTML_GEN_ROOT/hack/colorize/pygmentize.sh" "${NAMES[@]}"
