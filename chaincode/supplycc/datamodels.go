package main

import (
	"time"
)

// CultivationData - Cultivation block structure
type CultivationData struct {
	BatchID      string `json:"id"`
	FarmerName   string `json:"farmer-name"`
	FarmAddress  string `json:"farm-address"`
	ExporterName string `json:"exporter-name"`
	ImporerName  string `json:"importer-name"`
}

// FarmInspectorData - FarmInspector block structure
type FarmInspectorData struct {
	BatchID        string `json:"id"`
	SeedType       string `json:"seed-type"`
	CoffeeFamily   string `json:"coffee-family"`
	FertilizerUsed string `json:"fertilizer-used"`
}

// HarvesterData - Harvester block structure
type HarvesterData struct {
	BatchID     string  `json:"id"`
	CropVariety string  `json:"crop-variety"`
	Temperature float32 `json:"temperature"` // in deg C
	Humidity    float32 `json:"humidity"`    // in %
}

// ExporterData - Exporter block structure
type ExporterData struct {
	BatchID              string     `json:"id"`
	Quantity             float32    `json:"quantity"` // in kg
	DestinationAddress   string     `json:"destination-address"`
	ShipName             string     `json:"ship-name"`
	ShipNo               string     `json:"ship-no"`
	DepartureDateTime    *time.Time `json:"departure-date-time"`
	EstimatedArrivalTime *time.Time `json:"estimated-arrival-time"`
	ExporterID           int        `json:"exporter-id"`
}

// ImporterData - Importer block structure
type ImporterData struct {
	BatchID          string     `json:"id"`
	ShipName         string     `json:"ship-name"`
	ShipNo           string     `json:"ship-no"`
	ArrivalDateTime  *time.Time `json:"arrival-date-time"`
	TransportInfo    string     `json:"transport-info"`
	WarehouseName    string     `json:"warehouse-name"`
	WarehouseAddress string     `json:"warehouse-address"`
	ImporterID       int        `json:"importer-id"`
}

// ProcessorData - Processor block structure
type ProcessorData struct {
	BatchID          string     `json:"id"`
	Temperature      float32    `json:"temperature"`       // in deg C
	RoastingDuration int        `json:"roasting-duration"` // in seconds
	PackageDateTime  *time.Time `json:"package-date-time"`
	ProcessorName    string     `json:"processor-name"`
	WarehouseAddress string     `json:"warehouse-address"`
}
