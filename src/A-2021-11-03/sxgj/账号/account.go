package 账号

import (
	"A-2021-11-03/sxgj"
	"time"
)

// AccountInformation 账户信息
type AccountInformation struct {
	BasicAccountInformation   BasicAccountInformation   `1.json:"basic_account_information"`
	AccountCycleInformation   AccountCycleInformation   `1.json:"account_cycle_information"`
	AccountRelatedInformation AccountRelatedInformation `1.json:"account_related_information"`
}

// AccountCycleInformation 账户生命周期信息
type AccountCycleInformation struct {
	AccountOpeningInformation      AccountOpeningInformation      `1.json:"account_opening_information"`
	AccountCancellationInformation AccountCancellationInformation `1.json:"account_cancellation_information"`
	FreezeInformation              FreezeInformation              `1.json:"freeze_information"`
	UnfreezingInformation          UnfreezingInformation          `1.json:"unfreezing_information"`
}

// BasicAccountInformation 账户基本信息
type BasicAccountInformation struct {
	AccountSubjectRef    string    `1.json:"account_subject_ref"`
	AccountDepositoryRef string    `1.json:"account_depository_ref"`
	AccountNumber        string    `1.json:"account_number"`
	AccountType          int       `1.json:"account_type"`
	AccountPurpose       []int     `1.json:"account_purpose"`
	AccountStatus        int       `1.json:"account_status"`
	AccountCreateTime    time.Time `1.json:"account_create_time"`
}

// AccountOpeningInformation 开户信息
type AccountOpeningInformation struct {
	AccountEstablishDate             time.Time       `1.json:"account_establish_date"`
	AccountOpeningDate               time.Time       `1.json:"account_opening_date"`
	AccountOpeningDoc                sxgj.FileObject `1.json:"account_opening_doc"`
	AccountOpeningAgentName          string          `1.json:"account_opening_agent_name"`
	AccountOpeningAgentContactNumber string          `1.json:"account_opening_agent_contact_number"`
}

// AccountCancellationInformation 销户信息
type AccountCancellationInformation struct {
	AccountClosingDate               string          `1.json:"account_closing_date"`
	AccountClosingDoc                sxgj.FileObject `1.json:"account_closing_doc"`
	AccountClosingAgentName          string          `1.json:"account_closing_agent_name"`
	AccountClosingAgentContactNumber string          `1.json:"account_closing_agent_contact_number"`
}

// FreezeInformation 冻结信息
type FreezeInformation struct {
	AccountFrozenDate          string          `1.json:"account_frozen_date"`
	AccountFrozenDoc           sxgj.FileObject `1.json:"account_frozen_doc"`
	AccountFrozenApplicantName string          `1.json:"account_frozen_applicant_name"`
	AccountFrozenRemark        string          `1.json:"account_frozen_remark"`
}

// UnfreezingInformation 解冻信息
type UnfreezingInformation struct {
	AccountThawDate          string          `1.json:"account_thaw_date"`
	AccountThawDoc           sxgj.FileObject `1.json:"account_thaw_doc"`
	AccountThawApplicantName string          `1.json:"account_thaw_applicant_name"`
	AccountThawRemark        string          `1.json:"account_thaw_remark"`
}

// AccountRelatedInformation 账户关联信息
type AccountRelatedInformation struct {
	AccountAssociation          int    `1.json:"account_association,omitempty"`
	AccountAssociatedAccountRef string `1.json:"account_associated_account_ref,omitempty"`
}
