// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func Form() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div class=\"backdrop\" x-show=\"openForm\" x-cloak><div class=\"modal\"><form @submit.prevent=\"openForm = !openForm\" class=\"form\" x-data=\"{ openInput: &#39;&#39;, high: &#39;&#39;, low: &#39;&#39;, close: &#39;&#39;, volume: &#39;&#39;, adj: &#39;&#39; }\" hx-post=\"/add\" hx-swap=\"afterbegin\" hx-target=\"#rows\" @htmx:after-request=\"openInput = &#39;&#39;; high = &#39;&#39;; low = &#39;&#39;; close = &#39;&#39;; volume = &#39;&#39;; adj = &#39;&#39;;\"><header class=\"form-header\"><div class=\"form-title\">Add Stock Data for the Day</div><div class=\"form-close\" @click=\"openForm = !openForm\">&times;</div></header><div class=\"add-input\"><label for=\"open-input\">Opening Price</label> <input x-mask=\"999.99\" x-model=\"openInput\" placeholder=\"000.00\" type=\"text\" name=\"open\" id=\"open-input\"></div><div class=\"add-input\"><label for=\"high-input\">Highest Price</label> <input x-mask=\"999.99\" x-model=\"high\" placeholder=\"000.00\" type=\"text\" name=\"high\" id=\"high-input\"></div><div class=\"add-input\"><label for=\"low-input\">Lowest Price</label> <input x-mask=\"999.99\" x-model=\"low\" placeholder=\"000.00\" type=\"text\" name=\"low\" id=\"low-input\"></div><div class=\"add-input\"><label for=\"close-input\">Closing Price</label> <input x-mask=\"999.99\" x-model=\"close\" placeholder=\"000.00\" type=\"text\" name=\"close\" id=\"close-input\"></div><div class=\"add-input\"><label for=\"volume-input\">Volume</label> <input x-mask=\"9999999\" x-model=\"volume\" placeholder=\"0000000\" type=\"text\" name=\"volume\" id=\"volume-input\"></div><div class=\"add-input\"><label for=\"adjClose-input\">Adjacent Closing Price</label> <input x-mask=\"999.99\" x-model=\"adj\" placeholder=\"000.00\" type=\"text\" name=\"adjClose\" id=\"adjClose-input\"></div><footer class=\"form-footer\"><button :disabled=\"!openInput || !high || !low || !close || !volume || !adj\" class=\"button form-submit\" type=\"submit\">Submit Form</button></footer></form></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
