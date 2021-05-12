package examples

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func Json() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	floatB, _ := json.Marshal(1.23)
	fmt.Println(string(floatB))

	strB, _ := json.Marshal("string")
	fmt.Println(string(strB))

	sliceB, _ := json.Marshal([]string{"a", "b", "c"})
	fmt.Println(string(sliceB))

	mapB, _ := json.Marshal(map[string]int{"a": 1, "b": 2})
	fmt.Println(string(mapB))

	res1B, _ := json.Marshal(&response1{
		Page:   1,
		Fruits: []string{"apple", "orange"},
	})
	fmt.Println(string(res1B))

	res2B, _ := json.Marshal(&response2{
		Page:   1,
		Fruits: []string{"apple", "orange"},
	})
	fmt.Println(string(res2B))

	byt := []byte(`{"page":1,"fruits":["apple","orange"]}`)

	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err != nil {
		panic(err)
	}

	fmt.Println(data)
	fmt.Println(data["fruits"].([]interface{})[0])

	str := `{"page":1,"fruits":["apple","orange"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"a": 1, "b": 2}
	enc.Encode(d)
}
