//go:build !no_runtime_type_checking

package awsappmesh

import (
	"fmt"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

func (d *jsiiProxy_DataTfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	if mapKeyAttributeName == nil {
		return fmt.Errorf("parameter mapKeyAttributeName is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataTfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyList) validateGetParameters(index *float64) error {
	if index == nil {
		return fmt.Errorf("parameter index is required, but nil was provided")
	}

	return nil
}

func (d *jsiiProxy_DataTfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	if context == nil {
		return fmt.Errorf("parameter context is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_DataTfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyList) validateSetWrapsSetParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewDataTfVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
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

