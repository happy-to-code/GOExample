package 资金结算

import "A-2021-11-03/sxgj"

type CapitalSettlementInformation struct {
	BasicInformationCapitalSettlement BasicInformationCapitalSettlement `1.json:"basic_information_capital_settlement"`
	TransferorInformation             TransferorInformation             `1.json:"transferor_information"`
	TransfereeInformation             TransfereeInformation             `1.json:"transferee_information"`
}

type TransfereeInformation struct {
	SettlementInBankCode                     string  `1.json:"settlement_in_bank_code"`
	SettlementInBankName                     string  `1.json:"settlement_in_bank_name"`
	SettlementInBankAccount                  string  `1.json:"settlement_in_bank_account"`
	SettlementInAccountObjectRef             string  `1.json:"settlement_in_account_object_ref"`
	SettlementInAccountName                  string  `1.json:"settlement_in_account_name"`
	SettlementInAccountBalanceBeforeTransfer float64 `1.json:"settlement_in_account_balance_before_transfer"`
	SettlementInAccountBalanceAfterTransfer  float64 `1.json:"settlement_in_account_balance_after_transfer"`
}

type TransferorInformation struct {
	SettlementOutBankCode                     string  `1.json:"settlement_out_bank_code"`
	SettlementOutBankName                     string  `1.json:"settlement_out_bank_name"`
	SettlementOutBankAccount                  string  `1.json:"settlement_out_bank_account"`
	SettlementOutAccountObjectRef             string  `1.json:"settlement_out_account_object_ref"`
	SettlementOutAccountName                  string  `1.json:"settlement_out_account_name"`
	SettlementOutAccountBalanceBeforeTransfer float64 `1.json:"settlement_out_account_balance_before_transfer"`
	SettlementOutAccountBalanceAfterTransfer  float64 `1.json:"settlement_out_account_balance_after_transfer"`
}

type BasicInformationCapitalSettlement struct {
	SettlementSubjectRef                 string          `1.json:"settlement_subject_ref"`
	SettlementType                       int             `1.json:"settlement_type"`
	SettlementSerialNumber               string          `1.json:"settlement_serial_number"`
	SettlementTime                       string          `1.json:"settlement_time"`
	SettlementProductRef                 string          `1.json:"settlement_product_ref"`
	SettlementTransactionRef             string          `1.json:"settlement_transaction_ref"`
	SettlementCurrency                   string          `1.json:"settlement_currency"`
	SettlementValue                      float64         `1.json:"settlement_value"`
	SettlementNote                       string          `1.json:"settlement_note"`
	SettlementOperationDoc               sxgj.FileObject `1.json:"settlement_operation_doc"`
	SettlementInformationMaintenanceTime string          `1.json:"settlement_information_maintenance_time"`
}
