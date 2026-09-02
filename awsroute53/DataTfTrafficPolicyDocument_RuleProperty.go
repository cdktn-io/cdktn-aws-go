package awsroute53


// Experimental.
type DataTfTrafficPolicyDocument_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#id DataTfTrafficPolicyDocument#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// geo_proximity_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#geo_proximity_location DataTfTrafficPolicyDocument#geo_proximity_location}
	// Experimental.
	GeoProximityLocation interface{} `field:"optional" json:"geoProximityLocation" yaml:"geoProximityLocation"`
	// items block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#items DataTfTrafficPolicyDocument#items}
	// Experimental.
	Items interface{} `field:"optional" json:"items" yaml:"items"`
	// location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#location DataTfTrafficPolicyDocument#location}
	// Experimental.
	Location interface{} `field:"optional" json:"location" yaml:"location"`
	// primary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#primary DataTfTrafficPolicyDocument#primary}
	// Experimental.
	Primary *DataTfTrafficPolicyDocument_PrimaryProperty `field:"optional" json:"primary" yaml:"primary"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#region DataTfTrafficPolicyDocument#region}
	// Experimental.
	Region interface{} `field:"optional" json:"region" yaml:"region"`
	// secondary block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#secondary DataTfTrafficPolicyDocument#secondary}
	// Experimental.
	Secondary *DataTfTrafficPolicyDocument_SecondaryProperty `field:"optional" json:"secondary" yaml:"secondary"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/route53_traffic_policy_document#type DataTfTrafficPolicyDocument#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

