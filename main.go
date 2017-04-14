package main

import (
	"os"

	"github.com/reusee/lgo"
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
window.on_delete_event = function()
	return true
end

local term = Vte.Terminal.new()
--term:set_cursor_shape(Vte.CursorShape.IBEAM)
term:set_cursor_blink_mode(Vte.CursorBlinkMode.OFF)
term:set_font(Pango.FontDescription.from_string('xos4 Terminus 13'))
term:set_color_cursor(Gdk.RGBA().parse('#fcaf17'))
term:set_color_cursor_foreground(Gdk.RGBA().parse('black'))
term:set_scrollback_lines(-1)
--term:set_scroll_on_output(true)
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
term.on_child_exited = function()
	Sys_exit()
end
term.on_button_press_event = function(widget, ev)
	if ev.button == 3 then
		term:copy_clipboard()
	end
end

window:add(term)

window:show_all()

Gtk.main()
	`)
}
