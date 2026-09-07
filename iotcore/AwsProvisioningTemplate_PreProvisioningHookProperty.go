package iotcore


// Experimental.
type AwsProvisioningTemplate_PreProvisioningHookProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_provisioning_template#target_arn AwsProvisioningTemplate#target_arn}.
	// Experimental.
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_provisioning_template#payload_version AwsProvisioningTemplate#payload_version}.
	// Experimental.
	PayloadVersion *string `field:"optional" json:"payloadVersion" yaml:"payloadVersion"`
}

