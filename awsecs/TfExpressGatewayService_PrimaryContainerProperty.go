package awsecs


// Experimental.
type TfExpressGatewayService_PrimaryContainerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#image TfExpressGatewayService#image}.
	// Experimental.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#aws_logs_configuration TfExpressGatewayService#aws_logs_configuration}.
	// Experimental.
	AwsLogsConfiguration interface{} `field:"optional" json:"awsLogsConfiguration" yaml:"awsLogsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#command TfExpressGatewayService#command}.
	// Experimental.
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#container_port TfExpressGatewayService#container_port}.
	// Experimental.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// environment block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#environment TfExpressGatewayService#environment}
	// Experimental.
	Environment interface{} `field:"optional" json:"environment" yaml:"environment"`
	// repository_credentials block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#repository_credentials TfExpressGatewayService#repository_credentials}
	// Experimental.
	RepositoryCredentials interface{} `field:"optional" json:"repositoryCredentials" yaml:"repositoryCredentials"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_express_gateway_service#secret TfExpressGatewayService#secret}
	// Experimental.
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
}

