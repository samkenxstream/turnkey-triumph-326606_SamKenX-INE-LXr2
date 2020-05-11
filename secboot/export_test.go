// -*- Mode: Go; indent-tabs-mode: t -*-
// +build !nosecboot

/*
 * Copyright (C) 2020 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package secboot

import (
	"io"

	sb "github.com/snapcore/secboot"

	"github.com/snapcore/snapd/asserts"
)

type TPMSupport = tpmSupport

func MockSbConnectToDefaultTPM(f func() (*sb.TPMConnection, error)) (restore func()) {
	old := sbConnectToDefaultTPM
	sbConnectToDefaultTPM = f
	return func() {
		sbConnectToDefaultTPM = old
	}
}

func MockSbProvisionTPM(f func(tpm *sb.TPMConnection, mode sb.ProvisionMode, newLockoutAuth []byte) error) (restore func()) {
	old := sbProvisionTPM
	sbProvisionTPM = f
	return func() {
		sbProvisionTPM = old
	}
}

func MockSbAddEFISecureBootPolicyProfile(f func(profile *sb.PCRProtectionProfile, params *sb.EFISecureBootPolicyProfileParams) error) (restore func()) {
	old := sbAddEFISecureBootPolicyProfile
	sbAddEFISecureBootPolicyProfile = f
	return func() {
		sbAddEFISecureBootPolicyProfile = old
	}
}

func MockSbAddSystemdEFIStubProfile(f func(profile *sb.PCRProtectionProfile, params *sb.SystemdEFIStubProfileParams) error) (restore func()) {
	old := sbAddSystemdEFIStubProfile
	sbAddSystemdEFIStubProfile = f
	return func() {
		sbAddSystemdEFIStubProfile = old
	}
}

func MockSbAddSnapModelProfile(f func(profile *sb.PCRProtectionProfile, params *sb.SnapModelProfileParams) error) (restore func()) {
	old := sbAddSnapModelProfile
	sbAddSnapModelProfile = f
	return func() {
		sbAddSnapModelProfile = old
	}
}

func MockSbSealKeyToTPM(f func(tpm *sb.TPMConnection, key []byte, keyPath, policyUpdatePath string, params *sb.KeyCreationParams) error) (restore func()) {
	old := sbSealKeyToTPM
	sbSealKeyToTPM = f
	return func() {
		sbSealKeyToTPM = old
	}
}

func MockSbLockAccessToSealedKeys(f func(tpm *sb.TPMConnection) error) (restore func()) {
	old := sbLockAccessToSealedKeys
	sbLockAccessToSealedKeys = f
	return func() {
		sbLockAccessToSealedKeys = old
	}
}

func MockSbActivateVolumeWithTPMSealedKey(f func(tpm *sb.TPMConnection, volumeName, sourceDevicePath, keyPath string,
	pinReader io.Reader, options *sb.ActivateWithTPMSealedKeyOptions) (bool, error)) (restore func()) {
	old := sbActivateVolumeWithTPMSealedKey
	sbActivateVolumeWithTPMSealedKey = f
	return func() {
		sbActivateVolumeWithTPMSealedKey = old
	}
}

func MockSbMeasureSnapSystemEpochToTPM(f func(tpm *sb.TPMConnection, pcrIndex int) error) (restore func()) {
	old := sbMeasureSnapSystemEpochToTPM
	sbMeasureSnapSystemEpochToTPM = f
	return func() {
		sbMeasureSnapSystemEpochToTPM = old
	}
}

func MockSbMeasureSnapModelToTPM(f func(tpm *sb.TPMConnection, pcrIndex int, model *asserts.Model) error) (restore func()) {
	old := sbMeasureSnapModelToTPM
	sbMeasureSnapModelToTPM = f
	return func() {
		sbMeasureSnapModelToTPM = old
	}
}

func MockDevDiskByLabelDir(new string) (restore func()) {
	old := devDiskByLabelDir
	devDiskByLabelDir = new
	return func() {
		devDiskByLabelDir = old
	}
}