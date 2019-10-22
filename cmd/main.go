package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
	"tune123/pkg/tune123"
)

func main() {
	fmt.Println("Hi there!")
	/*
		err := Setup("VOLUME", "50")
		if err != nil {
		    fmt.Println("Error")
		}

		err = Setup("PLAYER", "STOP")
		if err != nil {
		    fmt.Println("Error")
		}
	*/

	/*CONFIG SECTION*/
	tune123.GLOBALCONFIG.AudioPath = []string{"D:\\TEST_MUSIC\\ROOT1", "D:\\TEST_MUSIC\\ROOT2"}
	tune123.GLOBALCONFIG.PlaylistPath = "./playlist"
	tune123.GLOBALCONFIG.InitPlay = false
	tune123.GLOBALCONFIG.InitVolume = 50
	err := tune123.InitGlobal()
	if err != nil {
		fmt.Println("Application initialization:", err)
		os.Exit(2)
	}
	tune123.GLOBALCONFIG.CreateJSONFile() // создадим конфиг если его нет
	tune123.GLOBALCONFIG.WriteJSON()

	/*END CONFIG SECTION*/

	/**/
	tune123.GLOBALDATABASE.BuildDatabase()
	tune123.GLOBALPLAYLIST.Add(tune123.GLOBALDATABASE.Rec[1].ID)
	tune123.GLOBALPLAYLIST.Add(tune123.GLOBALDATABASE.Rec[2].ID)
	tune123.GLOBALPLAYLIST.Add(tune123.GLOBALDATABASE.Rec[10].ID)
	tune123.GLOBALPLAYLIST.Add(tune123.GLOBALDATABASE.Rec[12].ID)
	//tune123.GLOBALPLAYLIST.Remove(tune123.GLOBALDATABASE.Rec[1].ID)

	fmt.Println("PLAYLIST:", tune123.GLOBALPLAYLIST)
	tune123.GLOBALPLAYLIST.FileName = "playlist.json"
	tune123.GLOBALPLAYLIST.FileNameM3u = "playlist01.m3u"
	tune123.GLOBALPLAYLIST.CreateJSONFile()
	tune123.GLOBALPLAYLIST.WriteJSON()
	err = tune123.GLOBALPLAYLIST.Write(true)
	if err != nil {
		fmt.Println(err)
	}
	/*
		var p croi.Player
		p.PlayerName = "MYPLAYER"
		p.Binary = "mplayer.exe"
		p.Path = "D:\\TEST\\MPLAYER"
		p.Command = make(chan string)
		go p.Exec()
		p.Command <- "volume 50 1"
		p.Command <- "get_file_name"
		time.Sleep(time.Second * 5)
		p.Command <- "get_meta_album"
		//time.Sleep(time.Second * 5)
		//p.Command <- "pt_step 1"
		time.Sleep(time.Second * 10)
		p.Command <- "Quit"

		/**/

	/*if err := croi.GLOBALTREE.TreeJSON(tune123.GLOBALDATABASE); err != nil {
		fmt.Println(err)
	    }

	    fmt.Println("GLOBALTREE:")

	    for _, r := range tune123.GLOBALTREE {
		fmt.Println(r.Key, r.Title)
	    }
	    err = tune123.ScanLevel(tune123.GLOBALDATABASE)
	    if err != nil {
		fmt.Println(err)
	    }*/

	//fmt.Println("CATALOG:::")
	//for _, r := range tune123.GLOGBALCATALOG.Files {
	//fmt.Println(r)
	//}

	str, err := tune123.GLOGBALCATALOG.PrintJSON()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("CATALOG-JSON:::", str)

	//HTTP server init
	tune123.HTTPServerConfig.SetManagerSrv("127.0.0.1", uint16(4400))
	var server tune123.WebCtl
	server.SetHost(net.ParseIP(tune123.HTTPServerConfig.ManagerSrvAddr()))
	server.SetPort(tune123.HTTPServerConfig.ManagerSrvPort())
	server.StartServe()

	WaitExit := false

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Printf("\nReceived %v, shutdown procedure initiated.\n\n", sig)
			tune123.Quit <- 1
			WaitExit = true
		}
	}()

	ticker := time.NewTicker(time.Second * 2)

	for range ticker.C {
		if !WaitExit {
			//fmt.Println("Tick!")
			continue
		}
		break
	}
	ticker.Stop()

}

// Применение настроек для микшера и плеера
func Setup(command string, parameter string) error {
	switch command {
	case "VOLUME":
		fmt.Println("Volume control:", parameter)
	case "PLAYER":
		fmt.Println("Player control:", parameter)
	}

	return nil
}
