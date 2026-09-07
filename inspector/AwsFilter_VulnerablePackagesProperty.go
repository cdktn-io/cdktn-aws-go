package inspector


// Experimental.
type AwsFilter_VulnerablePackagesProperty struct {
	// architecture block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#architecture AwsFilter#architecture}
	// Experimental.
	Architecture interface{} `field:"optional" json:"architecture" yaml:"architecture"`
	// epoch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#epoch AwsFilter#epoch}
	// Experimental.
	Epoch interface{} `field:"optional" json:"epoch" yaml:"epoch"`
	// file_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#file_path AwsFilter#file_path}
	// Experimental.
	FilePath interface{} `field:"optional" json:"filePath" yaml:"filePath"`
	// name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#name AwsFilter#name}
	// Experimental.
	Name interface{} `field:"optional" json:"name" yaml:"name"`
	// release block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#release AwsFilter#release}
	// Experimental.
	Release interface{} `field:"optional" json:"release" yaml:"release"`
	// source_lambda_layer_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#source_lambda_layer_arn AwsFilter#source_lambda_layer_arn}
	// Experimental.
	SourceLambdaLayerArn interface{} `field:"optional" json:"sourceLambdaLayerArn" yaml:"sourceLambdaLayerArn"`
	// source_layer_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#source_layer_hash AwsFilter#source_layer_hash}
	// Experimental.
	SourceLayerHash interface{} `field:"optional" json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#version AwsFilter#version}
	// Experimental.
	Version interface{} `field:"optional" json:"version" yaml:"version"`
}

