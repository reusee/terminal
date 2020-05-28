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
GObject = lgi.GObject
Pango = lgi.Pango
Vte = lgi.Vte

local window = Gtk.Window{type = Gtk.WindowType.TOPLEVEL}
window.on_delete_event = function()
	return true
end

local term = Vte.Terminal.new()
term:set_cursor_shape(Vte.CursorShape.BLOCK)
term:set_cursor_blink_mode(Vte.CursorBlinkMode.OFF)
term:set_font(Pango.FontDescription.from_string('Terminus (TTF) 12'))
--term:set_font(Pango.FontDescription.from_string('Cascadia Code 10'))
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
--term:set_cjk_ambiguous_width(2)
term:set_colors(

	Gdk.RGBA().parse('#F3F3F3'),
	Gdk.RGBA().parse('#555555'),
	{
		Gdk.RGBA().parse('#555555'),
		Gdk.RGBA().parse('#FF8272'),
		Gdk.RGBA().parse('#B4FA72'),
		Gdk.RGBA().parse('#fefdc2'),
		Gdk.RGBA().parse('#add5fe'),
		Gdk.RGBA().parse('#ff8ffd'),
		Gdk.RGBA().parse('#d0d1fe'),
		Gdk.RGBA().parse('#f3f3f3'),
		Gdk.RGBA().parse('#666666'),
		Gdk.RGBA().parse('#ffc4bd'),
		Gdk.RGBA().parse('#d6fcb9'),
		Gdk.RGBA().parse('#fefdd5'),
		Gdk.RGBA().parse('#c1e3fe'),
		Gdk.RGBA().parse('#ffb1fe'),
		Gdk.RGBA().parse('#e5e6fe'),
		Gdk.RGBA().parse('#feffff')
	}

)

term:spawn_sync(
	Vte.PtyFlags.DEFAULT,
	'/dev/shm',
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

--[[
local accel_group = Gtk.AccelGroup{}
window:add_accel_group(accel_group)
accel_group:connect(118, -- 'v'
  Gdk.ModifierType.CONTROL_MASK,
  Gtk.AccelFlags.LOCKED,
  GObject.Closure(function()
    term:paste_clipboard()
  end)
)
--]]

window:show_all()

Gtk.main()
	`)
}
