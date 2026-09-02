//go:build !no_runtime_type_checking

package awswaf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validatePutResponseHeaderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderProperty:
		value := value.(*[]*TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderProperty:
		value_ := value.([]*TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseResponseHeaderProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateSetCustomResponseBodyKeyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseProperty:
		val := val.(*TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseProperty:
		val_ := val.(TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponseProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateSetResponseCodeParameters(val *float64) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfWebAclRuleGroupAssociation_RuleGroupReferenceRuleActionOverrideActionToUseBlockCustomResponsePropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

