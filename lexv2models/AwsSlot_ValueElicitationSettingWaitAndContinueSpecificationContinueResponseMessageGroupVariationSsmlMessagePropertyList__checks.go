//go:build !no_runtime_type_checking

package lexv2models

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessagePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessagePropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessagePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessagePropertyList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty:
		val := val.(*[]*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty:
		val_ := val.([]*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessageProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessagePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessagePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessagePropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsSlot_ValueElicitationSettingWaitAndContinueSpecificationContinueResponseMessageGroupVariationSsmlMessagePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

