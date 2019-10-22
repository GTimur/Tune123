package tune123

import (
	"errors"
	"log"
	"os"
)

type Playlist struct {
	Name        string
	PlayList    Database
	FileName    string
	FileNameM3u string
}

func (p *Playlist) Add(recordId int64) error {

	r, err := GLOBALDATABASE.Record(recordId)
	if err != nil {
		return err
	}

	p.PlayList.Rec = append(p.PlayList.Rec, r)
	return nil
}

func (p *Playlist) Remove(recordID int64) error {

	index := -1

	for i, r := range p.PlayList.Rec {
		if r.ID != recordID {
			continue
		}
		index = i
	}

	if index == -1 {
		return errors.New("PLAYLIST: Нет данных для удаления")
	}

	// удаляем указанный элемент по индексу
	p.PlayList.Rec = append(p.PlayList.Rec[:index], p.PlayList.Rec[index+1:]...)

	return nil
}

func (p *Playlist) WriteJSON() error {
	err := WriteJSON(p, p.FileName)
	if err != nil {
		log.Println("WriteJSON [", p.FileName, "]:", err)
	}
	return err
}

func (p *Playlist) ReadJSON() error {
	err := ReadJSON(p, p.FileName)
	if err != nil {
		log.Println("ReadJSON [", p.FileName, "]:", err)
	}
	return err
}

func (p *Playlist) CreateJSONFile() error {
	err := CreateJSONFile(p, p.FileName)
	if err != nil {
		log.Println("CreateFile [", p.FileName, "]:", err)
	}
	return err
}

/*======================== m3u playlist ==========================*/
func (p *Playlist) Write(rewrite bool) error {
	//Создадим файл m3u и перезапишем если с таким именем уже имеется
	err := p.createM3uFile(rewrite)
	if err != nil {
		log.Println("CreateM3uFile [", p.FileNameM3u, "]:", err)
		return err
	}

	file, err := os.OpenFile(p.FileNameM3u, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	if err != nil {
		log.Printf("error writing m3u file: %v\n", err)
		return err
	}

	//пишем в файл
	for _, r := range p.PlayList.Rec {
		_, err = file.WriteString(r.FullPath + "\r\n")
		if err != nil {
			log.Println("Writing m3u error:", err)
			return err
		}
	}
	file.Sync()
	return err
}

func (p *Playlist) createM3uFile(rewrite bool) (err error) {

	if _, err = os.Stat(p.FileNameM3u); err == nil && !rewrite {
		//File exist and will not be rewrited
		//fmt.Println(os.IsExist(err),err)
		return err
	}
	file, err := os.Create(p.FileNameM3u)
	defer file.Close()
	return err
}

/*==================== end of m3u =================================*/
