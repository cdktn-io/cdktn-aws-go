//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfFunction_DurableConfigPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfFunction_DurableConfigPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfFunction_DurableConfigPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfFunction_DurableConfigPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfFunction_DurableConfigPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfFunction_DurableConfigPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfFunction_DurableConfigPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

