//go:build !no_runtime_type_checking

package awssagemakerai

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validatePutAutoParametersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesAutoParametersProperty:
		value := value.(*[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesAutoParametersProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesAutoParametersProperty:
		value_ := value.([]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesAutoParametersProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesAutoParametersProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validatePutCategoricalParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesCategoricalParameterRangesProperty:
		value := value.(*[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesCategoricalParameterRangesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesCategoricalParameterRangesProperty:
		value_ := value.([]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesCategoricalParameterRangesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesCategoricalParameterRangesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validatePutContinuousParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesContinuousParameterRangesProperty:
		value := value.(*[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesContinuousParameterRangesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesContinuousParameterRangesProperty:
		value_ := value.([]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesContinuousParameterRangesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesContinuousParameterRangesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validatePutIntegerParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesIntegerParameterRangesProperty:
		value := value.(*[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesIntegerParameterRangesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesIntegerParameterRangesProperty:
		value_ := value.([]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesIntegerParameterRangesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesIntegerParameterRangesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty:
		val := val.(*AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty:
		val_ := val.(AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsSagemakerHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

