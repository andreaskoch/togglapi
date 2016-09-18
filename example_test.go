package togglapi

import (
	"fmt"
	"os"
	"time"
)

// This example shows how you can use the Toggl API to print
// all your workspaces names, client names, project names and the time entries
// of the last month.
func Example() {

	baseURL := "https://www.toggl.com/api/v8"
	apiToken := "Toggl-API-Token"
	api := NewAPI(baseURL, apiToken)

	// print workspace names
	fmt.Println("Workspaces:")

	workspaces, workspacesError := api.GetWorkspaces()
	if workspacesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get workspaces: %s", workspacesError)
		return
	}

	for _, workspace := range workspaces {
		fmt.Println(workspace.Name)
	}

	fmt.Println("")

	// print client names
	fmt.Println("Clients:")

	clients, clientsError := api.GetClients()
	if clientsError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get clients: %s", clientsError)
		return
	}

	for _, client := range clients {
		fmt.Println(client.Name)
	}

	fmt.Println("")

	// print project names
	fmt.Println("Projects:")

	for _, workspace := range workspaces {

		projects, projectsError := api.GetProjects(workspace.ID)
		if projectsError != nil {
			fmt.Fprintf(os.Stderr, "Failed to get projects: %s", projectsError)
			return
		}

		for _, project := range projects {
			fmt.Println(project.Name)
		}
	}

	fmt.Println("")

	// print time entries
	fmt.Println("Time Entries:")

	stop := time.Now()
	start := stop.AddDate(0, -1, 0)

	timeEntries, timeEntriesError := api.GetTimeEntries(start, stop)
	if timeEntriesError != nil {
		fmt.Fprintf(os.Stderr, "Failed to get timeEntries: %s", timeEntriesError)
		return
	}

	for _, timeEntry := range timeEntries {
		fmt.Printf("%s - %s: %s\n", timeEntry.Start, timeEntry.Stop, timeEntry.Description)
	}

	fmt.Println("")
}
