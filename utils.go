package eventstream

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type stub struct {
	ID   interface{} `json:"id"`
	Data interface{} `json:"data"`
}

func readStub(name string) ([][]byte, error) {
	msgs := [][]byte{}
	stubs := []*stub{}
	file, err := os.Open("./testdata/" + name)

	if err != nil {
		return msgs, err
	}

	defer file.Close()

	body, err := ioutil.ReadAll(file)

	if err != nil {
		return msgs, err
	}

	err = json.Unmarshal(body, &stubs)

	for _, stub := range stubs {
		msg := "event: message\n"
		id, err := json.Marshal(stub.ID)

		if err != nil {
			return msgs, err
		}

		msg += "id: " + string(id) + "\n"
		data, err := json.Marshal(stub.Data)

		if err != nil {
			return msgs, err
		}

		msg += "data: " + string(data) + "\n"
		msgs = append(msgs, []byte(msg))
	}

	return msgs, err
}
