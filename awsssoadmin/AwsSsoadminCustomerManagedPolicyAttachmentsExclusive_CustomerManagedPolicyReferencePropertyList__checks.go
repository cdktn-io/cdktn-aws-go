//go:build !no_runtime_type_checking

package awsssoadmin

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferencePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferencePropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferencePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferencePropertyList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferenceProperty:
		val := val.(*[]*AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferenceProperty)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferenceProperty:
		val_ := val.([]*AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferenceProperty)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferenceProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferencePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferencePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferencePropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsSsoadminCustomerManagedPolicyAttachmentsExclusive_CustomerManagedPolicyReferencePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

