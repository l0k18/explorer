package utl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const NEW_ADDRESS = "http://localhost"
const DEFAULT_MAX_AGE = 60

func RegisterHandler(pattern string, handler http.HandlerFunc) {
	http.HandleFunc(pattern, EnableCors(handler))
}

func RedirectToSecureVersion(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, NEW_ADDRESS+r.URL.Path, http.StatusMovedPermanently)
}

func IsInsecureRequest(r *http.Request) bool {
	return r.TLS == nil && !strings.HasPrefix(r.Host, "localhost")
}

func isFormPostRequest(method string, w http.ResponseWriter) bool {
	if method != "POST" {
		http.Error(w, "post only", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func EnableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := GetOrigin(r)
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		sourceOrigin := GetSourceOrigin(r)
		if sourceOrigin == "" {
			next.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Access-Control-Expose-Headers", "AMP-Access-Control-Allow-Source-Origin")
		w.Header().Set("AMP-Access-Control-Allow-Source-Origin", sourceOrigin)
		next.ServeHTTP(w, r)
	}
}

func GetOrigin(r *http.Request) string {
	origin := r.Header.Get("Origin")
	if origin != "" {
		return origin
	}
	if r.Header.Get("amp-same-origin") == "true" {
		return GetSourceOrigin(r)
	}
	return "*"
}

func GetSourceOrigin(r *http.Request) string {
	return r.URL.Query().Get("__amp_source_origin")
}

func GetHost(r *http.Request) string {
	if r.TLS == nil {
		return "http://" + r.Host
	}
	return "https://" + r.Host
}

func SetVary(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		canonical(w.Header(), "vary")
		add(w.Header(), "vary", "Accept")
		add(w.Header(), "vary", "AMP-Cache-Transform")
		h.ServeHTTP(w, r)
	})
}

func SetContentTypeJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func SendJsonResponse(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SetContentTypeJson(w)
	w.Write(jsonData)
}

func SendJsonError(w http.ResponseWriter, code int, data interface{}) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SetContentTypeJson(w)
	w.WriteHeader(code)
	w.Write(jsonData)
}

func SendAmpListItems(w http.ResponseWriter, data ...interface{}) {
	SendJsonResponse(w, map[string]interface{}{
		"items": data,
	})
}

func SendJsonFile(w http.ResponseWriter, filePath string) {
	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	SetContentTypeJson(w)
	w.Write(jsonData)
}

func SetDefaultMaxAge(w http.ResponseWriter) {
	SetMaxAge(w, DEFAULT_MAX_AGE)
}

func SetMaxAge(w http.ResponseWriter, age int) {
	w.Header().Set("cache-control", fmt.Sprintf("max-age=%d, public, must-revalidate", age))
}

func onlyPost(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "post only", http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	}
}

func canonical(h http.Header, key string) {
	v := h[http.CanonicalHeaderKey(key)]
	var a []string
	for _, vv := range v {
		if vv != "" {
			a = append(a, stringMap(strings.Split(vv, ","), strings.TrimSpace)...)
		}
	}
	if len(a) != 0 {
		h[http.CanonicalHeaderKey(key)] = []string{strings.Join(a, ", ")}
	}
}

func add(h http.Header, key string, value string) {
	v := h[http.CanonicalHeaderKey(key)]
	if len(v) == 0 {
		h[http.CanonicalHeaderKey(key)] = []string{value}
	} else {
		a := stringMap(strings.Split(v[0], ","), strings.TrimSpace)
		for _, vv := range a {
			if vv == value {
				return
			}
		}
		h[http.CanonicalHeaderKey(key)] = []string{strings.Join(append(a, value), ", ")}
	}
}

func stringMap(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
