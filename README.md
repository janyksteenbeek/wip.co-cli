# wip.co CLI

This CLI tool allows users to interact with the WIP.co GraphQL API for managing todos. It supports adding new todos and posting completed todos.

## Install

1. Download the latest release from the [releases page](https://github.com/janyksteenbeek/wipco-cli/releases).
2. Rename the binary to `wip` (or anything else you'd like).
3. Place the binary in your `$PATH` (e.g. `/usr/local/bin`).
4. Set up your API key (see below).

## Usage

### Setting the API Key
```bash
wip -apikey YOUR_API_KEY
```
This command will save your API key to `~/.wipcoconfig` for future use. You can fetch your API key [here](https://wip.co/api).

### Adding a completed todo
```bash
wip "Your completed todo message here"
```
Adds a new completed item. This is the default command, so there is no flag required.

### Adding a Todo
```bash
wip -todo "Your todo message here"
```
Adds a new todo with the provided message.

## Configuration
The tool saves the API key in a file located at `~/.wipcoconfig`. Please find the above section on setting the API key for more information.

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Sponsor
If you like the work I do, please consider [sponsoring me on GitHub Sponsors](https://github.com/sponsors/janyksteenbeek) or [buy me a coffee through Ko-fi](https://ko-fi.com/janyk/shop).

## Joining WIP.co / invitations
If you'd like to join WIP.co, you can apply for an invitation [here](https://wip.co/applicants/new).

