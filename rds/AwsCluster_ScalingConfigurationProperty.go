package rds


// Experimental.
type AwsCluster_ScalingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#auto_pause AwsCluster#auto_pause}.
	// Experimental.
	AutoPause interface{} `field:"optional" json:"autoPause" yaml:"autoPause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#max_capacity AwsCluster#max_capacity}.
	// Experimental.
	MaxCapacity *float64 `field:"optional" json:"maxCapacity" yaml:"maxCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#min_capacity AwsCluster#min_capacity}.
	// Experimental.
	MinCapacity *float64 `field:"optional" json:"minCapacity" yaml:"minCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#seconds_before_timeout AwsCluster#seconds_before_timeout}.
	// Experimental.
	SecondsBeforeTimeout *float64 `field:"optional" json:"secondsBeforeTimeout" yaml:"secondsBeforeTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#seconds_until_auto_pause AwsCluster#seconds_until_auto_pause}.
	// Experimental.
	SecondsUntilAutoPause *float64 `field:"optional" json:"secondsUntilAutoPause" yaml:"secondsUntilAutoPause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#timeout_action AwsCluster#timeout_action}.
	// Experimental.
	TimeoutAction *string `field:"optional" json:"timeoutAction" yaml:"timeoutAction"`
}

