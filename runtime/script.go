package runtime

import "encoding/json"

func CallScript[T any](call *PHPProcess, rootPath string, code []byte, out T) (T, error) {
	output, err := call.Exec(rootPath, code)
	if err == nil {
		err = json.NewDecoder(output).Decode(&out)
	}
	return out, err
}
