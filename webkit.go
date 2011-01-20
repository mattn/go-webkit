package webkit

/*
#ifndef uintptr
#define uintptr unsigned int*
#endif
#include <webkit/webkit.h>
#include <unistd.h>
#include <stdlib.h>
#include <stdint.h>
#include <stdarg.h>
#include <string.h>
#include <stdio.h>
#include <pthread.h>

static inline void free_string(char* s) { free(s); }

static inline gchar* to_gcharptr(const char* s) { return (gchar*)s; }

static inline char* to_charptr(const gchar* s) { return (char*)s; }

static WebKitWebView* to_WebKitWebView(void* w) { return WEBKIT_WEB_VIEW(w); }

static void* _webkit_web_view_new() {
	return webkit_web_view_new();
}
*/
import "C"
import "gtk"
import "glib"
import "unsafe"

func bool2gboolean(b bool) C.gboolean {
	if b {
		return C.gboolean(1)
	}
	return C.gboolean(0)
}
func gboolean2bool(b C.gboolean) bool {
	if b != 0 {
		return true
	}
	return false
}

type WebKitWebView struct {
	gtk.GtkWidget
}

func WebView() *WebKitWebView {
	return &WebKitWebView{gtk.GtkWidget{gtk.FromNative(C._webkit_web_view_new())}}
}
func (v *WebKitWebView) LoadUri(uri string) {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	C.webkit_web_view_load_uri(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), C.to_gcharptr(ptr))
}
// TODO
func (v *WebKitWebView) GetTitle() string {
	return C.GoString(C.to_charptr(C.webkit_web_view_get_title(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))))
}
func (v *WebKitWebView) GetUri() string {
	return C.GoString(C.to_charptr(C.webkit_web_view_get_uri(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))))
}
//WEBKIT_API void webkit_web_view_set_maintains_back_forward_list (WebKitWebView *web_view, gboolean flag);
//WEBKIT_API WebKitWebBackForwardList *webkit_web_view_get_back_forward_list (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_go_to_back_forward_item (WebKitWebView *web_view, WebKitWebHistoryItem *item);
func (v *WebKitWebView) CanGoBack() bool {
	return gboolean2bool(C.webkit_web_view_can_go_back(C.to_WebKitWebView(unsafe.Pointer(v.Widget))))
}
func (v *WebKitWebView) CanGoBackOrForward(steps int) bool {
	return gboolean2bool(C.webkit_web_view_can_go_back_or_forward(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), C.gint(steps)))
}
func (v *WebKitWebView) CanGoForward() bool {
	return gboolean2bool(C.webkit_web_view_can_go_forward(C.to_WebKitWebView(unsafe.Pointer(v.Widget))))
}
func (v *WebKitWebView) GoBack() {
	C.webkit_web_view_go_back(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) GoBackOrForward(steps int) {
	C.webkit_web_view_go_back_or_forward(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), C.gint(steps))
}
func (v *WebKitWebView) GoForward() {
	C.webkit_web_view_go_forward(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) StopLoading() {
	C.webkit_web_view_stop_loading(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) Open(uri string) {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	C.webkit_web_view_open(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), C.to_gcharptr(ptr))
}
func (v *WebKitWebView) Reload() {
	C.webkit_web_view_reload(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) ReloadBypassCache() {
	C.webkit_web_view_reload_bypass_cache(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) LoadString(content, mime_type, encoding, base_uri string) {
	pcontent := C.CString(content)
	defer C.free_string(pcontent)
	pmime_type := C.CString(mime_type)
	defer C.free_string(pmime_type)
	pencoding := C.CString(encoding)
	defer C.free_string(pencoding)
	pbase_uri := C.CString(base_uri)
	defer C.free_string(pbase_uri)
	C.webkit_web_view_load_string(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), C.to_gcharptr(pcontent), C.to_gcharptr(pmime_type), C.to_gcharptr(pencoding), C.to_gcharptr(pbase_uri))
}
func (v *WebKitWebView) LoadHtmlString(content, base_uri string) {
	pcontent := C.CString(content)
	defer C.free_string(pcontent)
	pbase_uri := C.CString(base_uri)
	defer C.free_string(pbase_uri)
	C.webkit_web_view_load_html_string(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), C.to_gcharptr(pcontent), C.to_gcharptr(pbase_uri))
}
//WEBKIT_API void webkit_web_view_load_request (WebKitWebView *web_view, WebKitNetworkRequest *request);
//WEBKIT_API gboolean webkit_web_view_search_text (WebKitWebView *web_view, const gchar *text, gboolean case_sensitive, gboolean forward, gboolean wrap);
//WEBKIT_API guint webkit_web_view_mark_text_matches (WebKitWebView *web_view, const gchar *string, gboolean case_sensitive, guint limit);
func (v *WebKitWebView) SetHighlightTextMatches(highlight bool) {
	C.webkit_web_view_set_highlight_text_matches(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), bool2gboolean(highlight))
}
func (v *WebKitWebView) UnmarkTextMatches() {
	C.webkit_web_view_unmark_text_matches(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
//WEBKIT_API WebKitWebFrame * webkit_web_view_get_main_frame (WebKitWebView *web_view);
//WEBKIT_API WebKitWebFrame * webkit_web_view_get_focused_frame (WebKitWebView *web_view);
func (v *WebKitWebView) ExecuteScript(script string) {
	pscript := C.CString(script)
	defer C.free_string(pscript)
	C.webkit_web_view_execute_script(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), C.to_gcharptr(pscript))
}
func (v *WebKitWebView) CanCutClipboard() bool {
	return gboolean2bool(C.webkit_web_view_can_cut_clipboard(C.to_WebKitWebView(unsafe.Pointer(v.Widget))))
}
func (v *WebKitWebView) CanCopyClipboard() bool {
	return gboolean2bool(C.webkit_web_view_can_copy_clipboard(C.to_WebKitWebView(unsafe.Pointer(v.Widget))))
}
func (v *WebKitWebView) CanPasteCilpboard() bool {
	return gboolean2bool(C.webkit_web_view_can_paste_clipboard(C.to_WebKitWebView(unsafe.Pointer(v.Widget))))
}
func (v *WebKitWebView) CutClipboard() {
	C.webkit_web_view_cut_clipboard(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) CopyClipboard() {
	C.webkit_web_view_copy_clipboard(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) PasteCilpboard() {
	C.webkit_web_view_paste_clipboard(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) DeleteSelection() {
	C.webkit_web_view_delete_selection(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) HasSelection() bool {
	return gboolean2bool(C.webkit_web_view_has_selection(C.to_WebKitWebView(unsafe.Pointer(v.Widget))))
}
func (v *WebKitWebView) SelectAll() {
	C.webkit_web_view_select_all(C.to_WebKitWebView(unsafe.Pointer(v.Widget)))
}
func (v *WebKitWebView) GetEditable() bool {
	return gboolean2bool(C.webkit_web_view_get_editable(C.to_WebKitWebView(unsafe.Pointer(v.Widget))))
}
func (v *WebKitWebView) SetEditable(flag bool) {
	C.webkit_web_view_set_editable(C.to_WebKitWebView(unsafe.Pointer(v.Widget)), bool2gboolean(flag))
}
//WEBKIT_API GtkTargetList * webkit_web_view_get_copy_target_list (WebKitWebView *web_view);
//WEBKIT_API GtkTargetList * webkit_web_view_get_paste_target_list (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_set_settings (WebKitWebView *web_view, WebKitWebSettings *settings);
//WEBKIT_API WebKitWebSettings * webkit_web_view_get_settings (WebKitWebView *web_view);
//WEBKIT_API WebKitWebInspector * webkit_web_view_get_inspector (WebKitWebView *web_view);
//WEBKIT_API WebKitWebWindowFeatures* webkit_web_view_get_window_features (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_can_show_mime_type (WebKitWebView *web_view, const gchar *mime_type);
//WEBKIT_API gboolean webkit_web_view_get_transparent (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_set_transparent (WebKitWebView *web_view, gboolean flag);
//WEBKIT_API gfloat webkit_web_view_get_zoom_level (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_set_zoom_level (WebKitWebView *web_view, gfloat zoom_level);
//WEBKIT_API void webkit_web_view_zoom_in (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_zoom_out (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_get_full_content_zoom (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_set_full_content_zoom (WebKitWebView *web_view, gboolean full_content_zoom);
//WEBKIT_API SoupSession* webkit_get_default_session (void);
//WEBKIT_API const gchar* webkit_web_view_get_encoding (WebKitWebView * webView);
//WEBKIT_API void webkit_web_view_set_custom_encoding (WebKitWebView * webView, const gchar * encoding);
//WEBKIT_API const char* webkit_web_view_get_custom_encoding (WebKitWebView * webView);
//WEBKIT_API void webkit_web_view_move_cursor (WebKitWebView * webView, GtkMovementStep step, gint count);
//WEBKIT_API WebKitLoadStatus webkit_web_view_get_load_status (WebKitWebView *web_view);
//WEBKIT_API gdouble webkit_web_view_get_progress (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_undo (WebKitWebView *webView);
//WEBKIT_API gboolean webkit_web_view_can_undo (WebKitWebView *webView);
//WEBKIT_API void webkit_web_view_redo (WebKitWebView *webView);
//WEBKIT_API gboolean webkit_web_view_can_redo (WebKitWebView *webView);
//WEBKIT_API void webkit_web_view_set_view_source_mode (WebKitWebView *web_view, gboolean view_source_mode);
//WEBKIT_API gboolean webkit_web_view_get_view_source_mode (WebKitWebView *web_view);
//WEBKIT_API WebKitHitTestResult* webkit_web_view_get_hit_test_result (WebKitWebView *webView, GdkEventButton *event);
//WEBKIT_API G_CONST_RETURN gchar * webkit_web_view_get_icon_uri (WebKitWebView *webView);
//WEBKIT_API void webkit_set_cache_model (WebKitCacheModel cache_model);
//WEBKIT_API WebKitCacheModel webkit_get_cache_model (void);

type SoupURI struct {
	glib.ObjectLike
}

func SoupUri(uri string) *SoupURI {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	return &SoupURI{ glib.GObject{ unsafe.Pointer(C.soup_uri_new(ptr)) }}
}

func (v *SoupURI) Free() {
	C.soup_uri_free((*C.SoupURI)(v.GetInternalValue()))
}


type SoupSession struct {
	glib.GObject
}

func GetDefaultSession() *SoupSession {
	return &SoupSession{glib.GObject{unsafe.Pointer(C.webkit_get_default_session())}}
}
