package form4

import "encoding/xml"

const (
	SecForm4Tag  = "sec-form4-tag"
	SecForm4Data = "sec-form4-data"
)

type FileContent struct {
	Document DocumentInfo
	Reporter ReportingOwnerInfo
	Issuer   IssuerInfo
	Xml      XmlBody
}

const (
	XmlReportingOwnerStartLine = "<REPORTING-OWNER>"
	XmlReportingOwnerEndLine   = "</REPORTING-OWNER>"

	XmlIssuerStartLine = "<ISSUER>"
	XmlIssuerEndLine   = "</ISSUER>"

	XmlStartLine = "<XML>"
	XmlEndLine   = "</XML>"
)

type DocumentInfo struct {
	AccessionNumber     string `sec-form4-data:"ACCESSION-NUMBER"`
	Type                string `sec-form4-data:"TYPE"`
	PublicDocumentCount string `sec-form4-data:"PUBLIC-DOCUMENT-COUNT"`
	Period              string `sec-form4-data:"PERIOD"`
	FilingDate          string `sec-form4-data:"FILING-DATE"`
	DateChange          string `sec-form4-data:"DATE-OF-FILING-DATE-CHANGE"`
}

const (
	XmlReportingOwnerInfoStartLine = "<OWNER-DATA>"
	XmlReportingOwnerInfoEndLine   = "</OWNER-DATA>"

	XmlReportingFilingInfoStartLine = "<FILING-VALUES>"
	XmlReportingFilingInfoEndLine   = "</FILING-VALUES>"

	XmlReportingMailingInfoStartLine = "<MAIL-ADDRESS>"
	XmlReportingMailingInfoEndLine   = "</MAIL-ADDRESS>"
)

type ReportingOwnerInfo struct {
	OwnerData      OwnerDataInfo
	FilingValues   FilingValuesInfo
	MailingAddress AddressInfo
}

type OwnerDataInfo struct {
	ConformedName string `sec-form4-data:"CONFORMED-NAME"`
	Cik           string `sec-form4-data:"CIK"`
}

type FilingValuesInfo struct {
	FormType   string `sec-form4-data:"FORM-TYPE"`
	Act        string `sec-form4-data:"ACT"`
	FileNumber string `sec-form4-data:"FILE-NUMBER"`
	FilmNumber string `sec-form4-data:"FILM-NUMBER"`
}

type AddressInfo struct {
	Street string `sec-form4-data:"STREET1"`
	City   string `sec-form4-data:"CITY"`
	State  string `sec-form4-data:"STATE"`
	Zip    string `sec-form4-data:"ZIP"`
	Phone  string `sec-form4-data:"PHONE"`
}

const (
	XmlCompanyStartLine = "<COMPANY-DATA>"
	XmlCompanyEndLine   = "</COMPANY-DATA>"

	XmlBusinessStartLine = "<BUSINESS-ADDRESS>"
	XmlBusinessEndLine   = "</BUSINESS-ADDRESS>"

	XmlMailingStartLine = "<MAIL-ADDRESS>"
	XmlMailingEndLine   = "</MAIL-ADDRESS>"

	XmlFormerStartLine = "<FORMER-COMPANY>"
	XmlFormerEndLine   = "</FORMER-COMPANY>"
)

type IssuerInfo struct {
	CompanyData     CompanyDataInfo
	BusinessAddress AddressInfo
	MailAddress     AddressInfo
	FormerCompany   FormerCompanyInfo
}

type CompanyDataInfo struct {
	ConformedName        string `sec-form4-data:"CONFORMED-NAME"`
	Cik                  string `sec-form4-data:"CIK"`
	AssignedSic          string `sec-form4-data:"ASSIGNED-SIC"`
	IrsNumber            string `sec-form4-data:"IRS-NUMBER"`
	StateOfIncorporation string `sec-form4-data:"STATE-OF-INCORPORATION"`
	FiscalYearEnd        string `sec-form4-data:"FISCAL-YEAR-END"`
}

type FormerCompanyInfo struct {
	FormerConformedName string `sec-form4-data:"FORMER-CONFORMED-NAME"`
	DateChanged         string `sec-form4-data:"DATE-CHANGED"`
}

type XmlBody struct {
	XMLName           xml.Name `xml:"XML"`
	Text              string   `xml:",chardata"`
	OwnershipDocument struct {
		Text                  string `xml:",chardata"`
		SchemaVersion         string `xml:"schemaVersion"`
		DocumentType          string `xml:"documentType"`
		PeriodOfReport        string `xml:"periodOfReport"`
		NotSubjectToSection16 string `xml:"notSubjectToSection16"`
		Issuer                struct {
			Text                string `xml:",chardata"`
			IssuerCik           string `xml:"issuerCik"`
			IssuerName          string `xml:"issuerName"`
			IssuerTradingSymbol string `xml:"issuerTradingSymbol"`
		} `xml:"issuer"`
		ReportingOwner struct {
			Text             string `xml:",chardata"`
			ReportingOwnerId struct {
				Text         string `xml:",chardata"`
				RptOwnerCik  string `xml:"rptOwnerCik"`
				RptOwnerName string `xml:"rptOwnerName"`
			} `xml:"reportingOwnerId"`
			ReportingOwnerAddress struct {
				Text                     string `xml:",chardata"`
				RptOwnerStreet1          string `xml:"rptOwnerStreet1"`
				RptOwnerStreet2          string `xml:"rptOwnerStreet2"`
				RptOwnerCity             string `xml:"rptOwnerCity"`
				RptOwnerState            string `xml:"rptOwnerState"`
				RptOwnerZipCode          string `xml:"rptOwnerZipCode"`
				RptOwnerStateDescription string `xml:"rptOwnerStateDescription"`
			} `xml:"reportingOwnerAddress"`
			ReportingOwnerRelationship struct {
				Text              string `xml:",chardata"`
				IsDirector        string `xml:"isDirector"`
				IsOfficer         string `xml:"isOfficer"`
				IsTenPercentOwner string `xml:"isTenPercentOwner"`
				IsOther           string `xml:"isOther"`
				OfficerTitle      string `xml:"officerTitle"`
				OtherText         string `xml:"otherText"`
			} `xml:"reportingOwnerRelationship"`
		} `xml:"reportingOwner"`
		NonDerivativeTable struct {
			Text                 string `xml:",chardata"`
			NonDerivativeHolding []struct {
				Text          string `xml:",chardata"`
				SecurityTitle struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"securityTitle"`
				PostTransactionAmounts struct {
					Text                            string `xml:",chardata"`
					SharesOwnedFollowingTransaction struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value"`
					} `xml:"sharesOwnedFollowingTransaction"`
				} `xml:"postTransactionAmounts"`
				OwnershipNature struct {
					Text                      string `xml:",chardata"`
					DirectOrIndirectOwnership struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value"`
					} `xml:"directOrIndirectOwnership"`
					NatureOfOwnership struct {
						Text       string `xml:",chardata"`
						Value      string `xml:"value"`
						FootnoteId struct {
							Text string `xml:",chardata"`
							ID   string `xml:"id,attr"`
						} `xml:"footnoteId"`
					} `xml:"natureOfOwnership"`
				} `xml:"ownershipNature"`
			} `xml:"nonDerivativeHolding"`
			NonDerivativeTransaction struct {
				Text          string `xml:",chardata"`
				SecurityTitle struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"securityTitle"`
				TransactionDate struct {
					Text  string `xml:",chardata"`
					Value string `xml:"value"`
				} `xml:"transactionDate"`
				TransactionCoding struct {
					Text                string `xml:",chardata"`
					TransactionFormType string `xml:"transactionFormType"`
					TransactionCode     string `xml:"transactionCode"`
					EquitySwapInvolved  string `xml:"equitySwapInvolved"`
				} `xml:"transactionCoding"`
				TransactionAmounts struct {
					Text              string `xml:",chardata"`
					TransactionShares struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value"`
					} `xml:"transactionShares"`
					TransactionPricePerShare struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value"`
					} `xml:"transactionPricePerShare"`
					TransactionAcquiredDisposedCode struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value"`
					} `xml:"transactionAcquiredDisposedCode"`
				} `xml:"transactionAmounts"`
				PostTransactionAmounts struct {
					Text                            string `xml:",chardata"`
					SharesOwnedFollowingTransaction struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value"`
					} `xml:"sharesOwnedFollowingTransaction"`
				} `xml:"postTransactionAmounts"`
				OwnershipNature struct {
					Text                      string `xml:",chardata"`
					DirectOrIndirectOwnership struct {
						Text  string `xml:",chardata"`
						Value string `xml:"value"`
					} `xml:"directOrIndirectOwnership"`
				} `xml:"ownershipNature"`
			} `xml:"nonDerivativeTransaction"`
		} `xml:"nonDerivativeTable"`
		DerivativeTable string `xml:"derivativeTable"`
		Footnotes       struct {
			Text     string `xml:",chardata"`
			Footnote []struct {
				Text string `xml:",chardata"`
				ID   string `xml:"id,attr"`
			} `xml:"footnote"`
		} `xml:"footnotes"`
		Remarks        string `xml:"remarks"`
		OwnerSignature struct {
			Text          string `xml:",chardata"`
			SignatureName string `xml:"signatureName"`
			SignatureDate string `xml:"signatureDate"`
		} `xml:"ownerSignature"`
	} `xml:"ownershipDocument"`
}
