//go:build !no_runtime_type_checking

package waf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutAllQueryArgumentsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchAllQueryArgumentsProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutBodyParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchBodyProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutCookiesParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchCookiesProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchCookiesProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchCookiesProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchCookiesProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchCookiesProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutHeaderOrderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeaderOrderProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutHeadersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchHeadersProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutJa3FingerprintParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa3FingerprintProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutJa4FingerprintParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJa4FingerprintProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutJsonBodyParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchJsonBodyProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutMethodParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchMethodProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutQueryStringParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchQueryStringProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutSingleHeaderParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleHeaderProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutSingleQueryArgumentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchSingleQueryArgumentProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutUriFragmentParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriFragmentProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validatePutUriPathParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathProperty:
		value := value.(*[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathProperty:
		value_ := value.([]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchUriPathProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchProperty:
		val := val.(*AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchProperty:
		val_ := val.(AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementFieldToMatchPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

