//go:build !no_runtime_type_checking

package waf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutAllowParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty:
		value := value.(*[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty:
		value_ := value.([]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutBlockParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty:
		value := value.(*[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty:
		value_ := value.([]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutCaptchaParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty:
		value := value.(*[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty:
		value_ := value.([]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutChallengeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty:
		value := value.(*[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty:
		value_ := value.([]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutCountParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty:
		value := value.(*[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty:
		value_ := value.([]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty:
		val := val.(*AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty:
		val_ := val.(AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

