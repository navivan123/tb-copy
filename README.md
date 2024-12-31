# tb-copy
an attempt at a go version of the bot used by twitch streamer kenta https://github.com/komdog/KomDog-Twitch-Bot

## Installation Requirements
### MPG123 Instructions
#### Windows
Download MPG123: https://www.mpg123.de/download/win64/1.32.10/ \
Add MPG123's folder location to the PATH environment variable
#### Linux
Use pacman -S mpg123 to install mpg123 on arch.  
Use apt for Ubuntu... etc etc etc

It should automatically add it to the PATH and the app will use it.


### Environment Instructions

#### Linux:
Create a file .env on app directory called .env
List out your api key on the file like so:
ELABS_API="<API_KEY_HERE>"

#### Windows:
Set an environment variable named "ELABS_API" with the ELABS API KEY you use on your account
