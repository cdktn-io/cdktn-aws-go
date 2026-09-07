//go:build !no_runtime_type_checking

package lexv2models

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty:
		val := val.(*[]*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty:
		val_ := val.([]*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	if wrapsSet == nil {
		return fmt.Errorf("parameter wrapsSet is required, but nil was provided")
	}

	return nil
}

