//go:build !no_runtime_type_checking

package awssagemakerai

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validatePutAutoParametersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesAutoParametersProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validatePutCategoricalParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesCategoricalParameterRangesProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesCategoricalParameterRangesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesCategoricalParameterRangesProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesCategoricalParameterRangesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesCategoricalParameterRangesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validatePutContinuousParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesContinuousParameterRangesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validatePutIntegerParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesIntegerParameterRangesProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesIntegerParameterRangesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesIntegerParameterRangesProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesIntegerParameterRangesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesIntegerParameterRangesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesProperty:
		val := val.(*TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesProperty:
		val_ := val.(TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfHyperParameterTuningJob_TrainingJobDefinitionsHyperParameterRangesPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

