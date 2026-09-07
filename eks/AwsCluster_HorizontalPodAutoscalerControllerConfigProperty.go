package eks


// Experimental.
type AwsCluster_HorizontalPodAutoscalerControllerConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/eks_cluster#horizontal_pod_autoscaler_sync_period AwsCluster#horizontal_pod_autoscaler_sync_period}.
	// Experimental.
	HorizontalPodAutoscalerSyncPeriod *string `field:"optional" json:"horizontalPodAutoscalerSyncPeriod" yaml:"horizontalPodAutoscalerSyncPeriod"`
}

