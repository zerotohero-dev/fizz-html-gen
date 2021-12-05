# \\,
#  \\\,^,.,,.                     Zero to Hero
# ,;7~((\))`;;,,               <zerotohero.dev>
#  ,(@') ;)`))\;;',    stay up to date, be curious: learn
#   )  . ),((  ))\;,
#  /;`,,/7),)) )) )\,,
# (& )`   (,((,((;( ))\,

D="$FIZZ_HTML_GEN_ROOT/dist"
F="about"

D="$D/$F"
mkdir -p "$D"

NAMES=("$F/doc.go")

"$FIZZ_HTML_GEN_ROOT/hack/colorize/pygmentize.sh" "${NAMES[@]}"
