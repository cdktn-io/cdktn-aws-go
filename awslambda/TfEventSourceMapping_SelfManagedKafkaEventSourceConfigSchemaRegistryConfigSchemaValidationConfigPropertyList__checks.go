//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigProperty:
		val := val.(*[]*TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigProperty)
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	case []*TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigProperty:
		val_ := val.([]*TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigProperty)
		val := &val_
		for idx_97dfc6, v := range *val {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter val[%#v]", idx_97dfc6) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *[]*TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfEventSourceMapping_SelfManagedKafkaEventSourceConfigSchemaRegistryConfigSchemaValidationConfigPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

