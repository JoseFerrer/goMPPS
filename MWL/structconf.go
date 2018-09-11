package main

// xml2JSON Struct for get xml information
type xml2JSON struct {
  LastName             string `xml:"apellidos"`
	StudyTime            string `xml:"hora_estudio"`
	ProcedureName        string `xml:"nombre_procedimiento"`
	Weight               string `xml:"peso"`
	BirthDate            string `xml:"fecha_nacimiento"`
	Age                  string `xml:"edad"`
	Names                string `xml:"nombres"`
	StudyNumbers         string `xml:"numero_estudio"`
	Modality             string `xml:"modalidad"`
	ProcedureCode        string `xml:"codigo_procedimiento"`
	Date                 string `xml:"fecha_actual"`
	PatientID            string `xml:"id_paciente"`
	Sex                  string `xml:"sexo"`
	ModalityType         string `xml:"tipo_modalidad"`
	IDPhysician          string `xml:"cmp_medico_referente"`
}

// ConfigMPPS Struct of configuration MPPS C-FIND
type ConfigMPPS struct {
	MwlPath      string
	DBMwl        string
	Index		     string
	Executable   string
	PacsAETitle  string
	PacsIP       string
	PacsPort     string
	NroTag       string
	OptionsTags  string
	JSONMppsPath string
  ElapsedTime  int
}

// IndObj Struct for read, write or delete index
type IndObj struct {
	Names				 string
	DateCreation string
}
