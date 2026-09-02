//go:build no_runtime_type_checking

package awssigner

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TfSigningJob_SignedObjectPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (t *jsiiProxy_TfSigningJob_SignedObjectPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TfSigningJob_SignedObjectPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TfSigningJob_SignedObjectPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TfSigningJob_SignedObjectPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TfSigningJob_SignedObjectPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTfSigningJob_SignedObjectPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

