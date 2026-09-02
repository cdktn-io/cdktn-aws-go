//go:build !no_runtime_type_checking

package awseks

import (
	"fmt"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeControllerManagerConfigPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeControllerManagerConfigPropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeControllerManagerConfigPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeControllerManagerConfigPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeControllerManagerConfigPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeControllerManagerConfigPropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataTfClusterVersions_ClusterVersionsControlPlaneScalingTiersControlPlaneComponentConfigOverridesKubeControllerManagerConfigPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

