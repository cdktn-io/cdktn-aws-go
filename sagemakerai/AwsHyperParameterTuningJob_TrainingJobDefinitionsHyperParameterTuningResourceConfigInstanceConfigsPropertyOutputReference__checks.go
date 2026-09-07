//go:build !no_runtime_type_checking

package sagemakerai

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateSetInstanceCountParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateSetInstanceTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsProperty:
		val := val.(*AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsProperty:
		val_ := val.(AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReference) validateSetVolumeSizeInGbParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterTuningResourceConfigInstanceConfigsPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

