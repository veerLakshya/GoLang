package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("start")
	//defer fmt.Println("middle")
	//defer fmt.Println("end")
	// defer things are executed in lifo order

	//res, err := http.Get("https://www.google.com/robots.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	////defer res.Body.Close()
	//robots, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%s", robots)

	// Panics
	//fmt.Println("start")
	//panic("Panic happened, so lower line wont be executed")
	//fmt.Println("end")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello go"))
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}
	/*
		*DEFER
			-used to delay execution of a statement until function exites
			-useful to group open and close function together
				-be careful in loops
			-Runs in LIFO order if there are multiple in defers
			-Arguments evaluated at time defer is executed, not at time of function execution

		*PANIC
			-when there is err, we can either return err or panic
			-panic occurs when program cannot continue at all
				-Don't use when file cant be opened, unless it is critical
				-Use for unrecoverable events - cannot obtain TCP port for web server
			-Function will stop executing
				-Deferred functions will still execute
			-if nothing handles panic, program will exit

		*RECOVER
			-used to recover from panics
			-only useful in deferred functions
			-current function will not attempt to continue, but higher functions in call stack will
	*/
}
