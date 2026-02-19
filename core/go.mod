module togo

require tasks v1.0.0

require db v1.0.0

require github.com/lib/pq v1.10.9 // indirect

replace tasks => ./tasks

replace db => ./db

go 1.25.4
