//go:build !no_runtime_type_checking

package awsvpc

import (
	"fmt"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderDestinationPortRangesPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderDestinationPortRangesPropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderDestinationPortRangesPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderDestinationPortRangesPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderDestinationPortRangesPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderDestinationPortRangesPropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfEc2NetworkInsightsAnalysis_ForwardPathComponentsInboundHeaderDestinationPortRangesPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

