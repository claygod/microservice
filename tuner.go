package main

// Microservice
// Tuner
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"errors"
	//"flag"
	//"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	// DELIMITER_COMMAND - is used to bind the section name and key name (command line)
	DELIMITER_COMMAND string = "/"
	// DELIMITER_PARAM - is used to separate key and value (command line)
	DELIMITER_PARAM string = "="
	// ARGUMENT_NAME_CONF_FILE - an argument path name of the configuration file on the command line
	ARGUMENT_NAME_CONF_FILE = "confile"
)

// NewTuner - create a new Tuner-struct
func NewTuner(path string) (*Tuner, error) {
	t := &Tuner{}
	argCommLine, err := t.parseCommandLine()
	if err != nil {
		return nil, err
	}
	if s, ok := argCommLine[ARGUMENT_NAME_CONF_FILE]; ok {
		path = s
	}
	if _, err := toml.DecodeFile(path, t); err != nil {
		return nil, err
	}
	if err := t.env(argCommLine); err != nil {
		return nil, err
	}
	return t, nil
}

// Tuner structure
//Configuring using the configuration file, environment variables, and command-line variables.
//The default configuration is loaded from the specified file ( `config.toml`).
//The configuration file can be changed from the command line like this:
//
//	`yourservice -confile config.toml`
//
//If the operating system environment variables are set up, they have a higher priority than variables from the configuration file. Command line //parameters are most important priority. To change a parameter in the command line you need to specify its name in the form of a section name and //parameter name (with capital letters!). Here is an example to change the port and the name:
//
//	'yourservice -Main/Port 85 -Main/Name Happy`
//
type Tuner struct {
	Config
}

func (t *Tuner) parseCommandLine() (map[string]string, error) {
	argCommLine := make(map[string]string)
	ln := len(os.Args)
	//if ln%2 != 1 && ln != 1 {
	//	return nil, errors.New("Command Line Error (check the number of arguments)")
	//}
	for i := 1; i < ln-1; i++ {
		key := strings.TrimLeft(os.Args[i], "-")
		if len(key) == len(os.Args[i])-1 {
			i++
			argCommLine[key] = os.Args[i]
		} // else {
		//	return nil, errors.New("Command Line Error (no dashes)")
		//}
	}
	return argCommLine, nil
}

func (t *Tuner) env(argCommLine map[string]string) error {
	for _, str := range os.Environ() {
		parEnv := strings.Split(str, DELIMITER_PARAM)
		parKey := strings.Split(parEnv[0], DELIMITER_COMMAND)
		if len(parKey) == 2 {
			if err := t.reflecting(parKey[0], parKey[1], parEnv[1]); err != nil {
				return err
			}
		}
	}
	for k, str := range argCommLine {
		parKey := strings.Split(k, DELIMITER_COMMAND)
		if len(parKey) == 2 {
			if err := t.reflecting(parKey[0], parKey[1], str); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *Tuner) reflecting(key1 string, key2 string, str string) error {
	t1 := reflect.TypeOf(*t)
	v1 := reflect.ValueOf(t)
	v1 = reflect.Indirect(v1)
	if m, ok := t1.FieldByName(key1); ok {
		v2 := v1.FieldByName(key1).Addr()
		v2 = reflect.Indirect(v2)
		if _, okk := m.Type.FieldByName(key2); okk {
			v3 := v2.FieldByName(key2).Addr()
			v3 = reflect.Indirect(v3)
			return t.switchType(v3, str)
		}
	}
	return nil
}

func (t *Tuner) switchType(v3 reflect.Value, str string) error {
	tp := v3.Type()
	switch tp {
	case reflect.TypeOf(""):
		v3.SetString(str)
	case reflect.TypeOf(1):
		num, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return err
		}
		v3.SetInt(num)
	case reflect.TypeOf(0.1):
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
		v3.SetFloat(num)
	default:
		return errors.New("Not supported by type!")
	}
	return nil
}
