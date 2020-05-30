package main

import (
	"context"
	"fmt"
	"log"

	grafana "github.com/grafana-tools/sdk"
)

// Grafana object
type Grafana struct {
	address  string
	port     int64
	user     string
	password string
	client   *grafana.Client
}

// NewGrafana init API connection
func NewGrafana(address string, port int64, user, password string) *Grafana {
	client := grafana.NewClient(fmt.Sprintf("http://%s:%d", address, port), fmt.Sprintf("%s:%s", user, password), grafana.DefaultHTTPClient)
	if client == nil {
		log.Fatal("Fail to init grafana client")
		return nil
	}

	g := Grafana{
		address:  address,
		port:     port,
		user:     user,
		password: password,
		client:   client,
	}
	return &g
}

// GetAllDashboards get all the dashboards
func (g *Grafana) GetAllDashboards() ([]grafana.FoundBoard, error) {
	boards, err := g.client.Search(context.Background(), grafana.SearchType(grafana.SearchTypeDashboard))
	if err != nil {
		return nil, err
	}
	return boards, nil
}

// GetAllPanels get all panels for a dashboard
func (g *Grafana) GetAllPanels(uid string) ([]*grafana.Panel, error) {
	board, _, err := g.client.GetDashboardByUID(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	return board.Panels, nil
}

func main() {
	g := NewGrafana("127.0.0.1", 3000, "admin", "admin")
	boards, err := g.GetAllDashboards()
	if err != nil {
		log.Fatalf("Fail to get all dashboards: %s", err)
	}

	var tuid string
	tname := "Prometheus"
	for _, board := range boards {
		if board.Title == tname {
			tuid = board.UID
			break
		}
	}

	panels, err := g.GetAllPanels(tuid)
	for _, panel := range panels {
		fmt.Println("Panel ID", panel.ID)
		fmt.Println("Panel Title", panel.Title)
		fmt.Println("Panel Type", panel.Type)
	}
}
