//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfFunction_VpcConfigPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfFunction_VpcConfigPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfFunction_VpcConfigPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfFunction_VpcConfigPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfFunction_VpcConfigPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfFunction_VpcConfigPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfFunction_VpcConfigPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

