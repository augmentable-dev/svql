## svql

Use SQLite to query CSV files via the [CSV virtual table](https://www.sqlite.org/csv.html).

`gcc -g -fPIC -shared csv.c -o csv.dylib`

`go build main.go`

`./main -file /path/to/csv -query "SELECT * FROM csv LIMIT 10"`