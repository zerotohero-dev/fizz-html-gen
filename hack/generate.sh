#
#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,
#

# Source code folders to generate html files off of.
FOLDERS=(
  "/about"
  "/concepts/001-doubly-linked-list"
  "/concepts/002-binary-search"
  "/concepts/003-dfs-bfs"
  "/concepts/005-preorder-traversal"
  "/concepts/006-inorder-traversal"
  "/concepts/007-inorder-successor"
  "/concepts/008-postorder-traversal"
  "/warm-up/001-two-sum"
  "/warm-up/009-palindrome-number"
  "/warm-up/013-roman-to-integer"
  "/warm-up/014-longest-common-prefix"
  "/warm-up/020-valid-parentheses"
  "/pro/208-implement-trie"
  "/concepts"
)

echo "1. Preflight (env) 🦄"

if [ -z "$FIZZ_BUZZ_DATA_ROOT" ]
then
   echo "FIZZ_BUZZ_DATA_ROOT is empty!"; exit 1;
fi

if [ -z "$FIZZ_HTML_GEN_ROOT" ]
then
   echo "FIZZ_HTML_GEN_ROOT is empty!"; exit 1;
fi

if [ -z "$FIZZ_WEB_ROOT" ]
then
   echo "FIZZ_WEB_ROOT is empty!"; exit 1;
fi

# Check data root exists.
echo "2. Preflight (dir/data) 🦄"
cd "$FIZZ_BUZZ_DATA_ROOT" || {
  echo "Failed to cd into FIZZ_BUZZ_DATA_ROOT"; exit 1;
}

# Format all oft the source code in data root first.
echo "3. Format 🦄"
for FOLDER in "${FOLDERS[@]}"
do
  cd "$FIZZ_BUZZ_DATA_ROOT$FOLDER" || {
    echo "Failed to cd into data folder '$FOLDER'"; exit 1;
  }
  go fmt ./...
done

# Remove the ./dist folder for a clean generation.
echo "4. Cleanup dist 🦄"
cd "$FIZZ_HTML_GEN_ROOT" || {
  echo "Failed to cd into html gen root"; exit 1;
}
rm -rf dist

# For each source code folder run the corresponding generator.
echo "5. Colorize 🦄"
for FOLDER in "${FOLDERS[@]}"
do
  cd "$FIZZ_HTML_GEN_ROOT/src$FOLDER" || {
    echo "Failed to cd into hack folder"; exit 1;
  }
  ./colorize.sh || {
    echo "Failed to colorize folder '$FOLDER'"; exit 1;
  }
done

# Check the generator root exists.
echo "6. Preflight (dir/generate) 🦄"
cd "$FIZZ_HTML_GEN_ROOT" || {
  echo "Failed to cd into html gen root"; exit 1;
}

# Make additional mutations on the generated files
# on top of what `pygmentize` has already done.
echo "7. Mutate 🦄"
go run cmd/main.go

# Check the dist folder has been (re)created.
echo "8. Preflight (dir/dist) 🦄"
cd "$FIZZ_HTML_GEN_ROOT/dist" || {
  echo "Failed to cd into the dist folder"; exit 1;
}

echo "9. Copy artifacts (/usr/local/share) 🦄"
rm -rf /usr/local/share/fizz/dist
cp -R "$FIZZ_HTML_GEN_ROOT/dist" "/usr/local/share/fizz" || {
  echo "Failed to copy generated files into /usr/local/share"; exit 1;
}

echo "10. Copy artifacts (FIZZ_WEB_ROOT) 🦄"
rm -rf "$FIZZ_WEB_ROOT/usr/local/share/fizz"
mkdir -p "$FIZZ_WEB_ROOT/usr/local/share/fizz"
cp -R "$FIZZ_HTML_GEN_ROOT/dist" "$FIZZ_WEB_ROOT/usr/local/share/fizz" || {
 echo "Failed to copy generated files into FIZZ_WEB_ROOT/usr/local/share"; exit 1;
}

echo "🦄 Everything Is Awesome 🦄"
