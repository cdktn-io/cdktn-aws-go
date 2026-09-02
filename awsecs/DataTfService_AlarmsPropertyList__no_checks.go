//go:build no_runtime_type_checking

package awsecs

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfService_AlarmsPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfService_AlarmsPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfService_AlarmsPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfService_AlarmsPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfService_AlarmsPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfService_AlarmsPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfService_AlarmsPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

