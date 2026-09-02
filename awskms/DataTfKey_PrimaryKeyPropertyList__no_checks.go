//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfKey_PrimaryKeyPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfKey_PrimaryKeyPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfKey_PrimaryKeyPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfKey_PrimaryKeyPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfKey_PrimaryKeyPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfKey_PrimaryKeyPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfKey_PrimaryKeyPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

