# Roster Service

The Leagueify roster service is responsible for managing player registration,
team formation, and roster updates. The Leagueify roster service uses 
[Go][go-download] using version `1.23.0`.

## Key Functions

- Register players and maintain player profiles
- Create and manage teams within leagues.
- Assign players to teams.
- Track roster changes

## Development

### Prerequisites

- [Docker][docker-download] is installed and running.

### Getting Started

Local development of the Leagueify roster service uses docker for a consistent
development environment. Running the Leagueify roster service locally can be
accomplished using commands found in the [Makefile][repo-makefile]. During local
development changes will trigger a live reload, eliminating the requirement to
restart the docker image manually. This is accomplished by using the wonderful
tool [air][github-air]. The most common commands have been outlined below:

```bash
# start leagueify docker image
make start

# stop leagueify docker image removing attached volumes
make clean
```

The Leagueify roster service is ready for development once the banner output is
visible within the terminal. The banner blelow was created using the
[Text to ASCII Art Generator by Patorjk][patorjk-taag].

```
leagueify-roster-1  |
leagueify-roster-1  | '||'      '||''''|      |      ..|'''.|  '||'  '|' '||''''|  '||' '||''''| '||' '|'
leagueify-roster-1  |  ||        ||  .       |||    .|'     '   ||    |   ||  .     ||   ||  .     || |
leagueify-roster-1  |  ||        ||''|      |  ||   ||    ....  ||    |   ||''|     ||   ||''|      ||
leagueify-roster-1  |  ||        ||        .''''|.  '|.    ||   ||    |   ||        ||   ||         ||
leagueify-roster-1  | .||.....| .||.....| .|.  .||.  ''|...'|    '|..'   .||.....| .||. .||.       .||.
leagueify-roster-1  |
leagueify-roster-1  | '||''|.    ..|''||    .|'''.|  |''||''| '||''''|  '||''|.
leagueify-roster-1  |  ||   ||  .|'    ||   ||..  '     ||     ||  .     ||   ||
leagueify-roster-1  |  ||''|'   ||      ||   ''|||.     ||     ||''|     ||''|'
leagueify-roster-1  |  ||   |.  '|.     || .     '||    ||     ||        ||   |.
leagueify-roster-1  | .||.  '|'  ''|...|'  |'....|'    .||.   .||.....| .||.  '|'
leagueify-roster-1  |
```

[docker-download]: https://www.docker.com/get-started/
[github-air]: https://github.com/air-verse/air
[go-download]: https://go.dev/dl/
[patorjk-taag]: https://patorjk.com/software/taag/#p=display&f=Kban&t=LEAGUEIFY%0AROSTER
[repo-makefile]: ./Makefile
