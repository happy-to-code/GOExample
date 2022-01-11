package 财务报表

type FinancialStatement struct {
	BasicInfo                       BasicInfo                       `json:"basic_info"`
	OperatingInstitutionSuitable    OperatingInstitutionSuitable    `json:"operating_institution_suitable"`
	ConsolidatedFinancialStatements ConsolidatedFinancialStatements `json:"consolidated_financial_statements"`
}
type BasicInfo struct {
	ReporterSubjectRef         string `json:"reporter_subject_ref"`
	SubjectMarketRole          int64  `json:"subject_market_role"`
	PeriodStartdate            string `json:"period_startdate"`
	PeriodEnddate              string `json:"period_enddate"`
	PeriodType                 int    `json:"period_type"`
	FormatType                 int    `json:"format_type"`
	AuditInstitution           string `json:"audit_institution"`
	AuditInstitutionSubjectRef string `json:"audit_institution_subject_ref"`
	AuditOpinion               string `json:"audit_opinion"`
}
type OperatingInstitutionSuitable struct {
	IncomeStatementBrief   IncomeStatementBrief   `json:"income_statement_brief"`
	BalanceSheetBrief      BalanceSheetBrief      `json:"balance_sheet_brief"`
	CashFlowStatementBrief CashFlowStatementBrief `json:"cash_flow_statement_brief"`
}
type IncomeStatementBrief struct {
	TotalBusinessRevenue                 float64 `json:"total_business_revenue"`
	MainBusinessRevenue                  float64 `json:"main_business_revenue"`
	IncludingFinancingServiceRevenue     float64 `json:"including_financing_service_revenue"`
	ListingServiceRevenue                float64 `json:"listing_service_revenue"`
	TransferCommissionRevenue            float64 `json:"transfer_commission_revenue"`
	RegistrationDepositoryServiceRevenue float64 `json:"registration_depository_service_revenue"`
	IntermediateServiceRevenue           float64 `json:"intermediate_service_revenue"`
	OtherSalesRevenue                    float64 `json:"other_sales_revenue"`
	IncludingGovernmentAllowance         float64 `json:"including_government_allowance"`
	MembershipFees                       float64 `json:"membership_fees"`
	OtherIncome                          float64 `json:"other_income"`
	OperatingCosts                       float64 `json:"operating_costs"`
	GrossProfit                          float64 `json:"gross_profit"`
	GA                                   float64 `json:"ga"`
	MaterialsExpenses                    float64 `json:"materials_expenses"`
	EmployeeCompensation                 float64 `json:"employee_compensation"`
	DepreciationAndAmortization          float64 `json:"depreciation_and_amortization"`
	ResearchAndDevelopmentCosts          float64 `json:"research_and_development_costs"`
	OtherOperatingExpenses               float64 `json:"other_operating_expenses"`
	OperatingProfit                      float64 `json:"operating_profit"`
	InterestExpenses                     float64 `json:"interest_expenses"`
	InterestIncome                       float64 `json:"interest_income"`
	EquityInvestmentGainOrLoss           float64 `json:"equity_investment_gain_or_loss"`
	GainOrLossOfExtraordinaryItems       float64 `json:"gain_or_loss_of_extraordinary_items"`
	OtherExtraordinaryItems              float64 `json:"other_extraordinary_items"`
	ProfitBeforeExtraordinaryItems       float64 `json:"profit_before_extraordinary_items"`
	EarningsBeforeTax                    float64 `json:"earnings_before_tax"`
	IncomeTax                            float64 `json:"income_tax"`
	MinorityInterestIncome               float64 `json:"minority_interest_income"`
	NPContinuingOperations               float64 `json:"np_continuing_operations"`
	NPDiscontinuedOperations             float64 `json:"np_discontinued_operations"`
	NP                                   float64 `json:"np"`
	OCICommonEquity                      float64 `json:"oci_common_equity"`
	NPCommonEquity                       float64 `json:"np_common_equity"`
}

type BalanceSheetBrief struct {
	Cashandcashequivalents                          float64 `json:"cashandcashequivalents"`
	Tradingfinancialassets                          float64 `json:"tradingfinancialassets"`
	Financialassetsfairvalueaccumulated             float64 `json:"financialassetsfairvalueaccumulated"`
	Financialderivativesassets                      float64 `json:"financialderivativesassets"`
	OthersTinvestment                               float64 `json:"others_tinvestment"`
	AR                                              float64 `json:"ar"`
	ARandnR                                         float64 `json:"a_randn_r"`
	Otherreceivables                                float64 `json:"otherreceivables"`
	inventories                                     float64 `json:"inventories"`
	OthercA                                         float64 `json:"otherc_a"`
	TotalcA                                         float64 `json:"totalc_a"`
	TotalfA                                         float64 `json:"totalf_a"`
	Accumulateddepreciationanddecrease              float64 `json:"accumulateddepreciationanddecrease"`
	NetworthfA                                      float64 `json:"networthf_a"`
	PPE                                             float64 `json:"ppe"`
	Realestateinvestment                            float64 `json:"realestateinvestment"`
	CIP                                             float64 `json:"cip"`
	OtherfA                                         float64 `json:"otherf_a"`
	Equityinvestment                                float64 `json:"equityinvestment"`
	goodwill                                        float64 `json:"goodwill"`
	TotalnCA                                        float64 `json:"totaln_ca"`
	Totalassets                                     float64 `json:"totalassets"`
	APandnP                                         float64 `json:"a_pandn_p"`
	Taxpayable                                      float64 `json:"taxpayable"`
	Tradingfinancialliabilities                     float64 `json:"tradingfinancialliabilities"`
	Tradingfinancialliabilitiesfairvalueaccumulated float64 `json:"tradingfinancialliabilitiesfairvalueaccumulated"`
	Financialderivativesliabilities                 float64 `json:"financialderivativesliabilities"`
	LTloanscurrentperiod                            float64 `json:"l_tloanscurrentperiod"`
	OthercL                                         float64 `json:"otherc_l"`
	TotalcL                                         float64 `json:"totalc_l"`
	LTloans                                         float64 `json:"l_tloans"`
	OthernCL                                        float64 `json:"othern_cl"`
	TotalnCL                                        float64 `json:"totaln_cl"`
	Totalliabilities                                float64 `json:"totalliabilities"`
	Numberofpreferred                               float64 `json:"numberofpreferred"`
	Numberofcommonequity                            float64 `json:"numberofcommonequity"`
	reserve                                         float64 `json:"reserve"`
	Equityinvestmentpremium                         float64 `json:"equityinvestmentpremium"`
	Retainedearnings                                float64 `json:"retainedearnings"`
	OCI                                             float64 `json:"oci"`
	Comprehensiveearningscommonequity               float64 `json:"comprehensiveearningscommonequity"`
	Parentcompanyequity                             float64 `json:"parentcompanyequity"`
	Minorityprofit                                  float64 `json:"minorityprofit"`
	Optionrights                                    float64 `json:"optionrights"`
	Totalliabilitiesandtotalequity                  float64 `json:"totalliabilitiesandtotalequity"`
}

type CashFlowStatementBrief struct {
	Changeinworkingcapital    float64 `json:"changeinworkingcapital"`
	Othernoncashadjustments   float64 `json:"othernoncashadjustments"`
	NetoCF                    float64 `json:"neto_cf"`
	Capitalexpenditures       float64 `json:"capitalexpenditures"`
	CashreceivedfAsale        float64 `json:"cashreceivedf_asale"`
	Inventoryincrease         float64 `json:"inventoryincrease"`
	Investmentdecrease        float64 `json:"investmentdecrease"`
	Otherinvestingcashinflow  float64 `json:"otherinvestingcashinflow"`
	Netinvestingcashflow      float64 `json:"netinvestingcashflow"`
	Debtincrease              float64 `json:"debtincrease"`
	Debtdecrease              float64 `json:"debtdecrease"`
	Sharesincrease            float64 `json:"sharesincrease"`
	Sharesdecrease            float64 `json:"sharesdecrease"`
	Totaldividendpayments     float64 `json:"totaldividendpayments"`
	Othernetfinancingcashflow float64 `json:"othernetfinancingcashflow"`
	Netfinancingcashflow      float64 `json:"netfinancingcashflow"`
}

type ConsolidatedFinancialStatements struct {
	AdjustmentNotes               AdjustmentNotes               `json:"adjustment_notes"`
	ConsolidatedBalanceSheet      ConsolidatedBalanceSheet      `json:"consolidated_balance_sheet"`
	ConsolidatedIncomeStatement   ConsolidatedIncomeStatement   `json:"consolidated_income_statement"`
	ConsolidatedCashFlowStatement ConsolidatedCashFlowStatement `json:"consolidated_cash_flow_statement"`
}
type AdjustmentNotes struct {
	Displaycurrency       float64 `json:"displaycurrency"`
	Originalcurrency      float64 `json:"originalcurrency"`
	Exchangerate          float64 `json:"exchangerate"`
	Exchangeratetype      float64 `json:"exchangeratetype"`
	Taxrate               float64 `json:"taxrate"`
	Taxrateexplanation    float64 `json:"taxrateexplanation"`
	Adjustmentreason      float64 `json:"adjustmentreason"`
	Adjustmentexplanation float64 `json:"adjustmentexplanation"`
	Announcementdate      float64 `json:"announcementdate"`
	Datasource            float64 `json:"datasource"`
}

type ConsolidatedBalanceSheet struct {
	CA                         CA                         `json:"ca"`
	NCA                        NCA                        `json:"nca"`
	CL                         CL                         `json:"cl"`
	NCL                        NCL                        `json:"ncl"`
	OwnersOrShareholdersEquity OwnersOrShareholdersEquity `json:"owners_or_shareholders_equity"`
}
type CA struct {
	Monetarycapital                       float64 `json:"monetarycapital"`
	Tradingfinancialassets                float64 `json:"tradingfinancialassets"`
	Financialderivativeassets             float64 `json:"financialderivativeassets"`
	NRandaR                               float64 `json:"n_randa_r"`
	NR                                    float64 `json:"nr"`
	AR                                    float64 `json:"ar"`
	Accountsreceivablefinancing           float64 `json:"accountsreceivablefinancing"`
	Prepayment                            float64 `json:"prepayment"`
	Totalotherreceivables                 float64 `json:"totalotherreceivables"`
	Dividendreceivable                    float64 `json:"dividendreceivable"`
	Interestreceivable                    float64 `json:"interestreceivable"`
	Otherreceivables                      float64 `json:"otherreceivables"`
	Purchaseofrepurchasingfinancialassets float64 `json:"purchaseofrepurchasingfinancialassets"`
	Inventories                           float64 `json:"inventories"`
	Includingconsumablebiologicalassets   float64 `json:"includingconsumablebiologicalassets"`
	Contractassets                        float64 `json:"contractassets"`
	Disposableassetsholding               float64 `json:"disposableassetsholding"`
	NCAwithinoneyear                      float64 `json:"nc_awithinoneyear"`
	Unamortizedexpenditures               float64 `json:"unamortizedexpenditures"`
	OthercA                               float64 `json:"otherc_a"`
	OtherfinancialcA                      float64 `json:"otherfinancialc_a"`
	DiffcAspecialitems                    float64 `json:"diffc_aspecialitems"`
	DiffcAbalanceitems                    float64 `json:"diffc_abalanceitems"`
	TotalcA                               float64 `json:"totalc_a"`
}

type NCA struct {
	Issuedloanandadvance           float64 `json:"issuedloanandadvance"`
	FVOCI                          float64 `json:"fvoci"`
	AMC                            float64 `json:"amc"`
	Debtinvestments                float64 `json:"debtinvestments"`
	Otherdebtinvestments           float64 `json:"otherdebtinvestments"`
	AFS                            float64 `json:"afs"`
	Otherequityinstrument          float64 `json:"otherequityinstrument"`
	HTM                            float64 `json:"htm"`
	Othernoncurrentfinancialassets float64 `json:"othernoncurrentfinancialassets"`
	LTreceivables                  float64 `json:"l_treceivables"`
	LTequityinvestment             float64 `json:"l_tequityinvestment"`
	Realestateinvestment           float64 `json:"realestateinvestment"`
	TotalfA                        float64 `json:"totalf_a"`
	FA                             float64 `json:"fa"`
	DisposalfA                     float64 `json:"disposalf_a"`
	TotalcIP                       float64 `json:"totalc_ip"`
	Constructioninprocess          float64 `json:"constructioninprocess"`
	Engineeringmaterials           float64 `json:"engineeringmaterials"`
	Regeneratingbiologicalassets   float64 `json:"regeneratingbiologicalassets"`
	Oilandgasreserves              float64 `json:"oilandgasreserves"`
	Rightofuseassets               float64 `json:"rightofuseassets"`
	IA                             float64 `json:"ia"`
	Developmentcosts               float64 `json:"developmentcosts"`
	Goodwill                       float64 `json:"goodwill"`
	LTunamortizedexpenditures      float64 `json:"l_tunamortizedexpenditures"`
	Deferredincometaxasset         float64 `json:"deferredincometaxasset"`
	OthernCA                       float64 `json:"othern_ca"`
	DiffnCAspecialitems            float64 `json:"diffn_c_aspecialitems"`
	DiffnCAbalanceitems            float64 `json:"diffn_c_abalanceitems"`
	TotalnCA                       float64 `json:"totaln_ca"`
	Difftotalassetsspecialitems    float64 `json:"difftotalassetsspecialitems"`
	Difftotalassetsbalanceitems    float64 `json:"difftotalassetsbalanceitems"`
	Totalassets                    float64 `json:"totalassets"`
}

type CL struct {
	STborrowing                        float64 `json:"s_tborrowing"`
	Tradingfinancialliabilities        float64 `json:"tradingfinancialliabilities"`
	Financialderivativeliabilities     float64 `json:"financialderivativeliabilities"`
	APandnP                            float64 `json:"a_pandn_p"`
	NP                                 float64 `json:"np"`
	AP                                 float64 `json:"ap"`
	Unearnedrevenue                    float64 `json:"unearnedrevenue"`
	Liabilitiesofcontracts             float64 `json:"liabilitiesofcontracts"`
	Feesandcommissionpayable           float64 `json:"feesandcommissionpayable"`
	Salariespayable                    float64 `json:"salariespayable"`
	Varioustaxpayables                 float64 `json:"varioustaxpayables"`
	Totalotherpayable                  float64 `json:"totalotherpayable"`
	Interestpayable                    float64 `json:"interestpayable"`
	Dividendpayable                    float64 `json:"dividendpayable"`
	Otherpayables                      float64 `json:"otherpayables"`
	Classifiedasliabilitiesheldforsale float64 `json:"classifiedasliabilitiesheldforsale"`
	NCLwithinoneyear                   float64 `json:"nc_lwithinoneyear"`
	Accruedexpenditures                float64 `json:"accruedexpenditures"`
	DeferredincomecL                   float64 `json:"deferredincomec_l"`
	SLbondpayable                      float64 `json:"s_lbondpayable"`
	OthercL                            float64 `json:"otherc_l"`
	OtherfinancialcL                   float64 `json:"otherfinancialc_l"`
	DiffcLspecialitems                 float64 `json:"diffc_lspecialitems"`
	DiffcLbalanceitems                 float64 `json:"diffc_lbalanceitems"`
	TotalcL                            float64 `json:"totalc_l"`
}

type NCL struct {
	LTborrowing                  float64 `json:"l_tborrowing"`
	Bondpayable                  float64 `json:"bondpayable"`
	Leaseliabilities             float64 `json:"leaseliabilities"`
	TotallTpayables              float64 `json:"totall_tpayables"`
	LTpayable                    float64 `json:"l_tpayable"`
	Specificitempayable          float64 `json:"specificitempayable"`
	LTsalariespayable            float64 `json:"l_tsalariespayable"`
	Accruedliabilities           float64 `json:"accruedliabilities"`
	Deferredincometaxliabilities float64 `json:"deferredincometaxliabilities"`
	DeferredincomenCL            float64 `json:"deferredincomen_cl"`
	OthernCL                     float64 `json:"othern_cl"`
	DiffnCLspecialitems          float64 `json:"diffn_c_lspecialitems"`
	DiffnCLbalanceitems          float64 `json:"diffn_c_lbalanceitems"`
	TotalnCL                     float64 `json:"totaln_cl"`
	Diffliabilitiesspecialitems  float64 `json:"diffliabilitiesspecialitems"`
	Diffliabilitiesbalanceitems  float64 `json:"diffliabilitiesbalanceitems"`
	Totalliabilities             float64 `json:"totalliabilities"`
}

type OwnersOrShareholdersEquity struct {
	Paidincapital                                  float64 `json:"paidincapital"`
	Otherequityinstruments                         float64 `json:"otherequityinstruments"`
	Includingpreferenceshares                      float64 `json:"includingpreferenceshares"`
	Includingperpetualdebt                         float64 `json:"includingperpetualdebt"`
	Capitalreserve                                 float64 `json:"capitalreserve"`
	Lesstreasurystock                              float64 `json:"lesstreasurystock"`
	OCI                                            float64 `json:"oci"`
	Specificreserve                                float64 `json:"specificreserve"`
	Surplusreserve                                 float64 `json:"surplusreserve"`
	Genericriskreserve                             float64 `json:"genericriskreserve"`
	Undistributedprofit                            float64 `json:"undistributedprofit"`
	Converteddifferenceonforeigncurrencystatements float64 `json:"converteddifferenceonforeigncurrencystatements"`
	Unidentifiedinvestmentloss                     float64 `json:"unidentifiedinvestmentloss"`
	Diffshareholdersequityspecialitems             float64 `json:"diffshareholdersequityspecialitems"`
	Diffshareholdersequitybalanceitems             float64 `json:"diffshareholdersequitybalanceitems"`
	Parentcompanyequity                            float64 `json:"parentcompanyequity"`
	Minorityprofit                                 float64 `json:"minorityprofit"`
	Totalownersequity                              float64 `json:"totalownersequity"`
	Diffliabilitiesandownersequityspecialitems     float64 `json:"diffliabilitiesandownersequityspecialitems"`
	Diffliabilitiesandownersequitybalanceitems     float64 `json:"diffliabilitiesandownersequitybalanceitems"`
	Totalliabilitiesandownersequity                float64 `json:"totalliabilitiesandownersequity"`
}

type ConsolidatedIncomeStatement struct {
	Totaloperatingincome                    float64 `json:"totaloperatingincome"`
	Operatingrevenue                        float64 `json:"operatingrevenue"`
	Otherfinancialoperatingrevenue          float64 `json:"otherfinancialoperatingrevenue"`
	Totaloperatingcosts                     float64 `json:"totaloperatingcosts"`
	Operatingcosts                          float64 `json:"operatingcosts"`
	Product                                 float64 `json:"product"`
	Area                                    float64 `json:"area"`
	Taxandassociatedchargeforprimeoperating float64 `json:"taxandassociatedchargeforprimeoperating"`
	Sellingexpenses                         float64 `json:"sellingexpenses"`
	Administrativeexpenses                  float64 `json:"administrativeexpenses"`
	Researchanddevelopmentexpenses          float64 `json:"researchanddevelopmentexpenses"`
	Financialexpense                        float64 `json:"financialexpense"`
	Includinginterestexpense                float64 `json:"includinginterestexpense"`
	Lessinterestincome                      float64 `json:"lessinterestincome"`
	Otherfinancialoperatingcosts            float64 `json:"otherfinancialoperatingcosts"`
	Addotherinvestmentincome                float64 `json:"addotherinvestmentincome"`
	Netinvestmentincome                     float64 `json:"netinvestmentincome"`
	Inclincomefrominvestmenttojointventures float64 `json:"inclincomefrominvestmenttojointventures"`
	GainsfromthederecognitionaMC            float64 `json:"gainsfromthederecognitiona_mc"`
	Netexposurehedginggainsandlosses        float64 `json:"netexposurehedginggainsandlosses"`
	FVGL                                    float64 `json:"fvgl"`
	Impairmentlossesonassets                float64 `json:"impairmentlossesonassets"`
	Impairmentlossesongoodwill              float64 `json:"impairmentlossesongoodwill"`
	Gainondisposalofassets                  float64 `json:"gainondisposalofassets"`
	Exchangnetincome                        float64 `json:"exchangnetincome"`
	Diffoperatingprofitspecialitems         float64 `json:"diffoperatingprofitspecialitems"`
	Diffoperatingprofitbalanceitems         float64 `json:"diffoperatingprofitbalanceitems"`
	Operatingprofit                         float64 `json:"operatingprofit"`
	Addnonoperatingrevenue                  float64 `json:"addnonoperatingrevenue"`
	Lessnonoperatingexpenses                float64 `json:"lessnonoperatingexpenses"`
	IncludingnetlossesondisposalnCA         float64 `json:"includingnetlossesondisposaln_ca"`
	Difftotalprofitspecialitems             float64 `json:"difftotalprofitspecialitems"`
	Difftotalprofitbalanceitems             float64 `json:"difftotalprofitbalanceitems"`
	Totalprofit                             float64 `json:"totalprofit"`
	Lessincometax                           float64 `json:"lessincometax"`
	Addunidentifiedinvestmentloss           float64 `json:"addunidentifiedinvestmentloss"`
	DiffnPspecialitems                      float64 `json:"diffn_pspecialitems"`
	DiffnPbalanceitems                      float64 `json:"diffn_pbalanceitems"`
	NP                                      float64 `json:"np"`
	NPcontinuingoperations                  float64 `json:"n_pcontinuingoperations"`
	NPdiscontinuedoperations                float64 `json:"n_pdiscontinuedoperations"`
	Minorityinterestincome                  float64 `json:"minorityinterestincome"`
	NPparentcompany                         float64 `json:"n_pparentcompany"`
	OCI                                     float64 `json:"oci"`
	TCI                                     float64 `json:"tci"`
	TCIminorityshareholders                 float64 `json:"tc_iminorityshareholders"`
	TCIcommonshareholders                   float64 `json:"tc_icommonshareholders"`
	EPS                                     float64 `json:"eps"`
	BEPS                                    float64 `json:"beps"`
	DEPS                                    float64 `json:"deps"`
}

type ConsolidatedCashFlowStatement struct {
	CFO                      CFO                      `json:"cfo"`
	CFI                      CFI                      `json:"cfi"`
	FinancingCashFlow        FinancingCashFlow        `json:"financing_cash_flow"`
	SupplementaryInformation SupplementaryInformation `json:"supplementary_information"`
}

type CFO struct {
	Cashreceivedfromsalesofgoodsorrenderingservices float64 `json:"cashreceivedfromsalesofgoodsorrenderingservices"`
	Refundsoftaxesandlevies                         float64 `json:"refundsoftaxesandlevies"`
	OthercFOreceived                                float64 `json:"otherc_f_oreceived"`
	CFOin                                           float64 `json:"cf_oin"`
	DiffcFOinspecialitems                           float64 `json:"diffc_f_oinspecialitems"`
	DiffcFOinbalanceitems                           float64 `json:"diffc_f_oinbalanceitems"`
	TotalcFOin                                      float64 `json:"totalc_f_oin"`
	Cashpaidforcommoditiesandservices               float64 `json:"cashpaidforcommoditiesandservices"`
	Cashpaidbehalfofemployees                       float64 `json:"cashpaidbehalfofemployees"`
	Varioustaxespaid                                float64 `json:"varioustaxespaid"`
	OthercFOpaid                                    float64 `json:"otherc_f_opaid"`
	CFOout                                          float64 `json:"cf_oout"`
	DiffcFOoutspecialitems                          float64 `json:"diffc_f_ooutspecialitems"`
	DiffcFOoutbalanceitems                          float64 `json:"diffc_f_ooutbalanceitems"`
	TotalcFOout                                     float64 `json:"totalc_f_oout"`
	DiffcFObalanceitems                             float64 `json:"diffc_f_obalanceitems"`
	CFO                                             float64 `json:"cfo"`
}

type CFI struct {
	Cashreceiptsfrominvestmentswithdrawal                   float64 `json:"cashreceiptsfrominvestmentswithdrawal"`
	Cashreceiptsfromreturnoninvestments                     float64 `json:"cashreceiptsfromreturnoninvestments"`
	NetcashdisposalfAiAlTA                                  float64 `json:"netcashdisposalf_ai_al_ta"`
	Netcashdisposalofsubsidiaryorotherbusinessunit          float64 `json:"netcashdisposalofsubsidiaryorotherbusinessunit"`
	OthercFI                                                float64 `json:"otherc_fi"`
	DiffcFIinspecialitems                                   float64 `json:"diffc_f_iinspecialitems"`
	DiffcFIinbalanceitems                                   float64 `json:"diffc_f_iinbalanceitems"`
	TotalcFI                                                float64 `json:"totalc_fi"`
	CashpurchasefAiAlTA                                     float64 `json:"cashpurchasef_ai_al_ta"`
	Cashpaymentforinvestment                                float64 `json:"cashpaymentforinvestment"`
	Netcashpurchaseofsubsidiariesandotherassociatedentities float64 `json:"netcashpurchaseofsubsidiariesandotherassociatedentities"`
	Otherinvestingcashoutflow                               float64 `json:"otherinvestingcashoutflow"`
	DiffcFIoutspecialitems                                  float64 `json:"diffc_f_ioutspecialitems"`
	DiffcFIoutbalanceitems                                  float64 `json:"diffc_f_ioutbalanceitems"`
	TotalcFIout                                             float64 `json:"totalc_f_iout"`
	DiffcFIbalanceitems                                     float64 `json:"diffc_f_ibalanceitems"`
	CFI                                                     float64 `json:"cfi"`
}

type FinancingCashFlow struct {
	Cashreceivedforattractedinvestment                       float64 `json:"cashreceivedforattractedinvestment"`
	Cashreceivedinvestmentminorityshareholderstosubsidiaries float64 `json:"cashreceivedinvestmentminorityshareholderstosubsidiaries"`
	Cashraisedfromloans                                      float64 `json:"cashraisedfromloans"`
	Cashreceivedfromotherfinancingactivities                 float64 `json:"cashreceivedfromotherfinancingactivities"`
	Cashproceedsofbondissue                                  float64 `json:"cashproceedsofbondissue"`
	DiffcFFinspecialitems                                    float64 `json:"diffc_f_finspecialitems"`
	DiffcFFinbalanceitems                                    float64 `json:"diffc_f_finbalanceitems"`
	TotalcFFin                                               float64 `json:"totalc_f_fin"`
	Cashdebtrepayment                                        float64 `json:"cashdebtrepayment"`
	Cashdistributionofdividendprofitorinterestpayment        float64 `json:"cashdistributionofdividendprofitorinterestpayment"`
	Dividendprofitpaidminorityshareholdersfromsubsidiaries   float64 `json:"dividendprofitpaidminorityshareholdersfromsubsidiaries"`
	Includingotherfinancingcashoutflow                       float64 `json:"includingotherfinancingcashoutflow"`
	DiffcFFoutspecialitems                                   float64 `json:"diffc_f_foutspecialitems"`
	DiffcFFoutbalanceitems                                   float64 `json:"diffc_f_foutbalanceitems"`
	TotalcFFout                                              float64 `json:"totalc_f_fout"`
	DiffcFFbalanceitems                                      float64 `json:"diffc_f_fbalanceitems"`
	CFF                                                      float64 `json:"cff"`
	Effectofforeignexchangeratechangesoncash                 float64 `json:"effectofforeignexchangeratechangesoncash"`
	Netincreasecashandcashequivalentsdirectspecialitems      float64 `json:"netincreasecashandcashequivalentsdirectspecialitems"`
	Netincreasecashandcashequivalentsdirectbalanceitems      float64 `json:"netincreasecashandcashequivalentsdirectbalanceitems"`
	Netincreaseofcashandcashequivalents                      float64 `json:"netincreaseofcashandcashequivalents"`
	Balanceofcashandcashequivalentsbeginningoftheperiod      float64 `json:"balanceofcashandcashequivalentsbeginningoftheperiod"`
	Balanceofcashandcashequivalentsendoftheperiod            float64 `json:"balanceofcashandcashequivalentsendoftheperiod"`
}

type SupplementaryInformation struct {
	NP                                        float64 `json:"np"`
	Depreciationreserveofassets               float64 `json:"depreciationreserveofassets"`
	Impairmentlossesongoodwill                float64 `json:"impairmentlossesongoodwill"`
	Depreciationcharge                        float64 `json:"depreciationcharge"`
	Amortizationofintangibleassets            float64 `json:"amortizationofintangibleassets"`
	AmortisationcostrOUassets                 float64 `json:"amortisationcostr_o_uassets"`
	AmortizationoflTdeferredexpenses          float64 `json:"amortizationofl_tdeferredexpenses"`
	Decreaseindeferredexpenses                float64 `json:"decreaseindeferredexpenses"`
	Increaseinaccruedexpenses                 float64 `json:"increaseinaccruedexpenses"`
	LossesondisposalfAiAlTA                   float64 `json:"lossesondisposalf_ai_al_ta"`
	LossesondiscardingfA                      float64 `json:"lossesondiscardingf_a"`
	Lossesonfairvaluechange                   float64 `json:"lossesonfairvaluechange"`
	Financialexpense                          float64 `json:"financialexpense"`
	Investmentlosses                          float64 `json:"investmentlosses"`
	Decreaseofdeferredincometaxasset          float64 `json:"decreaseofdeferredincometaxasset"`
	Increaseofdeferredincometaxliabilities    float64 `json:"increaseofdeferredincometaxliabilities"`
	Decreaseininventories                     float64 `json:"decreaseininventories"`
	Decreaseofoperatingreceivables            float64 `json:"decreaseofoperatingreceivables"`
	Increaseofoperatingpayables               float64 `json:"increaseofoperatingpayables"`
	Unidentifiedinvestmentloss                float64 `json:"unidentifiedinvestmentloss"`
	Otheritem                                 float64 `json:"otheritem"`
	DiffcFOindirectspecialitems               float64 `json:"diffc_f_oindirectspecialitems"`
	DiffcFOindirectbalanceitems               float64 `json:"diffc_f_oindirectbalanceitems"`
	CFOindirectbalanceitems                   float64 `json:"cf_oindirectbalanceitems"`
	Debtconvertedtosharecapital               float64 `json:"debtconvertedtosharecapital"`
	Convertiblebondsmaturewithinoneyear       float64 `json:"convertiblebondsmaturewithinoneyear"`
	FAacquiredunderfinanceleases              float64 `json:"f_aacquiredunderfinanceleases"`
	Cashendoftheperiod                        float64 `json:"cashendoftheperiod"`
	Cashbeginningoftheperiod                  float64 `json:"cashbeginningoftheperiod"`
	Cashequivalentsendoftheperiod             float64 `json:"cashequivalentsendoftheperiod"`
	Cashequivalentsbeginningoftheperiod       float64 `json:"cashequivalentsbeginningoftheperiod"`
	Diffnetincreasecashindirectspecialitems   float64 `json:"diffnetincreasecashindirectspecialitems"`
	Diffnetincreasecashindirectbalanceitems   float64 `json:"diffnetincreasecashindirectbalanceitems"`
	Netincreasecashandcashequivalentsindirect float64 `json:"netincreasecashandcashequivalentsindirect"`
}
