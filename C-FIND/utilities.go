package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
)

func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println("It is not possible to delete the file ", path, " by ", err)
	}
}

func readFile(filePath string) []byte {
	// Open our jsonFile
	jsonFile, err := os.OpenFile(filePath, os.O_RDWR, 0644)
		if err != nil {
			fmt.Println("Error to Read JSON ", filePath, " by ", err)
		}
		// defer the closing of our xmlFile so that we can parse it later on
		defer jsonFile.Close()
				

		// ********************* Read JSON File ****************************************
		// read our opened jsonFile as a byte array.
		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println("Error to Read JSON file and transfer as a byte array by ", err)
		}
		return byteValue
}

func extractMsn(msn string, key string) string {

	indx := strings.Index(msn, key)
	if indx != -1 {
		indxi := getIndx(msn, indx, "[")
		indxf := getIndx(msn, indxi, "]")
		if indxi == indxf-1 {
			return ""
		}
		return msn[indxi+1 : indxf]
	}
	return ""
}

func valResponse(msn string, key string) string {
	indx := strings.Index(msn, key)
	if indx == -1 {
		return "Study not found"
	}
	return ""

}

func cutMsn(msn string, key string) (string, int) {

	indx1 := strings.Index(msn, key)
	ix := indx1 + len(key)
	cutmessage := msn[ix:len(msn)]
	indx2 := strings.Index(cutmessage, key)
	if indx2 == -1 {
		indx2 = strings.Index(cutmessage, "status=0H")
	}
	indexf := ix + indx2 + len(key)
	return cutmessage[0:indx2], indexf

}

func getIndx(msn string, indx int, chr string) int {

	inx := 1
	for {
		chrt := msn[indx+inx]
		if string(chrt) == chr {
			break
		}
		inx++
	}

	return (indx + inx)
}

func queryCFind(path string) string {

	// Ejecutar FINDSCU
	cmd := exec.Command("cmd", "/C", path)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("cmd.Run() failedwith %s\n", err)
	}

	return string(out)
}
