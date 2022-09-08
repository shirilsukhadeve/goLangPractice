package main

import (
        "fmt"
        "math"
        "math/rand"
        "C"
        "time"
        "strconv"
    )

var baseColumns = [] string {"INT","SMALLINT","BIGINT","TINYINT", "VARCHAR(50)","DATE","DATETIME","DATETIMEX"}

var datatypes = [] string {"BASE","ARRAY","ROW"}

func createTableRecursion(createTableStatement *string, currLevel int, nestedness int, i int, parentDatatype string, nameTillNow string, columnNameVec *[]string ) {
    var choice string
    choice = datatypes[rand.Intn(len(datatypes))]
    if currLevel >= nestedness {
        choice = "BASE"
    }
    var colName string
    if parentDatatype != "ARRAY" {
        colName += "C_" + strconv.Itoa(currLevel) + "_" + strconv.Itoa(i)
        *createTableStatement += colName
        nameTillNow += "."
        nameTillNow += colName
        nameTillNow += "."
    }

    if choice == "BASE" {
        datatype := baseColumns[rand.Intn(len(baseColumns))]
        *createTableStatement += " " + datatype
        if nameTillNow[len(nameTillNow)-1] == '.' {
            nameTillNow = nameTillNow[:len(nameTillNow)-1]
        }
        *columnNameVec = append(*columnNameVec,nameTillNow)
    } else if choice == "ROW" {
        *createTableStatement += " ROW ( "
        if nameTillNow[len(nameTillNow)-1] == '.' {
            nameTillNow = nameTillNow[:len(nameTillNow)-1]
        }
        for a := 0; a < 5; a++ {
            createTableRecursion(createTableStatement,currLevel+1,nestedness,a,"ROW",nameTillNow,columnNameVec)
            if a < 4 {
                *createTableStatement += " , "
            }
        }
        *createTableStatement +=  " ) "
    } else {
        *createTableStatement += " ARRAY OF "
        a:=0
        createTableRecursion(createTableStatement,currLevel+1,nestedness,a,"ARRAY",nameTillNow,columnNameVec)
    }
}

// export getCreateTableStatement
func GetCreateTableStatement(numColumns int, nestedness int, currTime int64, createTable bool) string {
    rand.Seed(currTime)

    var createTableStatement string = "CREATE TABLE TABLE1 ( "
    var columnNameVec []string

    var base int = int (math.Round( float64((30/100) * numColumns) ))
    for i:=0; i < numColumns; i++ {
        choosen := datatypes[rand.Intn(len(datatypes))]
        var colName string
        if choosen == "BASE" {
            if base < 0 {
                i--
                continue
            }
            base--
            colName += "C_" + strconv.Itoa(0) + "_" + strconv.Itoa(i)
            datatype := baseColumns[rand.Intn(len(baseColumns))]
            createTableStatement += colName + " "+ datatype
            columnNameVec = append(columnNameVec,colName)
        } else if choosen == "ROW" {
            colName += "C_" + strconv.Itoa(0) + "_" + strconv.Itoa(i)
            createTableStatement += colName + " ROW ( "
            for a := 0; a < 5; a++ {
                createTableRecursion(&createTableStatement,1,nestedness,a,"ROW",colName,&columnNameVec)
                if a < 4 {
                    createTableStatement += " , "
                }
            }
            createTableStatement +=  ") "
        } else {
            colName += "C_" + strconv.Itoa(0) + "_" + strconv.Itoa(i)
            createTableStatement += colName + " ARRAY OF "
            a:=0
            createTableRecursion(&createTableStatement,1,nestedness,a,"ARRAY",colName,&columnNameVec)
        }
        if i < numColumns-1 {
            createTableStatement += " , "
        }
    }
    createTableStatement += " ) "

    if createTable {
        fmt.Println( createTableStatement )
        return createTableStatement
    } else {
        var alterTableStatement string = "ALTER TABLE TABLE1 ADD DATASOURCE AS TABLE1_PARQUET(\n"
        for i:=0; i < len(columnNameVec); i++ {
            alterTableStatement += columnNameVec[i] + " FROM COLUMN " + columnNameVec[i]
            if i < len(columnNameVec)-1 {
                alterTableStatement += ","
                alterTableStatement += "\n"
            }
        }
        alterTableStatement += " ) PARQUET ('no_verify=true&hdlfs:///shared/region.parquet') ENCODING 'UTF_8'";
        fmt.Println( alterTableStatement )
        return alterTableStatement
    }
}

func main() {
    var numColumns int
    var nestedness int
    fmt.Print("Number of Columns:")
    fmt.Scan(&numColumns)
    fmt.Print("Number of Nestedness:")
    fmt.Scan(&nestedness)
    currTime := time.Now().UTC().UnixNano()
    GetCreateTableStatement(numColumns,nestedness,currTime,true)
    GetCreateTableStatement(numColumns,nestedness,currTime,false)
}
