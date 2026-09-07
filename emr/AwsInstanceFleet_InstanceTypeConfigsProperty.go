package emr


// Experimental.
type AwsInstanceFleet_InstanceTypeConfigsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#instance_type AwsInstanceFleet#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#bid_price AwsInstanceFleet#bid_price}.
	// Experimental.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#bid_price_as_percentage_of_on_demand_price AwsInstanceFleet#bid_price_as_percentage_of_on_demand_price}.
	// Experimental.
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#configurations AwsInstanceFleet#configurations}
	// Experimental.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#ebs_config AwsInstanceFleet#ebs_config}
	// Experimental.
	EbsConfig interface{} `field:"optional" json:"ebsConfig" yaml:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_instance_fleet#weighted_capacity AwsInstanceFleet#weighted_capacity}.
	// Experimental.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

