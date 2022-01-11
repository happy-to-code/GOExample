package 账号

import (
	"A-2021-11-03/sxgj"
	"time"
)

// AccountInformation 账户信息
type AccountInformation struct {
	BasicAccountInformation   BasicAccountInformation   `json:"basic_account_information"`
	AccountCycleInformation   AccountCycleInformation   `json:"account_cycle_information"`
	AccountRelatedInformation AccountRelatedInformation `json:"account_related_information"`
}

// AccountCycleInformation 账户生命周期信息
type AccountCycleInformation struct {
	AccountOpeningInformation      AccountOpeningInformation      `json:"account_opening_information"`
	AccountCancellationInformation AccountCancellationInformation `json:"account_cancellation_information"`
	FreezeInformation              FreezeInformation              `json:"freeze_information"`
	UnfreezingInformation          UnfreezingInformation          `json:"unfreezing_information"`
}

// BasicAccountInformation 账户基本信息
type BasicAccountInformation struct {
	AccountSubjectRef    string    `json:"account_subject_ref"`
	AccountDepositoryRef string    `json:"account_depository_ref"`
	AccountNumber        string    `json:"account_number"`
	AccountType          int       `json:"account_type"`
	AccountPurpose       []int     `json:"account_purpose"`
	AccountStatus        int       `json:"account_status"`
	AccountCreateTime    time.Time `json:"account_create_time"`
}

// AccountOpeningInformation 开户信息
type AccountOpeningInformation struct {
	AccountEstablishDate             time.Time       `json:"account_establish_date"`
	AccountOpeningDate               time.Time       `json:"account_opening_date"`
	AccountOpeningDoc                sxgj.FileObject `json:"account_opening_doc"`
	AccountOpeningAgentName          string          `json:"account_opening_agent_name"`
	AccountOpeningAgentContactNumber string          `json:"account_opening_agent_contact_number"`
}

// AccountCancellationInformation 销户信息
type AccountCancellationInformation struct {
	AccountClosingDate               string          `json:"account_closing_date"`
	AccountClosingDoc                sxgj.FileObject `json:"account_closing_doc"`
	AccountClosingAgentName          string          `json:"account_closing_agent_name"`
	AccountClosingAgentContactNumber string          `json:"account_closing_agent_contact_number"`
}

// FreezeInformation 冻结信息
type FreezeInformation struct {
	AccountFrozenDate          string          `json:"account_frozen_date"`
	AccountFrozenDoc           sxgj.FileObject `json:"account_frozen_doc"`
	AccountFrozenApplicantName string          `json:"account_frozen_applicant_name"`
	AccountFrozenRemark        string          `json:"account_frozen_remark"`
}

// UnfreezingInformation 解冻信息
type UnfreezingInformation struct {
	AccountThawDate          string          `json:"account_thaw_date"`
	AccountThawDoc           sxgj.FileObject `json:"account_thaw_doc"`
	AccountThawApplicantName string          `json:"account_thaw_applicant_name"`
	AccountThawRemark        string          `json:"account_thaw_remark"`
}

// AccountRelatedInformation 账户关联信息
type AccountRelatedInformation struct {
	AccountAssociation          int    `json:"account_association,omitempty"`
	AccountAssociatedAccountRef string `json:"account_associated_account_ref,omitempty"`
}
