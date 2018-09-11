# goMPPS
Dicom MPPS API for Mirth Connect usign MirthConnect

## Modules


## Get c-find studies


## Get Modality WorkList


## Configuration JSON

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
