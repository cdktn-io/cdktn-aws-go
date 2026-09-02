//go:build !no_runtime_type_checking

package awsnetworkmanager

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsProperty:
		val := val.(*[]*DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsProperty)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsProperty:
		val_ := val.([]*DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsProperty)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataTfCoreNetworkPolicyDocument_AttachmentRoutingPolicyRulesConditionsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

