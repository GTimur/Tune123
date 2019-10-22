package tune123

var (
	GLOBALCONFIG   Config
	GLOBALMASK     = []string{"*.mp3"}
	GLOBALDATABASE Database
	GLOBALPLAYLIST Playlist
	GLOGBALCATALOG Catalog
)

func InitGlobal() error {
	GLOBALCONFIG.FileName = "config.json"
	GLOBALCONFIG.Log = LogFile{Filename: "player.log"}
	err := GLOBALCONFIG.Log.Init()
	if err != nil {
		return err
	}

	// Сканируем все указанные папки для создания каталога файлов и папок

	for _, r := range GLOBALCONFIG.AudioPath {
		if err := GLOGBALCATALOG.DirScan(r, GLOGBALCATALOG.MaxID); err != nil {
			return err
		}
	}

	return err
}
