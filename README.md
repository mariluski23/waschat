# Waschat

Waschat is a FOSS, decentrailzed, terminal chat client using WebSockets.

## Features

- Ran by an individual, not a monopoly
- Encrypted communication
- Decentralized
- Free and Open Source
- No Ads
- Runs on the terminal

## Installation

### From Source (Recommended) (MacOS, Linux and FreeBSD)

1. Clone the repository

    ```bash
    git clone https://github.com/mariluski23/waschat.git
    cd waschat
    ```

2. Go to the `src` folder:

    ```bash
    cd src 
    ```

3. Build the project:

    ```bash
    go build
    ```

4. Test:

    ```bash
    ./waschat
    ```

5. Move the binary to your desired location:

    ```bash
    sudo mv waschat /usr/local/bin
    ```

> [!NOTE]
> Waschat is not yet available for Windows.

## Usage

Waschat requires a GPG key to be set up in order to encrypt and decrypt messages.
Set one up using the following command:

```bash
gpg --full-generate-key
```

After setting up the key, run `waschat`. Now, copy and paste your public GPG key into the input field and press enter.
Now, choose a nickname and press enter. You're all set!

## Options

`-d`, `--debug`:         Enable debug mode.
`-h`, `--help`:          Show help message.
`-v`, `--version`:       Show version information.
`-i`, `incognito`:       Chat in *incognito* mode (No nickname, but also no encryption).

## License

This project is licensed under the GNU GPLv3 license. See the [LICENSE](LICENSE) file for more information.

## Misc

```plaintext


██╗    ██╗ █████╗ ███████╗ ██████╗██╗  ██╗ █████╗ ████████╗
██║    ██║██╔══██╗██╔════╝██╔════╝██║  ██║██╔══██╗╚══██╔══╝
██║ █╗ ██║███████║███████╗██║     ███████║███████║   ██║   
██║███╗██║██╔══██║╚════██║██║     ██╔══██║██╔══██║   ██║   
╚███╔███╔╝██║  ██║███████║╚██████╗██║  ██║██║  ██║   ██║   
 ╚══╝╚══╝ ╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝   
                                                           
```
