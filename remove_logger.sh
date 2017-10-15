find . -name \*.go | xargs sed -r "/(^\tdefer.*$)/d"
find . -name \*.go | xargs sed -r "/(^var file =.*$)/d"
