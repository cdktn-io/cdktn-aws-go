package eks


// Experimental.
type AwsCluster_KubeControllerManagerConfigProperty struct {
	// horizontal_pod_autoscaler_controller_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#horizontal_pod_autoscaler_controller_config AwsCluster#horizontal_pod_autoscaler_controller_config}
	// Experimental.
	HorizontalPodAutoscalerControllerConfig *AwsCluster_HorizontalPodAutoscalerControllerConfigProperty `field:"optional" json:"horizontalPodAutoscalerControllerConfig" yaml:"horizontalPodAutoscalerControllerConfig"`
}

