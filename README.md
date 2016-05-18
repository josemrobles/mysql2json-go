# mysql2json-go
Inspired by [Arturom's](https://github.com/arturom) awesome [mysql2json](https://github.com/arturom/mysql2json) python script. Connects to a MySQL database and prints JSON formatted records to STDOUT...but using Golang!


### Use this:
```bash
./app -host=server -port=3306 -user=hello -password=hola -database=bonjour -query="SELECT * FROM products LIMIT 1"
```

### Turn this
```sql
SELECT * FROM products;
```

```
+----+----------------+-------------------+
| id | name           | quantity          |
+----+----------------+-------------------+
|  1 | TLR 22 3.0     |                 1 |
+----+----------------+-------------------+
```

### Into this
```json
[{"id":"1","name":"TLR 22 3.0","quantity":"1"}]
```
