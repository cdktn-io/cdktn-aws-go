//go:build no_runtime_type_checking

package awssigner

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataTfSigningJob_SourcePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataTfSigningJob_SourcePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataTfSigningJob_SourcePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataTfSigningJob_SourcePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataTfSigningJob_SourcePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataTfSigningJob_SourcePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataTfSigningJob_SourcePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

