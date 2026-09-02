package awsemr


// Experimental.
type TfCluster_CoreInstanceFleetInstanceTypeConfigsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#instance_type TfCluster#instance_type}.
	// Experimental.
	InstanceType *string `field:"required" json:"instanceType" yaml:"instanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#bid_price TfCluster#bid_price}.
	// Experimental.
	BidPrice *string `field:"optional" json:"bidPrice" yaml:"bidPrice"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#bid_price_as_percentage_of_on_demand_price TfCluster#bid_price_as_percentage_of_on_demand_price}.
	// Experimental.
	BidPriceAsPercentageOfOnDemandPrice *float64 `field:"optional" json:"bidPriceAsPercentageOfOnDemandPrice" yaml:"bidPriceAsPercentageOfOnDemandPrice"`
	// configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#configurations TfCluster#configurations}
	// Experimental.
	Configurations interface{} `field:"optional" json:"configurations" yaml:"configurations"`
	// ebs_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#ebs_config TfCluster#ebs_config}
	// Experimental.
	EbsConfig interface{} `field:"optional" json:"ebsConfig" yaml:"ebsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/emr_cluster#weighted_capacity TfCluster#weighted_capacity}.
	// Experimental.
	WeightedCapacity *float64 `field:"optional" json:"weightedCapacity" yaml:"weightedCapacity"`
}

