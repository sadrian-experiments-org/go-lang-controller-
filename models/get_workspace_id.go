package models

import (
	"fmt"
	"net/http"
)

// TODO: Modify the function to expect the variables bellow as arguments
func GetWorkspaceID() {
	org_name := 
	workspace_name :=
	TOKEN :=  

	url := fmt.Sprintf(
		"https://app.terraform.io/api/v2/organizations/%s/workspaces/%s",
		org_name, 
		workspace_name
	)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Something went wrong while retrieving the workspace ID:",
		err
	)
		return
	}

	request.Header.Add("Authorization", "Bearer %s", TOKEN)
	request.Header.Add("Content-Type: application/vnd.api+json")


}