/*

Copyright 2017 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

*/

package parser

import (
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestParseAttrs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Parser Suite")
}

var _ = Describe("Attributes", func() {
	var attrs *LVAttributes
	var err error

	Context("rwi-aor", func() {
		BeforeEach(func() { attrs, err = parseAttrs("rwi-aor") })

		It("should fail to parse", func() {
			Expect(err).ToNot(BeNil())
		})
	})

	Context("rwi-aor---", func() {
		BeforeEach(func() { attrs, err = parseAttrs("rwi-aor---") })

		It("should parse", func() {
			Expect(err).To(BeNil())
		})

		It("should be raid", func() {
			Expect(attrs.Type).To(Equal(VolumeTypeRAID))
		})

		It("should be writable", func() {
			Expect(attrs.Permissions).To(Equal(VolumePermissionsWriteable))
		})

		It("should be allocated with inherited policy", func() {
			Expect(attrs.Allocation).To(Equal(VolumeAllocationInherited))
		})

		It("should not have fixed minor id", func() {
			Expect(attrs.FixedMinor).To(Equal(VolumeFixedMinorDisabled))
		})

		It("should be active", func() {
			Expect(attrs.State).To(Equal(VolumeStateActive))
		})

		It("should be open", func() {
			Expect(attrs.Open).To(Equal(VolumeOpenIsOpen))
		})

		It("should be raid target", func() {
			Expect(attrs.TargetType).To(Equal(VolumeTargetTypeRAID))
		})

		It("should be non-zeroing", func() {
			Expect(attrs.Zeroing).To(Equal(VolumeZeroingIsNonZeroing))
		})

		It("should be in good health", func() {
			Expect(attrs.Health).To(Equal(VolumeHealthOK))
		})

		It("should not be skipped during activation", func() {
			Expect(attrs.ActivationSkipped).To(Equal(VolumeActivationSkippedIsNotSkipped))
		})
	})
})

var _ = Describe("Logical Volume", func() {
	const line = "LVM2_LV_NAME='root'<:SEP:>LVM2_LV_SIZE='10737418240'<:SEP:>LVM2_LV_UUID='v2jVj9-HTY0-G5IR-9zyc-lMqc-iPdg-cSxYs2'<:SEP:>LVM2_LV_ATTR='rwi-aor---'<:SEP:>LVM2_COPY_PERCENT='100.00'<:SEP:>LVM2_LV_KERNEL_MAJOR='252'<:SEP:>LVM2_LV_KERNEL_MINOR='4'<:SEP:>LVM2_LV_TAGS='host,protected'"
	var vol *LV
	var err error

	Context(line, func() {
		BeforeEach(func() { vol, err = ParseLV(line) })

		It("should parse", func() {
			Expect(err).To(BeNil())
		})

		It("should convert to proto", func() {
			pb := vol.ToProto()
			pbText := strings.TrimSpace(proto.MarshalTextString(pb))

			Expect(pbText).To(Equal(strings.TrimSpace(`
name: "root"
size: 10737418240
uuid: "v2jVj9-HTY0-G5IR-9zyc-lMqc-iPdg-cSxYs2"
attributes: <
  type: RAID
  permissions: WRITEABLE
  allocation: INHERITED
  state: ACTIVE
  open: true
  target_type: RAID_TARGET
>
copy_percent: "100.00"
actual_dev_major_number: 252
actual_dev_minor_number: 4
tags: "host"
tags: "protected"
`)))
		})

		It("should have a name", func() {
			Expect(vol.Name).To(Equal("root"))
		})

		It("should have a size", func() {
			Expect(vol.Size).To(Equal(uint64(10737418240)))
		})

		It("should have an UUID", func() {
			Expect(vol.UUID).To(Equal("v2jVj9-HTY0-G5IR-9zyc-lMqc-iPdg-cSxYs2"))
		})

		It("should have copy percentage", func() {
			Expect(vol.CopyPercent).To(Equal("100.00"))
		})

		It("should have an actual major device number", func() {
			Expect(vol.ActualDevMajNumber).To(Equal(uint32(252)))
		})

		It("should have an actual minor device number", func() {
			Expect(vol.ActualDevMinNumber).To(Equal(uint32(4)))
		})

		It("should have tags list", func() {
			Expect(vol.Tags).To(Equal([]string{"host", "protected"}))
		})
	})
})
