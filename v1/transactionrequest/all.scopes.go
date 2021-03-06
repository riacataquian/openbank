// Code generated by "scopegen"; DO NOT EDIT.
package transactionrequest

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/transactionrequest.read"
	Scope_Write AuthScope = "https://auth.bnk.to/transactionrequest.write"
)

var AuthScopes = map[string][]AuthScope{
	"/transactionrequest.TransactionRequestService/AnswerTransactionRequestChallenge": []AuthScope{Scope_Write},
	"/transactionrequest.TransactionRequestService/CreateAccountOTPTransaction": []AuthScope{Scope_Write},
	"/transactionrequest.TransactionRequestService/CreateAccountTransaction": []AuthScope{Scope_Write},
	"/transactionrequest.TransactionRequestService/CreateCounterPartyTransaction": []AuthScope{Scope_Write},
	"/transactionrequest.TransactionRequestService/CreateFreeFormTransaction": []AuthScope{Scope_Write},
	"/transactionrequest.TransactionRequestService/CreateSEPATransaction": []AuthScope{Scope_Write},
	"/transactionrequest.TransactionRequestService/GetSupportedTransactionRequestTypes": []AuthScope{Scope_Read},
	"/transactionrequest.TransactionRequestService/GetTransactionRequestTypes": []AuthScope{Scope_Read},
	"/transactionrequest.TransactionRequestService/GetTransactionRequests": []AuthScope{Scope_Read},
	"/transactionrequest.TransactionRequestService/SaveHistoricTransaction": []AuthScope{Scope_Write},
}
