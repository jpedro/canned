package canned

import (
	"encoding/json"
	"os"

	"gopkg.in/yaml.v2"
)

func dump(can *Can) error {
	envDump := env("CANNED_DUMP", "")
	if envDump != "yes-pretty-please" {
		return nil
	}

	redacted := can
	for name := range can.Items {
		redacted.SetItem(name, "[redacted]")
	}

	err := dumpJson(redacted, can.file + ".json")
	if err != nil {
		return err
	}

	err = dumpYaml(redacted, can.file + ".yaml")
	if err != nil {
		return err
	}

	return nil
}

func dumpJson(data any, file string) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, bytes, 0644)
	if err != nil {
		return err
	}

	err = verifyDumpJson(file)
	if err != nil {
		return err
	}

	return nil
}

func dumpYaml(data any, file string) error {
	bytes, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, bytes, 0644)
	if err != nil {
		return err
	}

	err = verifyDumpYaml(file)
	if err != nil {
		return err
	}

	return nil
}

func verifyDumpJson(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var can *Can
	err = json.Unmarshal(data, &can)
	if err != nil {
		panic(err)
	}

	return err
}

func verifyDumpYaml(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var can *Can
	err = yaml.Unmarshal(data, &can)
	if err != nil {
		panic(err)
	}

	return err
}
