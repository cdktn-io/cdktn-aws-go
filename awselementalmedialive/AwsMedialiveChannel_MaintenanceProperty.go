package awselementalmedialive


// Experimental.
type AwsMedialiveChannel_MaintenanceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#maintenance_day AwsMedialiveChannel#maintenance_day}.
	// Experimental.
	MaintenanceDay *string `field:"required" json:"maintenanceDay" yaml:"maintenanceDay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/medialive_channel#maintenance_start_time AwsMedialiveChannel#maintenance_start_time}.
	// Experimental.
	MaintenanceStartTime *string `field:"required" json:"maintenanceStartTime" yaml:"maintenanceStartTime"`
}

