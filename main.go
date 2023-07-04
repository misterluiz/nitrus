package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Manifest struct {
	XMLName xml.Name `xml:"manifest"`
	Package string   `xml:"package,attr"`
	// Outros campos relevantes que você queira extrair do AndroidManifest.xml
}

func main() {

	fmt.Println("███╗░░██╗██╗████████╗██████╗░██╗░░░██╗░██████╗")
	fmt.Println("████╗░██║██║╚══██╔══╝██╔══██╗██║░░░██║██╔════╝")
	fmt.Println("██╔██╗██║██║░░░██║░░░██████╔╝██║░░░██║╚█████╗░")
	fmt.Println("██║╚████║██║░░░██║░░░██╔══██╗██║░░░██║░╚═══██╗")
	fmt.Println("██║░╚███║██║░░░██║░░░██║░░██║╚██████╔╝██████╔╝")
	fmt.Println("╚═╝░░╚══╝╚═╝░░░╚═╝░░░╚═╝░░╚═╝░╚═════╝░╚═════╝░")

	if len(os.Args) < 2 {
		fmt.Println("Por favor, especifique o caminho para o arquivo APK.")
		return
	}

	apkPath := os.Args[1]

	cmd := exec.Command("jadx", "-d", "output", apkPath)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Aplicativo decompilado com sucesso.")

	manifestPath := "output/resources/AndroidManifest.xml"

	manifestFile, err := os.Open(manifestPath)
	if err != nil {
		log.Fatal(err)
	}
	defer manifestFile.Close()

	byteValue, err := ioutil.ReadAll(manifestFile)
	if err != nil {
		log.Fatal(err)
	}

	var manifest Manifest
	err = xml.Unmarshal(byteValue, &manifest)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Package:", manifest.Package)

	debuggableString := `android:debuggable="true"`
	if strings.Contains(string(byteValue), debuggableString) {
		fmt.Println("Debuggable: true")
	} else {
		fmt.Println("Debuggable: false")
	}

	clearText := `android:usesCleartextTraffic="true"`
	if strings.Contains(string(byteValue), clearText) {
		fmt.Println("CleartextTraffic: true")
	} else {
		fmt.Println("CleartextTraffic: false")
	}

	backup := `android:allowBackup="true"`
	if strings.Contains(string(byteValue), backup) {
		fmt.Println("AllowBackup: true")
	} else {
		fmt.Println("AllowBackup: false")
	}

	// Extraia outras informações relevantes do arquivo AndroidManifest.xml conforme necessário
}
