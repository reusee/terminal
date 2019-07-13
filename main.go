package main

/*
#cgo LDFLAGS: -lmimalloc
*/
import "C"

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
--term:set_cjk_ambiguous_width(2)
term:set_colors(

	-- solarized light
	--Gdk.RGBA().parse('#657B83'),
	--Gdk.RGBA().parse('#FDF6E3'),
	--{
	--	Gdk.RGBA().parse('#073642'),
	--	Gdk.RGBA().parse('#DC322F'),
	--	Gdk.RGBA().parse('#859900'),
	--	Gdk.RGBA().parse('#B58900'),
	--	Gdk.RGBA().parse('#268BD2'),
	--	Gdk.RGBA().parse('#D33682'),
	--	Gdk.RGBA().parse('#2AA198'),
	--	Gdk.RGBA().parse('#EEE8D5'),
	--	Gdk.RGBA().parse('#002B36'),
	--	Gdk.RGBA().parse('#CB4B16'),
	--	Gdk.RGBA().parse('#586E75'),
	--	Gdk.RGBA().parse('#657B83'),
	--	Gdk.RGBA().parse('#839496'),
	--	Gdk.RGBA().parse('#6C71C4'),
	--	Gdk.RGBA().parse('#93A1A1'),
	--	Gdk.RGBA().parse('#FDF6E3')
	--}

	-- "Hemisu Light"
	--Gdk.RGBA().parse('#444444'),
	--Gdk.RGBA().parse('#EFEFEF'),
	--{
	--	Gdk.RGBA().parse('#777777'),
	--	Gdk.RGBA().parse('#FF0055'),
	--	Gdk.RGBA().parse('#739100'),
	--	Gdk.RGBA().parse('#503D15'),
	--	Gdk.RGBA().parse('#538091'),
	--	Gdk.RGBA().parse('#5B345E'),
	--	Gdk.RGBA().parse('#538091'),
	--	Gdk.RGBA().parse('#999999'),
	--	Gdk.RGBA().parse('#999999'),
	--	Gdk.RGBA().parse('#D65E76'),
	--	Gdk.RGBA().parse('#9CC700'),
	--	Gdk.RGBA().parse('#947555'),
	--	Gdk.RGBA().parse('#9DB3CD'),
	--	Gdk.RGBA().parse('#A184A4'),
	--	Gdk.RGBA().parse('#85B2AA'),
	--	Gdk.RGBA().parse('#BABABA')
	--}

	-- Tomorrow
	--Gdk.RGBA().parse('#4D4D4C'),
	--Gdk.RGBA().parse('#E0E0E0'),
	--{
	--	Gdk.RGBA().parse('#000000'),
	--	Gdk.RGBA().parse('#C82828'),
	--	Gdk.RGBA().parse('#718C00'),
	--	Gdk.RGBA().parse('#EAB700'),
	--	Gdk.RGBA().parse('#4171AE'),
	--	Gdk.RGBA().parse('#8959A8'),
	--	Gdk.RGBA().parse('#3E999F'),
	--	Gdk.RGBA().parse('#FFFEFE'),
	--	Gdk.RGBA().parse('#000000'),
	--	Gdk.RGBA().parse('#C82828'),
	--	Gdk.RGBA().parse('#708B00'),
	--	Gdk.RGBA().parse('#E9B600'),
	--	Gdk.RGBA().parse('#4170AE'),
	--	Gdk.RGBA().parse('#8958A7'),
	--	Gdk.RGBA().parse('#3D999F'),
	--	Gdk.RGBA().parse('#FFFEFE')
	--}

	-- molokai
	Gdk.RGBA().parse('#BBBBBB'),
	Gdk.RGBA().parse('1b1d1e'),
	{
		Gdk.RGBA().parse('#232323'),
		Gdk.RGBA().parse('#7325FA'),
		Gdk.RGBA().parse('#23E298'),
		Gdk.RGBA().parse('#60D4DF'),
		Gdk.RGBA().parse('#D08010'),
		Gdk.RGBA().parse('#FF0087'),
		Gdk.RGBA().parse('#D0A843'),
		Gdk.RGBA().parse('#BBBBBB'),
		Gdk.RGBA().parse('#555555'),
		Gdk.RGBA().parse('#9D66F6'),
		Gdk.RGBA().parse('#5FE0B1'),
		Gdk.RGBA().parse('#6DF2FF'),
		Gdk.RGBA().parse('#FFAF00'),
		Gdk.RGBA().parse('#FF87AF'),
		Gdk.RGBA().parse('#FFCE51'),
		Gdk.RGBA().parse('#FFFFFF')
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

local accel_group = Gtk.AccelGroup{}
window:add_accel_group(accel_group)
accel_group:connect(118, -- 'v'
  Gdk.ModifierType.CONTROL_MASK,
  Gtk.AccelFlags.LOCKED,
  GObject.Closure(function()
    term:paste_clipboard()
  end)
)

window:show_all()

Gtk.main()
	`)
}
