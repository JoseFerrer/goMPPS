// Copyright 2018 Author: José FERRER VILLENA

// This program read a json file, do query c-find to PACS and save selected DICOM Tags in json files
// for Mirth Reading and query to Database.
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings" //borrar si es posible
	"time"
)

func main() {

	// ***************** Open configuration file and set MPPS options *******************************
	// Open JSON File Configuration
	jsonFile, err := os.Open(pathConfig)
	if err != nil {
		fmt.Println("JSON File Configuration is not found. ", err)
	}
	fmt.Println("Successfully Opened configuration.json")
	defer jsonFile.Close()
	// Read our opened jsonFile as a byte array
	fmt.Println("DEspues de defer")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our ConfigMPPS
	var confFile ConfigMPPS
	// we unmarshal our byteArray wich contains our
	// jsonFile's content into 'confFile' which we defined above
	json.Unmarshal(byteValue, &confFile)

	ejecutable := confFile.Executable

	// PACS Configuration
	pacsaetitle := confFile.PacsAETitle
	ip := confFile.PacsIP
	port := confFile.PacsPort

	// Entity Configuration
	entityaetitle := confFile.ENTITYAETitle
	entityip := confFile.ENTITYIP
	entityport := confFile.ENTITYPort

	Tag := confFile.NroTag
	optcfind := confFile.OptionsTags

	// revisar luego agregando el aetitle de la aplicacion
	comando1 := ejecutable + " " + "-b " + entityaetitle + "@" + entityip + ":" + entityport + " " + "-c " + pacsaetitle + "@" + ip + ":" + port + " -L SERIES" + " -m " + Tag + "="
	// *****************************************************************************************

	// ************************ Set Paths: MWL, MPPS, ElapsedTime ******************************
	dbmwl := confFile.DBMwl
	dbmpps := confFile.JSONMppsPath
	elapsedT := time.Duration(confFile.ElapsedTime) * time.Minute
	// *****************************************************************************************

	fmt.Println("Configuration Done.")

	num := 0
	datos := new(Tagdcm)

	for {

		// ************************** Read .json files MWL *************************************
		var files []string

		err := filepath.Walk(dbmwl, func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		})
		if err != nil {
			panic(err)
		}

		// *************************************************************************************
		// Analize only json files MWL
		for _, f := range files {
			nameFSlice := strings.Split(f, "DBMWL")
			nFile := nameFSlice[1]
			n2save := nFile[1:len(nFile)]
			filePath := f
			fmt.Println(filePath)
			fmt.Println(n2save)

			// ************************** Open json file ***************************************
			if filepath.Ext(filePath) == ".json" {

				byteValue := readFile(filePath)
				// we initialize our array
				var jsondata DBmwl
				// we unmarshal our byteArray which contains our
				// jsonFiles content into 'jsondata' which we defined above
				json.Unmarshal(byteValue, &jsondata)
				var studyID string
				studyID = jsondata.StudyNumbers

				// ************************ Query C-FIND ****************************************
				// Comando completed
				comando := comando1 + studyID + " " + optcfind

				// Get C-Find cutting response
				RespStdout := queryCFind(comando)
				fmt.Println("The command has been executed: ", comando)

				// ********************** Dicom validation query *********************************
				validResp := valResponse(RespStdout, respFromPacs)


				// Revisar logica del valResponse no le hace a todos
				if validResp != "Study not found" {

					strMsn := RespStdout
					var n int
					var strCut string

					nSeries := strings.Count(strMsn, respFromPacs)

					for i := 0; i < nSeries; i++ {
						if i == 0 {
							strCut = strMsn
						} else {
							strCut = strMsn[n-(len(respFromPacs)) : len(strMsn)]

						}
						slice_Msn, m := cutMsn(strCut, respFromPacs)
						n = m + n - (len(respFromPacs) + 1)
						// ********************** Extract Data ***************************************
						// Get Struct dicom data tags and save jsonmpps
						// Get AccessionNumber ("0008,0050")
						datos.AccessionNumber = extractMsn(slice_Msn, TagAccession)
						// Get SeriesInstanceUID ("0020,000E")
						datos.SeriesInstanceUID = extractMsn(slice_Msn, TagSeriesInUID)
						// Get StudyInstanceUID ("0020,000D")
						datos.StudyInstanceUID = extractMsn(slice_Msn, TagStudyInsUID)
						// Get SeriesTime ("0008,0031")
						datos.SeriesTime = extractMsn(slice_Msn, TagSeriesTime)
						// Get SeriesDate ("0008,0021")
						datos.SeriesDate = extractMsn(slice_Msn, TagSeriesDate)
						// Get StudyDescription ("0008,1030")
						datos.StudyDescription = extractMsn(slice_Msn, TagStudyDesc)
						// Get SeriesDescription ("0008,103E")
						datos.SeriesDescription = extractMsn(slice_Msn, TagSeriesDesc)
						// Get SeriesNumber ("0020,0011")
						datos.SeriesNumber = extractMsn(slice_Msn, TagSeriesN)
						// Get BodyPartExamined ("0018,0015")
						datos.BodyPartExamined = extractMsn(slice_Msn, TagBodyPart)
						// Get NumberOfSeriesRelatedInstances ("0020,1209")
						datos.NumberOfSeriesRelatedInstances = extractMsn(slice_Msn, TagNumberSRI)
						// Get StationName ()
						datos.StationName = extractMsn(slice_Msn, TagStationName)

						//***************** Save JSON MPPS Data ***********************************
						tagsJSON, _ := json.Marshal(datos)
						nameF := strings.Split(n2save, ".")
						err = ioutil.WriteFile(dbmpps+nameF[0]+"_"+strconv.Itoa(i)+".json", tagsJSON, 0644)
						if err != nil {
							fmt.Println("Error to Write MPPS file by ", err)
						}
					}

					//  *********** Delete jsonmwl MPPSStatus and store jsonmpps **************
					fmt.Println(dbmwl + n2save)
					deleteFile(dbmwl + n2save)

					num++

				} else {
					fmt.Println("MWL Study not found.")
				}

			} else {
				fmt.Println("No JSON Modality WorkList found.")
			}
		}

		// Set time to wait: 5 Minutes
		fmt.Println("Waiting some Minutes")
		time.Sleep(elapsedT)
	}

}
