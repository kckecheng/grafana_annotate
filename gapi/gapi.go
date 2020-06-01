// Package gapi Grafana annotation and dependent functions encapsulation
package gapi

import (
	"context"
	"fmt"
	"time"

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

// NewGrafana init Grafana connection
func NewGrafana(address string, port int64, user, password string) (*Grafana, error) {
	client := grafana.NewClient(fmt.Sprintf("http://%s:%d", address, port), fmt.Sprintf("%s:%s", user, password), grafana.DefaultHTTPClient)
	if client == nil {
		return nil, fmt.Errorf("Could not initial a client to Grafana")
	}

	g := Grafana{
		address:  address,
		port:     port,
		user:     user,
		password: password,
		client:   client,
	}
	return &g, nil
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

// GetAnnotattions list all annotations
func (g *Grafana) GetAnnotattions(dashboardIds, panelIds []uint, from, to time.Time, tags []string) ([]grafana.AnnotationResponse, error) {
	var params []grafana.GetAnnotationsParams

	if dashboardIds != nil && len(dashboardIds) > 0 {
		for _, dashboardId := range dashboardIds {
			params = append(params, grafana.WithDashboard(dashboardId))
		}
	}

	if panelIds != nil && len(panelIds) > 0 {
		for _, panelId := range panelIds {
			params = append(params, grafana.WithPanel(panelId))
		}
	}

	var tzero time.Time
	if from != tzero {
		params = append(params, grafana.WithStartTime(from))
	}
	if to != tzero {
		params = append(params, grafana.WithEndTime(to))
	}

	if tags != nil && len(tags) > 0 {
		for _, tag := range tags {
			params = append(params, grafana.WithTag(tag))
		}
	}

	return g.client.GetAnnotations(context.Background(), params...)
}

// CreateAnnotation create an annotation
func (g *Grafana) CreateAnnotation(dashboardId, panelId uint, from, to int64, tags []string, text string) (*grafana.StatusMessage, error) {
	areq := grafana.CreateAnnotationRequest{
		DashboardID: dashboardId,
		PanelID:     panelId,
		Text:        text,
	}

	if from != 0 {
		areq.Time = from
	}
	if to != 0 {
		areq.TimeEnd = to
	}

	if tags != nil && len(tags) > 0 {
		areq.Tags = tags
	}

	msg, err := g.client.CreateAnnotation(context.Background(), areq)
	if err != nil {
		fmt.Printf("Fail to create annotation with error: %s", err)
		return nil, err
	}

	return &msg, nil
}
