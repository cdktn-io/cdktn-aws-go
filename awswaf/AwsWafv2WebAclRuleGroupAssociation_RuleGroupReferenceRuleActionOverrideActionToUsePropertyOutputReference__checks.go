//go:build !no_runtime_type_checking

package awswaf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutAllowParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty:
		value := value.(*[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty:
		value_ := value.([]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseAllowProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutBlockParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty:
		value := value.(*[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty:
		value_ := value.([]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutCaptchaParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty:
		value := value.(*[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty:
		value_ := value.([]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCaptchaProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutChallengeParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty:
		value := value.(*[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty:
		value_ := value.([]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseChallengeProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validatePutCountParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty:
		value := value.(*[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty:
		value_ := value.([]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseCountProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty:
		val := val.(*AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty:
		val_ := val.(AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsWafv2WebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUsePropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

