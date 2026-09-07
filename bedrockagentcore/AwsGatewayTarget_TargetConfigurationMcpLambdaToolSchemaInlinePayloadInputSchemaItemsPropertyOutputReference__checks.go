//go:build !no_runtime_type_checking

package bedrockagentcore

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validatePutItemsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsItemsProperty:
		value := value.(*[]*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsItemsProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsItemsProperty:
		value_ := value.([]*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsItemsProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsItemsProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validatePutPropertyParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyProperty:
		value := value.(*[]*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyProperty:
		value_ := value.([]*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateSetDescriptionParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsProperty:
		val := val.(*AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsProperty:
		val_ := val.(AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReference) validateSetTypeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsGatewayTarget_TargetConfigurationMcpLambdaToolSchemaInlinePayloadInputSchemaItemsPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

