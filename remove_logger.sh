find . -name \*.go | xargs sed -r "/(^\tdefer.*$)/d"
