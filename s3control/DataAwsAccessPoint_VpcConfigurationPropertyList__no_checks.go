//go:build no_runtime_type_checking

package s3control

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsAccessPoint_VpcConfigurationPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsAccessPoint_VpcConfigurationPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsAccessPoint_VpcConfigurationPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsAccessPoint_VpcConfigurationPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsAccessPoint_VpcConfigurationPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsAccessPoint_VpcConfigurationPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsAccessPoint_VpcConfigurationPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

