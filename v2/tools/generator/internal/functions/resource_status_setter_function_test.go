/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package functions

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/astmodel"
	"github.com/Azure/azure-service-operator/v2/tools/generator/internal/test"
)

// TestGolden_NewResourceStatusSetterFunction_GeneratesExpectedCode tests the code generated by the SetStatus() function
func TestGolden_NewResourceStatusSetterFunction_GeneratesExpectedCode(t *testing.T) {
	g := NewGomegaWithT(t)
	idFactory := astmodel.NewIdentifierFactory()

	// Create our upstream type
	personSpec2020 := test.CreateSpec(test.Pkg2020, "Person", test.FullNameProperty, test.KnownAsProperty, test.FamilyNameProperty)
	personStatus2020 := test.CreateStatus(test.Pkg2020, "Person")
	person2020 := test.CreateResource(test.Pkg2020, "Person", personSpec2020, personStatus2020)

	// Create Resource Conversion Functions
	rt, ok := astmodel.AsResourceType(person2020.Type())
	g.Expect(ok).To(BeTrue())

	setStatus := NewResourceStatusSetterFunction(rt, idFactory)
	injector := astmodel.NewFunctionInjector()
	modified, err := injector.Inject(person2020, setStatus)
	g.Expect(err).To(Succeed())

	test.AssertSingleTypeDefinitionGeneratesExpectedCode(t, "NewResourceStatusSetterFunction", modified, test.DiffWith(person2020))
}
