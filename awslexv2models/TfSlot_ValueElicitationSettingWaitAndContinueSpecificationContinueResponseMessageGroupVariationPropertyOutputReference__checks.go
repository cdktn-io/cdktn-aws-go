//go:build !no_runtime_type_checking

package awslexv2models

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validatePutCustomPayloadParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationCustomPayloadProperty:
		value := value.(*[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationCustomPayloadProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationCustomPayloadProperty:
		value_ := value.([]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationCustomPayloadProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationCustomPayloadProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validatePutImageResponseCardParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationImageResponseCardProperty:
		value := value.(*[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationImageResponseCardProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationImageResponseCardProperty:
		value_ := value.([]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationImageResponseCardProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationImageResponseCardProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validatePutPlainTextMessageParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPlainTextMessageProperty:
		value := value.(*[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPlainTextMessageProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPlainTextMessageProperty:
		value_ := value.([]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPlainTextMessageProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPlainTextMessageProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validatePutSsmlMessageParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty:
		value := value.(*[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty:
		value_ := value.([]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty:
		val := val.(*TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty:
		val_ := val.(TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

