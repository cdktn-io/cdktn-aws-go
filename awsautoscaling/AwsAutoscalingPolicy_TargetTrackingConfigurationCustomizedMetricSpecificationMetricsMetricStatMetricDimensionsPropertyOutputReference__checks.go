//go:build !no_runtime_type_checking

package awsautoscaling

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
	switch val.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter val must be one of the allowed types: *string, *float64; received %#v (a %T)", val, val)
	}

	return nil
}

func (j *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsProperty:
		val := val.(*AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsProperty:
		val_ := val.(AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReference) validateSetValueParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsAutoscalingPolicy_TargetTrackingConfigurationCustomizedMetricSpecificationMetricsMetricStatMetricDimensionsPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if complexObjectIndex == nil {
		return fmt.Errorf("parameter complexObjectIndex is required, but nil was provided")
	}

	if complexObjectIsFromSet == nil {
		return fmt.Errorf("parameter complexObjectIsFromSet is required, but nil was provided")
	}

	return nil
}

