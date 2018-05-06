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

local css_provider = Gtk.CssProvider()
css_provider:load_from_data([[
vte-terminal {
    padding: 10px 20px;
}
]])
Gtk.StyleContext.add_provider_for_screen(
	Gdk.Screen.get_default(),
	css_provider,
	Gtk.STYLE_PROVIDER_PRIORITY_USER)

local window = Gtk.Window{type = Gtk.WindowType.TOPLEVEL}
window.on_delete_event = function()
	return true
end

local term = Vte.Terminal.new()
term:set_cursor_shape(Vte.CursorShape.BLOCK)
term:set_cursor_blink_mode(Vte.CursorBlinkMode.OFF)
term:set_font(Pango.FontDescription.from_string('xos4 Terminus 12'))
term:set_color_cursor(Gdk.RGBA().parse('#fcaf17'))
term:set_color_cursor_foreground(Gdk.RGBA().parse('black'))
term:set_scrollback_lines(-1)
term:set_scroll_on_output(false)
term:set_scroll_on_keystroke(true)
term:set_rewrap_on_resize(true)
term:set_encoding('UTF-8')
term:set_allow_bold(true)
term:set_allow_hyperlink(true)
term:set_mouse_autohide(true)
term:set_cjk_ambiguous_width(2)
term:set_colors(
	nil,
	nil,
	{

		-- solarized dark
		--Gdk.RGBA().parse('#002b36'),
		--Gdk.RGBA().parse('#073642'),
		--Gdk.RGBA().parse('#586e75'),
		--Gdk.RGBA().parse('#657b83'),
		--Gdk.RGBA().parse('#839496'),
		--Gdk.RGBA().parse('#93a1a1'),
		--Gdk.RGBA().parse('#eee8d5'),
		--Gdk.RGBA().parse('#fdf6e3')

	}
)

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
