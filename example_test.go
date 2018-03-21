package batchputs_test

import (
	"fmt"
	"os"

	"time"

	"github.com/zanjs/batchputs"
)

/*
With this example, We created 30k records with 3 columns each, and inserts it into database in batch. and then we updates 20k records.

semaphoreci.com runs this example for inserting 30k records and updates 20k records totally less than 2 seconds.

```
=== RUN   ExamplePut_perf
--- PASS: ExamplePut_perf (1.73s)
```

[![Build Status](https://semaphoreci.com/api/v1/theplant/batchputs/branches/master/badge.svg)](https://semaphoreci.com/theplant/batchputs)

*/
func ExamplePut_perf() {
	db := openAndMigrate()
	// with table
	// CREATE TABLE countries
	// (
	//     code VARCHAR(50) PRIMARY KEY NOT NULL,
	//     short_name TEXT,
	//     special_notes TEXT,
	//     region TEXT,
	//     income_group TEXT,
	//     count INTEGER,
	//     avg_age NUMERIC
	// );

	rows := [][]interface{}{}
	for i := 0; i < 30000; i++ {
		rows = append(rows, []interface{}{
			fmt.Sprintf("CODE_%d", i),
			fmt.Sprintf("short name %d", i),
			i,
		})
	}
	columns := []string{"code", "short_name", "count"}
	dialect := os.Getenv("DB_DIALECT")
	if len(dialect) == 0 {
		dialect = "postgres"
	}

	start := time.Now()
	err := batchputs.Put(db.DB(), dialect, "countries", "code", columns, rows)
	if err != nil {
		panic(err)
	}
	duration := time.Since(start)
	fmt.Println("Inserts 30000 records using less than 3 seconds:", duration.Seconds() < 3)

	rows = [][]interface{}{}
	for i := 0; i < 20000; i++ {
		rows = append(rows, []interface{}{
			fmt.Sprintf("CODE_%d", i),
			fmt.Sprintf("short name %d", i),
			i + 1,
		})
	}
	start = time.Now()
	err = batchputs.Put(db.DB(), dialect, "countries", "code", columns, rows)
	if err != nil {
		panic(err)
	}
	duration = time.Since(start)
	fmt.Println("Updates 20000 records using less than 3 seconds:", duration.Seconds() < 3)

	//Output:
	// Inserts 30000 records using less than 3 seconds: true
	// Updates 20000 records using less than 3 seconds: true

}
