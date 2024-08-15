# dotenv-Rowan
#### A library in Golang that can load environment variables from .env files.

## What are .env files?
The .env file is used in projects to store configuration settings, environment variables, and sensitive information securely. 
It provides a convenient way to manage and organize various parameters that your project needs without hard-coding them directly into your code.


## Example of .env file
```
USERNAME=rowancanchill
PASSWORD=rowancantchill
```
## How to Use

### 1. Import library
```
import (
    dotenv "github.com/codescalersinternships/dotenv-Rowan"
)
```

### 2. Use Load Function
### a. With no parameters
By default, this reads from `.env` file present on your current directory path
```
err := dotenv.Load()
```
### b. With any number of file paths
By default, this reads from `.env` file present on your current directory path
```
envFiles := []string {"db.env", "client.env", "server.env"}
err := dotenv.Load(envFiles...)
```