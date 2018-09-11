# goMPPS
Dicom MPPS API for Mirth Connect usign [MirthConnect](https://www.nextgen.com/products-and-services/integration-engine) and [dcm4chee tool kit](https://sourceforge.net/projects/dcm4che/files/dcm4chee-arc-light5/)

## Modules

* MWL: Query Modality WorkList Module to Database
* C-FIND: Query c-Find Module for MPPS

## Get c-find studies



## Get Modality WorkList



## Configuration JSON

This JSON provide to Modules the configuration for Modules

* MWLPath: File Path where Mirth Connect save the XML file coming from RIS
* DBMWL: Modality WorkList JSON database
* Index: Study Index database
* executable: Path of dcm4chee findscu
* PACSAETitle: AETitle of PACS target
* PACSIP: PACS target IP
* PACSPort: PACS target port
* NroTag: DICOM Tag for matching key
* OptionsTags: DICOMs Tag for Specify returning key
* JSONMPPSPath: MPPS JSON database
* ElapsedTime: Time for Queries to Database

```
{
	"MWLPath": "/Users/Name/NNN/xml_dcm/",
	"DBMWL": "/Users/Name/NNN/DB_MWL/",
	"Index": "/Users/Name/NNN/Index/",
	"executable": "/Users/Name/dcm4che-5.14.0/bin/findscu ",
	"PACSAETitle": "DCM4CHEE",
	"PACSIP": "192.168.1.36",
	"PACSPort": "11112",
	"NroTag": "00100010",
	"OptionsTags": "-r 00080050 -r 0020000E -r 0020000D -r 00080031 -r 00080021 -r 00081030 -r 0008103E -r 00200011 -r 00180015 -r 00201209 -r 00081010",
	"JSONMPPSPath": "/Users/Name/NNN/DB_MPPS/",
	"ElapsedTime": 1
}
```

## License

This project is licensed under the MIT License
