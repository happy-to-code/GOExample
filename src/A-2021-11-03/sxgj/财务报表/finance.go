package 财务报表

type FinancialStatement struct {
	BasicInfo                       BasicInfo                       `1.json:"basic_info"`
	OperatingInstitutionSuitable    OperatingInstitutionSuitable    `1.json:"operating_institution_suitable"`
	ConsolidatedFinancialStatements ConsolidatedFinancialStatements `1.json:"consolidated_financial_statements"`
}
type BasicInfo struct {
	ReporterSubjectRef         string `1.json:"reporter_subject_ref"`
	SubjectMarketRole          int64  `1.json:"subject_market_role"`
	PeriodStartdate            string `1.json:"period_startdate"`
	PeriodEnddate              string `1.json:"period_enddate"`
	PeriodType                 int    `1.json:"period_type"`
	FormatType                 int    `1.json:"format_type"`
	AuditInstitution           string `1.json:"audit_institution"`
	AuditInstitutionSubjectRef string `1.json:"audit_institution_subject_ref"`
	AuditOpinion               string `1.json:"audit_opinion"`
}
type OperatingInstitutionSuitable struct {
	IncomeStatementBrief   IncomeStatementBrief   `1.json:"income_statement_brief"`
	BalanceSheetBrief      BalanceSheetBrief      `1.json:"balance_sheet_brief"`
	CashFlowStatementBrief CashFlowStatementBrief `1.json:"cash_flow_statement_brief"`
}
type IncomeStatementBrief struct {
	TotalBusinessRevenue                 float64 `1.json:"total_business_revenue"`
	MainBusinessRevenue                  float64 `1.json:"main_business_revenue"`
	IncludingFinancingServiceRevenue     float64 `1.json:"including_financing_service_revenue"`
	ListingServiceRevenue                float64 `1.json:"listing_service_revenue"`
	TransferCommissionRevenue            float64 `1.json:"transfer_commission_revenue"`
	RegistrationDepositoryServiceRevenue float64 `1.json:"registration_depository_service_revenue"`
	IntermediateServiceRevenue           float64 `1.json:"intermediate_service_revenue"`
	OtherSalesRevenue                    float64 `1.json:"other_sales_revenue"`
	IncludingGovernmentAllowance         float64 `1.json:"including_government_allowance"`
	MembershipFees                       float64 `1.json:"membership_fees"`
	OtherIncome                          float64 `1.json:"other_income"`
	OperatingCosts                       float64 `1.json:"operating_costs"`
	GrossProfit                          float64 `1.json:"gross_profit"`
	GA                                   float64 `1.json:"ga"`
	MaterialsExpenses                    float64 `1.json:"materials_expenses"`
	EmployeeCompensation                 float64 `1.json:"employee_compensation"`
	DepreciationAndAmortization          float64 `1.json:"depreciation_and_amortization"`
	ResearchAndDevelopmentCosts          float64 `1.json:"research_and_development_costs"`
	OtherOperatingExpenses               float64 `1.json:"other_operating_expenses"`
	OperatingProfit                      float64 `1.json:"operating_profit"`
	InterestExpenses                     float64 `1.json:"interest_expenses"`
	InterestIncome                       float64 `1.json:"interest_income"`
	EquityInvestmentGainOrLoss           float64 `1.json:"equity_investment_gain_or_loss"`
	GainOrLossOfExtraordinaryItems       float64 `1.json:"gain_or_loss_of_extraordinary_items"`
	OtherExtraordinaryItems              float64 `1.json:"other_extraordinary_items"`
	ProfitBeforeExtraordinaryItems       float64 `1.json:"profit_before_extraordinary_items"`
	EarningsBeforeTax                    float64 `1.json:"earnings_before_tax"`
	IncomeTax                            float64 `1.json:"income_tax"`
	MinorityInterestIncome               float64 `1.json:"minority_interest_income"`
	NPContinuingOperations               float64 `1.json:"np_continuing_operations"`
	NPDiscontinuedOperations             float64 `1.json:"np_discontinued_operations"`
	NP                                   float64 `1.json:"np"`
	OCICommonEquity                      float64 `1.json:"oci_common_equity"`
	NPCommonEquity                       float64 `1.json:"np_common_equity"`
}

type BalanceSheetBrief struct {
	Cashandcashequivalents                          float64 `1.json:"cashandcashequivalents"`
	Tradingfinancialassets                          float64 `1.json:"tradingfinancialassets"`
	Financialassetsfairvalueaccumulated             float64 `1.json:"financialassetsfairvalueaccumulated"`
	Financialderivativesassets                      float64 `1.json:"financialderivativesassets"`
	OthersTinvestment                               float64 `1.json:"others_tinvestment"`
	AR                                              float64 `1.json:"ar"`
	ARandnR                                         float64 `1.json:"a_randn_r"`
	Otherreceivables                                float64 `1.json:"otherreceivables"`
	inventories                                     float64 `1.json:"inventories"`
	OthercA                                         float64 `1.json:"otherc_a"`
	TotalcA                                         float64 `1.json:"totalc_a"`
	TotalfA                                         float64 `1.json:"totalf_a"`
	Accumulateddepreciationanddecrease              float64 `1.json:"accumulateddepreciationanddecrease"`
	NetworthfA                                      float64 `1.json:"networthf_a"`
	PPE                                             float64 `1.json:"ppe"`
	Realestateinvestment                            float64 `1.json:"realestateinvestment"`
	CIP                                             float64 `1.json:"cip"`
	OtherfA                                         float64 `1.json:"otherf_a"`
	Equityinvestment                                float64 `1.json:"equityinvestment"`
	goodwill                                        float64 `1.json:"goodwill"`
	TotalnCA                                        float64 `1.json:"totaln_ca"`
	Totalassets                                     float64 `1.json:"totalassets"`
	APandnP                                         float64 `1.json:"a_pandn_p"`
	Taxpayable                                      float64 `1.json:"taxpayable"`
	Tradingfinancialliabilities                     float64 `1.json:"tradingfinancialliabilities"`
	Tradingfinancialliabilitiesfairvalueaccumulated float64 `1.json:"tradingfinancialliabilitiesfairvalueaccumulated"`
	Financialderivativesliabilities                 float64 `1.json:"financialderivativesliabilities"`
	LTloanscurrentperiod                            float64 `1.json:"l_tloanscurrentperiod"`
	OthercL                                         float64 `1.json:"otherc_l"`
	TotalcL                                         float64 `1.json:"totalc_l"`
	LTloans                                         float64 `1.json:"l_tloans"`
	OthernCL                                        float64 `1.json:"othern_cl"`
	TotalnCL                                        float64 `1.json:"totaln_cl"`
	Totalliabilities                                float64 `1.json:"totalliabilities"`
	Numberofpreferred                               float64 `1.json:"numberofpreferred"`
	Numberofcommonequity                            float64 `1.json:"numberofcommonequity"`
	reserve                                         float64 `1.json:"reserve"`
	Equityinvestmentpremium                         float64 `1.json:"equityinvestmentpremium"`
	Retainedearnings                                float64 `1.json:"retainedearnings"`
	OCI                                             float64 `1.json:"oci"`
	Comprehensiveearningscommonequity               float64 `1.json:"comprehensiveearningscommonequity"`
	Parentcompanyequity                             float64 `1.json:"parentcompanyequity"`
	Minorityprofit                                  float64 `1.json:"minorityprofit"`
	Optionrights                                    float64 `1.json:"optionrights"`
	Totalliabilitiesandtotalequity                  float64 `1.json:"totalliabilitiesandtotalequity"`
}

type CashFlowStatementBrief struct {
	Changeinworkingcapital    float64 `1.json:"changeinworkingcapital"`
	Othernoncashadjustments   float64 `1.json:"othernoncashadjustments"`
	NetoCF                    float64 `1.json:"neto_cf"`
	Capitalexpenditures       float64 `1.json:"capitalexpenditures"`
	CashreceivedfAsale        float64 `1.json:"cashreceivedf_asale"`
	Inventoryincrease         float64 `1.json:"inventoryincrease"`
	Investmentdecrease        float64 `1.json:"investmentdecrease"`
	Otherinvestingcashinflow  float64 `1.json:"otherinvestingcashinflow"`
	Netinvestingcashflow      float64 `1.json:"netinvestingcashflow"`
	Debtincrease              float64 `1.json:"debtincrease"`
	Debtdecrease              float64 `1.json:"debtdecrease"`
	Sharesincrease            float64 `1.json:"sharesincrease"`
	Sharesdecrease            float64 `1.json:"sharesdecrease"`
	Totaldividendpayments     float64 `1.json:"totaldividendpayments"`
	Othernetfinancingcashflow float64 `1.json:"othernetfinancingcashflow"`
	Netfinancingcashflow      float64 `1.json:"netfinancingcashflow"`
}

type ConsolidatedFinancialStatements struct {
	AdjustmentNotes               AdjustmentNotes               `1.json:"adjustment_notes"`
	ConsolidatedBalanceSheet      ConsolidatedBalanceSheet      `1.json:"consolidated_balance_sheet"`
	ConsolidatedIncomeStatement   ConsolidatedIncomeStatement   `1.json:"consolidated_income_statement"`
	ConsolidatedCashFlowStatement ConsolidatedCashFlowStatement `1.json:"consolidated_cash_flow_statement"`
}
type AdjustmentNotes struct {
	Displaycurrency       float64 `1.json:"displaycurrency"`
	Originalcurrency      float64 `1.json:"originalcurrency"`
	Exchangerate          float64 `1.json:"exchangerate"`
	Exchangeratetype      float64 `1.json:"exchangeratetype"`
	Taxrate               float64 `1.json:"taxrate"`
	Taxrateexplanation    float64 `1.json:"taxrateexplanation"`
	Adjustmentreason      float64 `1.json:"adjustmentreason"`
	Adjustmentexplanation float64 `1.json:"adjustmentexplanation"`
	Announcementdate      float64 `1.json:"announcementdate"`
	Datasource            float64 `1.json:"datasource"`
}

type ConsolidatedBalanceSheet struct {
	CA                         CA                         `1.json:"ca"`
	NCA                        NCA                        `1.json:"nca"`
	CL                         CL                         `1.json:"cl"`
	NCL                        NCL                        `1.json:"ncl"`
	OwnersOrShareholdersEquity OwnersOrShareholdersEquity `1.json:"owners_or_shareholders_equity"`
}
type CA struct {
	Monetarycapital                       float64 `1.json:"monetarycapital"`
	Tradingfinancialassets                float64 `1.json:"tradingfinancialassets"`
	Financialderivativeassets             float64 `1.json:"financialderivativeassets"`
	NRandaR                               float64 `1.json:"n_randa_r"`
	NR                                    float64 `1.json:"nr"`
	AR                                    float64 `1.json:"ar"`
	Accountsreceivablefinancing           float64 `1.json:"accountsreceivablefinancing"`
	Prepayment                            float64 `1.json:"prepayment"`
	Totalotherreceivables                 float64 `1.json:"totalotherreceivables"`
	Dividendreceivable                    float64 `1.json:"dividendreceivable"`
	Interestreceivable                    float64 `1.json:"interestreceivable"`
	Otherreceivables                      float64 `1.json:"otherreceivables"`
	Purchaseofrepurchasingfinancialassets float64 `1.json:"purchaseofrepurchasingfinancialassets"`
	Inventories                           float64 `1.json:"inventories"`
	Includingconsumablebiologicalassets   float64 `1.json:"includingconsumablebiologicalassets"`
	Contractassets                        float64 `1.json:"contractassets"`
	Disposableassetsholding               float64 `1.json:"disposableassetsholding"`
	NCAwithinoneyear                      float64 `1.json:"nc_awithinoneyear"`
	Unamortizedexpenditures               float64 `1.json:"unamortizedexpenditures"`
	OthercA                               float64 `1.json:"otherc_a"`
	OtherfinancialcA                      float64 `1.json:"otherfinancialc_a"`
	DiffcAspecialitems                    float64 `1.json:"diffc_aspecialitems"`
	DiffcAbalanceitems                    float64 `1.json:"diffc_abalanceitems"`
	TotalcA                               float64 `1.json:"totalc_a"`
}

type NCA struct {
	Issuedloanandadvance           float64 `1.json:"issuedloanandadvance"`
	FVOCI                          float64 `1.json:"fvoci"`
	AMC                            float64 `1.json:"amc"`
	Debtinvestments                float64 `1.json:"debtinvestments"`
	Otherdebtinvestments           float64 `1.json:"otherdebtinvestments"`
	AFS                            float64 `1.json:"afs"`
	Otherequityinstrument          float64 `1.json:"otherequityinstrument"`
	HTM                            float64 `1.json:"htm"`
	Othernoncurrentfinancialassets float64 `1.json:"othernoncurrentfinancialassets"`
	LTreceivables                  float64 `1.json:"l_treceivables"`
	LTequityinvestment             float64 `1.json:"l_tequityinvestment"`
	Realestateinvestment           float64 `1.json:"realestateinvestment"`
	TotalfA                        float64 `1.json:"totalf_a"`
	FA                             float64 `1.json:"fa"`
	DisposalfA                     float64 `1.json:"disposalf_a"`
	TotalcIP                       float64 `1.json:"totalc_ip"`
	Constructioninprocess          float64 `1.json:"constructioninprocess"`
	Engineeringmaterials           float64 `1.json:"engineeringmaterials"`
	Regeneratingbiologicalassets   float64 `1.json:"regeneratingbiologicalassets"`
	Oilandgasreserves              float64 `1.json:"oilandgasreserves"`
	Rightofuseassets               float64 `1.json:"rightofuseassets"`
	IA                             float64 `1.json:"ia"`
	Developmentcosts               float64 `1.json:"developmentcosts"`
	Goodwill                       float64 `1.json:"goodwill"`
	LTunamortizedexpenditures      float64 `1.json:"l_tunamortizedexpenditures"`
	Deferredincometaxasset         float64 `1.json:"deferredincometaxasset"`
	OthernCA                       float64 `1.json:"othern_ca"`
	DiffnCAspecialitems            float64 `1.json:"diffn_c_aspecialitems"`
	DiffnCAbalanceitems            float64 `1.json:"diffn_c_abalanceitems"`
	TotalnCA                       float64 `1.json:"totaln_ca"`
	Difftotalassetsspecialitems    float64 `1.json:"difftotalassetsspecialitems"`
	Difftotalassetsbalanceitems    float64 `1.json:"difftotalassetsbalanceitems"`
	Totalassets                    float64 `1.json:"totalassets"`
}

type CL struct {
	STborrowing                        float64 `1.json:"s_tborrowing"`
	Tradingfinancialliabilities        float64 `1.json:"tradingfinancialliabilities"`
	Financialderivativeliabilities     float64 `1.json:"financialderivativeliabilities"`
	APandnP                            float64 `1.json:"a_pandn_p"`
	NP                                 float64 `1.json:"np"`
	AP                                 float64 `1.json:"ap"`
	Unearnedrevenue                    float64 `1.json:"unearnedrevenue"`
	Liabilitiesofcontracts             float64 `1.json:"liabilitiesofcontracts"`
	Feesandcommissionpayable           float64 `1.json:"feesandcommissionpayable"`
	Salariespayable                    float64 `1.json:"salariespayable"`
	Varioustaxpayables                 float64 `1.json:"varioustaxpayables"`
	Totalotherpayable                  float64 `1.json:"totalotherpayable"`
	Interestpayable                    float64 `1.json:"interestpayable"`
	Dividendpayable                    float64 `1.json:"dividendpayable"`
	Otherpayables                      float64 `1.json:"otherpayables"`
	Classifiedasliabilitiesheldforsale float64 `1.json:"classifiedasliabilitiesheldforsale"`
	NCLwithinoneyear                   float64 `1.json:"nc_lwithinoneyear"`
	Accruedexpenditures                float64 `1.json:"accruedexpenditures"`
	DeferredincomecL                   float64 `1.json:"deferredincomec_l"`
	SLbondpayable                      float64 `1.json:"s_lbondpayable"`
	OthercL                            float64 `1.json:"otherc_l"`
	OtherfinancialcL                   float64 `1.json:"otherfinancialc_l"`
	DiffcLspecialitems                 float64 `1.json:"diffc_lspecialitems"`
	DiffcLbalanceitems                 float64 `1.json:"diffc_lbalanceitems"`
	TotalcL                            float64 `1.json:"totalc_l"`
}

type NCL struct {
	LTborrowing                  float64 `1.json:"l_tborrowing"`
	Bondpayable                  float64 `1.json:"bondpayable"`
	Leaseliabilities             float64 `1.json:"leaseliabilities"`
	TotallTpayables              float64 `1.json:"totall_tpayables"`
	LTpayable                    float64 `1.json:"l_tpayable"`
	Specificitempayable          float64 `1.json:"specificitempayable"`
	LTsalariespayable            float64 `1.json:"l_tsalariespayable"`
	Accruedliabilities           float64 `1.json:"accruedliabilities"`
	Deferredincometaxliabilities float64 `1.json:"deferredincometaxliabilities"`
	DeferredincomenCL            float64 `1.json:"deferredincomen_cl"`
	OthernCL                     float64 `1.json:"othern_cl"`
	DiffnCLspecialitems          float64 `1.json:"diffn_c_lspecialitems"`
	DiffnCLbalanceitems          float64 `1.json:"diffn_c_lbalanceitems"`
	TotalnCL                     float64 `1.json:"totaln_cl"`
	Diffliabilitiesspecialitems  float64 `1.json:"diffliabilitiesspecialitems"`
	Diffliabilitiesbalanceitems  float64 `1.json:"diffliabilitiesbalanceitems"`
	Totalliabilities             float64 `1.json:"totalliabilities"`
}

type OwnersOrShareholdersEquity struct {
	Paidincapital                                  float64 `1.json:"paidincapital"`
	Otherequityinstruments                         float64 `1.json:"otherequityinstruments"`
	Includingpreferenceshares                      float64 `1.json:"includingpreferenceshares"`
	Includingperpetualdebt                         float64 `1.json:"includingperpetualdebt"`
	Capitalreserve                                 float64 `1.json:"capitalreserve"`
	Lesstreasurystock                              float64 `1.json:"lesstreasurystock"`
	OCI                                            float64 `1.json:"oci"`
	Specificreserve                                float64 `1.json:"specificreserve"`
	Surplusreserve                                 float64 `1.json:"surplusreserve"`
	Genericriskreserve                             float64 `1.json:"genericriskreserve"`
	Undistributedprofit                            float64 `1.json:"undistributedprofit"`
	Converteddifferenceonforeigncurrencystatements float64 `1.json:"converteddifferenceonforeigncurrencystatements"`
	Unidentifiedinvestmentloss                     float64 `1.json:"unidentifiedinvestmentloss"`
	Diffshareholdersequityspecialitems             float64 `1.json:"diffshareholdersequityspecialitems"`
	Diffshareholdersequitybalanceitems             float64 `1.json:"diffshareholdersequitybalanceitems"`
	Parentcompanyequity                            float64 `1.json:"parentcompanyequity"`
	Minorityprofit                                 float64 `1.json:"minorityprofit"`
	Totalownersequity                              float64 `1.json:"totalownersequity"`
	Diffliabilitiesandownersequityspecialitems     float64 `1.json:"diffliabilitiesandownersequityspecialitems"`
	Diffliabilitiesandownersequitybalanceitems     float64 `1.json:"diffliabilitiesandownersequitybalanceitems"`
	Totalliabilitiesandownersequity                float64 `1.json:"totalliabilitiesandownersequity"`
}

type ConsolidatedIncomeStatement struct {
	Totaloperatingincome                    float64 `1.json:"totaloperatingincome"`
	Operatingrevenue                        float64 `1.json:"operatingrevenue"`
	Otherfinancialoperatingrevenue          float64 `1.json:"otherfinancialoperatingrevenue"`
	Totaloperatingcosts                     float64 `1.json:"totaloperatingcosts"`
	Operatingcosts                          float64 `1.json:"operatingcosts"`
	Product                                 float64 `1.json:"product"`
	Area                                    float64 `1.json:"area"`
	Taxandassociatedchargeforprimeoperating float64 `1.json:"taxandassociatedchargeforprimeoperating"`
	Sellingexpenses                         float64 `1.json:"sellingexpenses"`
	Administrativeexpenses                  float64 `1.json:"administrativeexpenses"`
	Researchanddevelopmentexpenses          float64 `1.json:"researchanddevelopmentexpenses"`
	Financialexpense                        float64 `1.json:"financialexpense"`
	Includinginterestexpense                float64 `1.json:"includinginterestexpense"`
	Lessinterestincome                      float64 `1.json:"lessinterestincome"`
	Otherfinancialoperatingcosts            float64 `1.json:"otherfinancialoperatingcosts"`
	Addotherinvestmentincome                float64 `1.json:"addotherinvestmentincome"`
	Netinvestmentincome                     float64 `1.json:"netinvestmentincome"`
	Inclincomefrominvestmenttojointventures float64 `1.json:"inclincomefrominvestmenttojointventures"`
	GainsfromthederecognitionaMC            float64 `1.json:"gainsfromthederecognitiona_mc"`
	Netexposurehedginggainsandlosses        float64 `1.json:"netexposurehedginggainsandlosses"`
	FVGL                                    float64 `1.json:"fvgl"`
	Impairmentlossesonassets                float64 `1.json:"impairmentlossesonassets"`
	Impairmentlossesongoodwill              float64 `1.json:"impairmentlossesongoodwill"`
	Gainondisposalofassets                  float64 `1.json:"gainondisposalofassets"`
	Exchangnetincome                        float64 `1.json:"exchangnetincome"`
	Diffoperatingprofitspecialitems         float64 `1.json:"diffoperatingprofitspecialitems"`
	Diffoperatingprofitbalanceitems         float64 `1.json:"diffoperatingprofitbalanceitems"`
	Operatingprofit                         float64 `1.json:"operatingprofit"`
	Addnonoperatingrevenue                  float64 `1.json:"addnonoperatingrevenue"`
	Lessnonoperatingexpenses                float64 `1.json:"lessnonoperatingexpenses"`
	IncludingnetlossesondisposalnCA         float64 `1.json:"includingnetlossesondisposaln_ca"`
	Difftotalprofitspecialitems             float64 `1.json:"difftotalprofitspecialitems"`
	Difftotalprofitbalanceitems             float64 `1.json:"difftotalprofitbalanceitems"`
	Totalprofit                             float64 `1.json:"totalprofit"`
	Lessincometax                           float64 `1.json:"lessincometax"`
	Addunidentifiedinvestmentloss           float64 `1.json:"addunidentifiedinvestmentloss"`
	DiffnPspecialitems                      float64 `1.json:"diffn_pspecialitems"`
	DiffnPbalanceitems                      float64 `1.json:"diffn_pbalanceitems"`
	NP                                      float64 `1.json:"np"`
	NPcontinuingoperations                  float64 `1.json:"n_pcontinuingoperations"`
	NPdiscontinuedoperations                float64 `1.json:"n_pdiscontinuedoperations"`
	Minorityinterestincome                  float64 `1.json:"minorityinterestincome"`
	NPparentcompany                         float64 `1.json:"n_pparentcompany"`
	OCI                                     float64 `1.json:"oci"`
	TCI                                     float64 `1.json:"tci"`
	TCIminorityshareholders                 float64 `1.json:"tc_iminorityshareholders"`
	TCIcommonshareholders                   float64 `1.json:"tc_icommonshareholders"`
	EPS                                     float64 `1.json:"eps"`
	BEPS                                    float64 `1.json:"beps"`
	DEPS                                    float64 `1.json:"deps"`
}

type ConsolidatedCashFlowStatement struct {
	CFO                      CFO                      `1.json:"cfo"`
	CFI                      CFI                      `1.json:"cfi"`
	FinancingCashFlow        FinancingCashFlow        `1.json:"financing_cash_flow"`
	SupplementaryInformation SupplementaryInformation `1.json:"supplementary_information"`
}

type CFO struct {
	Cashreceivedfromsalesofgoodsorrenderingservices float64 `1.json:"cashreceivedfromsalesofgoodsorrenderingservices"`
	Refundsoftaxesandlevies                         float64 `1.json:"refundsoftaxesandlevies"`
	OthercFOreceived                                float64 `1.json:"otherc_f_oreceived"`
	CFOin                                           float64 `1.json:"cf_oin"`
	DiffcFOinspecialitems                           float64 `1.json:"diffc_f_oinspecialitems"`
	DiffcFOinbalanceitems                           float64 `1.json:"diffc_f_oinbalanceitems"`
	TotalcFOin                                      float64 `1.json:"totalc_f_oin"`
	Cashpaidforcommoditiesandservices               float64 `1.json:"cashpaidforcommoditiesandservices"`
	Cashpaidbehalfofemployees                       float64 `1.json:"cashpaidbehalfofemployees"`
	Varioustaxespaid                                float64 `1.json:"varioustaxespaid"`
	OthercFOpaid                                    float64 `1.json:"otherc_f_opaid"`
	CFOout                                          float64 `1.json:"cf_oout"`
	DiffcFOoutspecialitems                          float64 `1.json:"diffc_f_ooutspecialitems"`
	DiffcFOoutbalanceitems                          float64 `1.json:"diffc_f_ooutbalanceitems"`
	TotalcFOout                                     float64 `1.json:"totalc_f_oout"`
	DiffcFObalanceitems                             float64 `1.json:"diffc_f_obalanceitems"`
	CFO                                             float64 `1.json:"cfo"`
}

type CFI struct {
	Cashreceiptsfrominvestmentswithdrawal                   float64 `1.json:"cashreceiptsfrominvestmentswithdrawal"`
	Cashreceiptsfromreturnoninvestments                     float64 `1.json:"cashreceiptsfromreturnoninvestments"`
	NetcashdisposalfAiAlTA                                  float64 `1.json:"netcashdisposalf_ai_al_ta"`
	Netcashdisposalofsubsidiaryorotherbusinessunit          float64 `1.json:"netcashdisposalofsubsidiaryorotherbusinessunit"`
	OthercFI                                                float64 `1.json:"otherc_fi"`
	DiffcFIinspecialitems                                   float64 `1.json:"diffc_f_iinspecialitems"`
	DiffcFIinbalanceitems                                   float64 `1.json:"diffc_f_iinbalanceitems"`
	TotalcFI                                                float64 `1.json:"totalc_fi"`
	CashpurchasefAiAlTA                                     float64 `1.json:"cashpurchasef_ai_al_ta"`
	Cashpaymentforinvestment                                float64 `1.json:"cashpaymentforinvestment"`
	Netcashpurchaseofsubsidiariesandotherassociatedentities float64 `1.json:"netcashpurchaseofsubsidiariesandotherassociatedentities"`
	Otherinvestingcashoutflow                               float64 `1.json:"otherinvestingcashoutflow"`
	DiffcFIoutspecialitems                                  float64 `1.json:"diffc_f_ioutspecialitems"`
	DiffcFIoutbalanceitems                                  float64 `1.json:"diffc_f_ioutbalanceitems"`
	TotalcFIout                                             float64 `1.json:"totalc_f_iout"`
	DiffcFIbalanceitems                                     float64 `1.json:"diffc_f_ibalanceitems"`
	CFI                                                     float64 `1.json:"cfi"`
}

type FinancingCashFlow struct {
	Cashreceivedforattractedinvestment                       float64 `1.json:"cashreceivedforattractedinvestment"`
	Cashreceivedinvestmentminorityshareholderstosubsidiaries float64 `1.json:"cashreceivedinvestmentminorityshareholderstosubsidiaries"`
	Cashraisedfromloans                                      float64 `1.json:"cashraisedfromloans"`
	Cashreceivedfromotherfinancingactivities                 float64 `1.json:"cashreceivedfromotherfinancingactivities"`
	Cashproceedsofbondissue                                  float64 `1.json:"cashproceedsofbondissue"`
	DiffcFFinspecialitems                                    float64 `1.json:"diffc_f_finspecialitems"`
	DiffcFFinbalanceitems                                    float64 `1.json:"diffc_f_finbalanceitems"`
	TotalcFFin                                               float64 `1.json:"totalc_f_fin"`
	Cashdebtrepayment                                        float64 `1.json:"cashdebtrepayment"`
	Cashdistributionofdividendprofitorinterestpayment        float64 `1.json:"cashdistributionofdividendprofitorinterestpayment"`
	Dividendprofitpaidminorityshareholdersfromsubsidiaries   float64 `1.json:"dividendprofitpaidminorityshareholdersfromsubsidiaries"`
	Includingotherfinancingcashoutflow                       float64 `1.json:"includingotherfinancingcashoutflow"`
	DiffcFFoutspecialitems                                   float64 `1.json:"diffc_f_foutspecialitems"`
	DiffcFFoutbalanceitems                                   float64 `1.json:"diffc_f_foutbalanceitems"`
	TotalcFFout                                              float64 `1.json:"totalc_f_fout"`
	DiffcFFbalanceitems                                      float64 `1.json:"diffc_f_fbalanceitems"`
	CFF                                                      float64 `1.json:"cff"`
	Effectofforeignexchangeratechangesoncash                 float64 `1.json:"effectofforeignexchangeratechangesoncash"`
	Netincreasecashandcashequivalentsdirectspecialitems      float64 `1.json:"netincreasecashandcashequivalentsdirectspecialitems"`
	Netincreasecashandcashequivalentsdirectbalanceitems      float64 `1.json:"netincreasecashandcashequivalentsdirectbalanceitems"`
	Netincreaseofcashandcashequivalents                      float64 `1.json:"netincreaseofcashandcashequivalents"`
	Balanceofcashandcashequivalentsbeginningoftheperiod      float64 `1.json:"balanceofcashandcashequivalentsbeginningoftheperiod"`
	Balanceofcashandcashequivalentsendoftheperiod            float64 `1.json:"balanceofcashandcashequivalentsendoftheperiod"`
}

type SupplementaryInformation struct {
	NP                                        float64 `1.json:"np"`
	Depreciationreserveofassets               float64 `1.json:"depreciationreserveofassets"`
	Impairmentlossesongoodwill                float64 `1.json:"impairmentlossesongoodwill"`
	Depreciationcharge                        float64 `1.json:"depreciationcharge"`
	Amortizationofintangibleassets            float64 `1.json:"amortizationofintangibleassets"`
	AmortisationcostrOUassets                 float64 `1.json:"amortisationcostr_o_uassets"`
	AmortizationoflTdeferredexpenses          float64 `1.json:"amortizationofl_tdeferredexpenses"`
	Decreaseindeferredexpenses                float64 `1.json:"decreaseindeferredexpenses"`
	Increaseinaccruedexpenses                 float64 `1.json:"increaseinaccruedexpenses"`
	LossesondisposalfAiAlTA                   float64 `1.json:"lossesondisposalf_ai_al_ta"`
	LossesondiscardingfA                      float64 `1.json:"lossesondiscardingf_a"`
	Lossesonfairvaluechange                   float64 `1.json:"lossesonfairvaluechange"`
	Financialexpense                          float64 `1.json:"financialexpense"`
	Investmentlosses                          float64 `1.json:"investmentlosses"`
	Decreaseofdeferredincometaxasset          float64 `1.json:"decreaseofdeferredincometaxasset"`
	Increaseofdeferredincometaxliabilities    float64 `1.json:"increaseofdeferredincometaxliabilities"`
	Decreaseininventories                     float64 `1.json:"decreaseininventories"`
	Decreaseofoperatingreceivables            float64 `1.json:"decreaseofoperatingreceivables"`
	Increaseofoperatingpayables               float64 `1.json:"increaseofoperatingpayables"`
	Unidentifiedinvestmentloss                float64 `1.json:"unidentifiedinvestmentloss"`
	Otheritem                                 float64 `1.json:"otheritem"`
	DiffcFOindirectspecialitems               float64 `1.json:"diffc_f_oindirectspecialitems"`
	DiffcFOindirectbalanceitems               float64 `1.json:"diffc_f_oindirectbalanceitems"`
	CFOindirectbalanceitems                   float64 `1.json:"cf_oindirectbalanceitems"`
	Debtconvertedtosharecapital               float64 `1.json:"debtconvertedtosharecapital"`
	Convertiblebondsmaturewithinoneyear       float64 `1.json:"convertiblebondsmaturewithinoneyear"`
	FAacquiredunderfinanceleases              float64 `1.json:"f_aacquiredunderfinanceleases"`
	Cashendoftheperiod                        float64 `1.json:"cashendoftheperiod"`
	Cashbeginningoftheperiod                  float64 `1.json:"cashbeginningoftheperiod"`
	Cashequivalentsendoftheperiod             float64 `1.json:"cashequivalentsendoftheperiod"`
	Cashequivalentsbeginningoftheperiod       float64 `1.json:"cashequivalentsbeginningoftheperiod"`
	Diffnetincreasecashindirectspecialitems   float64 `1.json:"diffnetincreasecashindirectspecialitems"`
	Diffnetincreasecashindirectbalanceitems   float64 `1.json:"diffnetincreasecashindirectbalanceitems"`
	Netincreasecashandcashequivalentsindirect float64 `1.json:"netincreasecashandcashequivalentsindirect"`
}
