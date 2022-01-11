package 监管

import "A-2021-11-03/sxgj"

type Regulation struct {
	BasicInfo                 BasicInfo
	BusinessRuleExecutionInfo BusinessRuleExecutionInfo
	PolicyImplementationInfo  PolicyImplementationInfo
	MajorEventsMeasuresInfo   MajorEventsMeasuresInfo
	RegulatoryExecutionInfo   RegulatoryExecutionInfo
}

type BasicInfo struct {
	OperatingInstitutionSubjectRef string `json:"operating_institution_subject_ref"`
	InfoType                       int    `json:"info_type"`
}

type BusinessRuleExecutionInfo struct {
	PublishTime            string          `json:"publish_time"`
	EffectiveTime          string          `json:"effective_time"`
	ReporterNameSubjectRef string          `json:"reporter_name_subject_ref"`
	ContactDetails         string          `json:"contact_details"`
	ReportTime             string          `json:"report_time"`
	ExecutionTime          string          `json:"execution_time"`
	BusinessRuleType       int             `json:"business_rule_type"`
	FileName               sxgj.FileObject `json:"file_name"`
	FileContent            string          `json:"file_content"`
	Implementation         string          `json:"implementation"`
	Remark                 string          `json:"remark"`
}

type PolicyImplementationInfo struct {
	ReleasingOrganizationName       string          `json:"releasing_organization_name"`
	ReleasingOrganizationSubjectRef string          `json:"releasing_organization_subject_ref"`
	DocumentNumber                  string          `json:"document_number"`
	FileType                        string          `json:"file_type"`
	FileName                        sxgj.FileObject `json:"file_name"`
	PublishTime                     string          `json:"publish_time"`
	EffectiveTime                   string          `json:"effective_time"`
	Implementation                  string          `json:"implementation"`
	ExpiryTime                      string          `json:"expiry_time"`
	Remark                          string          `json:"remark"`
}

type MajorEventsMeasuresInfo struct {
	ReporterNameSubjectRef        string `json:"reporter_name_subject_ref"`
	ContactDetails                string `json:"contact_details"`
	MajorEvents                   string `json:"major_events"`
	ReportTime                    string `json:"report_time"`
	OccurringTime                 string `json:"occurring_time"`
	OccurringPlace                string `json:"occurring_place"`
	OccurringCause                string `json:"occurring_cause"`
	CurrentStatus                 string `json:"current_status"`
	PossibleConsequencesInfluence string `json:"possible_consequences_influence"`
	TakenMeasures                 string `json:"taken_measures"`
	NextMeasures                  string `json:"next_measures"`
	Remark                        string `json:"remark"`
}

type RegulatoryExecutionInfo struct {
	ReporterNameSubjectRef string          `json:"reporter_name_subject_ref"`
	ContactDetails         string          `json:"contact_details"`
	ReportTime             string          `json:"report_time"`
	RegulatoryMeasuresFile sxgj.FileObject `json:"regulatory_measures_file"`
	RegulatorName          string          `json:"regulator_name"`
	RegulatorSubjectRef    string          `json:"regulator_subject_ref"`
	RegulationType         int             `json:"regulation_type"`
	SelfRegulatoryType     int             `json:"self_regulatory_type"`
	RegulatedObjectName    string          `json:"regulated_object_name"`
	RegulatedObjectRef     string          `json:"regulated_object_ref"`
	ViolatingAction        string          `json:"violating_action"`
	SanctionContent        string          `json:"sanction_content"`
	ExecutionTime          string          `json:"execution_time"`
	Implementation         string          `json:"implementation"`
	Remark                 string          `json:"remark"`
}
