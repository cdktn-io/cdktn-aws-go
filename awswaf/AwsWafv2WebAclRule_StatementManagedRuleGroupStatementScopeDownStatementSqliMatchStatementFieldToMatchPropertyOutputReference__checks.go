//go:build !no_runtime_type_checking

package awswaf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutAllQueryArgumentsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchAllQueryArgumentsProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchAllQueryArgumentsProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchAllQueryArgumentsProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchAllQueryArgumentsProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchAllQueryArgumentsProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutBodyParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchBodyProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchBodyProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchBodyProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchBodyProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchBodyProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutCookiesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchCookiesProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchCookiesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchCookiesProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchCookiesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchCookiesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutHeaderOrderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeaderOrderProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeaderOrderProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeaderOrderProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeaderOrderProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeaderOrderProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutHeadersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeadersProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeadersProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeadersProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeadersProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchHeadersProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutJa3FingerprintParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa3FingerprintProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa3FingerprintProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa3FingerprintProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa3FingerprintProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa3FingerprintProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutJa4FingerprintParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa4FingerprintProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa4FingerprintProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa4FingerprintProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa4FingerprintProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJa4FingerprintProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutJsonBodyParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJsonBodyProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJsonBodyProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJsonBodyProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJsonBodyProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchJsonBodyProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutMethodParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchMethodProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchMethodProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchMethodProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchMethodProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchMethodProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutQueryStringParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchQueryStringProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchQueryStringProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchQueryStringProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchQueryStringProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchQueryStringProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutSingleHeaderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleHeaderProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleHeaderProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleHeaderProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleHeaderProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleHeaderProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutSingleQueryArgumentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleQueryArgumentProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleQueryArgumentProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleQueryArgumentProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleQueryArgumentProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchSingleQueryArgumentProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutUriFragmentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriFragmentProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriFragmentProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriFragmentProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriFragmentProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriFragmentProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validatePutUriPathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriPathProperty:
		value := value.(*[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriPathProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriPathProperty:
		value_ := value.([]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriPathProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchUriPathProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchProperty:
		val := val.(*AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchProperty:
		val_ := val.(AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsWafv2WebAclRule_StatementManagedRuleGroupStatementScopeDownStatementSqliMatchStatementFieldToMatchPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

