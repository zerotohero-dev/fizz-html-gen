# \\,
#  \\\,^,.,,.                     Zero to Hero
# ,;7~((\))`;;,,               <zerotohero.dev>
#  ,(@') ;)`))\;;',    stay up to date, be curious: learn
#   )  . ),((  ))\;,
#  /;`,,/7),)) )) )\,,
# (& )`   (,((,((;( ))\,

DIST_PATH="$FIZZ_HTML_GEN_ROOT"/dist
mkdir -p "$DIST_PATH/$FOLDER"

for FILENAME in "$@"
do
  # echo "»» colorize »» $DIST_PATH/$FILENAME.html"
  pygmentize -f html -l go -O full,style=borland \
    -o "$DIST_PATH/$FILENAME.html" \
    "$FIZZ_BUZZ_DATA_ROOT/$FILENAME" || {
      echo "Failed to pygmentize $FILENAME"; exit 1;
    }
done