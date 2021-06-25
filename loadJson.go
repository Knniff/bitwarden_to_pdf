package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/jung-kurt/gofpdf"
)

func readJson(path string, data *BWData) {
	input, _ := ioutil.ReadFile(path)

	json.Unmarshal(input, &data)

}

func makePDF(path string) {
	var data BWData
	readJson(path, &data)

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	for i := 0; i < len(data.Items); i++ {

		if data.Items[i].Type == 1 {
			if i%5 == 0 && i != 0 {
				pdf.AddPage()
			}
			pdf.SetFont("Arial", "B", 16)
			pdf.Cellf(100, 10, "Service:")
			pdf.SetX(100)
			pdf.SetFont("Arial", "", 16)
			pdf.Cellf(100, 10, data.Items[i].Name)
			pdf.SetX(0)
			pdf.SetY(pdf.GetY() + 10)
			pdf.SetFont("Arial", "B", 16)
			pdf.Cellf(100, 10, "Benutzername:")
			pdf.SetX(100)
			pdf.SetFont("Arial", "", 16)
			pdf.Cellf(100, 10, data.Items[i].Login.Username)
			pdf.SetX(0)
			pdf.SetY(pdf.GetY() + 10)
			pdf.SetFont("Arial", "B", 16)
			pdf.Cellf(100, 10, "Passwort:")
			pdf.SetX(100)
			pdf.SetFont("Arial", "", 16)
			pdf.Write(10, data.Items[i].Login.Password)
			pdf.SetX(0)
			pdf.SetY(pdf.GetY() + 30)
		}
	}
	pdf.OutputFileAndClose("passwords.pdf")
}

type BWData struct {
	Folders []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"folders"`
	Items []struct {
		ID             string `json:"id"`
		OrganizationID string `json:"organizationId"`
		FolderID       string `json:"folderId"`
		Type           int    `json:"type"`
		Name           string `json:"name"`
		Notes          string `json:"notes"`
		Favorite       bool   `json:"favorite"`
		Fields         []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
			Type  int    `json:"type"`
		} `json:"fields"`
		SecureNote struct {
			Type int `json:"type"`
		} `json:"secureNote,omitempty"`
		CollectionIds []string `json:"collectionIds"`
		Card          struct {
			CardholderName string `json:"cardholderName"`
			Brand          string `json:"brand"`
			Number         string `json:"number"`
			ExpMonth       string `json:"expMonth"`
			ExpYear        string `json:"expYear"`
			Code           string `json:"code"`
		} `json:"card,omitempty"`
		Identity struct {
			Title          string `json:"title"`
			FirstName      string `json:"firstName"`
			MiddleName     string `json:"middleName"`
			LastName       string `json:"lastName"`
			Address1       string `json:"address1"`
			Address2       string `json:"address2"`
			Address3       string `json:"address3"`
			City           string `json:"city"`
			State          string `json:"state"`
			PostalCode     string `json:"postalCode"`
			Country        string `json:"country"`
			Company        string `json:"company"`
			Email          string `json:"email"`
			Phone          string `json:"phone"`
			Ssn            string `json:"ssn"`
			Username       string `json:"username"`
			PassportNumber string `json:"passportNumber"`
			LicenseNumber  string `json:"licenseNumber"`
		} `json:"identity,omitempty"`
		Login struct {
			Uris []struct {
				Match string `json:"match"`
				URI   string `json:"uri"`
			} `json:"uris"`
			Username string `json:"username"`
			Password string `json:"password"`
			Totp     string `json:"totp"`
		} `json:"login,omitempty"`
	} `json:"items"`
}
