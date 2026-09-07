//go:build no_runtime_type_checking

package signer

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataAwsSigningJob_SourcePropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataAwsSigningJob_SourcePropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataAwsSigningJob_SourcePropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataAwsSigningJob_SourcePropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataAwsSigningJob_SourcePropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataAwsSigningJob_SourcePropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataAwsSigningJob_SourcePropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

