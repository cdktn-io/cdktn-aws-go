//go:build no_runtime_type_checking

package awssigner

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfSigningJob_SignedObjectS3PropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfSigningJob_SignedObjectS3PropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfSigningJob_SignedObjectS3PropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfSigningJob_SignedObjectS3PropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfSigningJob_SignedObjectS3PropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfSigningJob_SignedObjectS3PropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfSigningJob_SignedObjectS3PropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

