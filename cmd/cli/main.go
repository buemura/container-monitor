package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/jedib0t/go-pretty/v6/table"
)

const (
	colorReset = "\033[0m"
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
)

func main() {
	// Create a new Docker client
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	for {
		// List running containers
		containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
		if err != nil {
			panic(err)
		}

		// Clear the screen
		fmt.Print("\033[H\033[2J")

		// Create a new table writer
		t := table.NewWriter()
		renderTableHeader(t)

		// Render table rows
		for _, container := range containers {
			renderTableRow(t, container)
		}

		// Render the table to a buffer and print it
		buf := new(bytes.Buffer)
		t.SetOutputMirror(buf)
		t.Render()
		fmt.Println(buf.String())

		// Wait for 1 second
		time.Sleep(1 * time.Second)
	}
}

/**
 * @description Renders the header of the table with the specified columns.
 * @param t table.Writer - The table writer to render the header to.
 */
func renderTableHeader(t table.Writer) {
	t.SetStyle(table.StyleLight)
	t.Style().Options.SeparateRows = true
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"ID", "Image", "Port", "Created", "Status", "Names"})
}

/**
 * @description Renders a single row of the table with information about a container.
 * @param t table.Writer - The table writer to render the row to.
 * @param c types.Container - The container object containing information about the container.
 * @return nil - The function does not return a value.
 */
func renderTableRow(t table.Writer, c types.Container) {
	// Container ID
	containerID := c.ID[:12]

	// Container Created At
	createdTimeUTC := time.Unix(c.Created, 0)
	createdInRFC3339 := createdTimeUTC.Format("2006-01-02 15:04:05")

	// Container Port when private is ::
	var port types.Port
	for _, p := range c.Ports {
		if p.IP == "::" {
			port = p
		}
	}

	// Set color based on the status
	var status string
	if strings.Contains(c.Status, "Up") {
		status = fmt.Sprintf("%s%s%s", colorGreen, c.Status, colorReset)
	} else {
		status = fmt.Sprintf("%s%s%s", colorRed, c.Status, colorReset)
	}

	t.AppendRow(table.Row{containerID, c.Image, fmt.Sprintf("%s%d->%d", port.IP, port.PrivatePort, port.PublicPort), createdInRFC3339, status, c.Names[0]})
}
