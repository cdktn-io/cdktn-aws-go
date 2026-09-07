package fsx


// Experimental.
type AwsFileCache_DataRepositoryAssociationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#data_repository_path AwsFileCache#data_repository_path}.
	// Experimental.
	DataRepositoryPath *string `field:"required" json:"dataRepositoryPath" yaml:"dataRepositoryPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#file_cache_path AwsFileCache#file_cache_path}.
	// Experimental.
	FileCachePath *string `field:"required" json:"fileCachePath" yaml:"fileCachePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#data_repository_subdirectories AwsFileCache#data_repository_subdirectories}.
	// Experimental.
	DataRepositorySubdirectories *[]*string `field:"optional" json:"dataRepositorySubdirectories" yaml:"dataRepositorySubdirectories"`
	// nfs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#nfs AwsFileCache#nfs}
	// Experimental.
	Nfs interface{} `field:"optional" json:"nfs" yaml:"nfs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_file_cache#tags AwsFileCache#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}

