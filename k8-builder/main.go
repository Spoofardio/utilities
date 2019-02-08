package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Entry point for the program
func main() {
	templateFile := os.Args[1]
	configFile := os.Args[2]

	template := getFileContents(templateFile)
	config := getFileContents(configFile)
	buildOutputFolder("target/")

	fmt.Println("Template File: " + templateFile)
	fmt.Println("Configuration File: " + configFile)

	output := buildTemplateWithConfig(config, template)
	outputFile := createOutputFile(templateFile, output)

	fmt.Println("Output File: " + outputFile)

}

// func main() {
// 	release := os.Args[1]
// 	configFile := os.Args[2] + ".conf"
// 	config := getFileContents(configFile)
// 	buildOutputFolder("target/")
// 	files, err := ioutil.ReadDir("./" + release + "/")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for _, f := range files {
// 		fmt.Println(f.Name())
// 	}
// }

// Returns a string of the filled out template
func buildTemplateWithConfig(config string, template string) string {
	output := template
	configScanner := bufio.NewScanner(strings.NewReader(config))
	for configScanner.Scan() {
		line := configScanner.Text()
		if line[0] != '#' {
			i := strings.Index(line, "=")
			key := fmt.Sprintf("{{%s}}", line[:i])
			value := line[i+1:]
			output = strings.Replace(output, key, value, -1)
		}
	}
	return output
}

// Creates the yaml file and saves it in the output folder.
func createOutputFile(file string, contents string) string {
	fileName := file[0 : len(file)-5]
	if strings.Contains(fileName, "/") {
		fileName = file[strings.LastIndex(fileName, "/")+1 : len(file)-5]
	}
	fileType := file[len(file)-4:]
	filePath := fmt.Sprintf("target/%s.%s", fileName, fileType)
	os.Remove(filePath)
	err := ioutil.WriteFile(filePath, []byte(contents), 0644)
	if err != nil {
		fmt.Println("Failed to save generated file: " + filePath)
		os.Exit(0)
	}
	return filePath
}

// Creates the output folder where all finished yaml files will be put.
func buildOutputFolder(outputFolder string) {
	_, err := os.Stat(outputFolder)
	if os.IsNotExist(err) {
		errMkdir := os.MkdirAll(outputFolder, 0755)
		if errMkdir != nil {
			panic("Failed to create the buildfolder. Check directory permissions.")
		}
	}
}

// Returns a string that contains everything in the specified file.
func getFileContents(file string) string {
	templateBytes, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("File not found: " + file)
		os.Exit(0)
	}
	return string(templateBytes)
}
