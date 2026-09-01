//go:build !no_runtime_type_checking

package awskinesisfirehose

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validatePutParametersParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsParametersProperty:
		value := value.(*[]*AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsParametersProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsParametersProperty:
		value_ := value.([]*AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsParametersProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsParametersProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsProperty:
		val := val.(*AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsProperty:
		val_ := val.(AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReference) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsKinesisFirehoseDeliveryStream_HttpEndpointConfigurationProcessingConfigurationProcessorsPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

