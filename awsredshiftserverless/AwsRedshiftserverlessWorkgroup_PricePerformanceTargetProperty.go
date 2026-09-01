package awsredshiftserverless


// Experimental.
type AwsRedshiftserverlessWorkgroup_PricePerformanceTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_workgroup#enabled AwsRedshiftserverlessWorkgroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_workgroup#level AwsRedshiftserverlessWorkgroup#level}.
	// Experimental.
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

