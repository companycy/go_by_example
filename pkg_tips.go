
package db
import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)
func init {
	// initialization code here
}


package main

// In such a situation, we can use a blank identifier ( _ ) as the package alias name, so the compiler ignores the error of not using the package identifier, but will still invoke the init function.
import (
	_ "mywebapp/libs/mongodb/db"
	"fmt"
	"log"
)
func main() {
	//implementation here
}


package main
import (
	mongo "mywebapp/libs/mongodb/db"
	mysql "mywebapp/libs/mysql/db"
)
func main() {
	mongodata :=mongo.Get() //calling method of package  "mywebapp/libs/mongodb/db"
	sqldata:=mysql.Get() //calling method of package "mywebapp/libs/mysql/db"
	fmt.Println(mongodata )
	fmt.Println(sqldata )
}
