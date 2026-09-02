package awsinspector


// Experimental.
type TfFilter_FilterCriteriaProperty struct {
	// aws_account_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#aws_account_id TfFilter#aws_account_id}
	// Experimental.
	AwsAccountId interface{} `field:"optional" json:"awsAccountId" yaml:"awsAccountId"`
	// code_repository_project_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#code_repository_project_name TfFilter#code_repository_project_name}
	// Experimental.
	CodeRepositoryProjectName interface{} `field:"optional" json:"codeRepositoryProjectName" yaml:"codeRepositoryProjectName"`
	// code_repository_provider_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#code_repository_provider_type TfFilter#code_repository_provider_type}
	// Experimental.
	CodeRepositoryProviderType interface{} `field:"optional" json:"codeRepositoryProviderType" yaml:"codeRepositoryProviderType"`
	// code_vulnerability_detector_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#code_vulnerability_detector_name TfFilter#code_vulnerability_detector_name}
	// Experimental.
	CodeVulnerabilityDetectorName interface{} `field:"optional" json:"codeVulnerabilityDetectorName" yaml:"codeVulnerabilityDetectorName"`
	// code_vulnerability_detector_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#code_vulnerability_detector_tags TfFilter#code_vulnerability_detector_tags}
	// Experimental.
	CodeVulnerabilityDetectorTags interface{} `field:"optional" json:"codeVulnerabilityDetectorTags" yaml:"codeVulnerabilityDetectorTags"`
	// code_vulnerability_file_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#code_vulnerability_file_path TfFilter#code_vulnerability_file_path}
	// Experimental.
	CodeVulnerabilityFilePath interface{} `field:"optional" json:"codeVulnerabilityFilePath" yaml:"codeVulnerabilityFilePath"`
	// component_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#component_id TfFilter#component_id}
	// Experimental.
	ComponentId interface{} `field:"optional" json:"componentId" yaml:"componentId"`
	// component_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#component_type TfFilter#component_type}
	// Experimental.
	ComponentType interface{} `field:"optional" json:"componentType" yaml:"componentType"`
	// ec2_instance_image_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ec2_instance_image_id TfFilter#ec2_instance_image_id}
	// Experimental.
	Ec2InstanceImageId interface{} `field:"optional" json:"ec2InstanceImageId" yaml:"ec2InstanceImageId"`
	// ec2_instance_subnet_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ec2_instance_subnet_id TfFilter#ec2_instance_subnet_id}
	// Experimental.
	Ec2InstanceSubnetId interface{} `field:"optional" json:"ec2InstanceSubnetId" yaml:"ec2InstanceSubnetId"`
	// ec2_instance_vpc_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ec2_instance_vpc_id TfFilter#ec2_instance_vpc_id}
	// Experimental.
	Ec2InstanceVpcId interface{} `field:"optional" json:"ec2InstanceVpcId" yaml:"ec2InstanceVpcId"`
	// ecr_image_architecture block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ecr_image_architecture TfFilter#ecr_image_architecture}
	// Experimental.
	EcrImageArchitecture interface{} `field:"optional" json:"ecrImageArchitecture" yaml:"ecrImageArchitecture"`
	// ecr_image_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ecr_image_hash TfFilter#ecr_image_hash}
	// Experimental.
	EcrImageHash interface{} `field:"optional" json:"ecrImageHash" yaml:"ecrImageHash"`
	// ecr_image_in_use_count block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ecr_image_in_use_count TfFilter#ecr_image_in_use_count}
	// Experimental.
	EcrImageInUseCount interface{} `field:"optional" json:"ecrImageInUseCount" yaml:"ecrImageInUseCount"`
	// ecr_image_last_in_use_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ecr_image_last_in_use_at TfFilter#ecr_image_last_in_use_at}
	// Experimental.
	EcrImageLastInUseAt interface{} `field:"optional" json:"ecrImageLastInUseAt" yaml:"ecrImageLastInUseAt"`
	// ecr_image_pushed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ecr_image_pushed_at TfFilter#ecr_image_pushed_at}
	// Experimental.
	EcrImagePushedAt interface{} `field:"optional" json:"ecrImagePushedAt" yaml:"ecrImagePushedAt"`
	// ecr_image_registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ecr_image_registry TfFilter#ecr_image_registry}
	// Experimental.
	EcrImageRegistry interface{} `field:"optional" json:"ecrImageRegistry" yaml:"ecrImageRegistry"`
	// ecr_image_repository_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ecr_image_repository_name TfFilter#ecr_image_repository_name}
	// Experimental.
	EcrImageRepositoryName interface{} `field:"optional" json:"ecrImageRepositoryName" yaml:"ecrImageRepositoryName"`
	// ecr_image_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#ecr_image_tags TfFilter#ecr_image_tags}
	// Experimental.
	EcrImageTags interface{} `field:"optional" json:"ecrImageTags" yaml:"ecrImageTags"`
	// epss_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#epss_score TfFilter#epss_score}
	// Experimental.
	EpssScore interface{} `field:"optional" json:"epssScore" yaml:"epssScore"`
	// exploit_available block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#exploit_available TfFilter#exploit_available}
	// Experimental.
	ExploitAvailable interface{} `field:"optional" json:"exploitAvailable" yaml:"exploitAvailable"`
	// finding_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#finding_arn TfFilter#finding_arn}
	// Experimental.
	FindingArn interface{} `field:"optional" json:"findingArn" yaml:"findingArn"`
	// finding_status block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#finding_status TfFilter#finding_status}
	// Experimental.
	FindingStatus interface{} `field:"optional" json:"findingStatus" yaml:"findingStatus"`
	// finding_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#finding_type TfFilter#finding_type}
	// Experimental.
	FindingType interface{} `field:"optional" json:"findingType" yaml:"findingType"`
	// first_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#first_observed_at TfFilter#first_observed_at}
	// Experimental.
	FirstObservedAt interface{} `field:"optional" json:"firstObservedAt" yaml:"firstObservedAt"`
	// fix_available block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#fix_available TfFilter#fix_available}
	// Experimental.
	FixAvailable interface{} `field:"optional" json:"fixAvailable" yaml:"fixAvailable"`
	// inspector_score block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#inspector_score TfFilter#inspector_score}
	// Experimental.
	InspectorScore interface{} `field:"optional" json:"inspectorScore" yaml:"inspectorScore"`
	// lambda_function_execution_role_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#lambda_function_execution_role_arn TfFilter#lambda_function_execution_role_arn}
	// Experimental.
	LambdaFunctionExecutionRoleArn interface{} `field:"optional" json:"lambdaFunctionExecutionRoleArn" yaml:"lambdaFunctionExecutionRoleArn"`
	// lambda_function_last_modified_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#lambda_function_last_modified_at TfFilter#lambda_function_last_modified_at}
	// Experimental.
	LambdaFunctionLastModifiedAt interface{} `field:"optional" json:"lambdaFunctionLastModifiedAt" yaml:"lambdaFunctionLastModifiedAt"`
	// lambda_function_layers block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#lambda_function_layers TfFilter#lambda_function_layers}
	// Experimental.
	LambdaFunctionLayers interface{} `field:"optional" json:"lambdaFunctionLayers" yaml:"lambdaFunctionLayers"`
	// lambda_function_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#lambda_function_name TfFilter#lambda_function_name}
	// Experimental.
	LambdaFunctionName interface{} `field:"optional" json:"lambdaFunctionName" yaml:"lambdaFunctionName"`
	// lambda_function_runtime block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#lambda_function_runtime TfFilter#lambda_function_runtime}
	// Experimental.
	LambdaFunctionRuntime interface{} `field:"optional" json:"lambdaFunctionRuntime" yaml:"lambdaFunctionRuntime"`
	// last_observed_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#last_observed_at TfFilter#last_observed_at}
	// Experimental.
	LastObservedAt interface{} `field:"optional" json:"lastObservedAt" yaml:"lastObservedAt"`
	// network_protocol block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#network_protocol TfFilter#network_protocol}
	// Experimental.
	NetworkProtocol interface{} `field:"optional" json:"networkProtocol" yaml:"networkProtocol"`
	// port_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#port_range TfFilter#port_range}
	// Experimental.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// related_vulnerabilities block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#related_vulnerabilities TfFilter#related_vulnerabilities}
	// Experimental.
	RelatedVulnerabilities interface{} `field:"optional" json:"relatedVulnerabilities" yaml:"relatedVulnerabilities"`
	// resource_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#resource_id TfFilter#resource_id}
	// Experimental.
	ResourceId interface{} `field:"optional" json:"resourceId" yaml:"resourceId"`
	// resource_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#resource_tags TfFilter#resource_tags}
	// Experimental.
	ResourceTags interface{} `field:"optional" json:"resourceTags" yaml:"resourceTags"`
	// resource_type block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#resource_type TfFilter#resource_type}
	// Experimental.
	ResourceType interface{} `field:"optional" json:"resourceType" yaml:"resourceType"`
	// severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#severity TfFilter#severity}
	// Experimental.
	Severity interface{} `field:"optional" json:"severity" yaml:"severity"`
	// title block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#title TfFilter#title}
	// Experimental.
	Title interface{} `field:"optional" json:"title" yaml:"title"`
	// updated_at block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#updated_at TfFilter#updated_at}
	// Experimental.
	UpdatedAt interface{} `field:"optional" json:"updatedAt" yaml:"updatedAt"`
	// vendor_severity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#vendor_severity TfFilter#vendor_severity}
	// Experimental.
	VendorSeverity interface{} `field:"optional" json:"vendorSeverity" yaml:"vendorSeverity"`
	// vulnerability_id block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#vulnerability_id TfFilter#vulnerability_id}
	// Experimental.
	VulnerabilityId interface{} `field:"optional" json:"vulnerabilityId" yaml:"vulnerabilityId"`
	// vulnerability_source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#vulnerability_source TfFilter#vulnerability_source}
	// Experimental.
	VulnerabilitySource interface{} `field:"optional" json:"vulnerabilitySource" yaml:"vulnerabilitySource"`
	// vulnerable_packages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#vulnerable_packages TfFilter#vulnerable_packages}
	// Experimental.
	VulnerablePackages interface{} `field:"optional" json:"vulnerablePackages" yaml:"vulnerablePackages"`
}

