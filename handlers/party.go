package handlers

import (
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

type User struct {
	Email string
	Name  string
}

type Party struct {
	EventName    string
	Users        []User
	Transactions []Transaction
}

type PartyService interface {
	CreateParty(name string, users []User) error
	GetParty(id string) (Party, error)
	UpdateUser(id string, user User) error
	DeleteParty(id string) error
}

type PartyServiceImpl struct {
	partyCollection *models.Collection
}

func NewPartyService(app *pocketbase.PocketBase) *PartyServiceImpl {
	partyCollection, err := app.Dao().FindCollectionByNameOrId("parties")
	if err != nil {
		app.Logger().Error("parties collection not established", "err", err)
	}
	return &PartyServiceImpl{partyCollection: partyCollection}
}

func (p *PartyServiceImpl) CreateParty(name string, users []User) error {
	return nil
}
