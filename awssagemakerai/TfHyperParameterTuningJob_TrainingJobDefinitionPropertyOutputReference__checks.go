//go:build !no_runtime_type_checking

package awssagemakerai

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutAlgorithmSpecificationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionAlgorithmSpecificationProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutCheckpointConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionCheckpointConfigProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionCheckpointConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionCheckpointConfigProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionCheckpointConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionCheckpointConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutHyperParameterRangesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterRangesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutHyperParameterTuningResourceConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterTuningResourceConfigProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterTuningResourceConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterTuningResourceConfigProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterTuningResourceConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionHyperParameterTuningResourceConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutInputDataConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionInputDataConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutOutputDataConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionOutputDataConfigProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionOutputDataConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionOutputDataConfigProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionOutputDataConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionOutputDataConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutResourceConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionResourceConfigProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionResourceConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionResourceConfigProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionResourceConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionResourceConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutRetryStrategyParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionRetryStrategyProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionRetryStrategyProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionRetryStrategyProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionRetryStrategyProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionRetryStrategyProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutStoppingConditionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionStoppingConditionProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionStoppingConditionProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionStoppingConditionProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionStoppingConditionProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionStoppingConditionProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutTuningObjectiveParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionTuningObjectiveProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionTuningObjectiveProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionTuningObjectiveProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionTuningObjectiveProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionTuningObjectiveProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validatePutVpcConfigParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHyperParameterTuningJob_TrainingJobDefinitionVpcConfigProperty:
		value := value.(*[]*TfHyperParameterTuningJob_TrainingJobDefinitionVpcConfigProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHyperParameterTuningJob_TrainingJobDefinitionVpcConfigProperty:
		value_ := value.([]*TfHyperParameterTuningJob_TrainingJobDefinitionVpcConfigProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHyperParameterTuningJob_TrainingJobDefinitionVpcConfigProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetDefinitionNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetEnableInterContainerTrafficEncryptionParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktn.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetEnableManagedSpotTrainingParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktn.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetEnableNetworkIsolationParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktn.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetEnvironmentParameters(val *map[string]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfHyperParameterTuningJob_TrainingJobDefinitionProperty:
		val := val.(*TfHyperParameterTuningJob_TrainingJobDefinitionProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfHyperParameterTuningJob_TrainingJobDefinitionProperty:
		val_ := val.(TfHyperParameterTuningJob_TrainingJobDefinitionProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfHyperParameterTuningJob_TrainingJobDefinitionProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetRoleArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetStaticHyperParametersParameters(val *map[string]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfHyperParameterTuningJob_TrainingJobDefinitionPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

