package s3control


// Experimental.
type AwsMultiRegionAccessPointRoutes_RouteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point_routes#bucket AwsMultiRegionAccessPointRoutes#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point_routes#region AwsMultiRegionAccessPointRoutes#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point_routes#traffic_dial_percentage AwsMultiRegionAccessPointRoutes#traffic_dial_percentage}.
	// Experimental.
	TrafficDialPercentage *float64 `field:"required" json:"trafficDialPercentage" yaml:"trafficDialPercentage"`
}

