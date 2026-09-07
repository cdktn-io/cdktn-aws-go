//go:build !no_runtime_type_checking

package lexv2models

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validatePutSlotParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentSlotProperty:
		value := value.(*[]*AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentSlotProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentSlotProperty:
		value_ := value.([]*AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentSlotProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentSlotProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentProperty:
		val := val.(*AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentProperty:
		val_ := val.(AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateSetNameParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsIntent_FulfillmentCodeHookPostFulfillmentStatusSpecificationSuccessNextStepIntentPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

