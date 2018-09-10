package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func deleteFile(path string) {
	err := os.Remove(path)
	if err != nil {
		fmt.Println("It is not possible to delete the file ", path)
		fmt.Println(err)
	}
}

func extractMsn(msn string, key string) string {

	indx := strings.Index(msn, key)
	// fmt.Println(indx)
	if indx != -1 {
		indxi := getIndx(msn, indx, "[")
		//fmt.Println("Inicio: ", indxi)
		indxf := getIndx(msn, indxi, "]")
		//fmt.Println("Fin: ", indxf)
		if indxi == indxf-1 {
			return "Tag empty"
		}
		//fmt.Println(string(msn[indxi:indxf]))
		return msn[indxi+1 : indxf]
	}
	return "No Tag in DICOM object"
}

func valResponse(msn string, key string) string {
	indx := strings.Index(msn, key)
	if indx == -1 {
		return "Study not found"
	}
	return ""

}

func cutMsn(msn string, key string) string {

	indx := strings.Index(msn, key)
	ix := indx + len(key)
	return msn[ix:len(msn)]

}

func getIndx(msn string, indx int, chr string) int {

	inx := 1
	//fmt.Println("El primer index: ", indx)
	for {
		chrt := msn[indx+inx]
		if string(chrt) == chr {
			//fmt.Println("Valor: ", indx + inx)
			//fmt.Println(string(msn[indx +inx]))
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
		log.Fatalf("cmd.Run() failedwith %s\n", err)
	}

	// Obtener la salida de FINDSCU
	msndcm := string(out)
	lastTag := "NumberOfSeriesRelatedInstances"
	mensajedcm := cutMsn(msndcm, lastTag)
	return mensajedcm
}
