package location


// Experimental.
type AwsGeofenceCollection_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/location_geofence_collection#create AwsGeofenceCollection#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/location_geofence_collection#delete AwsGeofenceCollection#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/location_geofence_collection#update AwsGeofenceCollection#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

