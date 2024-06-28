package schema

import (
	"errors"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// Config object
type Config struct {
	Name         string
	IndexServer  string
	SearchServer string
	Fields       map[string]Field
	FieldsSort   []string //记录字段原生排序
}

type Setting struct {
	Schema *Schema
	Logger *Schema
	Conf   *Config
}

// LoadConf from file
func LoadConf(file string) (*Setting, error) {
	cfg := new(Config)
	err := loadIniFile(file, cfg)
	if err != nil {
		return nil, err
	}
	return cfg.checkValid()
}

func (config *Config) checkValid() (*Setting, error) {
	if config.Name == "" {
		return nil, errors.New("Missing the name of project")
	}
	setting := &Setting{}

	if config.IndexServer == "" {
		config.IndexServer = "127.0.0.1:8383"
	}
	if config.SearchServer == "" {
		config.SearchServer = "127.0.0.1:8384"
	}
	if strings.Index(config.IndexServer, ":") == 0 {
		config.IndexServer = "127.0.0.1" + config.IndexServer
	}
	if strings.Index(config.SearchServer, ":") == 0 {
		config.SearchServer = "127.0.0.1" + config.SearchServer
	}
	if m, _ := regexp.Match("^[1-9]\\d{1,4}$", []byte(config.IndexServer)); m {
		config.IndexServer = "127.0.0.1:" + config.IndexServer
	}
	if m, _ := regexp.Match("^[1-9]\\d{1,4}$", []byte(config.SearchServer)); m {
		config.IndexServer = "127.0.0.1:" + config.SearchServer
	}
	setting.Conf = config
	sch, err := newSchema(config.Fields, config.FieldsSort)
	if err != nil {
		return nil, err
	}
	setting.Schema = sch

	logCfg := new(Config)
	err = loadIni(logger, logCfg)
	if err != nil {
		return nil, err
	}
	lsch, err1 := newSchema(logCfg.Fields, config.FieldsSort)
	if err1 != nil {
		return nil, err1
	}
	setting.Logger = lsch
	return setting, nil
}

// 兼容讯搜原生配置
func loadIniFile(file string, cfg *Config) error {
	fileData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	return loadIni(string(fileData), cfg)
}

func loadIni(data string, cfg *Config) error {
	cfg.Fields = map[string]Field{}
	dataLins := strings.Split(data, "\n")
	lastField := ""          //上次字段名
	fieldsSort := []string{} //记录字段顺序
	for _, dataLin := range dataLins {
		dataLin = strings.TrimSpace(dataLin)
		if dataLin != "" {
			if strings.HasPrefix(dataLin, "project.name") {
				cfg.Name = SplitVal(dataLin)
			}
			if strings.HasPrefix(dataLin, "server.index") {
				cfg.IndexServer = SplitVal(dataLin)
			}
			if strings.HasPrefix(dataLin, "server.search") {
				cfg.SearchServer = SplitVal(dataLin)
			}
			//处理字段
			if strings.HasPrefix(dataLin, "[") {
				lastField = strings.ReplaceAll(dataLin, "[", "")
				lastField = strings.ReplaceAll(lastField, "]", "")
				fieldsSort = append(fieldsSort, lastField)
				cfg.Fields[lastField] = Field{}
			}
			if strings.HasPrefix(dataLin, "type") {
				if field, ok := cfg.Fields[lastField]; ok {
					field.Type = SplitVal(dataLin)
					cfg.Fields[lastField] = field
				}
			}
			if strings.HasPrefix(dataLin, "index") {
				if field, ok := cfg.Fields[lastField]; ok {
					field.Index = SplitVal(dataLin)
					cfg.Fields[lastField] = field
				}
			}
			if strings.HasPrefix(dataLin, "phrase") {
				if field, ok := cfg.Fields[lastField]; ok {
					field.Phrase = SplitVal(dataLin)
					cfg.Fields[lastField] = field
				}
			}
			if strings.HasPrefix(dataLin, "weight") {
				if field, ok := cfg.Fields[lastField]; ok {
					field.Weight = Str2uint16(SplitVal(dataLin))
					cfg.Fields[lastField] = field
				}
			}
			if strings.HasPrefix(dataLin, "fid") {
				if field, ok := cfg.Fields[lastField]; ok {
					field.Fid = Str2uint8(SplitVal(dataLin))
					cfg.Fields[lastField] = field
				}
			}

		}
	}
	cfg.FieldsSort = fieldsSort
	return nil
}

// 分割行
func SplitVal(data string) string {
	dataLins := strings.Split(data, "=")
	if len(dataLins) == 2 {
		return strings.TrimSpace(dataLins[1])
	}
	return data
}

// 字符串转unit16
func Str2uint16(str string) uint16 {
	num, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0
	}
	u16 := uint16(num)
	return u16
}

// 字符串转unit8
func Str2uint8(str string) uint8 {
	num, err := strconv.ParseUint(str, 10, 16)
	if err != nil {
		return 0
	}
	u8 := uint8(num)
	return u8
}
