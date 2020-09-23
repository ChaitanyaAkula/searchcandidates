package searchcandidates

import (
	
	"log"
    "fmt"
	"github.com/ChaitanyaAkula/gittyjobsdb"
	//"github.com/ChaitanyaAkula/newsearches"
	"database/sql"
)

var Companyid string
var IDslice []string

func GetSearchCandidates(keyword string,loc string)[]string{
	
	db:=dbconnection.Connection()
	defer db.Close()
	x := keyword
	location := loc

	if x=="" && location==""{

		result1, err1 := db.Query("select gittyaccountid from people order by gittyaccountid desc")
		if err1 != nil {

			log.Fatal(err1)
		}
		IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Companyid)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Companyid)
		}


	}
	if x=="" && location!=""{
		fmt.Println("test location:",location)
		result1, err1 := db.Query("select gittyaccountid from people where match(location) against(? IN boolean mode) order by gittyaccountid desc", location)
		if err1 != nil {

			log.Fatal(err1)
		}
		IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Companyid)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Companyid)
		}
	}
	if x!=""{
	
		if location==""{
			var result1 *sql.Rows=nil
			fmt.Println("test search",x)
		result1,_= db.Query("select gittyaccountid from people where match(companyname,location,role,headline,aboutme,salary,employmenttype,totalexp,jobcategory,worktypes) against(? IN boolean mode) order by gittyaccountid desc", x)
		
		if result1==nil{
			fmt.Println("gittyestaccount")
			result1,_= db.Query("select idgittyaccount from gittyaccount where match(fullname,emailid) against(? IN natural language mode)  order by idgittyaccount desc", x)
		
		}
		IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Companyid)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Companyid)
		}
		}
		if location!=""{
			var result1 *sql.Rows=nil
			fmt.Println("test message:",x,location)
			result1,_= db.Query("select gittyaccountid from people where match(companyname,location,role,headline,aboutme,salary,employmenttype,totalexp,jobcategory,worktypes) against(? IN boolean mode) and match(location) against(? IN boolean mode) order by gittyaccountid desc", x,location)
			if result1==nil{
				fmt.Println("gittyestaccount")
				result1,_= db.Query("select idgittyaccount from gittyaccount where match(fullname,emailid) against(? IN natural language mode) and  order by idgittyaccount desc", x)
			
			}
		
		IDslice=nil
		for result1.Next() {

			err2 := result1.Scan(&Companyid)
			if err2 != nil {
				log.Fatal(err2)
			}
			IDslice=append(IDslice,Companyid)
		}
	}

	

}

	return IDslice
}