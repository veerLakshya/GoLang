package advanced

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func serialization_deserialization() {
	user := User{
		Name:  "Lakshya",
		Email: "lak@ex.com",
	}

	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user, string(jsonData))

	var user1 User
	err = json.Unmarshal(jsonData, &user1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(user1)

	jsonData1 := `{"name": "John", "email": "john@ex.com"}`

	reader := strings.NewReader(jsonData1)
	decoder := json.NewDecoder(reader)

	var user2 User
	err = decoder.Decode(&user2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("from decoder: ", user2)

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)

	err = encoder.Encode(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Encoded string: ", buf.String())

}

/*

marshal 			- in memory
decoder/encoder 	- streams

json.marshal 			-> in-memory json
						-> one-shot operation
						-> entire json must fit in memory

json.NewDecoder().Decode()	-> reads json from an io.Reader
							-> HTTP request body
							-> decodes as data arrives
							-> does not need full json in memory


json.NewEncoder(w).Encode(res)	-> writes json to an io.Writer
								-> HTTP Response body
								-> appends a newline \n
								-> streams output
*/
