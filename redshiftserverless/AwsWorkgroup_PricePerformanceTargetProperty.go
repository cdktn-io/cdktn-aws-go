package redshiftserverless


// Experimental.
type AwsWorkgroup_PricePerformanceTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_workgroup#enabled AwsWorkgroup#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshiftserverless_workgroup#level AwsWorkgroup#level}.
	// Experimental.
	Level *float64 `field:"optional" json:"level" yaml:"level"`
}

