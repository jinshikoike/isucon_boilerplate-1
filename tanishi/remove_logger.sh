ls *.go | grep -E -v "logger.*" | xargs sed -r "/(^\tdefer MeasureEnd.*$)/d"
ls *.go | grep -E -v "logger.*" | xargs sed -r "/(^var file =.*$)/d"
