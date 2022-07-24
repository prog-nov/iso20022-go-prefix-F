package iso20022

// Set of characteristics shared by all individual transactions included in the message.
type GroupHeader3 struct {

	// Point to point reference assigned by the instructing party and sent to the next party in the chain to unambiguously identify the message.
	//
	// Usage: The instructing party has to make sure that 'MessageIdentification' is unique per instructed party for a pre-agreed period.
	MessageIdentification *Max35Text `xml:"MsgId"`

	// Date and time at which a (group of) payment instruction(s) was created by the instructing party.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the initiating party is allowed to initiate transactions from the account specified in the initiation.
	//
	// Usage: the content is not of a technical nature, but reflects the organisational structure at the initiating side. The authorisation element can typically be used in relay scenarios, payment initiations, payment returns or payment reversals that are initiated on behalf of a different party than the initiating party.
	Authorisation []*Max128Text `xml:"Authstn,omitempty"`

	// Identifies whether a single entry per individual transaction or a batch entry for the sum of the amounts of all transactions in the message is requested.
	BatchBooking *BatchBookingIndicator `xml:"BtchBookg,omitempty"`

	// Number of individual transactions contained in the message.
	NumberOfTransactions *Max15NumericText `xml:"NbOfTxs"`

	// Total of all individual amounts included in the message, irrespective of currencies.
	ControlSum *DecimalNumber `xml:"CtrlSum,omitempty"`

	// Total amount of money transferred between instructing agent and instructed agent.
	TotalInterbankSettlementAmount *CurrencyAndAmount `xml:"TtlIntrBkSttlmAmt,omitempty"`

	// Date on which the amount of money ceases to be available to the agent that owes it and when the amount of money becomes available to the agent to which it is due.
	InterbankSettlementDate *ISODate `xml:"IntrBkSttlmDt,omitempty"`

	// Specifies the details on how the settlement of the transaction(s) between the instructing agent and the instructed agent is completed.
	SettlementInformation *SettlementInformation2 `xml:"SttlmInf"`

	// Set of elements used to further specify the type of transaction.
	PaymentTypeInformation *PaymentTypeInformation4 `xml:"PmtTpInf,omitempty"`

	// Agent that instructs the next party in the chain to carry out the (set of) instruction(s).
	InstructingAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstgAgt,omitempty"`

	// Agent that is instructed by the previous party in the chain to carry out the (set of) instruction(s).
	InstructedAgent *BranchAndFinancialInstitutionIdentification3 `xml:"InstdAgt,omitempty"`
}

func (g *GroupHeader3) SetMessageIdentification(value string) {
	g.MessageIdentification = (*Max35Text)(&value)
}

func (g *GroupHeader3) SetCreationDateTime(value string) {
	g.CreationDateTime = (*ISODateTime)(&value)
}

func (g *GroupHeader3) AddAuthorisation(value string) {
	g.Authorisation = append(g.Authorisation, (*Max128Text)(&value))
}

func (g *GroupHeader3) SetBatchBooking(value string) {
	g.BatchBooking = (*BatchBookingIndicator)(&value)
}

func (g *GroupHeader3) SetNumberOfTransactions(value string) {
	g.NumberOfTransactions = (*Max15NumericText)(&value)
}

func (g *GroupHeader3) SetControlSum(value string) {
	g.ControlSum = (*DecimalNumber)(&value)
}

func (g *GroupHeader3) SetTotalInterbankSettlementAmount(value, currency string) {
	g.TotalInterbankSettlementAmount = NewCurrencyAndAmount(value, currency)
}

func (g *GroupHeader3) SetInterbankSettlementDate(value string) {
	g.InterbankSettlementDate = (*ISODate)(&value)
}

func (g *GroupHeader3) AddSettlementInformation() *SettlementInformation2 {
	g.SettlementInformation = new(SettlementInformation2)
	return g.SettlementInformation
}

func (g *GroupHeader3) AddPaymentTypeInformation() *PaymentTypeInformation4 {
	g.PaymentTypeInformation = new(PaymentTypeInformation4)
	return g.PaymentTypeInformation
}

func (g *GroupHeader3) AddInstructingAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructingAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructingAgent
}

func (g *GroupHeader3) AddInstructedAgent() *BranchAndFinancialInstitutionIdentification3 {
	g.InstructedAgent = new(BranchAndFinancialInstitutionIdentification3)
	return g.InstructedAgent
}
