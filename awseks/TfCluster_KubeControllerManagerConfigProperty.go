package awseks


// Experimental.
type TfCluster_KubeControllerManagerConfigProperty struct {
	// horizontal_pod_autoscaler_controller_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#horizontal_pod_autoscaler_controller_config TfCluster#horizontal_pod_autoscaler_controller_config}
	// Experimental.
	HorizontalPodAutoscalerControllerConfig *TfCluster_HorizontalPodAutoscalerControllerConfigProperty `field:"optional" json:"horizontalPodAutoscalerControllerConfig" yaml:"horizontalPodAutoscalerControllerConfig"`
}

