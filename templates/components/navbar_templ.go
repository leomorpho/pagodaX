// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"github.com/mikestefanello/pagoda/pkg/controller"
)

func Navbar(page *controller.Page) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<nav class=\"bg-primary text-primary-content border-b-2 dark:border-0 border-solid border-slate-500\" hx-target=\"#main-content\" hx-select=\"#main-content\" hx-swap=\"outerHTML\" hx-push-url=\"true\"><div class=\"container mx-auto px-4 py-2 flex justify-between items-center\"><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = templ.URL(page.ToURL("home"))
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"flex items-center space-x-3\"><img src=\"/files/favicon.png\" class=\"h-8\" alt=\"Logo\"> <span class=\"text-2xl font-semibold\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(page.AppName)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/navbar.templ`, Line: 17, Col: 55}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></a><!-- Search Bar --><!-- Multi-Level Dropdown --><div class=\"relative items-center\" x-data=\"{ open: false }\" x-cloak><button @click=\"open = !open\" class=\"ml-3 text-sm rounded-full hover-brightness px-3 py-2 flex items-center justify-between\">Services <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\" fill=\"currentColor\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" d=\"M5.22 8.22a.75.75 0 0 1 1.06 0L10 11.94l3.72-3.72a.75.75 0 1 1 1.06 1.06l-4.25 4.25a.75.75 0 0 1-1.06 0L5.22 9.28a.75.75 0 0 1 0-1.06Z\" clip-rule=\"evenodd\"></path></svg></button><div x-show=\"open\" @click.away=\"open = false\" class=\"absolute right-0 mt-5 py-2 w-48 rounded-md shadow-xl z-20 bg-primary\"><a @click=\"open = false\" href=\"#\" class=\"block px-4 py-2 text-sm hover-brightness\">Web Development</a> <a @click=\"open = false\" href=\"#\" class=\"block px-4 py-2 text-sm hover-brightness\">SEO Services</a> <a @click=\"open = false\" href=\"#\" class=\"block px-4 py-2 text-sm hover-brightness\">Digital Marketing</a></div></div><!-- Example Menu Items --><div class=\"hidden md:flex items-center space-x-4\"><a hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(page.ToURL("home")))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-sm hover-brightness px-3 py-2 rounded-full\">Home</a> <a hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(page.ToURL("about")))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-sm hover-brightness px-3 py-2 rounded-full\">About</a> <a hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(page.ToURL("contact")))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"text-sm hover-brightness px-3 py-2 rounded-full\">Contact</a></div><div class=\"items-center justify-center\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = ThemeToggle().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div><div class=\"relative items-center\" x-data=\"{ open: false }\" x-cloak><button type=\"button\" @click=\"open = !open\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if page.IsAuth {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"flex text-sm rounded-full md:me-0 focus:ring-4 focus:ring-gray-300 dark:focus:ring-gray-600\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(" class=\"ml-3 text-sm rounded-full hover-brightness px-3 py-2 flex items-center justify-between\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if page.IsAuth {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img class=\"w-8 h-8 rounded-full\" src=\"https://api.multiavatar.com/Binx Bond.png\" alt=\"User photo\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("Login <svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 20 20\" fill=\"currentColor\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" d=\"M5.22 8.22a.75.75 0 0 1 1.06 0L10 11.94l3.72-3.72a.75.75 0 1 1 1.06 1.06l-4.25 4.25a.75.75 0 0 1-1.06 0L5.22 9.28a.75.75 0 0 1 0-1.06Z\" clip-rule=\"evenodd\"></path></svg>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</button><!-- Dropdown Menu --><div x-show=\"open\" @click.away=\"open = false\" class=\"absolute right-0 mt-5 py-2 w-48 rounded-md shadow-xl z-20 bg-primary\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if page.IsAuth {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"px-4 py-3\"><span class=\"text-sm\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(page.AuthUser.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/navbar.templ`, Line: 100, Col: 49}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span> <span class=\"text-sm\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(page.AuthUser.Email)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `templates/components/navbar.templ`, Line: 101, Col: 50}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></div><ul class=\"py-2\"><li><a href=\"#\" class=\"block px-4 py-2 text-sm hover-brightness\">Dashboard</a></li><li><a href=\"#\" class=\"block px-4 py-2 text-sm hover-brightness\">Settings</a></li><li><a href=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 templ.SafeURL = templ.URL(page.ToURL("logout"))
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var6)))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block px-4 py-2 text-sm hover-brightness\">Sign out</a></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<ul class=\"py-2\"><li><a @click=\"open = false\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(page.ToURL("login")))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block px-4 py-2 text-sm hover-brightness\">Login</a></li><li><a @click=\"open = false\" hx-get=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(page.ToURL("register")))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"block px-4 py-2 text-sm hover-brightness\">Register</a></li></ul>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></nav>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

func search(page *controller.Page) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var7 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var7 == nil {
			templ_7745c5c3_Var7 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"relative\" x-data=\"{ modal: false }\" hx-ext=\"debug\"><input class=\"border border-gray-300 rounded-lg p-2 w-full lg:max-w-xs\" type=\"search\" placeholder=\"Search...\" @click=\"modal = true\"><div x-show=\"modal\" class=\"fixed inset-0 bg-black bg-opacity-50 z-40 flex items-center justify-center\" x-transition:enter=\"ease-out duration-300\" x-transition:enter-start=\"opacity-0\" x-transition:enter-end=\"opacity-100\" x-transition:leave=\"ease-in duration-200\" x-transition:leave-start=\"opacity-100\" x-transition:leave-end=\"opacity-0\" @click.away=\"modal = false\" @keydown.escape.window=\"modal = false\" x-cloak><div class=\"modal-content bg-white shadow-lg rounded-lg max-w-2xl w-full p-5 mx-4\" x-transition:enter=\"ease-out duration-300\" x-transition:enter-start=\"opacity-0 transform scale-90\" x-transition:enter-end=\"opacity-100 transform scale-100\" x-transition:leave=\"ease-in duration-200\" x-transition:leave-start=\"opacity-100 transform scale-100\" x-transition:leave-end=\"opacity-0 transform scale-90\" @click.stop><div class=\"flex justify-between items-center mb-4\"><h2 class=\"text-2xl font-semibold\">Search</h2><button class=\"p-4\" @click=\"modal = false\" aria-label=\"close\"><svg xmlns=\"http://www.w3.org/2000/svg\" viewBox=\"0 0 24 24\" fill=\"currentColor\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" d=\"M12 2.25c-5.385 0-9.75 4.365-9.75 9.75s4.365 9.75 9.75 9.75 9.75-4.365 9.75-9.75S17.385 2.25 12 2.25Zm-1.72 6.97a.75.75 0 1 0-1.06 1.06L10.94 12l-1.72 1.72a.75.75 0 1 0 1.06 1.06L12 13.06l1.72 1.72a.75.75 0 1 0 1.06-1.06L13.06 12l1.72-1.72a.75.75 0 1 0-1.06-1.06L12 10.94l-1.72-1.72Z\" clip-rule=\"evenodd\"></path></svg></button></div><input hx-get=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(page.ToURL("search")))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" hx-trigger=\"keyup changed delay:300ms\" hx-target=\"#abcdefg\" hx-swap=\"innerHTML\" hx-push-url=\"false\" name=\"query\" class=\"border border-gray-300 rounded-lg p-2 w-full\" type=\"search\" placeholder=\"Search...\" x-ref=\"input\"><div class=\"mt-4\" id=\"abcdefg\">YO</div></div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
