package awsiotcore


// Experimental.
type AwsIotProvisioningTemplate_PreProvisioningHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_provisioning_template#target_arn AwsIotProvisioningTemplate#target_arn}.
	// Experimental.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_provisioning_template#payload_version AwsIotProvisioningTemplate#payload_version}.
	// Experimental.
	PayloadVersion *string `field:"optional" json:"payloadVersion" yaml:"payloadVersion"`
}

