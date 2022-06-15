// Code generated by girgen. DO NOT EDIT.

package flatpak

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	externglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <flatpak.h>
// #include <glib-object.h>
import "C"

// glib.Type values for flatpak-installed-ref.go.
var GTypeInstalledRef = externglib.Type(C.flatpak_installed_ref_get_type())

func init() {
	externglib.RegisterGValueMarshalers([]externglib.TypeMarshaler{
		{T: GTypeInstalledRef, F: marshalInstalledRef},
	})
}

// InstalledRefOverrider contains methods that are overridable.
type InstalledRefOverrider interface {
}

type InstalledRef struct {
	_ [0]func() // equal guard
	Ref
}

var (
	_ externglib.Objector = (*InstalledRef)(nil)
)

func classInitInstalledReffer(gclassPtr, data C.gpointer) {
	C.g_type_class_add_private(gclassPtr, C.gsize(unsafe.Sizeof(uintptr(0))))

	goffset := C.g_type_class_get_instance_private_offset(gclassPtr)
	*(*C.gpointer)(unsafe.Add(unsafe.Pointer(gclassPtr), goffset)) = data

}

func wrapInstalledRef(obj *externglib.Object) *InstalledRef {
	return &InstalledRef{
		Ref: Ref{
			Object: obj,
		},
	}
}

func marshalInstalledRef(p uintptr) (interface{}, error) {
	return wrapInstalledRef(externglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// AppdataContentRating returns the content rating field from the appdata. This
// is a potentially empty mapping of content rating attribute IDs to values, to
// be interpreted by the semantics of the content rating type (see
// flatpak_installed_ref_get_appdata_content_rating_type()).
//
// The function returns the following values:
//
//    - hashTable (optional): content rating or NULL.
//
func (self *InstalledRef) AppdataContentRating() map[string]string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.GHashTable          // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_appdata_content_rating(_arg0)
	runtime.KeepAlive(self)

	var _hashTable map[string]string // out

	if _cret != nil {
		_hashTable = make(map[string]string, gextras.HashTableSize(unsafe.Pointer(_cret)))
		gextras.MoveHashTable(unsafe.Pointer(_cret), false, func(k, v unsafe.Pointer) {
			ksrc := *(**C.gchar)(k)
			vsrc := *(**C.gchar)(v)
			var kdst string // out
			var vdst string // out
			kdst = C.GoString((*C.gchar)(unsafe.Pointer(ksrc)))
			vdst = C.GoString((*C.gchar)(unsafe.Pointer(vsrc)))
			_hashTable[kdst] = vdst
		})
	}

	return _hashTable
}

// AppdataContentRatingType returns the content rating type from the appdata.
// For example, oars-1.0 or oars-1.1.
//
// The function returns the following values:
//
//    - utf8 (optional): content rating type or NULL.
//
func (self *InstalledRef) AppdataContentRatingType() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_appdata_content_rating_type(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// AppdataLicense returns the license field from the appdata.
//
// The function returns the following values:
//
//    - utf8: license or NULL.
//
func (self *InstalledRef) AppdataLicense() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_appdata_license(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// AppdataName returns the name field from the appdata.
//
// The returned string is localized.
//
// The function returns the following values:
//
//    - utf8: name or NULL.
//
func (self *InstalledRef) AppdataName() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_appdata_name(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// AppdataSummary returns the summary field from the appdata.
//
// The returned string is localized.
//
// The function returns the following values:
//
//    - utf8: summary or NULL.
//
func (self *InstalledRef) AppdataSummary() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_appdata_summary(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// AppdataVersion returns the default version field from the appdata.
//
// The function returns the following values:
//
//    - utf8: version or NULL.
//
func (self *InstalledRef) AppdataVersion() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_appdata_version(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// DeployDir gets the deploy dir of the ref.
//
// The function returns the following values:
//
//    - utf8: deploy dir.
//
func (self *InstalledRef) DeployDir() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_deploy_dir(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Eol returns the end-of-life reason string, or NULL if the ref is not
// end-of-lifed.
//
// The function returns the following values:
//
//    - utf8: end-of-life reason or NULL.
//
func (self *InstalledRef) Eol() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_eol(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// EolRebase returns the end-of-life rebased ref, or NULL if the ref is not
// end-of-lifed.
//
// The function returns the following values:
//
//    - utf8: end-of-life rebased ref or NULL.
//
func (self *InstalledRef) EolRebase() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_eol_rebase(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// InstalledSize returns the installed size of the ref.
//
// The function returns the following values:
//
//    - guint64: installed size.
//
func (self *InstalledRef) InstalledSize() uint64 {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret C.guint64              // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_installed_size(_arg0)
	runtime.KeepAlive(self)

	var _guint64 uint64 // out

	_guint64 = uint64(_cret)

	return _guint64
}

// IsCurrent returns whether the ref is current.
//
// The function returns the following values:
//
//    - ok: TRUE if the ref is current.
//
func (self *InstalledRef) IsCurrent() bool {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret C.gboolean             // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_is_current(_arg0)
	runtime.KeepAlive(self)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LatestCommit gets the latest commit of the ref.
//
// The function returns the following values:
//
//    - utf8 (optional): latest commit.
//
func (self *InstalledRef) LatestCommit() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_latest_commit(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// Origin gets the origin of the ref.
//
// The function returns the following values:
//
//    - utf8: origin.
//
func (self *InstalledRef) Origin() string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret *C.char                // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_origin(_arg0)
	runtime.KeepAlive(self)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// Subpaths returns the subpaths that are installed, or NULL if all files
// installed.
//
// The function returns the following values:
//
//    - utf8s: strv, or NULL.
//
func (self *InstalledRef) Subpaths() []string {
	var _arg0 *C.FlatpakInstalledRef // out
	var _cret **C.char               // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))

	_cret = C.flatpak_installed_ref_get_subpaths(_arg0)
	runtime.KeepAlive(self)

	var _utf8s []string // out

	{
		var i int
		var z *C.char
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// LoadAppdata loads the compressed xml appdata for this ref (if it exists).
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//
// The function returns the following values:
//
//    - bytes containing the compressed appdata file, or NULL if an error
//      occurred.
//
func (self *InstalledRef) LoadAppdata(ctx context.Context) (*glib.Bytes, error) {
	var _arg0 *C.FlatpakInstalledRef // out
	var _arg1 *C.GCancellable        // out
	var _cret *C.GBytes              // in
	var _cerr *C.GError              // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.flatpak_installed_ref_load_appdata(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)

	var _bytes *glib.Bytes // out
	var _goerr error       // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytes, _goerr
}

// LoadMetadata loads the metadata file for this ref.
//
// The function takes the following parameters:
//
//    - ctx (optional): #GCancellable.
//
// The function returns the following values:
//
//    - bytes containing the metadata file, or NULL if an error occurred.
//
func (self *InstalledRef) LoadMetadata(ctx context.Context) (*glib.Bytes, error) {
	var _arg0 *C.FlatpakInstalledRef // out
	var _arg1 *C.GCancellable        // out
	var _cret *C.GBytes              // in
	var _cerr *C.GError              // in

	_arg0 = (*C.FlatpakInstalledRef)(unsafe.Pointer(externglib.InternObject(self).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg1 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}

	_cret = C.flatpak_installed_ref_load_metadata(_arg0, _arg1, &_cerr)
	runtime.KeepAlive(self)
	runtime.KeepAlive(ctx)

	var _bytes *glib.Bytes // out
	var _goerr error       // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _bytes, _goerr
}