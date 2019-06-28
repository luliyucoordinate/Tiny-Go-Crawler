package font

import (
	"regexp"
	"encoding/base64"
	"os"
	"os/exec"
	"log"
	"github.com/donnie4w/dom4g"
	"strconv"
	"hash/crc32"
	"fmt"
)


var fontName = "default.ttf"

func GetFontFile(contents []byte) error {
	re := regexp.MustCompile(`base64,(.*?)'`)
	bs4 := re.FindSubmatch(contents)

	if len(bs4) < 2 {
		return fmt.Errorf("Font: there's no font on this page.")
	}
	fontName = strconv.FormatUint(uint64(crc32.ChecksumIEEE(bs4[1])), 10)
	bs4ToStr, err := base64.StdEncoding.DecodeString(string(bs4[1]))
	if err != nil {
		log.Printf("Font: error get font: %v", err)
	}

	if _, err := os.Stat(fontName); !os.IsExist(err) {
		fontFile, err := os.Create(fontName)
		defer fontFile.Close()
		if err != nil {
			log.Printf("Font: write font error: %v", err)
		}
		fontFile.Write(bs4ToStr)
	}
	return nil
}

func BuildFontFile() {
	cmd := exec.Command("ttx", fontName)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Font: build font error: %v", err)
	}
}

func GetFontMap(contents []byte) (map[string]string, error) {
	err := GetFontFile(contents)
	if err != nil {
		return nil, err
	}

	BuildFontFile()
	defer os.Remove(fontName)
	defer os.Remove(fontName+".ttx")

	res := make(map[string]string)
	file, err := os.Open(fontName + ".ttx")
	if err != nil {
		log.Printf("Font: get font map error: %v", err)
	}
	ele, err := dom4g.LoadByStream(file)
	if err != nil {
		log.Printf("Font: get font map error: %v", err)
	}

	eles := ele.Node("cmap").Node("cmap_format_4").Nodes("map")

	for _, v := range eles {
		code, _ := v.AttrValue("code")
		name, _ := v.AttrValue("name")

		num, err := strconv.Atoi(name[len(name)-2:])

		if err != nil {
			log.Printf("Font: get font map error: %v", err)
		}
		res[code] = strconv.Itoa(num-1)
	}

	return res, nil
}