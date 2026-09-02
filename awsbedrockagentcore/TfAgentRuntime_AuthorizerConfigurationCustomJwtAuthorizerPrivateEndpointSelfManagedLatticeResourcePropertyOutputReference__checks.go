//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty:
		val := val.(*TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty:
		val_ := val.(TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourceProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateSetResourceConfigurationIdentifierParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfAgentRuntime_AuthorizerConfigurationCustomJwtAuthorizerPrivateEndpointSelfManagedLatticeResourcePropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

