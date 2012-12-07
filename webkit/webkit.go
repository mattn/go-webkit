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

//static GtkWidget* to_GtkWidget(void* w) { return GTK_WIDGET(w); }

static WebKitWebView* to_WebKitWebView(void* w) { return WEBKIT_WEB_VIEW(w); }

static WebKitWebFrame* to_WebKitWebFrame(void* w) { return WEBKIT_WEB_FRAME(w); }

static WebKitWebSettings* to_WebKitWebSettings(void* w) { return WEBKIT_WEB_SETTINGS(w); }
*/
// #cgo pkg-config: webkit-1.0
import "C"
import "github.com/mattn/go-gtk/gtk"
import "github.com/mattn/go-gtk/glib"
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

//-----------------------------------------------------------------------
// WebView
//-----------------------------------------------------------------------
type WebView struct {
	gtk.Widget
}

func (v *WebView) getWebView() *C.WebKitWebView {
	return C.to_WebKitWebView(unsafe.Pointer(v.GWidget))
}
func NewWebView() *WebView {
	return &WebView{*gtk.WidgetFromNative(unsafe.Pointer(C.webkit_web_view_new()))}
}
func (v *WebView) LoadUri(uri string) {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	C.webkit_web_view_load_uri(v.getWebView(), C.to_gcharptr(ptr))
}
func (v *WebView) GetTitle() string {
	return C.GoString(C.to_charptr(C.webkit_web_view_get_title(v.getWebView())))
}
func (v *WebView) GetUri() string {
	return C.GoString(C.to_charptr(C.webkit_web_view_get_uri(v.getWebView())))
}
func (v *WebView) SetMaintainsBackForwardList(flag bool) {
	C.webkit_web_view_set_maintains_back_forward_list(v.getWebView(), bool2gboolean(flag))
}
//WEBKIT_API WebKitWebBackForwardList *webkit_web_view_get_back_forward_list (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_go_to_back_forward_item (WebKitWebView *web_view, WebKitWebHistoryItem *item);
func (v *WebView) CanGoBack() bool {
	return gboolean2bool(C.webkit_web_view_can_go_back(v.getWebView()))
}
func (v *WebView) CanGoBackOrForward(steps int) bool {
	return gboolean2bool(C.webkit_web_view_can_go_back_or_forward(v.getWebView(), C.gint(steps)))
}
func (v *WebView) CanGoForward() bool {
	return gboolean2bool(C.webkit_web_view_can_go_forward(v.getWebView()))
}
func (v *WebView) GoBack() {
	C.webkit_web_view_go_back(v.getWebView())
}
func (v *WebView) GoBackOrForward(steps int) {
	C.webkit_web_view_go_back_or_forward(v.getWebView(), C.gint(steps))
}
func (v *WebView) GoForward() {
	C.webkit_web_view_go_forward(v.getWebView())
}
func (v *WebView) StopLoading() {
	C.webkit_web_view_stop_loading(v.getWebView())
}
func (v *WebView) Open(uri string) {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	C.webkit_web_view_open(v.getWebView(), C.to_gcharptr(ptr))
}
func (v *WebView) Reload() {
	C.webkit_web_view_reload(v.getWebView())
}
func (v *WebView) ReloadBypassCache() {
	C.webkit_web_view_reload_bypass_cache(v.getWebView())
}
func (v *WebView) LoadString(content, mime_type, encoding, base_uri string) {
	pcontent := C.CString(content)
	defer C.free_string(pcontent)
	pmime_type := C.CString(mime_type)
	defer C.free_string(pmime_type)
	pencoding := C.CString(encoding)
	defer C.free_string(pencoding)
	pbase_uri := C.CString(base_uri)
	defer C.free_string(pbase_uri)
	C.webkit_web_view_load_string(v.getWebView(), C.to_gcharptr(pcontent), C.to_gcharptr(pmime_type), C.to_gcharptr(pencoding), C.to_gcharptr(pbase_uri))
}
func (v *WebView) LoadHtmlString(content, base_uri string) {
	pcontent := C.CString(content)
	defer C.free_string(pcontent)
	pbase_uri := C.CString(base_uri)
	defer C.free_string(pbase_uri)
	C.webkit_web_view_load_html_string(v.getWebView(), C.to_gcharptr(pcontent), C.to_gcharptr(pbase_uri))
}
//WEBKIT_API void webkit_web_view_load_request (WebKitWebView *web_view, WebKitNetworkRequest *request);
func (v *WebView) SearchText(text string, case_sensitive bool, forward bool, wrap bool) bool {
	ptext := C.CString(text)
	defer C.free_string(ptext)
	return gboolean2bool(C.webkit_web_view_search_text(v.getWebView(), C.to_gcharptr(ptext), bool2gboolean(case_sensitive), bool2gboolean(forward), bool2gboolean(wrap)))
}
//WEBKIT_API guint webkit_web_view_mark_text_matches (WebKitWebView *web_view, const gchar *string, gboolean case_sensitive, guint limit);
func (v *WebView) SetHighlightTextMatches(highlight bool) {
	C.webkit_web_view_set_highlight_text_matches(v.getWebView(), bool2gboolean(highlight))
}
func (v *WebView) UnmarkTextMatches() {
	C.webkit_web_view_unmark_text_matches(v.getWebView())
}
func (v *WebView) GetMainFrame() *WebFrame {
	return &WebFrame{glib.GObject{unsafe.Pointer(C.webkit_web_view_get_main_frame(v.getWebView()))}}
}

func (v *WebView) GetFocusedFrame() *WebFrame {
	return &WebFrame{glib.GObject{unsafe.Pointer(C.webkit_web_view_get_focused_frame(v.getWebView()))}}
}

func (v *WebView) ExecuteScript(script string) {
	pscript := C.CString(script)
	defer C.free_string(pscript)
	C.webkit_web_view_execute_script(v.getWebView(), C.to_gcharptr(pscript))
}
func (v *WebView) CanCutClipboard() bool {
	return gboolean2bool(C.webkit_web_view_can_cut_clipboard(v.getWebView()))
}
func (v *WebView) CanCopyClipboard() bool {
	return gboolean2bool(C.webkit_web_view_can_copy_clipboard(v.getWebView()))
}
func (v *WebView) CanPasteCilpboard() bool {
	return gboolean2bool(C.webkit_web_view_can_paste_clipboard(v.getWebView()))
}
func (v *WebView) CutClipboard() {
	C.webkit_web_view_cut_clipboard(v.getWebView())
}
func (v *WebView) CopyClipboard() {
	C.webkit_web_view_copy_clipboard(v.getWebView())
}
func (v *WebView) PasteCilpboard() {
	C.webkit_web_view_paste_clipboard(v.getWebView())
}
func (v *WebView) DeleteSelection() {
	C.webkit_web_view_delete_selection(v.getWebView())
}
func (v *WebView) HasSelection() bool {
	return gboolean2bool(C.webkit_web_view_has_selection(v.getWebView()))
}
func (v *WebView) SelectAll() {
	C.webkit_web_view_select_all(v.getWebView())
}
func (v *WebView) GetEditable() bool {
	return gboolean2bool(C.webkit_web_view_get_editable(v.getWebView()))
}
func (v *WebView) SetEditable(flag bool) {
	C.webkit_web_view_set_editable(v.getWebView(), bool2gboolean(flag))
}
//WEBKIT_API GtkTargetList * webkit_web_view_get_copy_target_list (WebKitWebView *web_view);
//WEBKIT_API GtkTargetList * webkit_web_view_get_paste_target_list (WebKitWebView *web_view);
func (v *WebView) SetSettings(settings *WebSettings) {
	C.webkit_web_view_set_settings(v.getWebView(), C.to_WebKitWebSettings(settings.Object));
}
func (v *WebView) GetSettings() *WebSettings {
	return &WebSettings{glib.GObject{unsafe.Pointer(C.webkit_web_view_get_settings(v.getWebView()))}}
}
//WEBKIT_API WebKitWebInspector * webkit_web_view_get_inspector (WebKitWebView *web_view);
//WEBKIT_API WebKitWebWindowFeatures* webkit_web_view_get_window_features (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_can_show_mime_type (WebKitWebView *web_view, const gchar *mime_type);
func (v *WebView) GetTransparent() bool {
	return gboolean2bool(C.webkit_web_view_get_transparent(v.getWebView()))
}
func (v *WebView) SetTransparent(flag bool) {
	C.webkit_web_view_set_transparent(v.getWebView(), bool2gboolean(flag))
}
func (v *WebView) GetZoomLevel(zoom_level float64) float64 {
	return float64(C.webkit_web_view_get_zoom_level(v.getWebView()))
}
func (v *WebView) SetZoomLevel(zoom_level float64) {
	C.webkit_web_view_set_zoom_level(v.getWebView(), C.gfloat(zoom_level))
}
func (v *WebView) ZoomIn() {
	C.webkit_web_view_zoom_in(v.getWebView())
}
func (v *WebView) ZoomOut() {
	C.webkit_web_view_zoom_out(v.getWebView())
}
func (v *WebView) GetFullContentZoom() bool {
	return gboolean2bool(C.webkit_web_view_get_full_content_zoom(v.getWebView()))
}
func (v *WebView) SetFullContentZoom(full_content_zoom bool) {
	C.webkit_web_view_set_full_content_zoom(v.getWebView(), bool2gboolean(full_content_zoom))
}
func GetDefaultSession() *SoupSession {
	return &SoupSession{glib.GObject{unsafe.Pointer(C.webkit_get_default_session())}}
}
func (v *WebView) GetEncoding() string {
	return C.GoString(C.to_charptr(C.webkit_web_view_get_encoding(v.getWebView())))
}
func (v *WebView) SetCustomEncoding(encoding string) {
	pencoding := C.CString(encoding)
	defer C.free_string(pencoding)
	C.webkit_web_view_set_custom_encoding(v.getWebView(), C.to_gcharptr(pencoding))
}
func (v *WebView) GetCustomEncoding() string {
	return C.GoString(C.webkit_web_view_get_custom_encoding(v.getWebView()))
}
//WEBKIT_API void webkit_web_view_move_cursor (WebKitWebView * webView, GtkMovementStep step, gint count);
//WEBKIT_API WebKitLoadStatus webkit_web_view_get_load_status (WebKitWebView *web_view);
func (v *WebView) GetProgress() float64 {
	return float64(C.webkit_web_view_get_progress(v.getWebView()))
}
func (v *WebView) CanUndo() bool {
	return gboolean2bool(C.webkit_web_view_can_undo(v.getWebView()))
}
func (v *WebView) Undo() {
	C.webkit_web_view_undo(v.getWebView())
}
func (v *WebView) CanRedo() bool {
	return gboolean2bool(C.webkit_web_view_can_redo(v.getWebView()))
}
func (v *WebView) Redo() {
	C.webkit_web_view_redo(v.getWebView())
}
func (v *WebView) GetViewSourceMode() bool {
	return gboolean2bool(C.webkit_web_view_get_view_source_mode(v.getWebView()))
}
func (v *WebView) SetViewSourceMode(view_source_mode bool) {
	C.webkit_web_view_set_view_source_mode(v.getWebView(), bool2gboolean(view_source_mode))
}
//WEBKIT_API WebKitHitTestResult* webkit_web_view_get_hit_test_result (WebKitWebView *webView, GdkEventButton *event);
func (v *WebView) GetIconUri() string {
	return C.GoString(C.to_charptr(C.webkit_web_view_get_icon_uri(v.getWebView())))
}
//WEBKIT_API void webkit_set_cache_model (WebKitCacheModel cache_model);
//WEBKIT_API WebKitCacheModel webkit_get_cache_model (void);

//-----------------------------------------------------------------------
// SoupURI
//-----------------------------------------------------------------------
type SoupURI struct {
	glib.WrappedObject
	value *C.SoupURI
}

func SoupUri(uri string) *SoupURI {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	return &SoupURI{ nil, C.soup_uri_new(ptr) }
}

func (v *SoupURI) GetInternalValue() unsafe.Pointer {
	return unsafe.Pointer(v.value)
}

func (v *SoupURI) Free() {
	C.soup_uri_free(v.value)
}


//-----------------------------------------------------------------------
// SoupSession
//-----------------------------------------------------------------------
type SoupSession struct {
	glib.GObject
}

//-----------------------------------------------------------------------
// WebFrame
//-----------------------------------------------------------------------
type WebFrame struct {
	glib.GObject
}
func (v *WebFrame) getWebFrame() *C.WebKitWebFrame {
	return C.to_WebKitWebFrame(unsafe.Pointer(v.Object))
}

// WebKitWebFrame *    webkit_web_frame_find_frame         (WebKitWebFrame *frame,
//                                                          const gchar *name);
// WebKitWebDataSource * webkit_web_frame_get_data_source  (WebKitWebFrame *frame);
// JSGlobalContextRef  webkit_web_frame_get_global_context (WebKitWebFrame *frame);
func (v *WebFrame) GetHorizontalScrollbarPolicy() uint {
	return uint(C.webkit_web_frame_get_horizontal_scrollbar_policy(v.getWebFrame()))
}
// WebKitLoadStatus    webkit_web_frame_get_load_status    (WebKitWebFrame *frame);
func (v *WebFrame) GetName() string {
	return C.GoString(C.to_charptr(C.webkit_web_frame_get_name(v.getWebFrame())))
}
// WebKitNetworkResponse * webkit_web_frame_get_network_response(WebKitWebFrame *frame);
// WebKitWebFrame* webkit_web_frame_get_parent(WebKitWebFrame *frame);
// WebKitWebDataSource * webkit_web_frame_get_provisional_data_source(WebKitWebFrame *frame);
// WebKitSecurityOrigin * webkit_web_frame_get_security_origin(WebKitWebFrame *frame);
func (v *WebFrame) GetTitle() string {
	return C.GoString(C.to_charptr(C.webkit_web_frame_get_title(v.getWebFrame())))
}
func (v *WebFrame) GetUri() string {
	return C.GoString(C.to_charptr(C.webkit_web_frame_get_uri(v.getWebFrame())))
}
func (v *WebFrame) GetVerticalScrollbarPolicy() uint {
	return uint(C.webkit_web_frame_get_vertical_scrollbar_policy(v.getWebFrame()))
}
func (v *WebFrame) GetWebView() *WebView {
	//return &WebKitWebView{gtk.GtkWidget{gtk.WidgetFromNative(unsafe.Pointer(C.webkit_web_frame_get_web_view(v.getWebFrame())))}}
	return nil
}
// void webkit_web_frame_load_alternate_string(WebKitWebFrame *frame, const gchar *content, const gchar *base_url, const gchar *unreachable_url);
// void webkit_web_frame_load_request(WebKitWebFrame *frame, WebKitNetworkRequest *request);
func (v *WebFrame) LoadString(content, mime_type, encoding, base_uri string) {
	pcontent := C.CString(content)
	defer C.free_string(pcontent)
	pmime_type := C.CString(mime_type)
	defer C.free_string(pmime_type)
	pencoding := C.CString(encoding)
	defer C.free_string(pencoding)
	pbase_uri := C.CString(base_uri)
	defer C.free_string(pbase_uri)
	C.webkit_web_frame_load_string(v.getWebFrame(), C.to_gcharptr(pcontent), C.to_gcharptr(pmime_type), C.to_gcharptr(pencoding), C.to_gcharptr(pbase_uri))
}
func (v *WebFrame) LoadUri(uri string) {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	C.webkit_web_frame_load_uri(v.getWebFrame(), C.to_gcharptr(ptr))
}
func NewWebFrame(view *WebView) *WebFrame {
	return &WebFrame{glib.GObject{unsafe.Pointer(C.webkit_web_frame_new(view.getWebView()))}}
}
func (v *WebFrame) Print() {
	C.webkit_web_frame_print(v.getWebFrame())
}
// GtkPrintOperationResult  webkit_web_frame_print_full    (WebKitWebFrame *frame,
//                                                          GtkPrintOperation *operation,
//                                                          GtkPrintOperationAction action,
//                                                          GError **error);
func (v *WebFrame) Reload() {
	C.webkit_web_frame_reload(v.getWebFrame())
}
func (v *WebFrame) StopLoading() {
	C.webkit_web_frame_stop_loading(v.getWebFrame())
}

//-----------------------------------------------------------------------
// WebSettings
//-----------------------------------------------------------------------
type WebSettings struct {
	glib.GObject
}

func NewWebSettings() *WebSettings {
	return &WebSettings{glib.GObject{unsafe.Pointer(C.webkit_web_settings_new())}}
}

//-----------------------------------------------------------------------
// NetworkRequest
//-----------------------------------------------------------------------
type WebKitNetworkRequest struct {
	glib.GObject
}

func NetworkRequestFromNative(p unsafe.Pointer) WebKitNetworkRequest {
	return WebKitNetworkRequest{glib.GObject{p}}
}

func (nr *WebKitNetworkRequest)URL() string {
	return C.GoString(C.to_charptr(C.webkit_network_request_get_uri((*C.WebKitNetworkRequest)(nr.GObject.Object))))
}

func (nr *WebKitNetworkRequest)SetURL(url string) {
	ptr := C.CString(url)
	defer C.free_string(ptr)
	C.webkit_network_request_set_uri((*C.WebKitNetworkRequest)(nr.GObject.Object), C.to_gcharptr(ptr))
}
