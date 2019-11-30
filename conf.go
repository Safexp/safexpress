package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	Config struct
	sql include configures about SQL server infomation.

	sql.host SQL server localtion
	eg: 127.0.0.1:3306

	sql.user SQL server user name
	sql.user SQL server user password

	port listen port

	name server name (don't make sence)
*/
type config_t = struct {
	sql struct {
		host     string
		user     string
		password string
	}
	port int
	name string
}

/*
	Convert JSON obj path to config body.

	path: the way to json
	eg sql.user

	pi ps: the value will be change
*/
type opt = struct {
	path []string
	pi   *int
	ps   *string
}

// Global config
var config = config_t{}

// Example config json to config struct
var opts = []opt{
	{[]string{"sql", "host"}, nil, &config.sql.host},
	{[]string{"sql", "user"}, nil, &config.sql.user},
	{[]string{"sql", "password"}, nil, &config.sql.password},
	{[]string{"port"}, &config.port, nil},
	{[]string{"name"}, nil, &config.name},
}

// find path from obj and set value to pi or ps depend on type
func jsonTreeModify(obj map[string]interface{}, names []string, pi *int, ps *string) (err error) {

	switch o := obj[names[0]].(type) {
	// set value if it is an int or string
	// TODO: Need to handle errors
	// May have null point or wrong type
	case string:
		*ps = o
	case float64:
		*pi = int(o)
	case map[string]interface{}:
		// go to next layer until it is an int or string
		return jsonTreeModify(o, names[1:], pi, ps)
	default:
		// TODO: Should do more types
		return err
	}
	return nil
}

// cast JSON object (interface{}) to go type
// may not use anymore
func jsonTreeExec(obj map[string]interface{}, names []string, cmd func(interface{})) (err error) {

	switch o := obj[names[0]].(type) {
	case string:
		cmd(o)
	case float64:
		cmd(o)
	case map[string]interface{}:
		return jsonTreeExec(o, names[1:], cmd)
	default:
		return err
	}
	return nil
}

/*
	load config from file by given file name.
	TODO: Need to handle errors
*/
func LoadConfig(filename string) (c config_t) {

	// Open config file first
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Can't open config file reason : %s\n", err)
	}
	defer file.Close()

	// Read json file in to memo
	jsonCont, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Can't load config file reason : %s\n", err)
	}

	// Unmarshal json to go obj
	con := make(map[string]interface{})
	err = json.Unmarshal(jsonCont, &con)
	if err != nil {
		fmt.Print(err)
	}

	// analy go obj
	for _, v := range opts {
		jsonTreeModify(con, v.path, v.pi, v.ps)
	}

	// for dbg
	fmt.Print(config)

	return
}
