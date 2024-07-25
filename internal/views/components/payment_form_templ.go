// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func PaymentForm(price, buildTime string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<form hx-post=\"/payment\" hx-trigger=\"submit\" hx-target-error=\"\" hx-target=\"this\" class=\"flex w-full gap-6 flex-col\" id=\"payment-form\"><p class=\"text-lg font-semibold opacity-60 text-green-400\">Pay via Mpesa</p><div class=\"flex flex-col gap-8\"><div class=\"flex justify-between \"><div><label for=\"underline_select\" class=\"text-sm opacity-40\">Build time (est)</label> <input name=\"time\" disabled type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(buildTime)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/payment_form.templ`, Line: 24, Col: 23}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" id=\"underline_select\" class=\"block py-2.5 px-0 w-full text-sm  bg-transparent border-0 border-b-2  appearance-none text-gray-500 border-gray-900 focus:outline-none focus:ring-0 focus:border-gray-200 peer\"></div><div><label for=\"underline_select\" class=\"text-sm opacity-40\">Price (Ksh)</label> <input name=\"price\" disabled type=\"text\" value=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(price)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `internal/views/components/payment_form.templ`, Line: 35, Col: 19}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" id=\"underline_select\" class=\"block py-2.5 px-0 w-full text-sm  bg-transparent border-0 border-b-2  appearance-none text-gray-500 border-gray-900 focus:outline-none focus:ring-0 focus:border-gray-200 peer\"></div></div><div><label for=\"underline_select\" class=\"text-sm opacity-40\">Enter your phone number</label> <input required minlength=\"10\" maxlength=\"10\" name=\"phone\" id=\"phone-input\" type=\"text\" value=\"\" placeholder=\"0722...\" class=\"placeholder-zinc-600 block py-2.5 px-0 w-full text-sm text-gray-500 \n\t\t\t\tbg-transparent border-0 border-b-2 border-gray-200 appearance-none \n\t\t\t\tdark:text-gray-400 dark:border-gray-700 focus:outline-none focus:ring-0 focus:border-gray-200 peer\"></div></div><div class=\"bg-orange-500 bg-opacity-30 blur-3xl  drop-shadow-xl h-6 mt-8    \"></div><div class=\"flex absolute bottom-0 w-full gap-6\" id=\"btn-holder\"><button id=\"back-button\" class=\"w-full font-semibold  text-cyan-100 rounded-lg p-2 \n\t\t\t\t\t\t\tbg-none border border-orange-500 border-opacity-20\" type=\"button\">Back</button> <button id=\"submit-button\" class=\"w-full font-semibold  text-cyan-100 rounded-lg p-2 \n\t\t\t\t\t\t\tbg-green-800 bg-opacity-80\" type=\"submit\">Make payment</button></div></form>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
