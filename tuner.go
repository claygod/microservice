package main

// Microservice
// Tuner
// Copyright © 2016 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"flag"
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
)

// NewTuner - create a new Tuner-struct
func NewTuner(path string) (*Tuner, error) {
	var confile string
	flag.StringVar(&confile, "confile", path, "Config file") // microservice -confile=config.toml
	flag.Parse()
	s := &Tuner{}
	if _, err := toml.DecodeFile(confile, s); err != nil {
		return nil, err
	}
	s.env()
	return s, nil
}

/* Tuner structure
Configuring using the configuration file, environment variables, and command-line variables.
The default configuration is loaded from the specified file ( `config.toml`).
The configuration file can be changed from the command line like this:

	`yourservice confile=config.toml`

If the operating system environment variables are set up, they have a higher priority than variables from the configuration file. Command line parameters are most important priority. To change a parameter in the command line you need to specify its name in the form of a section name and parameter name (with capital letters!). Here is an example to change the port:

	'yourservice Main/Port=:85`

*/
type Tuner struct {
	Config
}

func (s *Tuner) env() {
	for _, str := range os.Environ() {
		parEnv := strings.Split(str, DELIMITER_PARAM)
		parKey := strings.Split(parEnv[0], DELIMITER_COMMAND)
		if len(parKey) == 2 {
			s.reflecting(parKey[0], parKey[1], parEnv[1])
		}
	}
	args := s.commandLineParser()
	for k, str := range args {
		parKey := strings.Split(k, DELIMITER_COMMAND)
		if len(parKey) == 2 {
			s.reflecting(parKey[0], parKey[1], str)
		}
	}
}

func (s *Tuner) reflecting(key1 string, key2 string, str string) {
	t1 := reflect.TypeOf(*s)
	v1 := reflect.ValueOf(s)
	v1 = reflect.Indirect(v1)
	if m, ok := t1.FieldByName(key1); ok {
		v2 := v1.FieldByName(key1).Addr()
		v2 = reflect.Indirect(v2)
		if _, okk := m.Type.FieldByName(key2); okk {
			v3 := v2.FieldByName(key2).Addr()
			v3 = reflect.Indirect(v3)
			tp := v3.Type()
			switch tp {
			case reflect.TypeOf(""):
				v3.SetString(str)
			case reflect.TypeOf(1):
				num, err := strconv.ParseInt(str, 10, 64)
				if err != nil {
					panic(err.Error())
				}
				v3.SetInt(num)
			case reflect.TypeOf(0.1):
				num, err := strconv.ParseFloat(str, 64)
				if err != nil {
					panic(err.Error())
				}
				v3.SetFloat(num)
			default:
				panic("Not supported by type!")
			}
		}
	}
}

func (s *Tuner) commandLineParser() map[string]string {
	var arg []string
	args := make(map[string]string)
	arg = os.Args
	for i := len(arg) - 1; i >= 0; i-- {
		arg[i] = strings.TrimLeft(arg[i], "-")
		twoSide := strings.Split(arg[i], DELIMITER_PARAM)
		key := twoSide[0]
		value := strings.Replace(arg[i], key+DELIMITER_PARAM, "", 1)
		args[key] = value
	}
	return args
}
