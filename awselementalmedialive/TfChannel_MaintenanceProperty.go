package awselementalmedialive


// Experimental.
type TfChannel_MaintenanceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#maintenance_day TfChannel#maintenance_day}.
	// Experimental.
	MaintenanceDay *string `field:"required" json:"maintenanceDay" yaml:"maintenanceDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#maintenance_start_time TfChannel#maintenance_start_time}.
	// Experimental.
	MaintenanceStartTime *string `field:"required" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
}

