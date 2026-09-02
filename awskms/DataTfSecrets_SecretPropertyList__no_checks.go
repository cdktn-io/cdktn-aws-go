//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfSecrets_SecretPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfSecrets_SecretPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfSecrets_SecretPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfSecrets_SecretPropertyList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_DataTfSecrets_SecretPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfSecrets_SecretPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfSecrets_SecretPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfSecrets_SecretPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

