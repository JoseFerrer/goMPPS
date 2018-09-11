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
	MwlPath      string
	DBMwl        string
	Index        string
	Executable   string
	PacsAETitle  string
	PacsIP       string
	PacsPort     string
	NroTag       string
	OptionsTags  string
	JSONMppsPath string
	ElapsedTime  int
}

// Tagdcm Struct of DICOM tags to get from C-FIND
type Tagdcm struct {
	IDSeries                       int    `json:"id_series"`
	AccessionNumber                string `json:"accessionNumber"`
	SeriesInstanceUID              string `json:"seriesinstanceUID"`
	StudyInstanceUID               string `json:"studyinstanceUID"`
	SeriesTime                     string `json:"seriestime"`
	SeriesDate                     string `json:"seriesdate"`
	StudyDescription               string `json:"studydescription"`
	SeriesDescription              string `json:"seriesdescription"`
	SeriesNumber                   string `json:"seriesnumber"`
	BodyPartExamined               string `json:"bodypartexamined"`
	NumberOfSeriesRelatedInstances string `json:"numberofseriesrelatedinstances"`
	StationName                    string `json:"stationname"`
}
