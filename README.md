# mkmod

mkmod is a command line tool to quickly create mod templates. It is still under heavy development and does not work yet.

## Installation

**IMPORTANT:** There is no automatic updater/notifier yet, so please make sure to check you're using the latest version
when finding any issues.
***
| Operating System | Architecture  | Command                                                                                                                     |
|------------------|---------------|-----------------------------------------------------------------------------------------------------------------------------|
| macOS            | Intel x86-64  | `sudo curl -L -o /usr/local/bin/mkmod https://snackbag.net/mkmod/latest-macOS-amd64 && sudo chmod a+x /usr/local/bin/mkmod` |
| macOS            | Apple Silicon | `sudo curl -L -o /usr/local/bin/mkmod https://snackbag.net/mkmod/latest-macOS-arm64 && sudo chmod a+x /usr/local/bin/mkmod` |
| Windows          | 64-bit x86-64 | `curl -L -o mkmod.exe https://snackbag.net/mkmod/latest-windows-amd64.exe && move /y mkmod.exe "%windir%\System32"`         |
| Windows          | 64-bit ARM    | `curl -L -o mkmod.exe https://snackbag.net/mkmod/latest-windows-arm64.exe && move /y mkmod.exe "%windir%\System32"`         |

**De-installation**\
macOS: `rm -f /usr/local/bin/mkmod`\
Windows: `del /f /q "%windir%\System32\mkmod.exe"`

## Build it yourself

Use `go build`\
Simpul :D
