// Copyright 2018 Author: José FERRER VILLENA
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This program read a json file and save in indice database

package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	scribble "github.com/nanobox-io/golang-scribble"
)

// Constant
const pathConfig = "configuration.json"

func deleteFiles(paths []string) {

	for i := 0; i < len(paths); i++ {
		path := paths[i]
		err := os.Remove(path)

		if err != nil {
			fmt.Println("It is not possible to delete the file ", path)
			fmt.Println(err, "\n")
		}
	}

}

func create_indx(path string, fname string, date string) {
	hola := IndObj{Names: fname, DateCreation: date}
	new, err := json.Marshal(hola)
	err = ioutil.WriteFile(path+fname, new, 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", hola)

}

func get_filename(path string, db_path string, ind_path string) []string {

	// Read files from path
	var nameFiles []string
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		file_path := path + f.Name()
		fmt.Println(file_path)
		if filepath.Ext(file_path) == ".xml" {

			// Open xmlFile
			xmlFile, _ := os.Open(file_path)
			// defer the closing of xmlFile
			defer xmlFile.Close()
			// read opened xmlFile as a byte array.
			byteValue, _ := ioutil.ReadAll(xmlFile)
			// we initialize our array
			var xmldata xml2JSON
			// we unmarshal the byteArray which contains the
			// xmlFiles content into 'xmldata' which we defined above
			xml.Unmarshal(byteValue, &xmldata)

			//**************** Set JSON Database **************************
			// create a new scribble database, providing a destination for the database to live
			db, _ := scribble.New(db_path, nil)
			// add some new MWL study to the database
			fileXML := strings.Split(f.Name(), ".")
			xmlID, _ := strconv.Atoi(fileXML[0])

			db.Write("DB_MWL", strconv.Itoa(xmlID), xml2JSON{LastName: xmldata.LastName,
				StudyTime:     xmldata.StudyTime,
				ProcedureName: xmldata.ProcedureName,
				Weight:        xmldata.Weight,
				BirthDate:     xmldata.BirthDate,
				Age:           xmldata.Age,
				Names:         xmldata.Names,
				StudyNumbers:  xmldata.StudyNumbers,
				Modality:      xmldata.Modality,
				ProcedureCode: xmldata.ProcedureCode,
				Date:          xmldata.Date,
				PatientID:     xmldata.PatientID,
				Sex:           xmldata.Sex,
				ModalityType:  xmldata.ModalityType,
				IDPhysician:   xmldata.IDPhysician})

			fmt.Printf("Database writing success\n")

			create_indx(ind_path, fileXML[0] + ".json", xmldata.Date)

			nameFiles = append(nameFiles, file_path)

		}

	}

	return nameFiles

}

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

	mwl_path := confFile.MwlPath
	db_mwl := confFile.DBMwl
	ind_path := confFile.Index

	for {
		files := get_filename(mwl_path, db_mwl, ind_path)
		if len(files) > 0 {
			deleteFiles(files)
		} else {
			fmt.Printf("No new MWL Studies\n")
		}

		// Set time to wait 5 Minutes
		time.Sleep(5 * time.Minute)
	}
}
