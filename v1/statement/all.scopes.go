// Code generated by "scopegen"; DO NOT EDIT.
package statement

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/statement.read"
	Scope_Write AuthScope = "https://auth.bnk.to/statement.write"
)

var AuthScopes = map[string][]AuthScope{
	"/statement.StatementService/GetStatement": []AuthScope{Scope_Read},
	"/statement.StatementService/GetStatements": []AuthScope{Scope_Read},
}
