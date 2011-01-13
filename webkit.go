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

static WebKitWebView* to_WebKitWebView(void* w) { return WEBKIT_WEB_VIEW(w); }

static void* _webkit_web_view_new() {
	return webkit_web_view_new();
}
*/
import "C"
import "gtk"
import "glib"
import "unsafe"

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
//WEBKIT_API G_CONST_RETURN gchar *webkit_web_view_get_title (WebKitWebView *web_view);
//WEBKIT_API G_CONST_RETURN gchar *webkit_web_view_get_uri (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_set_maintains_back_forward_list (WebKitWebView *web_view, gboolean flag);
//WEBKIT_API WebKitWebBackForwardList *webkit_web_view_get_back_forward_list (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_go_to_back_forward_item (WebKitWebView *web_view, WebKitWebHistoryItem *item);
//WEBKIT_API gboolean webkit_web_view_can_go_back (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_can_go_back_or_forward (WebKitWebView *web_view, gint steps);
//WEBKIT_API gboolean webkit_web_view_can_go_forward (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_go_back (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_go_back_or_forward (WebKitWebView *web_view, gint steps);
//WEBKIT_API void webkit_web_view_go_forward (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_stop_loading (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_open (WebKitWebView *web_view, const gchar *uri);
//WEBKIT_API void webkit_web_view_reload (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_reload_bypass_cache (WebKitWebView *web_view);
////WEBKIT_API void webkit_web_view_load_uri (WebKitWebView *web_view, const gchar *uri);
//WEBKIT_API void webkit_web_view_load_string (WebKitWebView *web_view, const gchar *content, const gchar *mime_type, const gchar *encoding, const gchar *base_uri);
//WEBKIT_API void webkit_web_view_load_html_string (WebKitWebView *web_view,  const gchar *content,  const gchar *base_uri);
//WEBKIT_API void webkit_web_view_load_request (WebKitWebView *web_view, WebKitNetworkRequest *request);
//WEBKIT_API gboolean webkit_web_view_search_text (WebKitWebView *web_view, const gchar *text, gboolean case_sensitive, gboolean forward, gboolean wrap);
//WEBKIT_API guint webkit_web_view_mark_text_matches (WebKitWebView *web_view, const gchar *string, gboolean case_sensitive, guint limit);
//WEBKIT_API void webkit_web_view_set_highlight_text_matches (WebKitWebView *web_view, gboolean highlight);
//WEBKIT_API void webkit_web_view_unmark_text_matches (WebKitWebView *web_view);
//WEBKIT_API WebKitWebFrame * webkit_web_view_get_main_frame (WebKitWebView *web_view);
//WEBKIT_API WebKitWebFrame * webkit_web_view_get_focused_frame (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_execute_script (WebKitWebView *web_view, const gchar *script);
//WEBKIT_API gboolean webkit_web_view_can_cut_clipboard (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_can_copy_clipboard (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_can_paste_clipboard (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_cut_clipboard (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_copy_clipboard (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_paste_clipboard (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_delete_selection (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_has_selection (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_select_all (WebKitWebView *web_view);
//WEBKIT_API gboolean webkit_web_view_get_editable (WebKitWebView *web_view);
//WEBKIT_API void webkit_web_view_set_editable (WebKitWebView *web_view, gboolean flag);
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

func SoupUriNew(uri string) *C.SoupURI {
	ptr := C.CString(uri)
	defer C.free_string(ptr)
	return C.soup_uri_new(ptr)
}

func SoupUriFree(soup_uri *C.SoupURI) {
	C.soup_uri_free(soup_uri)
}


type SoupSession struct {
	glib.GObject
}

func GetDefaultSession() *SoupSession {
	return &SoupSession{glib.GObject{unsafe.Pointer(C.webkit_get_default_session())}}
}
