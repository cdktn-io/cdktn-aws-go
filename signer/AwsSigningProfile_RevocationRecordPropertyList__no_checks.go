//go:build no_runtime_type_checking

package signer

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsSigningProfile_RevocationRecordPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsSigningProfile_RevocationRecordPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsSigningProfile_RevocationRecordPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsSigningProfile_RevocationRecordPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsSigningProfile_RevocationRecordPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsSigningProfile_RevocationRecordPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsSigningProfile_RevocationRecordPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

