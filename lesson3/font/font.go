package font

import (
	"regexp"
	"encoding/base64"
	"log"
	"os"
)


const fontName = "online_font.ttf"

func GetFontFile(contents []byte) {
	re := regexp.MustCompile(`base64,(.*?)\'`)
	bs4 := re.FindSubmatch(contents)

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
}