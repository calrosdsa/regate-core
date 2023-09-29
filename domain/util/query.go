package util

import (
	// "fmt"
	"fmt"
	"strconv"
	"strings"
)

func AppendQueries(condition,q string,columns []string) (string, error){
	qParts := make([]string, 0, 2)
	// args := make([]interface{}, 0, 2)
	for i := 0; i < len(columns	);i++{
			qParts = append(qParts,columns[i] + ` = $` + strconv.Itoa(i+1))
	}
	q += strings.Join(qParts, ",") + ` WHERE ` + condition +`= $` + strconv.Itoa(len(columns) +1) + `;`
     return q,nil
}

func ArrayToString(a []int, delim string) string {
    return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
    //return strings.Trim(strings.Join(strings.Split(fmt.Sprint(a), " "), delim), "[]")
    //return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(a)), delim), "[]")
}

// package util

// import (
// 	// "fmt"
// 	"strings"
// 	"strconv"
// )

// func AppendQueries(condition,q string,columns []string) (string, error){
// 	qParts := make([]string, len(columns))
// 	// args := make([]interface{}, 0, 2)
// 	for i := 0; i < len(columns	);i++{
// 		    qParts[i] = columns[i] + ` = $` + strconv.Itoa(i+1)
// 			// qParts = append(qParts,columns[i] + ` = $` + strconv.Itoa(i+1))
// 	}
// 	q += strings.Join(qParts, ",") + ` WHERE ` + condition +`= $` + strconv.Itoa(len(columns) +1 )
//      return q,nil
// }

