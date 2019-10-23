package tune123

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type File struct {
	ID       int64  `json:"key"` //номер файла
	ParentID int64  //номер родителя
	IsDir    bool   `json:"folder,omitempty"` //признак директории
	Name     string `json:"title"`
}

type Catalog struct {
	MaxID       int64
	MaxParentID int64
	Files       []File `json:"children,omitempty"`
}

//FindFiles - search files in folder by selected masks
func FindFiles(dir string, mask []string) (files map[string]string, err error) {
	var list []string
	files = make(map[string]string)

	for i := range mask {
		list, err = filepath.Glob(dir + "/" + strings.ToUpper(mask[i]))
		if err != nil {
			log.Println("FindFiles error: ", err)
			return nil, err
		}
		//files = append(files, list...)
		for _, f := range list {
			files[f] = mask[i]
		}
	}
	for i := range mask {
		list, err = filepath.Glob(dir + "/" + strings.ToLower(mask[i]))
		if err != nil {
			log.Println("FindFiles error: ", err)
			return nil, err
		}
		//files = append(files, list...)
		for _, f := range list {
			files[f] = mask[i]
		}
	}
	return files, err
}

//FindAllFiles - search files in all subdirectories by selected masks
func FindAllFiles(rootdir string, mask []string) (files map[string]string, err error) {
	var dirs []string
	files = make(map[string]string)

	dirs, err = FindAllDirs(rootdir, "")
	if err != nil {
		log.Fatalf("FindAllFiles error: %v", err)
	}

	for _, k := range dirs {
		f, err := FindFiles(k, mask)
		if err != nil {
			log.Fatalf("FindAllFiles error: %v", err)
		}

		for kk := range f {
			files[kk] = filepath.Dir(kk)
		}
	}
	return files, err
}

func FindAllDirs(rootdir string, subDirToSkip string) (files []string, err error) {

	/*	err = os.Chdir(rootdir)
		if err != nil {
		    fmt.Printf("error chdir the path %q: %v\n", rootdir, err)
		    return
		}*/

	err = filepath.Walk(rootdir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() && info.Name() == subDirToSkip {
			//skipping a dir
			return filepath.SkipDir
		}
		//save dir
		if info.IsDir() {
			files = append(files, path)
			return nil
		}
		return err
	})
	if err != nil {
		fmt.Printf("error walking the path %q: %v\n", rootdir, err)
		return
	}

	return
}

//DirScan - собирает каталог файлов и директорий внутри корневой директории
//parentID - текущий каталог родителя (номер)
func (c *Catalog) DirScan(rootpath string, parentID int64) error {
	files, err := ioutil.ReadDir(rootpath)
	if err != nil {
		return err
	}

	var f File //временная переменная для File

	for _, file := range files {
		if !file.IsDir() {
			//Если это файл
			f.IsDir = false
			f.Name = file.Name()
			f.ID = c.MaxID + 1
			f.ParentID = parentID
			c.MaxID++
			c.Files = append(c.Files, f)
			continue
		}
		//Если это директория
		f.IsDir = true
		f.Name = file.Name()
		f.ID = c.MaxID + 1
		c.MaxID++
		f.ParentID = parentID
		c.Files = append(c.Files, f)
		if err := c.DirScan(filepath.Join(rootpath, f.Name), f.ID); err != nil {
			return err
		}
	}
	return err
}

/* JSON PART*/
func (c *Catalog) PrintJSON() (out string, err error) {
	out, err = PrintJSON(c)
	if err != nil {
		log.Println("PrintJSON:", err)
	}
	return out, err
}

func (c *Catalog) WriteJSON() error {
	fileName := "Catalog.json"
	err := WriteJSON(c, fileName)
	if err != nil {
		log.Println("WriteJSON [", fileName, "]:", err)
	}
	return err
}

func (c *Catalog) ReadJSON() error {
	fileName := "Catalog.json"
	err := ReadJSON(c, fileName)
	if err != nil {
		log.Println("ReadJSON [", fileName, "]:", err)
	}
	return err
}

func (c *Catalog) CreateJSONFile() error {
	fileName := "Catalog.json"
	err := CreateJSONFile(c, fileName)
	if err != nil {
		log.Println("CreateJSONFile [", fileName, "]:", err)
	}
	return err
}

/* END JSON PART*/
