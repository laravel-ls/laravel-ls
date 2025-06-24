package runtime

import "encoding/json"

func CallScript[T any](proc Process, rootPath string, code []byte, out T) (T, error) {
	output, err := proc.Exec(rootPath, code)
	if err == nil {
		err = json.NewDecoder(output).Decode(&out)
	}
	return out, err
}
