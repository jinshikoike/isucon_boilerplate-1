ls *.go | grep -E -v "logger.*" | xargs sed -r 's/(^func main.*?$)/var file = Initialize()\n\1/' -i
ls *.go | grep -E -v "logger.*" | xargs sed -r 's/(^func.*?$)/\1\n\tdefer MeasureEnd(MeasureStart(""))/' -i

