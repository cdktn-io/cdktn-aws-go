//go:build no_runtime_type_checking

package signer

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsSigningJob_SignedObjectPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsSigningJob_SignedObjectPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsSigningJob_SignedObjectPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsSigningJob_SignedObjectPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsSigningJob_SignedObjectPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsSigningJob_SignedObjectPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsSigningJob_SignedObjectPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

