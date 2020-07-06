/*
Copyright © 2020 Kunal Kushwaha

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/kunalkushwaha/kubectl-design/pkg/cli"
)

var helperText = `
kubectl-design helps to generate kubernetes resource yaml definations using
command line options and opens it in your editor to edit it and use it.

A simple tool to save your time`

func main() {
	var command []string
	var stdout, stderr bytes.Buffer
	var filename, confirm string

	if len(os.Args) < 2 {
		fmt.Println("atleast one resource type is required.")
		printHelper()
		os.Exit(1)
	}

	resourceType := os.Args[1]
	if !validateResourceType(resourceType) {
		fmt.Printf("\"%s\" is not supported resource type\n", resourceType)
		os.Exit(1)
	}

	// flags for dry-run and yaml output
	defaultargs := []string{"--dry-run=client", "-oyaml"}
	// remaining flags that need to be passed to kubectl
	args := os.Args[2:]

	runOrCreate := "create"
	if resourceType == "pod" || resourceType == "po" {
		runOrCreate = "run"
		command = append(command, runOrCreate)
	} else {
		command = append(command, runOrCreate, resourceType)
	}

	command = append(command, defaultargs...)
	command = append(command, args...)

	// Execute the kubectl command on shell and fetch the yaml output on success.
	cmd := exec.Command("kubectl", command...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("%s", string(stderr.Bytes()))
	}

	// Open the yaml output in Editor
	editorText, err := cli.OpenYAMLInEditor(string(stdout.Bytes()))
	if err != nil {
		log.Fatalf("error in opening file in editor : %s", err)
	}

	fmt.Printf("Do you want to save modified output (Y/y/N/n)? : ")
	fmt.Scanf("%s", &confirm)
	if confirm == "Y" || confirm == "y" {
		// save the output in user specified file or dump on console.
		fmt.Printf("Enter file path to save YAML : ")
		fmt.Scanf("%s", &filename)

		if len(filename) > 1 {
			err = cli.SaveToFile(filename, editorText)
			if err != nil {
				fmt.Printf("Error is saving file : %v\n", err)
				os.Exit(1)
			}
		}
	}

}

func printHelper() {
	fmt.Printf("\nkubectl design [resource-name] [options]\n\n")
	fmt.Println(helperText)
}

func validateResourceType(resource string) bool {
	rc := true

	switch resource {
	case "pod":
	case "po":
	case "deploy":
	case "deployment":
	case "service":
	case "svc":
	case "job":
	case "cronjob":
	case "cj":
	case "cm":
	case "configmap":
	case "secret":
	case "clusterrole":
	case "clusterrolebinding":
	case "namespace":
	case "ns":
	case "poddisruptionbudget":
	case "pdb":
	case "priorityclass":
	case "pc":
	case "quota":
	case "role":
	case "rolebinding":
	case "serviceaccount":
	case "sa":

	default:
		rc = false
	}
	return rc
}
