package awsresiliencehubv2


// Experimental.
type TfPolicy_MultiRegionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#disaster_recovery_approach TfPolicy#disaster_recovery_approach}.
	// Experimental.
	DisasterRecoveryApproach *string `field:"required" json:"disasterRecoveryApproach" yaml:"disasterRecoveryApproach"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#rpo_in_minutes TfPolicy#rpo_in_minutes}.
	// Experimental.
	RpoInMinutes *float64 `field:"optional" json:"rpoInMinutes" yaml:"rpoInMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resiliencehubv2_policy#rto_in_minutes TfPolicy#rto_in_minutes}.
	// Experimental.
	RtoInMinutes *float64 `field:"optional" json:"rtoInMinutes" yaml:"rtoInMinutes"`
}

