// Package main implements a simple sample implementation of the Toggl API for go (github.com/andreaskoch/togglapi).
// The tool will print a list of all your workspaces, clients, project and some recent time entries to give
// you an idea how you can use the github.com/andreaskoch/togglapi package.
//
// Usage:
// go run main.go your-api-token
package main

import (
	"fmt"
	"os"
	"time"

	"github.com/andreaskoch/togglapi"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Please supply your API token as an argument\n")
		os.Exit(1)
	}

	apiToken := os.Args[1]
	baseURL := "https://www.toggl.com/api/v8"
	api := togglapi.NewAPI(baseURL, apiToken)

	// workspaces
	fmt.Println("Workspaces:")

	workspaces, workspacesError := api.GetWorkspaces()
	if workspacesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get workspaces: %s", workspacesError)
		os.Exit(1)
	}

	for _, workspace := range workspaces {
		fmt.Println(workspace.Name)
	}

	fmt.Println("")

	// clients
	fmt.Println("Clients:")

	clients, clientsError := api.GetClients()
	if clientsError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get clients: %s", clientsError)
		os.Exit(1)
	}

	for _, client := range clients {
		fmt.Println(client.Name)
	}

	fmt.Println("")

	// projects
	fmt.Println("Projects:")

	for _, workspace := range workspaces {

		projects, projectsError := api.GetProjects(workspace.ID)
		if projectsError != nil {
			fmt.Fprintf(os.Stderr, "Failed to get projects: %s", projectsError)
			os.Exit(1)
		}

		for _, project := range projects {
			fmt.Println(project.Name)
		}
	}

	fmt.Println("")

	// timeEntries
	fmt.Println("Time Entries:")

	stop := time.Now()
	start := stop.AddDate(0, -1, 0)

	timeEntries, timeEntriesError := api.GetTimeEntries(start, stop)
	if timeEntriesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get timeEntries: %s", timeEntriesError)
		os.Exit(1)
	}

	for _, timeEntry := range timeEntries {
		fmt.Printf("%s - %s: %s\n", timeEntry.Start, timeEntry.Stop, timeEntry.Description)
	}

	fmt.Println("")
}
