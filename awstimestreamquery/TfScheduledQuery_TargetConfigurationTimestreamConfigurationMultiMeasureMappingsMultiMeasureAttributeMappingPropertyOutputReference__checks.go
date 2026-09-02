//go:build !no_runtime_type_checking

package awstimestreamquery

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingProperty:
		val := val.(*TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingProperty:
		val_ := val.(TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateSetMeasureValueTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateSetSourceColumnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateSetTargetMultiMeasureAttributeNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfScheduledQuery_TargetConfigurationTimestreamConfigurationMultiMeasureMappingsMultiMeasureAttributeMappingPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

