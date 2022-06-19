package form4

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"reflect"
	"strings"
)

func GetParser(line string, ea []ExtractData) *ExtractData {
	for _, d := range ea {
		if strings.EqualFold(d.Start, line) {
			return &d
		}
	}
	return nil
}

var ExtractingDocument = []ExtractData{
	{XmlReportingOwnerStartLine, XmlReportingOwnerEndLine, extractReportingOwner},
	{XmlIssuerStartLine, XmlIssuerEndLine, extractReportIssuer},
	{XmlStartLine, XmlEndLine, extractReportXml},
}

type ExtractData struct {
	Start string
	End   string
	Proc  extractInfoFunc
}

type extractInfoFunc func(lines []string, data *FileContent) error

func extractReportXml(lines []string, data *FileContent) error {
	err := xml.Unmarshal([]byte(strings.Join(lines, "\n")), &data.Xml)
	if err != nil {
		log.Fatalf("parsing error: %v", err)
	}

	//fmt.Printf("xml content: %+v\n", data.Xml)
	return nil
}

func ExtractReportDocument(lines []string, data *FileContent) error {
	return extractItemProcess(lines, &data.Document)
}

func extractReportingOwner(lines []string, data *FileContent) error {
	rol := lines[1 : len(lines)-1]
	processLineByLineExtraction(rol, extractingReportingDataOwner, data)
	//fmt.Printf("Reporting owner: %v\n", Serialize(data.Reporter))
	return nil
}

var extractingReportingDataOwner = []ExtractData{
	{XmlReportingOwnerInfoStartLine, XmlReportingOwnerInfoEndLine, extractOwnerData},
	{XmlReportingFilingInfoStartLine, XmlReportingFilingInfoEndLine, extractFilingValue},
	{XmlReportingMailingInfoStartLine, XmlReportingMailingInfoEndLine, extractAddress},
}

func extractOwnerData(data []string, content *FileContent) error {
	return extractItemProcess(data, &content.Reporter.OwnerData)
}

func extractFilingValue(data []string, content *FileContent) error {
	return extractItemProcess(data, &content.Reporter.FilingValues)
}

func extractAddress(data []string, content *FileContent) error {
	return extractItemProcess(data, &content.Reporter.MailingAddress)
}

var extractingReportIssuer = []ExtractData{
	{XmlCompanyStartLine, XmlCompanyEndLine, extractCompanyData},
	{XmlBusinessStartLine, XmlBusinessEndLine, extractBusinessAddress},
	{XmlMailingStartLine, XmlMailingEndLine, extractMailAddress},
	{XmlFormerStartLine, XmlFormerEndLine, extractFormerCompany},
}

func extractReportIssuer(lines []string, data *FileContent) error {
	rol := lines[1 : len(lines)-1]
	processLineByLineExtraction(rol, extractingReportIssuer, data)
	//fmt.Printf("issuer: %v\n", Serialize(data.Issuer))
	return nil
}

func extractCompanyData(data []string, content *FileContent) error {
	return extractItemProcess(data, &content.Issuer.CompanyData)
}

func extractBusinessAddress(data []string, content *FileContent) error {
	return extractItemProcess(data, &content.Issuer.BusinessAddress)
}

func extractMailAddress(data []string, content *FileContent) error {
	return extractItemProcess(data, &content.Issuer.MailAddress)
}

func extractFormerCompany(data []string, content *FileContent) error {
	return extractItemProcess(data, &content.Issuer.FormerCompany)
}

func extractItemProcess(data []string, s interface{}) error {
	rol := data[1 : len(data)-1]
	tags := getStructTag(SecForm4Data, s)
	//fmt.Printf("tags: %v\n", Serialize(tags))
	for _, l := range rol {
		is := strings.Split(l, ">")
		tags[is[0][1:]] = is[1]
	}
	//fmt.Printf("tags: %v\n", Serialize(tags))
	setStructByTag(SecForm4Data, tags, s)
	//fmt.Printf("data: %v\n", Serialize(s))
	return nil

}

func getStructTag(tagName string, s interface{}) (tags map[string]string) {
	tags = make(map[string]string)
	t := reflect.TypeOf(s).Elem()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get(tagName)
		tags[tag] = ""
	}
	return
}

func setStructByTag(tagName string, tags map[string]string, ptr interface{}) {
	v := reflect.ValueOf(ptr)

	t := v.Type().Elem()
	//fmt.Println("Type:", t.Name())
	//fmt.Println("Kind:", t.Kind())
	v = v.Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fv := v.FieldByName(field.Name)
		tag := field.Tag.Get(tagName)
		fv.SetString(tags[tag])
	}
}

func processLineByLineExtraction(lines []string, eaa []ExtractData, data *FileContent) error {
	inTag := false
	var err error
	var buffer []string
	var tagHandler *ExtractData
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if !inTag {
			tagHandler = GetParser(l, eaa)
			if tagHandler != nil {
				inTag = true
				buffer = append(buffer, l)
			}
		} else {
			buffer = append(buffer, l)
			if strings.EqualFold(l, tagHandler.End) {
				err = tagHandler.Proc(buffer, data)
				buffer = nil
				inTag = false
			}
		}
	}
	return err
}

func Serialize(o interface{}) string {
	b, _ := json.Marshal(o)
	return string(b)
}
