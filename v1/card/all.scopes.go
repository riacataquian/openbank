// Code generated by "scopegen"; DO NOT EDIT.
package card

type AuthScope string

const (
	Scope_Read AuthScope = "https://auth.bnk.to/card.read"
	Scope_Write AuthScope = "https://auth.bnk.to/card.write"
)

var AuthScopes = map[string][]AuthScope{
	"/card.CardService/CreateCard": []AuthScope{Scope_Write},
	"/card.CardService/CreateCardAttribute": []AuthScope{Scope_Write},
	"/card.CardService/DeleteCard": []AuthScope{Scope_Write},
	"/card.CardService/GetCard": []AuthScope{Scope_Read},
	"/card.CardService/GetUserCards": []AuthScope{Scope_Read},
	"/card.CardService/UpdateCardAccessStatus": []AuthScope{Scope_Write},
	"/card.CardService/UpdateCardStatus": []AuthScope{Scope_Write},
}
