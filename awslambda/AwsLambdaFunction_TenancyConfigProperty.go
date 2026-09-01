package awslambda


// Experimental.
type AwsLambdaFunction_TenancyConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_function#tenant_isolation_mode AwsLambdaFunction#tenant_isolation_mode}.
	// Experimental.
	TenantIsolationMode *string `field:"required" json:"tenantIsolationMode" yaml:"tenantIsolationMode"`
}

