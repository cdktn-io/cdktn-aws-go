package awsvpclattice


// Experimental.
type TfTargetGroup_ConfigProperty struct {
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group#health_check TfTargetGroup#health_check}
	// Experimental.
	HealthCheck *TfTargetGroup_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group#ip_address_type TfTargetGroup#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group#lambda_event_structure_version TfTargetGroup#lambda_event_structure_version}.
	// Experimental.
	LambdaEventStructureVersion *string `field:"optional" json:"lambdaEventStructureVersion" yaml:"lambdaEventStructureVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group#port TfTargetGroup#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group#protocol TfTargetGroup#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group#protocol_version TfTargetGroup#protocol_version}.
	// Experimental.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_target_group#vpc_identifier TfTargetGroup#vpc_identifier}.
	// Experimental.
	VpcIdentifier *string `field:"optional" json:"vpcIdentifier" yaml:"vpcIdentifier"`
}

