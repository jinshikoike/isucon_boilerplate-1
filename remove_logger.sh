find . -name \*.go | xargs sed -r "/(^\tdefer MeasureEnd.*$)/d"
find . -name \*.go | xargs sed -r "/(^var file =.*$)/d"
