//go:build !no_runtime_type_checking

package awslexv2models

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validatePutDialogActionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepDialogActionProperty:
		value := value.(*[]*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepDialogActionProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepDialogActionProperty:
		value_ := value.([]*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepDialogActionProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepDialogActionProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validatePutIntentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentProperty:
		value := value.(*[]*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentProperty:
		value_ := value.([]*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepIntentProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepProperty:
		val := val.(*AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepProperty:
		val_ := val.(AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateSetSessionAttributesParameters(val *map[string]*string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsLexv2ModelsIntent_InitialResponseSettingCodeHookPostCodeHookSpecificationSuccessNextStepPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

