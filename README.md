# daylight

a command-line program for tracking sunrise and sunset times (Mac / Linux)

![img.png](img.png)

It tells you the sunrise, sunset, solar noon times and day length. It also projects these changes over the
next ten days.

`daylight` uses your IP-based location and timezone to tailor the results to your geometry. It works in (ant)arctic
locations, and you can override the IP location if you're travelling / on a VPN.

I love the sunlight and dread the long, dark winter evenings of Northern Europe. I often look up sunrise / sunset times
and count off the days until the dreary darkness is gone.

(IP lookup is powered by https://ipinfo.io. They provide a good service so please don't spam requests.)

## Installation

### Homebrew (MacOS and Linux)

```shell
# Add my tap (formula repository)
brew tap jbreckmckye/formulae
# Install daylight
brew install daylight
# Check it
daylight --help
```

### Manual installation (and Windows)

Pick up the executable for your system in the [releases](https://github.com/jbreckmckye/daylight/releases).
Unzip the package and put the program in a folder that's within your system `PATH`.

There is a Windows binary but I haven't been able to test this. Let me know if you get it working.

## Usage

```shell
# Today's data for your IP location
daylight

# Override the IP location and timezone (allows offline operation)
daylight --latitude="-33.92" --longitude="18.42" --timezone="Africa/Johannesburg"

# Short summary of the data
daylight --short

# Data for another date
daylight --date="2025-12-31"

# Disable the colour output
NO_COLOR=true daylight

# Show help
daylight --help
```

Daylight will attempt to adapt to your terminal background colour (dark vs light) but this might not work well for your
colour scheme. In this case you can use `NO_COLOR` to disable any colours

```
$ NO_COLOR=true daylight

Today's daylight                                                            
════════════════════════════════════════════════════════════════════════════
                                                                            
                                                                      
             Rises                 Noon:                 Sets:        
             05:41                 12:58                 20:14        
                                                                      
                                                                      
                                                                      
Day length                                                                  
════════════════════════════════════════════════════════════════════════════
                                                                            
           Daylight for:                    versus yesterday:         
          14 hrs, 32 mins                        +3m 39s              
                                                                           
 .................R-------------------------------------------S........... 
                                                                           
Ten day projection                                                          
════════════════════════════════════════════════════════════════════════════
                                                                            
      ┌────────────────┬───────────┬───────────┬─────────────────────┐      
      │      DATE      │  SUNRISE  │   SUNSET  │       LENGTH        │      
      ├────────────────┼───────────┼───────────┼─────────────────────┤      
      │   Sun Apr 27   │   05:39   │   20:16   │   14 hrs, 36 mins   │      
      │   Mon Apr 28   │   05:37   │   20:17   │   14 hrs, 40 mins   │      
      │   Tue Apr 29   │   05:35   │   20:19   │   14 hrs, 43 mins   │      
      │   Wed Apr 30   │   05:33   │   20:21   │   14 hrs, 47 mins   │      
      │   Thu May 01   │   05:32   │   20:22   │   14 hrs, 50 mins   │      
      │   Fri May 02   │   05:30   │   20:24   │   14 hrs, 54 mins   │      
      │   Sat May 03   │   05:28   │   20:26   │   14 hrs, 57 mins   │      
      │   Sun May 04   │   05:26   │   20:27   │   15 hrs, 1 mins    │      
      │   Mon May 05   │   05:24   │   20:29   │   15 hrs, 4 mins    │      
      │   Tue May 06   │   05:23   │   20:31   │   15 hrs, 8 mins    │      
      └────────────────┴───────────┴───────────┴─────────────────────┘      
                                                                            
Your stats                                                                  
════════════════════════════════════════════════════════════════════════════
                                                                            
 LOCATION  Latitude 51.51, Longitude -0.1257        IP ADDRESS  146.90.4.96 
                                                                            
https://github.com/jbreckmckye/daylight                                     

```

There is also a short summary mode

```
$ daylight --short
Rises:  05:41
Sets:   20:14
Length: 14 hrs, 32 mins
Change: +3m 39s
```

## Codebase

I wrote this project to learn Go, so don't expect anything too amazing. Feel free to raise issues.

The terminal UI library is [lipgloss](https://github.com/charmbracelet/lipgloss).

```
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀ ⠀⠀⠀⠀⢀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⢠⠀⠀⠀⠀⠀⠀⠀⠀⢰⣿⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠸⣷⣦⣀⠀⠀⠀⠀⠀⣿⣿⡇⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠙⣿⣿⣿⣦⠀⠠⠾⠿⣿⣷⠀⠀⠀⠀⠀⣠⣤⣄⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⢸⠟⢉⣠⣤⣶⡆⠀⣠⣈⠀⢀⣠⣴⣿⣿⠋⠀⠀⠀⠀
⠀⢀⡀⢀⣀⣀⣠⣤⡄⢀⣀⡘⣿⣿⣿⣷⣼⣿⣿⣷⡄⠹⣿⡿⠁⠀⠀⠀⠀⠀
⠀⠀⠻⠿⢿⣿⣿⣿⠁⣼⣿⣿⣿⣿⣿⣿⣿⣿⣿⣟⣁⠀⠋⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠈⠻⠀⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⡇⢰⣄⣀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⣠⡀⠀⣴⣿⣿⣿⣿⣿⣿⣿⡿⢿⡿⠀⣾⣿⣿⣿⣿⣶⡄⠀
⠀⠀⠀⠀⠀⢀⣾⣿⣷⡀⠻⣿⣿⡿⠻⣿⣿⣿⣿⠀⠀⠈⠉⠉⠉⠀⠀⠀⠀⠀
⠀⠀⠀⠀⣠⣾⡿⠟⠉⠉⠀⢀⡉⠁⠀⠛⠛⢉⣠⣴⣧⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠈⠉⠉⠀⠀⠀⠀⠀⢸⣿⣿⡿⠉⠀⠙⠿⣿⣿⣧⡀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠘⣿⣿⠁⠀⠀⠀⠀⠀⠙⠿⣷⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣿⠟⠀⠀⠀⠀⠀⠀⠀⠀ ⠃⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
```
