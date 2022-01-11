package 资金结算

import "A-2021-11-03/sxgj"

type CapitalSettlementInformation struct {
	BasicInformationCapitalSettlement BasicInformationCapitalSettlement `json:"basic_information_capital_settlement"`
	TransferorInformation             TransferorInformation             `json:"transferor_information"`
	TransfereeInformation             TransfereeInformation             `json:"transferee_information"`
}

type TransfereeInformation struct {
	SettlementInBankCode                     string  `json:"settlement_in_bank_code"`
	SettlementInBankName                     string  `json:"settlement_in_bank_name"`
	SettlementInBankAccount                  string  `json:"settlement_in_bank_account"`
	SettlementInAccountObjectRef             string  `json:"settlement_in_account_object_ref"`
	SettlementInAccountName                  string  `json:"settlement_in_account_name"`
	SettlementInAccountBalanceBeforeTransfer float64 `json:"settlement_in_account_balance_before_transfer"`
	SettlementInAccountBalanceAfterTransfer  float64 `json:"settlement_in_account_balance_after_transfer"`
}

type TransferorInformation struct {
	SettlementOutBankCode                     string  `json:"settlement_out_bank_code"`
	SettlementOutBankName                     string  `json:"settlement_out_bank_name"`
	SettlementOutBankAccount                  string  `json:"settlement_out_bank_account"`
	SettlementOutAccountObjectRef             string  `json:"settlement_out_account_object_ref"`
	SettlementOutAccountName                  string  `json:"settlement_out_account_name"`
	SettlementOutAccountBalanceBeforeTransfer float64 `json:"settlement_out_account_balance_before_transfer"`
	SettlementOutAccountBalanceAfterTransfer  float64 `json:"settlement_out_account_balance_after_transfer"`
}

type BasicInformationCapitalSettlement struct {
	SettlementSubjectRef                 string          `json:"settlement_subject_ref"`
	SettlementType                       int             `json:"settlement_type"`
	SettlementSerialNumber               string          `json:"settlement_serial_number"`
	SettlementTime                       string          `json:"settlement_time"`
	SettlementProductRef                 string          `json:"settlement_product_ref"`
	SettlementTransactionRef             string          `json:"settlement_transaction_ref"`
	SettlementCurrency                   string          `json:"settlement_currency"`
	SettlementValue                      float64         `json:"settlement_value"`
	SettlementNote                       string          `json:"settlement_note"`
	SettlementOperationDoc               sxgj.FileObject `json:"settlement_operation_doc"`
	SettlementInformationMaintenanceTime string          `json:"settlement_information_maintenance_time"`
}
