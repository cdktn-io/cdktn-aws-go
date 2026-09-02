//go:build !no_runtime_type_checking

package awswaf

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationPropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationPropertyList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationProperty:
		val := val.(*[]*TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationProperty)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationProperty:
		val_ := val.([]*TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationProperty)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationPropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfWebAclRule_StatementRateBasedStatementScopeDownStatementRegexPatternSetReferenceStatementTextTransformationPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

