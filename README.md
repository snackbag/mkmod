# mkmod

mkmod is a command line tool to quickly create mod templates. It is still under heavy development and may not support
everything yet. If you want a specific feature to be added, create a new issue! We'll be more than happy to add it :]

## Installation (or updating)

| Operating System | Architecture  | Command                                                                                                                     |
|------------------|---------------|-----------------------------------------------------------------------------------------------------------------------------|
| macOS            | Intel x86-64  | `sudo curl -L -o /usr/local/bin/mkmod https://snackbag.net/mkmod/latest-macOS-amd64 && sudo chmod a+x /usr/local/bin/mkmod` |
| macOS            | Apple Silicon | `sudo curl -L -o /usr/local/bin/mkmod https://snackbag.net/mkmod/latest-macOS-arm64 && sudo chmod a+x /usr/local/bin/mkmod` |
| Windows          | 64-bit x86-64 | `curl -L -o mkmod.exe https://snackbag.net/mkmod/latest-windows-amd64.exe && move /y mkmod.exe "%windir%\System32"`         |
| Windows          | 64-bit ARM    | `curl -L -o mkmod.exe https://snackbag.net/mkmod/latest-windows-arm64.exe && move /y mkmod.exe "%windir%\System32"`         |

**De-installation**\
macOS: `rm -f /usr/local/bin/mkmod`\
Windows: `del /f /q "%windir%\System32\mkmod.exe"`

## Usage

**Create a new mod or plugin**\
You can use `mkmod -name "<mod name>" <mod id> <java package> <main class name>` to generate a Fabric mod template for
the latest Minecraft version.

Example:\
`mkmod -name "Example Mod" examplemod me.jxsnack.example ExampleMod`

---

If you want to change platform or version, first make sure mkmod supports it. If it doesn't, but you want it to support
it, [create a new issue](https://github.com/snackbag/mkmod/issues). If it does exist, use the `-platform <platform>` and
`-version <version>` parameters. Example:\
`mkmod -platform fabric -version 1.20.1 -name "Example Mod" examplemod me.jxsnack.example ExampleMod`

mkmod will *always* create a new directory when making a mod template, and if there is already a directory with the same
name, it will not build. The directory name is the mod's name. mkmod also checks the mod id, package and main class name
if they fit the conventions/rules before creating the template. If you find any edge cases, please report them

<details>
<summary>Supported Platforms and Versions</summary>

Do you want support for any versions or mod loaders? Open a new issue! We'll be more than glad to add it

**Mod loader support**

- [X] [Fabric](https://fabricmc.net)
- [ ] [Quilt](https://quiltmc.org/)
- [ ] [NeoForge](https://neoforged.net/)
- [ ] [Forge](https://files.minecraftforge.net/net/minecraftforge/forge/)

**Plugin loader support** (planned for the future)

- [ ] [Paper](https://papermc.io/software/paper)
- [ ] [Spigot](https://www.spigotmc.org/)
- [ ] [Sponge](https://spongepowered.org/)

**Proxy support** (planned for the future)

- [ ] [Velocity](https://papermc.io/software/velocity)
- [ ] [Bungeecord](https://www.spigotmc.org/wiki/bungeecord/)

| Version | Fabric | Quilt    | NeoForge | Forge    |
|---------|--------|----------|----------|----------|
| 1.20.x  | ✅ Yes  | ❌ Coming | ❌ Coming | ❌ Coming |
| 1.21.x  | ✅ Yes  | ❌ Coming | ❌ Coming | ❌ Coming |

</details>

## Build it yourself

Use `go build` to build for your own OS (e.g. for testing/small changes (don't forget the AGPL-3.0 license
restrictions)\
Use `build.sh` to build for all operating systems

Simpul :D