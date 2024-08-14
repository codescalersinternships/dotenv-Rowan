package dotenv

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


// Load files and handles parsing and env setting
func Load(files ...string) error {
	var envVars map[string]string
	var err error
	if len(files) == 0 {
		files = append(files, ".env")
	}
	for i := 0; i < len(files); i++ {
		envVars, err = parse(files[i])
		if err != nil {
			return err	
		}
	}
	return osEnvSetter(envVars)
}

// parse unexposed externally, parses the file of key=value pairs into a map
func parse(filePath string) (map[string]string, error) {
	envVars := make(map[string]string)
	data, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// comment neglecting
		if line[0] == '#' {
			continue
		}
		keyValuePair := strings.Split(line, "=")
		if len(keyValuePair) != 2 {
			return nil, fmt.Errorf(".env file must have key=value only, problem at %v", line)
		}
		key := strings.TrimSpace(keyValuePair[0])
		value := strings.TrimSpace(keyValuePair[1])
		if(len(key) == 0 || len(value) == 0){
			return nil, fmt.Errorf("key or value not present, problem at %v", line)
		}
		envVars[key] = value
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return envVars, nil
}

// osEnvSetter sets os environment variables
func osEnvSetter(envVars map[string]string) error {
	for key, value := range envVars {
		fmt.Printf("key is %s and value is %s \n", key, value)
		err := os.Setenv(key, value)
		if err!= nil {
			return err
		}
	}
	return nil
}
