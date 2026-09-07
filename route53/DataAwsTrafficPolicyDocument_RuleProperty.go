package route53


// Experimental.
type DataAwsTrafficPolicyDocument_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#id DataAwsTrafficPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// geo_proximity_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#geo_proximity_location DataAwsTrafficPolicyDocument#geo_proximity_location}
	// Experimental.
	GeoProximityLocation interface{} `field:"optional" json:"geoProximityLocation" yaml:"geoProximityLocation"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#items DataAwsTrafficPolicyDocument#items}
	// Experimental.
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#location DataAwsTrafficPolicyDocument#location}
	// Experimental.
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#primary DataAwsTrafficPolicyDocument#primary}
	// Experimental.
	Primary *DataAwsTrafficPolicyDocument_PrimaryProperty `field:"optional" json:"primary" yaml:"primary"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#region DataAwsTrafficPolicyDocument#region}
	// Experimental.
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// secondary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#secondary DataAwsTrafficPolicyDocument#secondary}
	// Experimental.
	Secondary *DataAwsTrafficPolicyDocument_SecondaryProperty `field:"optional" json:"secondary" yaml:"secondary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#type DataAwsTrafficPolicyDocument#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

