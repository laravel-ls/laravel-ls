package runtime

import (
	"encoding/json"
	"fmt"
)

func CallScript[T any](proc Process, rootPath string, code []byte, out T) (T, error) {
	output, err := proc.Exec(rootPath, code)
	if err == nil {
		if err = json.NewDecoder(output).Decode(&out); err != nil {
			err = fmt.Errorf("json: %w", err)
		}
	}
	return out, err
}
