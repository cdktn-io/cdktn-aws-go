//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validatePutManagedVpcResourceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourceProperty:
		value := value.(*[]*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourceProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourceProperty:
		value_ := value.([]*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourceProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointManagedVpcResourceProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validatePutSelfManagedLatticeResourceParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty:
		value := value.(*[]*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty:
		value_ := value.([]*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointProperty:
		val := val.(*TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointProperty:
		val_ := val.(TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfGateway_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

