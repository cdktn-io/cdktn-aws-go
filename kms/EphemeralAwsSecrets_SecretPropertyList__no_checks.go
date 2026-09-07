//go:build no_runtime_type_checking

package kms

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_EphemeralAwsSecrets_SecretPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (e *jsiiProxy_EphemeralAwsSecrets_SecretPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (e *jsiiProxy_EphemeralAwsSecrets_SecretPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_EphemeralAwsSecrets_SecretPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_EphemeralAwsSecrets_SecretPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_EphemeralAwsSecrets_SecretPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_EphemeralAwsSecrets_SecretPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewEphemeralAwsSecrets_SecretPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

