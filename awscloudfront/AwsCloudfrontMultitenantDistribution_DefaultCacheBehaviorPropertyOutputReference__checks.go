//go:build !no_runtime_type_checking

package awscloudfront

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validatePutAllowedMethodsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsProperty:
		value := value.(*[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsProperty:
		value_ := value.([]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorAllowedMethodsProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validatePutFunctionAssociationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorFunctionAssociationProperty:
		value := value.(*[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorFunctionAssociationProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorFunctionAssociationProperty:
		value_ := value.([]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorFunctionAssociationProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorFunctionAssociationProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validatePutLambdaFunctionAssociationParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationProperty:
		value := value.(*[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationProperty:
		value_ := value.([]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorLambdaFunctionAssociationProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validatePutTrustedKeyGroupsParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsProperty:
		value := value.(*[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsProperty:
		value_ := value.([]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorTrustedKeyGroupsProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (a *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetCachePolicyIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetCompressParameters(val interface{}) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}
	switch val.(type) {
	case *bool:
		// ok
	case bool:
		// ok
	case cdktn.IResolvable:
		// ok
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: *bool, cdktn.IResolvable; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetFieldLevelEncryptionIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorProperty:
		val := val.(*AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorProperty:
		val_ := val.(AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetOriginRequestPolicyIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetRealtimeLogConfigArnParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetResponseHeadersPolicyIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetTargetOriginIdParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_AwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReference) validateSetViewerProtocolPolicyParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewAwsCloudfrontMultitenantDistribution_DefaultCacheBehaviorPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

