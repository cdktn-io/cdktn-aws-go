package awsecrpublic


// Experimental.
type DataTfImages_ImageIdsProperty struct {
	// Image digest.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecrpublic_images#image_digest DataTfImages#image_digest}
	// Experimental.
	ImageDigest *string `field:"optional" json:"imageDigest" yaml:"imageDigest"`
	// Image tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ecrpublic_images#image_tag DataTfImages#image_tag}
	// Experimental.
	ImageTag *string `field:"optional" json:"imageTag" yaml:"imageTag"`
}

