find . -name \*.go | xargs sed -r "s/(^func main.*?$)/var file = Initialize()\n\1/" -i
find . -name \*.go | xargs sed -r "s/(^func.*?$)/\1\n\tdefer MeasureEnd(MeasureStart(""))/" -i
