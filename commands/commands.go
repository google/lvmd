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

package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	opentracing "github.com/opentracing/opentracing-go"

	"golang.org/x/net/context"

	"github.com/google/lvmd/parser"
)

// ListLV lists lvm volumes
func ListLV(ctx context.Context, listspec string) ([]*parser.LV, error) {
	sp, _ := opentracing.StartSpanFromContext(ctx, "lvm.lvs")
	sp.SetTag("component", "lvm")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	cmd := exec.Command("lvs", "--units=b", "--separator=<:SEP:>", "--nosuffix", "--noheadings",
		"-o", "lv_name,lv_size,lv_uuid,lv_attr,copy_percent,lv_kernel_major,lv_kernel_minor,lv_tags", "--nameprefixes", "-a", listspec)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	outStr := strings.TrimSpace(string(out))
	outLines := strings.Split(outStr, "\n")
	lvs := make([]*parser.LV, len(outLines))
	for i, line := range outLines {
		line = strings.TrimSpace(line)
		lv, err := parser.ParseLV(line)
		if err != nil {
			return nil, err
		}
		lvs[i] = lv
	}
	return lvs, nil
}

// CreateLV creates a new volume
func CreateLV(ctx context.Context, vg string, name string, size uint64, mirrors uint32, tags []string) (string, error) {
	if size == 0 {
		return "", errors.New("size must be greater than 0")
	}

	sp, _ := opentracing.StartSpanFromContext(ctx, "lvm.lvcreate")
	sp.SetTag("component", "lvm")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	args := []string{"-v", "-n", name, "-L", fmt.Sprintf("%db", size)}
	if mirrors > 0 {
		args = append(args, "-m", fmt.Sprintf("%d", mirrors), "--nosync")
	}
	for _, tag := range tags {
		args = append(args, "--add-tag", tag)
	}
	args = append(args, vg)
	cmd := exec.Command("lvcreate", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// ProtectedTagName is a tag that prevents RemoveLV & RemoveVG from removing a volume
const ProtectedTagName = "protected"

// RemoveLV removes a volume
func RemoveLV(ctx context.Context, vg string, name string) (string, error) {
	lvs, err := ListLV(ctx, fmt.Sprintf("%s/%s", vg, name))
	if err != nil {
		return "", fmt.Errorf("failed to list LVs: %v", err)
	}
	if len(lvs) != 1 {
		return "", fmt.Errorf("expected 1 LV, got %d", len(lvs))
	}
	for _, tag := range lvs[0].Tags {
		if tag == ProtectedTagName {
			return "", errors.New("volume is protected")
		}
	}

	sp, _ := opentracing.StartSpanFromContext(ctx, "lvm.lvremove")
	sp.SetTag("component", "lvm")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	cmd := exec.Command("lvremove", "-v", "-f", fmt.Sprintf("%s/%s", vg, name))
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// CloneLV clones a volume via dd
func CloneLV(ctx context.Context, src, dest string) (string, error) {
	// FIXME(farcaller): bloody insecure. And broken.
	sp, _ := opentracing.StartSpanFromContext(ctx, "sys.dd")
	sp.SetTag("component", "dd")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	cmd := exec.Command("dd", fmt.Sprintf("if=%s", src), fmt.Sprintf("of=%s", dest), "bs=4M")
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func ListVG(ctx context.Context) ([]*parser.VG, error) {
	sp, _ := opentracing.StartSpanFromContext(ctx, "lvm.vgs")
	sp.SetTag("component", "lvm")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	cmd := exec.Command("vgs", "--units=b", "--separator=<:SEP:>", "--nosuffix", "--noheadings",
		"-o", "vg_name,vg_size,vg_free,vg_uuid,vg_tags", "--nameprefixes", "-a")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	outStr := strings.TrimSpace(string(out))
	outLines := strings.Split(outStr, "\n")
	vgs := make([]*parser.VG, len(outLines))
	for i, line := range outLines {
		line = strings.TrimSpace(line)
		vg, err := parser.ParseVG(line)
		if err != nil {
			return nil, err
		}
		vgs[i] = vg
	}
	return vgs, nil
}

func CreateVG(ctx context.Context, name string, physicalVolume string, tags []string) (string, error) {
	sp, _ := opentracing.StartSpanFromContext(ctx, "lvm.vgcreate")
	sp.SetTag("component", "lvm")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	args := []string{name, physicalVolume, "-v"}
	for _, tag := range tags {
		args = append(args, "--add-tag", tag)
	}
	cmd := exec.Command("vgcreate", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func RemoveVG(ctx context.Context, name string) (string, error) {
	vgs, err := ListVG(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to list VGs: %v", err)
	}
	var vg *parser.VG
	for _, v := range vgs {
		if v.Name == name {
			vg = v
			break
		}
	}
	if vg == nil {
		return "", fmt.Errorf("could not find vg to delete")
	}
	for _, tag := range vg.Tags {
		if tag == ProtectedTagName {
			return "", errors.New("volume is protected")
		}
	}

	sp, _ := opentracing.StartSpanFromContext(ctx, "lvm.vgremove")
	sp.SetTag("component", "lvm")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	cmd := exec.Command("vgremove", "-v", "-f", name)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func AddTagLV(ctx context.Context, vg string, name string, tags []string) (string, error) {
	sp, _ := opentracing.StartSpanFromContext(ctx, "lvm.addtaglv")
	sp.SetTag("component", "lvm")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	lvs, err := ListLV(ctx, fmt.Sprintf("%s/%s", vg, name))
	if err != nil {
		return "", fmt.Errorf("failed to list LVs: %v", err)
	}
	if len(lvs) != 1 {
		return "", fmt.Errorf("expected 1 LV, got %d", len(lvs))
	}

	args := make([]string, 0)

	for _, tag := range tags {
		args = append(args, "--addtag", tag)
	}

	args = append(args, fmt.Sprintf("%s/%s", vg, name))

	cmd := exec.Command("lvchange", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}

func RemoveTagLV(ctx context.Context, vg string, name string, tags []string) (string, error) {
	sp, _ := opentracing.StartSpanFromContext(ctx, "lvm.removetaglv")
	sp.SetTag("component", "lvm")
	sp.SetTag("span.kind", "client")
	defer sp.Finish()

	lvs, err := ListLV(ctx, fmt.Sprintf("%s/%s", vg, name))
	if err != nil {
		return "", fmt.Errorf("failed to list LVs: %v", err)
	}
	if len(lvs) != 1 {
		return "", fmt.Errorf("expected 1 LV, got %d", len(lvs))
	}

	args := make([]string, 0)

	for _, tag := range tags {
		args = append(args, "--deltag", tag)
	}

	args = append(args, fmt.Sprintf("%s/%s", vg, name))

	cmd := exec.Command("lvchange", args...)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
