//go:build !no_runtime_type_checking

package awslexv2models

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutFailureConditionalParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureConditionalProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutFailureNextStepParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStepProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStepProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStepProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStepProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureNextStepProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutFailureResponseParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationFailureResponseProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutSuccessConditionalParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessConditionalProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutSuccessNextStepParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutSuccessResponseParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessResponseProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutTimeoutConditionalParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutTimeoutNextStepParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStepProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStepProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStepProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStepProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutNextStepProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validatePutTimeoutResponseParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutResponseProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfIntent_PostFulfillmentStatusSpecificationProperty:
		val := val.(*TfIntent_PostFulfillmentStatusSpecificationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfIntent_PostFulfillmentStatusSpecificationProperty:
		val_ := val.(TfIntent_PostFulfillmentStatusSpecificationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfIntent_PostFulfillmentStatusSpecificationProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfIntent_PostFulfillmentStatusSpecificationPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfIntent_PostFulfillmentStatusSpecificationPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

