package office_converter

type DocumentFormat string

const (
	OfficeWord1997DOC        = "doc"
	OfficeWord2007DOCX       = "docx"
	OfficePowerPoint1997PPT  = "ppt"
	OfficePowerPoint2007PPTX = "pptx"
	OfficeExcel1997XLS       = "xls"
	OfficeExcel2007XLSX      = "xlsx"
	LibreOfficeODT           = "odt"
	PDF                      = "pdf"
	TEXT                     = "txt"
	IMAGE                    = "image"
)

var _officeTypes = []DocumentFormat{
	OfficeWord1997DOC,
	OfficeWord1997DOC,
	OfficeWord1997DOC,
	OfficeWord2007DOCX,
	OfficePowerPoint1997PPT,
	OfficePowerPoint2007PPTX,
	OfficeExcel1997XLS,
	OfficeExcel2007XLSX,
	LibreOfficeODT,
}

func (dt DocumentFormat) IsOffice() bool {
	for _, officeType := range _officeTypes {
		if dt == officeType {
			return true
		}
	}
	return false
}
