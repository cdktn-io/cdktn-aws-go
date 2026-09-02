//go:build !no_runtime_type_checking

package awswaf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateSetFallbackBehaviorParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentProperty:
		val := val.(*TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentProperty:
		val_ := val.(TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexMatchStatementFieldToMatchUriFragmentPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

