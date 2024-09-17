package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	CloudLabID           int    `json:"cloud_lab_id"`
	Name                 string `json:"name"`
	Code                 string `gorm:"unique" json:"code"`
	Type                 string `json:"type"`
	Description          string `json:"description"`
	MinimumQuantity      int    `gorm:"default:1" json:"minimum_quantity"`
	MaximumQuantity      *int   `json:"maximum_quantity"`
	PackingUnit          int    `gorm:"default:1" json:"packing_unit"`
	EnableCustomQuantity bool   `gorm:"default:false" json:"enable_custom_quantity"`
	EnableCustomFormat   bool   `gorm:"default:false" json:"enable_custom_format"`
	TimeToProduce        string `json:"time_to_produce"`
	RenamingRules        string `json:"renaming_rules"`
	OrderRules           string `json:"order_rules"`
	DefaultQuantity      int    `json:"default_quantity"`
	QuantitiesSelection  string `json:"quantities_selection"`
	PriceCalculationType string `json:"price_calculation_type"`
}
