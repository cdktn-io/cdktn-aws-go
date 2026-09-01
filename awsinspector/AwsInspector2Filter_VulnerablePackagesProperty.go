package awsinspector


// Experimental.
type AwsInspector2Filter_VulnerablePackagesProperty struct {
	// architecture block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#architecture AwsInspector2Filter#architecture}
	// Experimental.
	Architecture interface{} `field:"optional" json:"architecture" yaml:"architecture"`
	// epoch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#epoch AwsInspector2Filter#epoch}
	// Experimental.
	Epoch interface{} `field:"optional" json:"epoch" yaml:"epoch"`
	// file_path block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#file_path AwsInspector2Filter#file_path}
	// Experimental.
	FilePath interface{} `field:"optional" json:"filePath" yaml:"filePath"`
	// name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#name AwsInspector2Filter#name}
	// Experimental.
	Name interface{} `field:"optional" json:"name" yaml:"name"`
	// release block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#release AwsInspector2Filter#release}
	// Experimental.
	Release interface{} `field:"optional" json:"release" yaml:"release"`
	// source_lambda_layer_arn block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#source_lambda_layer_arn AwsInspector2Filter#source_lambda_layer_arn}
	// Experimental.
	SourceLambdaLayerArn interface{} `field:"optional" json:"sourceLambdaLayerArn" yaml:"sourceLambdaLayerArn"`
	// source_layer_hash block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#source_layer_hash AwsInspector2Filter#source_layer_hash}
	// Experimental.
	SourceLayerHash interface{} `field:"optional" json:"sourceLayerHash" yaml:"sourceLayerHash"`
	// version block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/inspector2_filter#version AwsInspector2Filter#version}
	// Experimental.
	Version interface{} `field:"optional" json:"version" yaml:"version"`
}

