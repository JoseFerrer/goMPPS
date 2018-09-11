package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {

	// ***************** Open configuration file and set MPPS options *******************************
	// Open JSON File Configuration
	jsonFile, err := os.Open(pathConfig)
	if err != nil {
		fmt.Println("JSON File Configuration is not found. ", err)
	}
	fmt.Println("Successfully Openend configuration.json")
	defer jsonFile.Close()
	// Read our opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// we initialize our ConfigMPPS
	var confFile ConfigMPPS
	// we unmarshal our byteArray wich contains our
	// jsonFile's content into 'confFile' which we defined above
	json.Unmarshal(byteValue, &confFile)

	ejecutable := confFile.Executable
	pacsaetitle := confFile.PacsAETitle
	ip := confFile.PacsIP
	port := confFile.PacsPort
	Tag := confFile.NroTag
	optcfind := confFile.OptionsTags

	comando1 := ejecutable + "-c " + pacsaetitle + "@" + ip + ":" + port + " -m " + Tag + "="
	// *****************************************************************************************

	// ************************ Set Paths: MWL, MPPS, ElapsedTime ******************************
	dbmwl := confFile.DBMwl
	dbmpps := confFile.JSONMppsPath
	indexdb := confFile.Index
	elapsedT := time.Duration(confFile.ElapsedTime) * time.Minute
	// *****************************************************************************************

	num := 0
	datos := new(Tagdcm)

	for {

		// ************************** Read .json files MWL *************************************
		files, err := ioutil.ReadDir(indexdb)
		if err != nil {
			log.Fatal(err)
		}
		// *************************************************************************************

		// Analize only json files MWL
		for _, f := range files {
			filePath := dbmwl + "DB_MWL/" + f.Name()

			// ************************** Open json file ***************************************
			if filepath.Ext(filePath) == ".json" {

				// Open our jsonFile
				jsonFile, _ := os.Open(filePath)
				// defer the closing of our xmlFile so that we can parse it later on
				defer jsonFile.Close()

				// ********************* Read JSON File ****************************************
				// read our opened jsonFile as a byte array.
				byteValue, _ := ioutil.ReadAll(jsonFile)

				// we initialize our array
				var jsondata DBmwl
				// we unmarshal our byteArray which contains our
				// jsonFiles content into 'jsondata' which we defined above
				json.Unmarshal(byteValue, &jsondata)
				var studyID string
				studyID = jsondata.StudyNumbers

				// ************************ Query C-FIND ****************************************
				// Comando completo
				comando := comando1 + studyID + " " + optcfind

				// Get C-Find cutting response
				RespStdout := queryCFind(comando)
				fmt.Println("The command has been executed: ", comando)

				// ********************** Dicom validation query *********************************
				validResp := valResponse(RespStdout, respFromPacs)

				// Revisar logica del valResponse no le hace a todos
				if validResp != "Study not found" {

					// ********************** Extract Data ***************************************
					// Get Struct dicom data tags and save jsonmpps
					// Id Series
					datos.IDSeries = int(num)
					// Get AccessionNumber ("0008,0050")
					datos.AccessionNumber = extractMsn(RespStdout, TagAccession)
					// Get SeriesInstanceUID ("0020,000E")
					datos.SeriesInstanceUID = extractMsn(RespStdout, TagSeriesInUID)
					// Get StudyInstanceUID ("0020,000D")
					datos.StudyInstanceUID = extractMsn(RespStdout, TagStudyInsUID)
					// Get SeriesTime ("0008,0031")
					datos.SeriesTime = extractMsn(RespStdout, TagSeriesTime)
					// Get SeriesDate ("0008,0021")
					datos.SeriesDate = extractMsn(RespStdout, TagSeriesDate)
					// Get StudyDescription ("0008,1030")
					datos.StudyDescription = extractMsn(RespStdout, TagStudyDesc)
					// Get SeriesDescription ("0008,103E")
					datos.SeriesDescription = extractMsn(RespStdout, TagSeriesDesc)
					// Get SeriesNumber ("0020,0011")
					datos.SeriesNumber = extractMsn(RespStdout, TagSeriesN)
					// Get BodyPartExamined ("0018,0015")
					datos.BodyPartExamined = extractMsn(RespStdout, TagBodyPart)
					// Get NumberOfSeriesRelatedInstances ("0020,1209")
					datos.NumberOfSeriesRelatedInstances = extractMsn(RespStdout, TagNumberSRI)

					//  *********** Delete jsonmwl MPPSStatus and store jsonmpps **************
					deleteFile(indexdb + f.Name())

					//***************** Save JSON MPPS Data ***********************************
					tagsJSON, _ := json.Marshal(datos)
					err = ioutil.WriteFile(dbmpps+f.Name(), tagsJSON, 0644)
					fmt.Printf("%+v", datos)

					num++

				} else {
					fmt.Println("MWL Study not found.")
				}

				// Set time to wait: 5 Minutes
				time.Sleep(elapsedT)

			}
		}
	}

}