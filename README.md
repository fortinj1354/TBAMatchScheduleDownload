# TBA Match Schedule Download
Downloads match schedules from The Blue Alliance and exports the result to a CSV. 
Adds a blank column after each team number to assign a team member to scout the match.

## Usage Guide

Get the schedule for an event:

`TBAMatchScheduleDownload.exe -event 2019gadal -key someapikey`

Get the schedule for an event and filter for a specific team number:

`TBAMatchScheduleDownload.exe -event 2019gadal -team 2974 -key someapikey`

View help:

`TBAMatchScheduleDownload.exe --help` 

The CSV will be placed in the directory the command is run in with a name based on the event and team number.
If there is no event schedule available the resulting CSV file will be empty.

## Command Line Arguments

- key (required)
  - API key for The Blue Alliance Read API v3 
  - https://www.thebluealliance.com/apidocs
- event (required)
  - Event code from The Blue Alliance.
  - Example: For https://www.thebluealliance.com/event/2019gadal the event code is 2019gadal
- team (optional)
  - Filter the schedule for a specific team number
  - Example: 2974

## Building

Windows builds available on the [Releases](https://github.com/fortinj1354/TBAMatchScheduleDownload/releases) page.

Built on Go 1.11.5, can be built for any OS by using the standard [build](https://golang.org/pkg/go/build/) command.
