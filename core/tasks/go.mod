module tasks

go 1.25.4

require db v1.0.0

require github.com/lib/pq v1.10.9 // indirect

replace db => ../db
