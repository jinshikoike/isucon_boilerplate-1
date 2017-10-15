find . -name \*.go | xargs sed -r "s/(^func.*?$)/\1\n\tdefer MeasureEnd(MeasureStart(), file)/" -i
