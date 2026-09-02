//go:build !no_runtime_type_checking

package awslexv2models

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validatePutConditionParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchConditionProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchConditionProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchConditionProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchConditionProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchConditionProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validatePutNextStepParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchNextStepProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchNextStepProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchNextStepProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchNextStepProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchNextStepProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validatePutResponseParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchResponseProperty:
		value := value.(*[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchResponseProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchResponseProperty:
		value_ := value.([]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchResponseProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchResponseProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchProperty:
		val := val.(*TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchProperty:
		val_ := val.(TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationTimeoutConditionalConditionalBranchPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

