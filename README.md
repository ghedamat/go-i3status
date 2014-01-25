# go-i3status

A Golang experiment that uses [i3bar-protocol](http://i3wm.org/docs/i3bar-protocol.html) to build an interactive i3 statusbar


## Usage

To get the default behaviour just install the *go-i3status* and set it as the `status_command` in your `.i3config`


## Customize
Clone the repo and edit main.go to add/remove widgets.


## Existing widgets

* I3statusWidget (echoes the original i3status)
* EchoWidget (test widget to echo click events)
* DateWidget
* TimerWidget (start/pause/stop timer)
* OnOffWidget (PantsOn PantsOff)


