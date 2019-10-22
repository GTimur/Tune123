package tune123

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type objJSON interface {
	WriteJSON() error
	ReadJSON() error
	CreateJSONFile() error
}

// WriteJSON - сохраняет файл конфигурации в формате JSON
func WriteJSON(o objJSON, filepath string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		log.Printf("error writing JSON file: %v\n", err)
		return err
	}

	//готовим данные JSON (конвертируем в экспортируемый вид)
	jsonConfig := o

	//пишем в файл
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(&jsonConfig)
	if err != nil {
		log.Printf("JSON encoder error: %v", err)
	}
	return err
}

//ReadJSON - читает конфигурационный json файл
func ReadJSON(o objJSON, filepath string) (err error) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Printf("error reading JSON file: %v\n", err)
		return err
	}

	//Готовим для импорта структуру JSON
	var jsonConfig objJSON

	//Читаем файл JSON
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&jsonConfig)
	if err != nil {
		log.Printf("JSON decoder error: %v", err)
		return err
	}

	o = jsonConfig

	return err
}

// CreateJSONFile - создает файл конфигурации если он не существует
func CreateJSONFile(o objJSON, filepath string) (err error) {
	if _, err = os.Stat(filepath); err == nil {
		//File exist and will not be rewrited
		//fmt.Println(os.IsExist(err),err)
		return err
	}
	file, err := os.Create(filepath)
	defer file.Close()
	return err
}

// PrintJSON - возвращает содержимое структуры в строке в JSON
func PrintJSON(o objJSON) (out string, err error) {
	//пишем в файл
	b, err := json.MarshalIndent(o, "", "\t")
	if err != nil {
		return "", err
	}
	//n := bytes.Index(b, []byte{0})
	out = fmt.Sprintf("%s", b)
	//out = string(b[:n])
	return out, err
}
