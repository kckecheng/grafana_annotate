package main

import (
	"fmt"
	"os"
	"time"

	grafana "github.com/grafana-tools/sdk"
	"github.com/kckecheng/grafana_annotate/gapi"
)

func main() {
	host := "127.0.0.1"
	port := 3000
	user := "admin"
	pass := "admin"
	title := "Prometheus"

	g, err := gapi.NewGrafana(host, int64(port), user, pass)
	if err != nil {
		fmt.Println(err)
	}

	boards, err := g.GetAllDashboards()
	if err != nil {
		fmt.Println("Fail to get all dashboards")
		os.Exit(1)
	}
	fmt.Println("*** Dashboards ***")
	for _, board := range boards {
		fmt.Printf("ID: %d\n", board.ID)
		fmt.Printf("UID: %s\n", board.UID)
		fmt.Printf("Title: %s\n", board.Title)
		fmt.Println()
	}

	var dashboard grafana.FoundBoard
	for _, board := range boards {
		if board.Title == title {
			dashboard = board
			break
		}
	}

	panels, err := g.GetAllPanels(dashboard.UID)
	if err != nil {
		fmt.Println("Fail to get all panels")
		os.Exit(1)
	}
	fmt.Println("*** Panels ***")
	for _, panel := range panels {
		fmt.Printf("ID: %d\n", panel.ID)
		fmt.Printf("Title: %s\n", panel.Title)
		fmt.Println()
	}

	now := time.Now()
	start := now.Add(-10*time.Minute).Unix() * 1000
	end := now.Add(-5*time.Minute).Unix() * 1000

	for _, panel := range panels {
		_, err := g.CreateAnnotation(dashboard.ID, panel.ID, start, end, []string{"tag1", "tag2"}, "Hello world")
		if err != nil {
			fmt.Printf("Fail to create annotation")
			os.Exit(1)
		}
	}

	var tzero time.Time
	annotations, err := g.GetAnnotattions(nil, nil, tzero, tzero, nil)
	if err != nil {
		fmt.Println("Fail to get annotations")
	}
	fmt.Println("*** Annotations ***")
	for _, annotation := range annotations {
		fmt.Printf("Dashboard: %d\n", annotation.DashboardID)
		fmt.Printf("Panel: %d\n", annotation.PanelID)
		fmt.Printf("Start Time: %d\n", annotation.Time)
		fmt.Printf("End Time: %d\n", annotation.TimeEnd)
		fmt.Printf("Tags: %v\n", annotation.Tags)
		fmt.Printf("Text: %s\n", annotation.Text)
		fmt.Println()
	}
}
