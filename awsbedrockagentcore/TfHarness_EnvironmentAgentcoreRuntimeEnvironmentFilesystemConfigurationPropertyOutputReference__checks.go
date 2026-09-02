//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validatePutEfsAccessPointParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationEfsAccessPointProperty:
		value := value.(*[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationEfsAccessPointProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationEfsAccessPointProperty:
		value_ := value.([]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationEfsAccessPointProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationEfsAccessPointProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validatePutS3FilesAccessPointParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationS3FilesAccessPointProperty:
		value := value.(*[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationS3FilesAccessPointProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationS3FilesAccessPointProperty:
		value_ := value.([]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationS3FilesAccessPointProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationS3FilesAccessPointProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validatePutSessionStorageParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationSessionStorageProperty:
		value := value.(*[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationSessionStorageProperty)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationSessionStorageProperty:
		value_ := value.([]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationSessionStorageProperty)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationSessionStorageProperty; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateSetInternalValueParameters(val interface{}) error {
	switch val.(type) {
	case cdktn.IResolvable:
		// ok
	case *TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationProperty:
		val := val.(*TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationProperty)
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	case TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationProperty:
		val_ := val.(TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationProperty)
		val := &val_
		if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
			return err
		}
	default:
		if !_jsii_.IsAnonymousProxy(val) {
			return fmt.Errorf("parameter val must be one of the allowed types: cdktn.IResolvable, *TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationProperty; received %#v (a %T)", val, val)
		}
	}

	return nil
}

func (j *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfHarness_EnvironmentAgentcoreRuntimeEnvironmentFilesystemConfigurationPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) error {
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

