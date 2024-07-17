// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	hostname "github.com/yosepkim/provider-akamai/internal/controller/edgehostname/hostname"
	domain "github.com/yosepkim/provider-akamai/internal/controller/gtm/domain"
	property "github.com/yosepkim/provider-akamai/internal/controller/property/property"
	providerconfig "github.com/yosepkim/provider-akamai/internal/controller/providerconfig"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		hostname.Setup,
		domain.Setup,
		property.Setup,
		providerconfig.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
