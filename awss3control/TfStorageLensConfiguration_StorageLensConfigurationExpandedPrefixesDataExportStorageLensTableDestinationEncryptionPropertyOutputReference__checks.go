//go:build !no_runtime_type_checking

package awss3control

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetStringAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validatePutSseKmsParameters(value *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionSseKmsProperty) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(value, func() string { return "parameter value" }); err != nil {
		return err
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validatePutSseS3Parameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
	case cdktn.IResolvable:
		// ok
	case *[]*TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionSseS3Property:
		value := value.(*[]*TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionSseS3Property)
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	case []*TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionSseS3Property:
		value_ := value.([]*TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionSseS3Property)
		value := &value_
		for idx_cd4240, v := range *value {
			if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter value[%#v]", idx_cd4240) }); err != nil {
				return err
			}
		}
	default:
		if !_jsii_.IsAnonymousProxy(value) {
			return fmt.Errorf("parameter value must be one of the allowed types: cdktn.IResolvable, *[]*TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionSseS3Property; received %#v (a %T)", value, value)
		}
	}

	return nil
}

func (t *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateSetComplexObjectIndexParameters(val interface{}) error {
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

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateSetComplexObjectIsFromSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateSetInternalValueParameters(val *TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionProperty) error {
	if err := _jsii_.ValidateStruct(val, func() string { return "parameter val" }); err != nil {
		return err
	}

	return nil
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_TfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReference) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewTfStorageLensConfiguration_StorageLensConfigurationExpandedPrefixesDataExportStorageLensTableDestinationEncryptionPropertyOutputReferenceParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) error {
	if terraformResource == nil {
		return fmt.Errorf("parameter terraformResource is required, but nil was provided")
	}

	if terraformAttribute == nil {
		return fmt.Errorf("parameter terraformAttribute is required, but nil was provided")
	}

	return nil
}

