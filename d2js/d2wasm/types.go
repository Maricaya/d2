//go:build js && wasm

package d2wasm

import (
	"oss.terrastruct.com/d2/d2ast"
	"oss.terrastruct.com/d2/d2graph"
	"oss.terrastruct.com/d2/d2target"
)

type WASMResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error *WASMError  `json:"error,omitempty"`
}

type WASMError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e *WASMError) Error() string {
	return e.Message
}

type RefRangesResponse struct {
	Ranges       []d2ast.Range `json:"ranges"`
	ImportRanges []d2ast.Range `json:"importRanges"`
}

type BoardPositionResponse struct {
	BoardPath []string `json:"boardPath"`
}

type CompileRequest struct {
	FS   map[string]string `json:"fs"`
	Opts *RenderOptions    `json:"options"`
}

type RenderOptions struct {
	Layout  *string `json:"layout"`
	Sketch  *bool   `json:"sketch"`
	ThemeID *int64  `json:"themeID"`
}

type CompileResponse struct {
	FS      map[string]string `json:"fs"`
	Diagram d2target.Diagram  `json:"diagram"`
	Graph   d2graph.Graph     `json:"graph"`
}

type CompletionResponse struct {
	Items []map[string]interface{} `json:"items"`
}

type RenderRequest struct {
	Diagram *d2target.Diagram `json:"diagram"`
	Opts    *RenderOptions    `json:"options"`
}
