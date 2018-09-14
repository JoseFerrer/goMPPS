package main

// DBmwl Struct of Database of Modality WorkList in xml
type DBmwl struct {
	LastName      string `json:"LastName"`
	StudyTime     string `json:"StudyTime"`
	ProcedureName string `json:"ProcedureName"`
	Weight        string `json:"Weight"`
	BirthDate     string `json:"BirthDate"`
	Age           string `json:"Age"`
	Names         string `json:"Names"`
	StudyNumbers  string `json:"StudyNumbers"`
	Modality      string `json:"Modality"`
	ProcedureCode string `json:"ProcedureCode"`
	Date          string `json:"Date"`
	PatientID     string `json:"PatientID"`
	Sex           string `json:"Sex"`
	ModalityType  string `json:"ModalityType"`
	IDPhysician   string `json:"IDPhysician"`
}

// ConfigMPPS Struct of configuration MPPS C-FIND
type ConfigMPPS struct {
	DBMwl        	string
	Executable   	string
	PacsAETitle  	string
	PacsIP       	string
	PacsPort     	string
	ENTITYAETitle string
	ENTITYIP			string
	ENTITYPort		string
	NroTag       	string
	OptionsTags  	string
	JSONMppsPath 	string
	ElapsedTime  	int
}

// Tagdcm Struct of DICOM tags to get from C-FIND
type Tagdcm struct {
	AccessionNumber                string `json:"AccessionNumber"`
	SeriesInstanceUID              string `json:"SeriesInstanceUID"`
	StudyInstanceUID               string `json:"StudyInstanceUID"`
	SeriesTime                     string `json:"seriesTime"`
	SeriesDate                     string `json:"SeriesDate"`
	StudyDescription               string `json:"StudyDescription"`
	SeriesDescription              string `json:"SeriesDescription"`
	SeriesNumber                   string `json:"SeriesNumber"`
	BodyPartExamined               string `json:"bodyPartExamined"`
	NumberOfSeriesRelatedInstances string `json:"numberOfSeriesRelatedInstances"`
	StationName                    string `json:"StationName"`
}
