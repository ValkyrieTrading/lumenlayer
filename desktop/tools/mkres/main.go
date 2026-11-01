// mkres writes a Windows resource object (.syso) holding the app icon and
// version information, using only the Go standard library. `go build` links
// any *.syso file in the package directory into the executable.
//
//	go run ./tools/mkres -ico icon.ico -arch amd64 -o resource_windows_amd64.syso
package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode/utf16"
)

var le = binary.LittleEndian

type resource struct {
	typ, id uint16
	data    []byte
}

func main() {
	icoPath := flag.String("ico", "icon.ico", "icon file")
	arch := flag.String("arch", "amd64", "amd64 or arm64")
	out := flag.String("o", "resource.syso", "output .syso")
	version := flag.String("version", "1.0.0.0", "file and product version")
	name := flag.String("name", "Lumen Layer Studio", "product name")
	exe := flag.String("exe", "LumenLayerStudio.exe", "original file name")
	flag.Parse()

	ico, err := os.ReadFile(*icoPath)
	check(err)
	images, group := parseICO(ico)
	var res []resource
	for i, img := range images {
		res = append(res, resource{3, uint16(i + 1), img}) // RT_ICON
	}
	res = append(res, resource{14, 1, group})                              // RT_GROUP_ICON
	res = append(res, resource{16, 1, versionInfo(*version, *name, *exe)}) // RT_VERSION

	machine, relType := uint16(0x8664), uint16(3) // IMAGE_FILE_MACHINE_AMD64, IMAGE_REL_AMD64_ADDR32NB
	if *arch == "arm64" {
		machine, relType = 0xAA64, 2 // IMAGE_FILE_MACHINE_ARM64, IMAGE_REL_ARM64_ADDR32NB
	}
	sec, relocs := layout(res)
	check(os.WriteFile(*out, coff(machine, relType, sec, relocs), 0o644))
	fmt.Printf("%s: %d icon images, %d bytes\n", *out, len(images), len(sec))
}

// parseICO splits an .ico file into its images and builds the RT_GROUP_ICON directory.
func parseICO(b []byte) (images [][]byte, group []byte) {
	if len(b) < 6 || le.Uint16(b[2:]) != 1 {
		check(fmt.Errorf("not an .ico file"))
	}
	n := int(le.Uint16(b[4:]))
	group = make([]byte, 6, 6+14*n)
	le.PutUint16(group[2:], 1)
	le.PutUint16(group[4:], uint16(n))
	for i := 0; i < n; i++ {
		e := b[6+16*i:]
		size, off := le.Uint32(e[8:]), le.Uint32(e[12:])
		images = append(images, b[off:off+size])
		planes, bpp := le.Uint16(e[4:]), le.Uint16(e[6:])
		if planes == 0 {
			planes = 1
		}
		if bpp == 0 {
			bpp = 32
		}
		g := make([]byte, 14)
		copy(g[0:4], e[0:4]) // width, height, colour count, reserved
		le.PutUint16(g[4:], planes)
		le.PutUint16(g[6:], bpp)
		le.PutUint32(g[8:], size)
		le.PutUint16(g[12:], uint16(i+1))
		group = append(group, g...)
	}
	return
}

// layout builds the three-level resource directory (type → id → language),
// the data entries and the data. It returns the section bytes and the offsets
// of the data-entry RVA fields that the linker must relocate.
func layout(res []resource) ([]byte, []uint32) {
	sort.Slice(res, func(i, j int) bool {
		if res[i].typ != res[j].typ {
			return res[i].typ < res[j].typ
		}
		return res[i].id < res[j].id
	})
	var types []uint16
	byType := map[uint16][]resource{}
	for _, r := range res {
		if len(byType[r.typ]) == 0 {
			types = append(types, r.typ)
		}
		byType[r.typ] = append(byType[r.typ], r)
	}
	dirSize := func(n int) uint32 { return 16 + 8*uint32(n) }
	off := dirSize(len(types))
	typeDir := map[uint16]uint32{}
	for _, t := range types {
		typeDir[t] = off
		off += dirSize(len(byType[t]))
	}
	idDir := make([]uint32, len(res))
	for i := range res {
		idDir[i] = off
		off += dirSize(1)
	}
	entry := make([]uint32, len(res))
	for i := range res {
		entry[i] = off
		off += 16
	}
	dataOff := make([]uint32, len(res))
	for i, r := range res {
		off = (off + 7) &^ 7
		dataOff[i] = off
		off += uint32(len(r.data))
	}
	sec := make([]byte, (off+7)&^7)
	writeDir := func(at uint32, n int) {
		le.PutUint16(sec[at+14:], uint16(n)) // NumberOfIdEntries
	}
	writeDir(0, len(types))
	for i, t := range types {
		le.PutUint32(sec[16+8*i:], uint32(t))
		le.PutUint32(sec[16+8*i+4:], 0x80000000|typeDir[t])
	}
	idx := 0
	for _, t := range types {
		writeDir(typeDir[t], len(byType[t]))
		for j, r := range byType[t] {
			p := typeDir[t] + 16 + 8*uint32(j)
			le.PutUint32(sec[p:], uint32(r.id))
			le.PutUint32(sec[p+4:], 0x80000000|idDir[idx+j])
		}
		idx += len(byType[t])
	}
	var relocs []uint32
	for i, r := range res {
		writeDir(idDir[i], 1)
		le.PutUint32(sec[idDir[i]+16:], 0x0409) // en-US
		le.PutUint32(sec[idDir[i]+20:], entry[i])
		le.PutUint32(sec[entry[i]:], dataOff[i]) // RVA once relocated
		le.PutUint32(sec[entry[i]+4:], uint32(len(r.data)))
		relocs = append(relocs, entry[i])
		copy(sec[dataOff[i]:], r.data)
	}
	return sec, relocs
}

// coff wraps the section in a minimal COFF object file with one section symbol.
func coff(machine, relType uint16, sec []byte, relocs []uint32) []byte {
	const hdr, shdr = 20, 40
	relOff := uint32(hdr + shdr + len(sec))
	symOff := relOff + 10*uint32(len(relocs))
	b := make([]byte, int(symOff)+18+4)
	le.PutUint16(b[0:], machine)
	le.PutUint16(b[2:], 1)
	le.PutUint32(b[8:], symOff)
	le.PutUint32(b[12:], 1)
	s := b[hdr:]
	copy(s[0:8], ".rsrc")
	le.PutUint32(s[16:], uint32(len(sec)))
	le.PutUint32(s[20:], hdr+shdr)
	le.PutUint32(s[24:], relOff)
	le.PutUint16(s[32:], uint16(len(relocs)))
	le.PutUint32(s[36:], 0x40000040) // initialised data, readable
	copy(b[hdr+shdr:], sec)
	for i, r := range relocs {
		p := relOff + 10*uint32(i)
		le.PutUint32(b[p:], r)
		le.PutUint32(b[p+4:], 0)
		le.PutUint16(b[p+8:], relType)
	}
	sym := b[symOff:]
	copy(sym[0:8], ".rsrc")
	le.PutUint16(sym[12:], 1) // section number
	sym[16] = 3               // IMAGE_SYM_CLASS_STATIC
	le.PutUint32(b[symOff+18:], 4)
	return b
}

// versionInfo builds a VS_VERSIONINFO block so Explorer and Task Manager show the app's name and version.
func versionInfo(version, name, exe string) []byte {
	v := [4]uint16{}
	for i, p := range strings.Split(version, ".") {
		if i < 4 {
			n, _ := strconv.Atoi(p)
			v[i] = uint16(n)
		}
	}
	fixed := make([]byte, 52)
	le.PutUint32(fixed[0:], 0xFEEF04BD)
	le.PutUint32(fixed[4:], 0x00010000)
	ms, ls := uint32(v[0])<<16|uint32(v[1]), uint32(v[2])<<16|uint32(v[3])
	le.PutUint32(fixed[8:], ms)
	le.PutUint32(fixed[12:], ls)
	le.PutUint32(fixed[16:], ms)
	le.PutUint32(fixed[20:], ls)
	le.PutUint32(fixed[24:], 0x3F)
	le.PutUint32(fixed[32:], 0x40004) // VOS_NT_WINDOWS32
	le.PutUint32(fixed[36:], 1)       // VFT_APP
	str := func(k, val string) []byte { u := utf16z(val); return node(k, 1, u, uint16(len(u)/2)) }
	table := node("040904B0", 1, nil, 0,
		str("CompanyName", name), str("FileDescription", name), str("FileVersion", version),
		str("InternalName", strings.TrimSuffix(exe, ".exe")), str("OriginalFilename", exe),
		str("ProductName", name), str("ProductVersion", version))
	trans := []byte{0x09, 0x04, 0xB0, 0x04}
	return node("VS_VERSION_INFO", 0, fixed, 52,
		node("StringFileInfo", 1, nil, 0, table),
		node("VarFileInfo", 1, nil, 0, node("Translation", 0, trans, 4)))
}

func node(key string, typ uint16, value []byte, valueLen uint16, children ...[]byte) []byte {
	b := append(make([]byte, 6), utf16z(key)...)
	b = pad4(b)
	b = append(b, value...)
	for _, c := range children {
		b = pad4(b)
		b = append(b, c...)
	}
	le.PutUint16(b[0:], uint16(len(b)))
	le.PutUint16(b[2:], valueLen)
	le.PutUint16(b[4:], typ)
	return b
}

func utf16z(s string) []byte {
	u := append(utf16.Encode([]rune(s)), 0)
	b := make([]byte, 2*len(u))
	for i, c := range u {
		le.PutUint16(b[2*i:], c)
	}
	return b
}

func pad4(b []byte) []byte {
	for len(b)%4 != 0 {
		b = append(b, 0)
	}
	return b
}

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "mkres:", err)
		os.Exit(1)
	}
}
