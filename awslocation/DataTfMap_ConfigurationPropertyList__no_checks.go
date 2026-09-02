//go:build no_runtime_type_checking

package awslocation

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfMap_ConfigurationPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfMap_ConfigurationPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfMap_ConfigurationPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfMap_ConfigurationPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfMap_ConfigurationPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfMap_ConfigurationPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfMap_ConfigurationPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

