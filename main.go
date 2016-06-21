package main

import (
	"github.com/reusee/lgo"
	"os"
)

func main() {
	lua := lgo.NewLua()

	lua.RegisterFunctions(map[string]interface{}{
		"Sys_exit": func() {
			os.Exit(0)
		},
	})

	lua.RunString(`
lgi = require 'lgi'
Gtk = lgi.require('Gtk', '3.0')
Gdk = lgi.Gdk
Pango = lgi.Pango
Vte = lgi.Vte

local window = Gtk.Window{type = Gtk.WindowType.TOPLEVEL}
window.on_destroy:connect(function()
	Sys_exit()
end)

local css_provider = Gtk.CssProvider()
css_provider:load_from_data([[
	VteTerminal {
		cursor-aspect-ratio: 0.8;
	}
]])
Gtk.StyleContext.add_provider_for_screen(
	Gdk.Screen.get_default(),
	css_provider,
	Gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

local term = Vte.Terminal.new()
--term:set_cursor_shape(Vte.CursorShape.IBEAM)
term:set_cursor_blink_mode(Vte.CursorBlinkMode.OFF)
term:set_font(Pango.FontDescription.from_string('Terminus 13'))
term:set_color_cursor(Gdk.RGBA().parse('#fcaf17'))
term:set_color_cursor_foreground(Gdk.RGBA().parse('black'))
term:set_scrollback_lines(-1)
term:set_scroll_on_output(true)
term:set_scroll_on_keystroke(true)
term:set_rewrap_on_resize(true)
term:set_encoding('UTF-8')
term:set_allow_bold(true)
term:spawn_sync(
	Vte.PtyFlags.DEFAULT,
	'.',
	{Vte.get_user_shell()},
	{},
	0,
	function() end,
	nil)
window:add(term)

window:show_all()

Gtk.main()
	`)
}
