package directoryservice

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRadiusSettingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#authentication_protocol AwsRadiusSettings#authentication_protocol}.
	// Experimental.
	AuthenticationProtocol *string `field:"required" json:"authenticationProtocol" yaml:"authenticationProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#directory_id AwsRadiusSettings#directory_id}.
	// Experimental.
	DirectoryId *string `field:"required" json:"directoryId" yaml:"directoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#display_label AwsRadiusSettings#display_label}.
	// Experimental.
	DisplayLabel *string `field:"required" json:"displayLabel" yaml:"displayLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#radius_port AwsRadiusSettings#radius_port}.
	// Experimental.
	RadiusPort *float64 `field:"required" json:"radiusPort" yaml:"radiusPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#radius_retries AwsRadiusSettings#radius_retries}.
	// Experimental.
	RadiusRetries *float64 `field:"required" json:"radiusRetries" yaml:"radiusRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#radius_servers AwsRadiusSettings#radius_servers}.
	// Experimental.
	RadiusServers *[]*string `field:"required" json:"radiusServers" yaml:"radiusServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#radius_timeout AwsRadiusSettings#radius_timeout}.
	// Experimental.
	RadiusTimeout *float64 `field:"required" json:"radiusTimeout" yaml:"radiusTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#shared_secret AwsRadiusSettings#shared_secret}.
	// Experimental.
	SharedSecret *string `field:"required" json:"sharedSecret" yaml:"sharedSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#id AwsRadiusSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#region AwsRadiusSettings#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#timeouts AwsRadiusSettings#timeouts}
	// Experimental.
	Timeouts *AwsRadiusSettings_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/directory_service_radius_settings#use_same_username AwsRadiusSettings#use_same_username}.
	// Experimental.
	UseSameUsername interface{} `field:"optional" json:"useSameUsername" yaml:"useSameUsername"`
}

