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
	OperatingInstitutionSubjectRef string `1.json:"operating_institution_subject_ref"`
	InfoType                       int    `1.json:"info_type"`
}

type BusinessRuleExecutionInfo struct {
	PublishTime            string          `1.json:"publish_time"`
	EffectiveTime          string          `1.json:"effective_time"`
	ReporterNameSubjectRef string          `1.json:"reporter_name_subject_ref"`
	ContactDetails         string          `1.json:"contact_details"`
	ReportTime             string          `1.json:"report_time"`
	ExecutionTime          string          `1.json:"execution_time"`
	BusinessRuleType       int             `1.json:"business_rule_type"`
	FileName               sxgj.FileObject `1.json:"file_name"`
	FileContent            string          `1.json:"file_content"`
	Implementation         string          `1.json:"implementation"`
	Remark                 string          `1.json:"remark"`
}

type PolicyImplementationInfo struct {
	ReleasingOrganizationName       string          `1.json:"releasing_organization_name"`
	ReleasingOrganizationSubjectRef string          `1.json:"releasing_organization_subject_ref"`
	DocumentNumber                  string          `1.json:"document_number"`
	FileType                        string          `1.json:"file_type"`
	FileName                        sxgj.FileObject `1.json:"file_name"`
	PublishTime                     string          `1.json:"publish_time"`
	EffectiveTime                   string          `1.json:"effective_time"`
	Implementation                  string          `1.json:"implementation"`
	ExpiryTime                      string          `1.json:"expiry_time"`
	Remark                          string          `1.json:"remark"`
}

type MajorEventsMeasuresInfo struct {
	ReporterNameSubjectRef        string `1.json:"reporter_name_subject_ref"`
	ContactDetails                string `1.json:"contact_details"`
	MajorEvents                   string `1.json:"major_events"`
	ReportTime                    string `1.json:"report_time"`
	OccurringTime                 string `1.json:"occurring_time"`
	OccurringPlace                string `1.json:"occurring_place"`
	OccurringCause                string `1.json:"occurring_cause"`
	CurrentStatus                 string `1.json:"current_status"`
	PossibleConsequencesInfluence string `1.json:"possible_consequences_influence"`
	TakenMeasures                 string `1.json:"taken_measures"`
	NextMeasures                  string `1.json:"next_measures"`
	Remark                        string `1.json:"remark"`
}

type RegulatoryExecutionInfo struct {
	ReporterNameSubjectRef string          `1.json:"reporter_name_subject_ref"`
	ContactDetails         string          `1.json:"contact_details"`
	ReportTime             string          `1.json:"report_time"`
	RegulatoryMeasuresFile sxgj.FileObject `1.json:"regulatory_measures_file"`
	RegulatorName          string          `1.json:"regulator_name"`
	RegulatorSubjectRef    string          `1.json:"regulator_subject_ref"`
	RegulationType         int             `1.json:"regulation_type"`
	SelfRegulatoryType     int             `1.json:"self_regulatory_type"`
	RegulatedObjectName    string          `1.json:"regulated_object_name"`
	RegulatedObjectRef     string          `1.json:"regulated_object_ref"`
	ViolatingAction        string          `1.json:"violating_action"`
	SanctionContent        string          `1.json:"sanction_content"`
	ExecutionTime          string          `1.json:"execution_time"`
	Implementation         string          `1.json:"implementation"`
	Remark                 string          `1.json:"remark"`
}
