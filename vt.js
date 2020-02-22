#!/usr/bin/gjs

imports.gi.versions.Gtk = '3.0'
const {Vte, Gtk, Pango, Gdk} = imports.gi

Gtk.init(null)

const win = new Gtk.Window({
  type: Gtk.WindowType.TOPLEVEL,
});
win.connect('delete-event', () => true)

const term = new Vte.Terminal({
  cursor_shape: Vte.CursorShape.BLOCK,
  cursor_blink_mode: Vte.CursorBlinkMode.OFF,
  scrollback_lines: -1,
  scroll_on_output: false,
  scroll_on_keystroke: true,
  rewrap_on_resize: true,
  encoding: 'UTF-8',
  allow_bold: true,
  bold_is_bright: true,
  allow_hyperlink: true,
  font_desc: Pango.FontDescription.from_string('Terminus (TTF) 12'),
  pointer_autohide: true,
})

term.spawn_sync(
	Vte.PtyFlags.DEFAULT,
	'/dev/shm',
	[Vte.get_user_shell()],
  [],
	0,
  () => { },
  null,
)
term.connect('child-exited', () => {
  Gtk.main_quit()
})
term.connect('button-press-event', (_, ev) => {
  if (ev.get_button()[1] == 3) {
    term.copy_clipboard()
  }
})

win.add(term)

win.show_all()
Gtk.main()

