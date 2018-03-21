package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type packageInfo struct {
	Name string
	Deps []string
}

func main() {
	var infoList []packageInfo
	var err error

	for _, name := range os.Args[1:] {
		infoList = append(infoList, packageInfo{name, []string{}})
	}

	// 一度目
	err, infoList = getPackageInfo(infoList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 二度目
	err, infoList = getPackageInfo(infoList)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// 重複を消す
	tmpMap := make(map[string]bool)
	for _, info := range infoList {
		tmpMap[info.Name] = true
		for _, dep := range info.Deps {
			tmpMap[dep] = true
		}
	}
	fmt.Println("依存パッケージ一覧")
	for result := range tmpMap {
		fmt.Println(result)
	}
}

func getPackageInfo(infoList []packageInfo) (error, []packageInfo) {
	tmpInfoList := infoList
	str := []string{"list", "-json"}

	for _, info := range tmpInfoList {
		str = append(str, info.Name)
		str = append(str, info.Deps...)
	}

	cmd := exec.Command("go", str...)

	r, _ := cmd.StdoutPipe()
	defer r.Close()

	cmd.Start()
	decoder := json.NewDecoder(r)
	for {
		var info packageInfo

		err := decoder.Decode(&info)

		if err != nil {

			if err == io.EOF {
				return nil, tmpInfoList
			}
			log.Printf("%v\n", err)
		}
		tmpInfoList = append(tmpInfoList, info)
	}
}
