package awsredshift

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfNamespaceRegistrationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_namespace_registration#consumer_identifier TfNamespaceRegistration#consumer_identifier}.
	// Experimental.
	ConsumerIdentifier *string `field:"required" json:"consumerIdentifier" yaml:"consumerIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_namespace_registration#namespace_type TfNamespaceRegistration#namespace_type}.
	// Experimental.
	NamespaceType *string `field:"required" json:"namespaceType" yaml:"namespaceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_namespace_registration#provisioned_cluster_identifier TfNamespaceRegistration#provisioned_cluster_identifier}.
	// Experimental.
	ProvisionedClusterIdentifier *string `field:"optional" json:"provisionedClusterIdentifier" yaml:"provisionedClusterIdentifier"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_namespace_registration#region TfNamespaceRegistration#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_namespace_registration#serverless_namespace_identifier TfNamespaceRegistration#serverless_namespace_identifier}.
	// Experimental.
	ServerlessNamespaceIdentifier *string `field:"optional" json:"serverlessNamespaceIdentifier" yaml:"serverlessNamespaceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_namespace_registration#serverless_workgroup_identifier TfNamespaceRegistration#serverless_workgroup_identifier}.
	// Experimental.
	ServerlessWorkgroupIdentifier *string `field:"optional" json:"serverlessWorkgroupIdentifier" yaml:"serverlessWorkgroupIdentifier"`
}

