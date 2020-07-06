package cli

import (
	"io/ioutil"
	"os"
	"os/exec"
)

const DefaultEditor = "vim"

// OpenFileInEditor opens filename in a text editor.
func OpenFileInEditor(filename string) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = DefaultEditor
	}

	// Get the full executable path for the editor.
	executable, err := exec.LookPath(editor)
	if err != nil {
		return err
	}

	cmd := exec.Command(executable, filename)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// OpenYAMLInEditor opens a temporary file with bytes in a text editor and
// returns the written bytes on success or an error on failure.
// It handles deletion of the temporary file behind the scenes.
func OpenYAMLInEditor(yamlInput string) ([]byte, error) {
	file, err := ioutil.TempFile(os.TempDir(), "*.yaml")
	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

	// Defer removal of the temporary file in case any of the next steps fail.
	defer os.Remove(filename)

	_, err = file.WriteString(yamlInput)
	if err != nil {
		return nil, err
	}
	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	if err = OpenFileInEditor(filename); err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}

// SaveToFile saves bytes to filename on disk
func SaveToFile(filename string, yaml []byte) error {
	err := ioutil.WriteFile(filename, yaml, 0644)
	return err
}
