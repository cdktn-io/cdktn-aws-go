package awsquicksight

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightIpRestrictionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_ip_restriction#enabled AwsQuicksightIpRestriction#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_ip_restriction#aws_account_id AwsQuicksightIpRestriction#aws_account_id}.
	// Experimental.
	AwsAccountId *string `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_ip_restriction#ip_restriction_rule_map AwsQuicksightIpRestriction#ip_restriction_rule_map}.
	// Experimental.
	IpRestrictionRuleMap *map[string]*string `field:"optional" json:"ipRestrictionRuleMap" yaml:"ipRestrictionRuleMap"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_ip_restriction#region AwsQuicksightIpRestriction#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_ip_restriction#vpc_endpoint_id_restriction_rule_map AwsQuicksightIpRestriction#vpc_endpoint_id_restriction_rule_map}.
	// Experimental.
	VpcEndpointIdRestrictionRuleMap *map[string]*string `field:"optional" json:"vpcEndpointIdRestrictionRuleMap" yaml:"vpcEndpointIdRestrictionRuleMap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_ip_restriction#vpc_id_restriction_rule_map AwsQuicksightIpRestriction#vpc_id_restriction_rule_map}.
	// Experimental.
	VpcIdRestrictionRuleMap *map[string]*string `field:"optional" json:"vpcIdRestrictionRuleMap" yaml:"vpcIdRestrictionRuleMap"`
}

